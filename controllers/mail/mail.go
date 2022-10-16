package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/smtp"
	"os"
)

var to = "メールアドレスを入力"

func sendMailByGmail(subject, html string, to []string) {
	auth := smtp.PlainAuth(
		"",
		os.Getenv("GMAIL_ACCOUNT"),
		os.Getenv("GMAIL_APP_PASSWORD"),
		"smtp.gmail.com",
	)

	msg := fmt.Sprintf("Subject: %s\n%s", subject, html)

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"",
		to,
		[]byte(msg),
	)
	if err != nil {
		fmt.Println(err)
	}
}

func sendHTMLMailByGmail(subject, templatePath string, to []string) {
	var body bytes.Buffer
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Fatal("メールテンプレートが見つからないため、メールを送信できませんでした。")
	}
	t.Execute(&body, struct{ Name string }{Name: "吉村弘明"})
	from := ""
	auth := smtp.PlainAuth(
		"",
		os.Getenv("GMAIL_ACCOUNT"),
		os.Getenv("GMAIL_APP_PASSWORD"),
		"smtp.gmail.com",
	)

	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"

	msg := fmt.Sprintf("Subject: %s\n%s\n\n%s", subject, headers, body.String())

	err = smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		from,
		to,
		[]byte(msg),
	)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	if os.Getenv("GMAIL_ACCOUNT") == "" || os.Getenv("GMAIL_APP_PASSWORD") == "" {
		log.Fatal("Gmailの環境変数が登録されていません")
	}
	sendHTMLMailByGmail(
		"このメールが届いた時点では本登録は完了しておりません",
		"./controllers/mail/template.html",
		[]string{to},
	)
}
