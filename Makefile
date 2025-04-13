# Project variables
PROTO_DIR := internal/proto
GO_PROTO_OUT := .
MAIN := cmd/server/main.go
BINARY := iceberg-grpc-server
MODULE := danylokravchenko/iceberg-grpc
FILES ?= $(shell find . -type f -name '*.go' -not -path "./vendor/*")
API_YAML := config/api.yaml
API_SPEC := https://github.com/apache/iceberg/raw/main/open-api/rest-catalog-open-api.yaml
API_GEN := api-gen

.PHONY: all init build run proto docker-up docker-down clean test fmt api-gen update-api

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

## Run tests
test:
	go test -v ./... -short

## Generate iceberg REST client
api-gen:
	openapi-generator generate -i $(API_SPEC) -g go -o $(API_GEN) \
		--additional-properties=packageName=icebergclient
	@sed -i '' 's/^go .*/go 1.24.2/g' $(API_GEN)/go.mod
	@sed -i '' "s#GIT_USER_ID/GIT_REPO_ID#$(MODULE)/icebergclient#g" $(API_GEN)/go.mod
	@find $(API_GEN)/test -type f -name '*.go' -exec sed -i '' "s#GIT_USER_ID/GIT_REPO_ID#$(MODULE)/icebergclient#g" {} \;
	@cd $(API_GEN) && go mod tidy

## Fetch the latest Iceberg OpenAPI Spec
update-api:
	@curl -o $(API_YAML) -fsSL $(API_SPEC)

## Start Docker services (MinIO, Nessie, gRPC server)
docker-up:
	docker-compose up --build

## Stop and remove Docker services
docker-down:
	docker-compose down

## Clean up generated files
clean:
	@rm -f $(BINARY)
	@rm -rf $(GO_PROTO_OUT)/*.pb.go
	@rm -f go.work*
	@rm -rf $(API_YAML)
	@rm -rf $(API_GEN)

## Initialize the project
init: clean update-api api-gen proto
	go work init . $(API_GEN)
	go work sync