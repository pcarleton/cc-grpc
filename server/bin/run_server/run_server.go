package main

import (
	"flag"
	"github.com/dgrijalva/jwt-go"
	"github.com/pcarleton/cc-grpc/auth"
	"github.com/pcarleton/cc-grpc/buildinfo"
	pb "github.com/pcarleton/cc-grpc/proto/api"
	server "github.com/pcarleton/cc-grpc/server"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

const (
	port = ":5001"
)

var cert string
var key string
var insecure bool

func init() {
	flag.StringVar(&cert, "cert", "certs/localhost.crt", "Path to the cert file to use for TLS")
	flag.StringVar(&key, "key", "certs/localhost.key", "TLS cert private key")
	flag.BoolVar(&insecure, "insecure", false, "Run without TLS")
}

func AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "missing context metadata")
	}
	if len(meta["token"]) != 1 {
		return nil, status.Errorf(codes.Unauthenticated, "no token sent")
	}

	token, err := auth.VerifyGoogleJwt(meta["token"][0])
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid token")
	}

	if err = token.Claims.Valid(); err != nil {
		log.Printf("Error: %s", err)
		return nil, status.Errorf(codes.Unauthenticated, "invalid claims")

	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		log.Printf("Invalid clamis: %+v", token.Claims)
		return nil, status.Errorf(codes.Unauthenticated, "uncastable claims")
	}
	email, ok := claims["email"].(string)

	// TODO: Don't hardcode valid users
	if email != "paulcarletonjr@gmail.com" {
		log.Printf("Unknown email: %s", email)
		return nil, status.Errorf(codes.Unauthenticated, "unknown email")

	}
	//name, ok := claims["name"].(string)

	//log.Printf(email)
	//log.Printf(name)

	return handler(ctx, req)
}

func main() {
	flag.Parse()
	log.Printf("Running version: %s", buildinfo.GitCommitID())
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcOptions := []grpc.ServerOption{
		grpc.UnaryInterceptor(AuthInterceptor),
	}

	if insecure {
		log.Printf("WARNING: Running without TLS")
	} else {
		creds, err := credentials.NewServerTLSFromFile(cert, key)
		if err != nil {
			panic(err)
		}
		grpcOptions = append(grpcOptions, grpc.Creds(creds))
	}

	s := grpc.NewServer(grpcOptions...)
	apiServer := server.NewServer()
	pb.RegisterApiServer(s, apiServer)
	// Register reflection service on gRPC server.
	reflection.Register(s)
	log.Printf("Starting server on %s", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
