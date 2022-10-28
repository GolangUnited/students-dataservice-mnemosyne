FROM golang:1.19.2-alpine3.16 AS builder

RUN go version

COPY . /mnemosyne/
WORKDIR /mnemosyne/

RUN go mod download
RUN GOOS=linux go build -o ./.bin/mnemosyne ./cmd/main.go

FROM migrate/migrate AS migrate

FROM alpine:latest

WORKDIR /app

COPY --from=builder /mnemosyne/.bin/mnemosyne .
COPY --from=builder /mnemosyne/swagger swagger/
COPY --from=builder /mnemosyne/configs/config.yml configs/config.yml
RUN touch .env

COPY --from=builder /mnemosyne/migrations/*.sql migrations/
COPY --from=builder /mnemosyne/app.sh .

COPY --from=migrate /migrate .

CMD /app/app.sh