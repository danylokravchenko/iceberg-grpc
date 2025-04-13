package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	service "github.com/danylokravchenko/iceberg-grpc/internal/grpc"
	"github.com/danylokravchenko/iceberg-grpc/internal/metrics"
	pb "github.com/danylokravchenko/iceberg-grpc/internal/pb/catalog"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"google.golang.org/grpc"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Load environment variables or configuration
	minioEndpoint := os.Getenv("MINIO_ENDPOINT")
	minioAccessKey := os.Getenv("MINIO_ACCESS_KEY")
	minioSecretKey := os.Getenv("MINIO_SECRET_KEY")
	icebergCatalogURL := os.Getenv("CATALOG_URL")

	// Initialize the MinIO client
	minioClient, err := minio.New(minioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(minioAccessKey, minioSecretKey, ""),
		Secure: false, // Assuming non-SSL, adjust as needed
	})
	if err != nil {
		log.Fatalf("Failed to create MinIO client: %v", err)
	}
	bucketName := "iceberg-data"

	// Initialize Prometheus metrics
	metrics.Initialize()
	go metrics.ExposeMetrics()

	grpcServer := grpc.NewServer()

	pb.RegisterCatalogServiceServer(grpcServer, service.NewIcebergService(minioClient, bucketName, icebergCatalogURL))

	// Listen for SIGINT and SIGTERM to gracefully shutdown the server
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// Start the gRPC server in a new goroutine
	go func() {
		listen, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("Failed to listen on port 50051: %v", err)
		}

		log.Println("Starting gRPC server on port 50051...")
		if err := grpcServer.Serve(listen); err != nil {
			log.Fatalf("Failed to start gRPC server: %v", err)
		}
	}()

	// Wait for an interrupt signal
	<-stop

	// Graceful shutdown
	log.Println("Shutting down the server...")
	grpcServer.GracefulStop()
	time.Sleep(2 * time.Second) // Give time for cleanup

	log.Println("Server gracefully stopped.")

}
