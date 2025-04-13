# Project variables
PROTO_DIR := internal/proto
GO_PROTO_OUT := catalog
MAIN := cmd/server/main.go
BINARY := iceberg-grpc-server
FILES ?= $(shell find . -type f -name '*.go' -not -path "./vendor/*")

.PHONY: all build run proto docker-up docker-down clean test fmt

## Generate protobuf files
# --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative
proto:
	protoc --go_out=$(GO_PROTO_OUT) --go-grpc_out=$(GO_PROTO_OUT) \
	       $(PROTO_DIR)/*.proto

 ## format the go source files
fmt:
	go fmt ./...
	goimports -w $(FILES)

## Build the Go binary
build:
	go build -o $(BINARY) $(MAIN)

## Run the gRPC server
run: build
	./$(BINARY)

test:
	go test -v ./... -short

## Start Docker services (MinIO, Nessie, gRPC server)
docker-up:
	docker-compose up --build

## Stop and remove Docker services
docker-down:
	docker-compose down

## Clean up generated files
clean:
	rm -f $(BINARY)
	rm -rf $(GO_PROTO_OUT)/*.pb.go