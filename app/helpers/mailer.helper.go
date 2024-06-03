package helpers

import (
	"crypto/tls"
	"fmt"

	"github.com/salamanderman234/outsourcing-api/app/providers"
	"gopkg.in/gomail.v2"
)

type mailerHelper struct{}

func (mailerHelper) SendEmail(to []string, subject string, msg string) error {
	Logger.Info(
		fmt.Sprintf("(Mail) Sending a new mail (to: %s, subject: %s)", to, subject),
	)
	dialer := providers.MailProvider.Dialer
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	mail := gomail.NewMessage()
	fmt.Println(to)
	mail.SetHeader("From", "outsourcingapp@gmail.com")
	mail.SetHeader("To", to...)
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/html", msg)
	return dialer.DialAndSend(mail)
}

var Mailer = mailerHelper{}
