package grpc_test

import (
	"context"
	"testing"
	"time"

	pb "github.com/danylokravchenko/iceberg-grpc/internal/pb/catalog"

	"google.golang.org/grpc"
)

func startTestGRPCServer(t *testing.T) (pb.CatalogServiceClient, func()) {

	// Connect client to server
	conn, err := grpc.NewClient("localhost:50051", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(3*time.Second))
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}

	client := pb.NewCatalogServiceClient(conn)

	// Return client and cleanup function
	return client, func() {
		conn.Close()
	}
}

func TestProcessMetadata(t *testing.T) {

	client, cleanup := startTestGRPCServer(t)
	defer cleanup()

	// Perform the gRPC call
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	req := &pb.MetadataRequest{
		TableName:    "test_table",
		MetadataJson: `{"format":"parquet"}`,
	}

	res, err := client.ProcessMetadata(ctx, req)
	if err != nil {
		t.Fatalf("gRPC call failed: %v", err)
	}

	if res.Status != "Queued" {
		t.Errorf("unexpected response: got %v, want %v", res.Status, "Queued")
	}
}
