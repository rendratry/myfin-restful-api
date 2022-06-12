package helper

import (
	"gopkg.in/gomail.v2"
	"log"
)

const CONFIG_SMTP_HOST = "smtp.titan.email"
const CONFIG_SMTP_PORT = 465
const CONFIG_SENDER_NAME = "MyFin <help@hellomyfin.engineer>"
const CONFIG_AUTH_EMAIL = "help@hellomyfin.engineer"
const CONFIG_AUTH_PASSWORD = "adminmyfin"

func SendOTP() {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", "rendraty1@gmail.com")
	mailer.SetHeader("Subject", "Test mail")
	mailer.SetBody("text/html", "Hello, <b>have a nice day</b>")

	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Mail sent!")
}
