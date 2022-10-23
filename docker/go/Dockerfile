# 開発環境 #########################################
FROM golang:1.19.1-alpine3.16 as dev

WORKDIR /app
COPY app ./app
COPY go.mod go.sum ./
COPY docker/go/entrypoint.sh ./docker/go/
RUN chmod +x ./docker/go/entrypoint.sh

RUN go install github.com/cosmtrek/air@latest
RUN go mod download

CMD [ "./docker/go/entrypoint.sh" ]
