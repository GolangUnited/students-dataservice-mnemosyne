FROM golang:1.19.1-alpine3.15 AS builder

RUN go version

COPY . /github.com/NEKETSKY/mnemosyne/
WORKDIR /github.com/NEKETSKY/mnemosyne/

RUN go mod download
RUN GOOS=linux go build -o ./.bin/app ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /github.com/NEKETSKY/mnemosyne/.bin/app .
COPY --from=0 /github.com/NEKETSKY/mnemosyne/configs/config.yml configs/config.yml

EXPOSE 8000

CMD ["./app"]