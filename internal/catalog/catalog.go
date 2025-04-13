package catalog

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
)

type MinioClient interface {
	PutObject(ctx context.Context, bucketName, objectName string, reader io.Reader, objectSize int64, opts minio.PutObjectOptions) (minio.UploadInfo, error)
}

type IcebergRestCatalog struct {
	MinioClient       MinioClient
	BucketName        string
	IcebergCatalogUrl string
}

func NewCatalog(minioClient *minio.Client, bucketName, icebergCatalogUrl string) *IcebergRestCatalog {
	ctx := context.Background()
	exists, err := minioClient.BucketExists(ctx, bucketName)
	if err != nil {
		log.Fatalf("Failed to check bucket: %v", err)
	}
	if !exists {
		err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
		if err != nil {
			log.Fatalf("Failed to create bucket: %v", err)
		}
	}
	return &IcebergRestCatalog{
		MinioClient:       minioClient,
		BucketName:        bucketName,
		IcebergCatalogUrl: icebergCatalogUrl,
	}
}

func (c *IcebergRestCatalog) WriteMetadataToMinIO(tableName, data string) (*minio.UploadInfo, error) {
	ctx := context.Background()
	objectName := fmt.Sprintf("metadata/%s/%s.json", tableName, uuid.New().String())
	metadataBytes := []byte(data)
	uploadInfo, err := c.MinioClient.PutObject(ctx, c.BucketName, objectName, bytes.NewReader(metadataBytes), int64(len(metadataBytes)), minio.PutObjectOptions{ContentType: "application/json"})
	if err != nil {
		return nil, fmt.Errorf("failed to upload metadata to MinIO: %v", err)
	}

	return &uploadInfo, nil
}

func (c *IcebergRestCatalog) CommitToIceberg(tableName, data string) error {
	url := fmt.Sprintf("%s/v1/namespaces/default/tables/%s/commit", c.IcebergCatalogUrl, tableName)
	payload := map[string]interface{}{
		"commitId": uuid.New().String(),
		"metadata": data,
	}
	jsonPayload, _ := json.Marshal(payload)
	resp, err := http.Post(url, "application/json", bytes.NewReader(jsonPayload))
	if err != nil || resp.StatusCode >= 300 {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		log.Printf("Commit to Iceberg failed for table %s: %v \n response: %s", tableName, err, bodyBytes)
	} else {
		log.Printf("Committed to Iceberg table: %s", tableName)
	}
	return err
}
