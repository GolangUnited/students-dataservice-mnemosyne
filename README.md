## See swagger for API endpoints
http://localhost:8000/swagger/index.html

## Requirements
- go 1.19
- docker & docker-compose

## Run Project

Use ```make run``` to build and run docker containers with application itself

## Build protobuf

### Install:
- go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
- go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
- go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
- go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
- https://docs.buf.build/installation

Use ```make buf``` to build protobuf files

[BloomRPC](https://github.com/bloomrpc/bloomrpc) - client for send gRPC requests
