# Use the official Golang image to build the Go application
FROM golang:1.24-alpine as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all the dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod tidy

# Copy the entire project
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o iceberg-grpc-server cmd/server/main.go

# Start a new stage to reduce the final image size
FROM alpine:latest

# Install the necessary libraries for running the application
RUN apk --no-cache add ca-certificates

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the pre-built binary from the builder stage
COPY --from=builder /app/iceberg-grpc-server .

# Expose the ports the application will use
EXPOSE 50051
EXPOSE 2112

# Command to run the executable
CMD ["./iceberg-grpc-server"]
