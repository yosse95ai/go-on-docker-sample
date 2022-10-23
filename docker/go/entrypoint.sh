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

# docker-composee up
# が実行されたとき
if [ $_ENV = "dev" ];then
  >&1 echo "> go run app/cmd/main.go"
  air -c .air.toml
# docker-compose -f docker-compose.yml -f production.yml up
# が実行されたとき
elif [ $_ENV = "prod" ];then
  >&1 echo "> go build app/cmd/main.go"
  go run app/cmd/main.go
  ./main
else
  >&2 echo "Missing _ENV: 環境変数 '_ENV' が設定されていません. - [Failed]"
fi
