package server

import (
	"cloud.google.com/go/storage"
	"fmt"
	"github.com/pcarleton/cc-grpc/lib"
	pb "github.com/pcarleton/cc-grpc/proto/api"
	"github.com/pcarleton/sheets"
	"golang.org/x/net/context"
	"io"
	"log"
)

type server struct {
	sheetsClient *sheets.Client
	config       *lib.Config
}

const (
	BUCKET_NAME  = "cashcoach-160218"
	SHEETS_CREDS = "tmp-client-secrets.json"
	CONFIG_YAML  = "tmp-config.yaml"
)

func NewServer() pb.ApiServer {
	sheetsClient, err := getSheetsClient()
	if err != nil {
		log.Printf("Unable to create sheets client: %s", err)
	}

	config, err := getServerConfig()
	if err != nil {
		log.Printf("Unable to load config: %s", err)
	}

	return &server{
		sheetsClient: sheetsClient,
		config:       config,
	}
}

func (s *server) GetHealth(ctx context.Context, request *pb.GetHealthRequest) (*pb.GetHealthResponse, error) {
	return &pb.GetHealthResponse{
		Status: pb.HealthStatus_OK,
	}, nil
}

func (s *server) CreateReport(ctx context.Context, request *pb.CreateReportRequest) (*pb.CreateReportResponse, error) {
	if s.sheetsClient != nil {
		s.testSheetsClient()
	} else {
		log.Printf("Skipping sheets client.")
	}

	email := "unset"
	if s.config != nil {
		email = s.config.Email
		log.Printf("Config email: %s", s.config.Email)
	} else {
		log.Printf("Config is nil.")
	}

	return &pb.CreateReportResponse{
		Result: fmt.Sprintf("Saw : %+v,  email: %s", *request, email),
	}, nil
}

func getServerConfig() (*lib.Config, error) {
	// TODO: Don't hard code this
	r, err := readBucketContents(BUCKET_NAME, CONFIG_YAML)
	if err != nil {
		return nil, fmt.Errorf("Unable to read credentials from gs://%s/%s : %s", BUCKET_NAME, SHEETS_CREDS, err)
	}
	defer r.Close()

	config, err := lib.NewConfig(r)
	return config, err
}

func getSheetsClient() (*sheets.Client, error) {
	// TODO: Don't hard code this
	r, err := readBucketContents(BUCKET_NAME, SHEETS_CREDS)
	if err != nil {
		return nil, fmt.Errorf("Unable to read credentials from gs://%s/%s : %s", BUCKET_NAME, SHEETS_CREDS, err)
	}
	defer r.Close()

	client, err := sheets.NewServiceAccountClient(r)
	return client, err
}

func (s *server) testSheetsClient() {
	files, err := s.sheetsClient.ListFiles("")
	if err != nil {
		log.Printf("Error listing files: %s", err)
		return
	}

	log.Printf("Successfully listed files! %+s", files)
}

func readBucketContents(bucketID, object string) (io.ReadCloser, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()
	bucket := client.Bucket(bucketID)

	rc, err := bucket.Object(object).NewReader(ctx)
	return rc, err
}
