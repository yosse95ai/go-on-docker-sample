FROM golang:1.19.1-alpine3.16

WORKDIR /app
COPY app ./app
COPY go.mod go.sum ./
COPY docker/go/entrypoint.sh ./docker/go/
RUN chmod +x ./docker/go/entrypoint.sh

RUN go mod download
