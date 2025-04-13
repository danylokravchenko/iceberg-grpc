package grpc

import (
	"context"

	pb "github.com/danylokravchenko/iceberg-grpc/internal/pb/catalog"

	catalog "github.com/danylokravchenko/iceberg-grpc/internal/catalog"
	metrics "github.com/danylokravchenko/iceberg-grpc/internal/metrics"

	"github.com/minio/minio-go/v7"
)

// server implements the CatalogServiceServer interface.
type IcebergService struct {
	pb.UnimplementedCatalogServiceServer

	catalog  *catalog.IcebergRestCatalog
	jobQueue chan *pb.MetadataRequest
}

func NewIcebergService(minioClient *minio.Client, bucketName, icebergCatalogURL string) *IcebergService {
	s := &IcebergService{
		catalog:  catalog.NewCatalog(minioClient, bucketName, icebergCatalogURL),
		jobQueue: make(chan *pb.MetadataRequest, 100),
	}
	go s.worker()
	return s
}

// Asynchronous gRPC request processing
func (s *IcebergService) ProcessMetadata(ctx context.Context, req *pb.MetadataRequest) (*pb.MetadataResponse, error) {
	select {
	case s.jobQueue <- req:
		metrics.IncrementQueueLength()
		return &pb.MetadataResponse{
			Status:  "Queued",
			Message: "Request enqueued for processing",
		}, nil
	default:
		return &pb.MetadataResponse{
			Status:  "Rejected",
			Message: "Server queue full, try again later",
		}, nil
	}
}

func (s *IcebergService) worker() {
	for req := range s.jobQueue {
		metrics.DecrementQueueLength()
		go s.process(req)
	}
}

func (s *IcebergService) process(req *pb.MetadataRequest) {
	if err := s.catalog.CommitToIceberg(req.TableName, req.MetadataJson); err != nil {
		metrics.IncrementProcessedRequests("FAILURE_CATALOG")
	}
	_, err := s.catalog.WriteMetadataToMinIO(req.TableName, req.MetadataJson)
	if err != nil {
		metrics.IncrementProcessedRequests("FAILURE_MINIO")
	}
	metrics.IncrementProcessedRequests("SUCCESS")
}
