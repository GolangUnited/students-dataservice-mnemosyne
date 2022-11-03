.PHONY:
run:
	docker-compose up --remove-orphans --build app

test:
	go test ./... -coverprofile cover.out

test-coverage:
	go tool cover -func cover.out | grep total | awk '{print $3}'

buf:
	cd ./proto && buf generate

migrate-create:
	migrate create -dir ./migrations -ext sql -seq migrate