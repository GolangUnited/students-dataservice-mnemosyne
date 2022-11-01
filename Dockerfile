FROM golang:1.19.2-alpine3.16 AS builder

RUN go version

COPY . /mnemosyne/
WORKDIR /mnemosyne/

RUN go mod download
RUN GOOS=linux go build -o ./.bin/mnemosyne ./cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /mnemosyne/.bin/mnemosyne .
COPY --from=builder /mnemosyne/swagger swagger/
COPY --from=builder /mnemosyne/configs/config.yml configs/config.yml
RUN touch .env

CMD /app/mnemosyne