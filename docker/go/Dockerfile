FROM golang:1.19.1-alpine3.16

WORKDIR /app
COPY app ./app
COPY go.mod go.sum ./

RUN go mod download

CMD ["go", "run", "app/cmd/main.go"]