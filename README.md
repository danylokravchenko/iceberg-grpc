# gRPC proxy-server to communicate with the Iceberg REST catalog and write metadata to MinIO

## Features

- Async gRPC server with Go routines
- MinIO S3-compatible metadata storage
- Iceberg REST catalog integration
- Prometheus metrics (`http://localhost:2112/metrics`)

## Usage

Generate Iceberg REST client:

```bash
make init
```

Run locally:

```bash
make run
```

Or inside docker:

```bash
docker-compose up --build
```

## Project structure

```log
iceberg-grpc-server/
├── api-gen                        # Generated from OpenAPI
├── cmd/
│   └── server/
│       └── main.go                # Main entry point
├── config/
│   └── rest-catalog-open-api.yaml # Iceberg REST OpenAPI specification
├── internal/
│   ├── catalog/
│   │   ├── catalog.go             # Catalog client
│   ├── grpc/
│   │   └── service.go             # gRPC service implementation
│   ├── metrics/
│   │   └── metrics.go             # Prometheus metric setup
├── proto/
│   └── catalog.proto              # Protobuf definition
├── go.mod
├── go.sum
├── Dockerfile
├── docker-compose.yml
└── README.md
```
