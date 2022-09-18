# Go on Docker

`Docker` に `Go`・`MySQL` を載せたサンプルプロジェクト

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

```bash
docker-compose up --build

# もしくは

docker-compose build
docker-compose up
```

### ２回目以降

```bash
# コンテナ起動
docker-compose up

# バックグラウンド起動
docker-compose up -d
```
