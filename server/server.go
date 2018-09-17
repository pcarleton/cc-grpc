package server

import (
	"cloud.google.com/go/storage"
	pb "github.com/pcarleton/cc-grpc/proto/api"
	"golang.org/x/net/context"
	"github.com/pcarleton/sheets"
  "fmt"
  "io"
	"log"
)

type server struct {
  sheetsClient *sheets.Client
}

const (
  BUCKET_NAME = "cashcoach-160218"
  SHEETS_CREDS = "tmp-client-secrets.json"
)

func NewServer() pb.ApiServer {
  sheetsClient, err := getSheetsClient()
  if err != nil {
    log.Printf("Unable to create sheets client: %s", err)
  }

	return &server{
    sheetsClient: sheetsClient,
  }
}

func (s *server) GetHealth(ctx context.Context, request *pb.GetHealthRequest) (*pb.GetHealthResponse, error) {
	return &pb.GetHealthResponse{
		Status: pb.HealthStatus_OK,
	}, nil
}

func (s *server) CreateReport(ctx context.Context, request *pb.CreateReportRequest) (*pb.CreateReportResponse, error) {
  s.testSheetsClient()

	return &pb.CreateReportResponse{
		Result: "HI",
	}, nil
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

