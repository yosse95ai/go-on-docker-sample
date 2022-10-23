#!/bin/ash

# データベースの開始を待ち
# データベースが完全に起動してから
# APIを立ち上げる

set -e

until nc -z db 3306; do
  >&2 echo "db container is unavailable - [Sleeping]"
  sleep 3
done
>&1 echo "db container is now fully activated. - [Running]"

if [ $_ENV = "dev" ];then
  >&1 echo "> go run app/cmd/main.go"
  air -c .air.toml
elif [ $_ENV = "prod" ];then
  >&1 echo "> go build app/cmd/main.go"
  go run app/cmd/main.go
  ./main
fi
