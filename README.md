## See swagger for API endpoints
http://localhost:8000/swagger/

## Requirements
- go 1.19
- docker & docker-compose

## Run Project

Use ```make run``` to build and run docker containers with application itself

## Build protobuf

### Install:
- go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
- go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
- go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
- go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
- go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-swagger@latest
- https://docs.buf.build/installation

Use ```make buf``` to build protobuf files

[BloomRPC](https://github.com/bloomrpc/bloomrpc) - client for send gRPC requests

## Endpoints

https://docs.google.com/spreadsheets/d/1--GgKECVO1CZut9QbEl9GInljdKJ68VsQJj74Hfo8lE/edit#gid=82299691