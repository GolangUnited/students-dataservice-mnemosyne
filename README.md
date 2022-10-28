## Run Project

Use ```make run``` to build and run docker containers with application itself

Use ```make migrate-create``` to create new migrate or run command ```migrate create -dir ./migrations -ext sql -seq <migrate name>```


### Example environment file ```.env```:

```
POSTGRES_HOST=localhost
POSTGRES_PORT=5436
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB_NAME=postgres
POSTGRES_SSL=disable
```

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

## See swagger for API endpoints
http://localhost:8000/swagger/

## Requirements
- go 1.19
- docker & docker-compose