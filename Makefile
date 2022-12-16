.PHONY:
run:
	docker-compose up --remove-orphans --build app

test:
	go test ./... -coverprofile cover.out

test-coverage:
	go tool cover -func cover.out | grep total | awk '{print $3}'

buf:
	cd ./proto && buf generate

buf-lint:
	buf lint

migrate-create:
	migrate create -dir ./migrations -ext sql -seq migrate

migrate-up:
	migrate -path ./migrations -database "postgresql://postgres:postgres@localhost:5436/postgres?sslmode=disable" -verbose up

migrate-down:
	migrate -path ./migrations -database "postgresql://postgres:postgres@localhost:5436/postgres?sslmode=disable" -verbose down
