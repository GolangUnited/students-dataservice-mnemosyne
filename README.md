## See swagger for API endpoints
http://localhost:8000/swagger/index.html

## Requirements
- go 1.19
- docker & docker-compose

## Run Project

Use ```make run``` to build and run docker containers with application itself

## Build protobuf

`go install github.com/bufbuild/buf@latest`

Use ```make buf``` to build protobuf files
