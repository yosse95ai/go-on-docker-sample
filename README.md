# Go on Docker

`Docker` に `Go`・`MySQL` を載せたサンプルプロジェクト

`Air`を導入してホットリローディングするようにしている

## クローン

```bash
# HTTPなら
git clone https://github.com/yosse95ai/go-on-docker-sample.git
cd go-on-docker-sample

# SSHなら
git clone git@github.com:yosse95ai/go-on-docker-sample.git
cd go-on-docker-sample
```

## 起動

<font color="#d88">※ docker-compose コマンドが使える環境が必要</font>

### 初回実行時

まずはネットワークを作成

```bash
docker network create go_app_network
```

コンテナをビルドして起動。

```bash
docker-compose up --build

# もしくは

docker-compose build
docker-compose up
```

初期データを挿入

```bash
docker exec go_api go run db/seeds.go
```

### ２回目以降

コンテナを起動してください。

```bash
# コンテナ起動
docker-compose up

# バックグラウンド起動
docker-compose up -d
```

内容に変更がある場合はコンテナをビルドし直してください。
