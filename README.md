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

## 準備

まずはネットワークを作成

```bash
docker network create go_app_network
```

`.env`ファイルを作成し、以下の内容を読み替えて記述してください。

```env
GMAIL_ACCOUNT="GmailのEメールアドレス"
GMAIL_APP_PASSWORD="上のアカウントで作成したアプリパスワード"
```

アプリパスワードの生成は以下を参照。

https://support.google.com/accounts/answer/185833?hl=ja

※ アプリパスワードの作成には 2 段階認証の設定が必要です。

## 起動

<font color="#d88">※ docker-compose コマンドが使える環境が必要</font>

コンテナをビルドして起動。

```bash
docker-compose up --build

# もしくは

docker-compose build
docker-compose up
```

### メールの送信

`controllers/mail/mail.go`の`12行目`の変数`to`の中身を好きなメールアドレスに書き換えてください。

書き換えが終わったら以下のコマンドを実行してください。

```bash
docker exec go_api go run controllers/mail/mail.go
```

Docker 関連の内容に変更がある場合はコンテナをビルドし直してください。
