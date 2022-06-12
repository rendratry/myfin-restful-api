package main

import (
	"gopkg.in/gomail.v2"
	"log"
	"myfin/helper"
	"testing"
)

func TestSendOTP(t *testing.T) {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", helper.CONFIG_SENDER_NAME)
	mailer.SetHeader("To", "rendratrykusuma@gmail.com")
	mailer.SetHeader("Subject", "Test mail")
	mailer.SetBody("text/html", "Mencoba")

	dialer := gomail.NewDialer(
		helper.CONFIG_SMTP_HOST,
		helper.CONFIG_SMTP_PORT,
		helper.CONFIG_AUTH_EMAIL,
		helper.CONFIG_AUTH_PASSWORD,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Mail sent!")
}
