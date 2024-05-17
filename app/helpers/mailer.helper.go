package helpers

import (
	"crypto/tls"
	"fmt"

	"github.com/salamanderman234/outsourcing-api/configs"
	"gopkg.in/gomail.v2"
)

var dialer = gomail.NewDialer(
	configs.MailerConfig.Host,
	configs.MailerConfig.Port,
	configs.MailerConfig.Email,
	configs.MailerConfig.Password,
)

type mailerHelper struct{}

func (mailerHelper) SendEmail(to []string, subject string, msg string) error {
	Logger.Info(
		fmt.Sprintf("(Mail) Sending a new mail (to: %s, subject: %s)", to, subject),
	)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	mail := gomail.NewMessage()
	mail.SetHeader("From", "TEST")
	mail.SetHeader("To", to...)
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/html", msg)
	return dialer.DialAndSend(mail)
}

var Mailer = mailerHelper{}
