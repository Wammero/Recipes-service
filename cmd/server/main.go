package main

import (
	"Recipes_service/cmd/migrate"
	"Recipes_service/internal/config"
	"Recipes_service/internal/repository"
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/jannden/golang-examples/grpc-with-rest/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedHelloServiceServer
}

func NewServer() *server {
	return &server{}
}

func main() {
	log.Println("🚀 Starting the application...")

	connStr, err := config.LoadDatabaseConnectionString()
	if err != nil {
		log.Fatalf("❌ Failed to load database connection string: %v", err)
	}

	log.Println("🔌 Connecting to the database...")
	repo, err := repository.New(connStr)
	if err != nil {
		log.Fatalf("❌ Failed to connect to the database: %v", err)
	}
	defer repo.Close()
	log.Println("✅ Successfully connected to the database!")

	applyMigrations(connStr)

	go startGRPCServer()
	startGRPCGateway()
}

func applyMigrations(connStr string) {
	log.Println("🔄 Checking and applying migrations...")
	m, err := migrate.CallMigrations(connStr)
	if err != nil {
		log.Fatalf("❌ Migration setup error: %v", err)
	}

	if err := m.Up(); err != nil {
		if err.Error() == "no change" {
			log.Println("✅ Migrations are already up to date.")
		} else {
			log.Fatalf("❌ Migration error: %v", err)
		}
	} else {
		log.Println("✅ Migrations successfully applied.")
	}
}

func startGRPCServer() {
	log.Println("🔧 Initializing gRPC server...")

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("❌ Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterHelloServiceServer(s, &server{})

	log.Println("🚀 Serving gRPC on 0.0.0.0:50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("❌ Failed to serve gRPC server: %v", err)
	}
}

func startGRPCGateway() {
	log.Println("🔧 Initializing gRPC-Gateway...")

	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "0.0.0.0:50051", grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("❌ Failed to dial gRPC server: %v", err)
	}

	gwmux := runtime.NewServeMux()
	if err := pb.RegisterHelloServiceHandler(ctx, gwmux, conn); err != nil {
		log.Fatalf("❌ Failed to register gRPC-Gateway: %v", err)
	}

	gwServer := &http.Server{
		Addr:    ":50052",
		Handler: gwmux,
	}

	log.Println("🌐 Serving gRPC-Gateway for REST on http://0.0.0.0:50052")
	if err := gwServer.ListenAndServe(); err != nil {
		log.Fatalf("❌ Failed to serve gRPC-Gateway server: %v", err)
	}
}
