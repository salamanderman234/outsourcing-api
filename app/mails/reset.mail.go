package mails

import (
	"github.com/salamanderman234/outsourcing-api/app/domains"
	"github.com/salamanderman234/outsourcing-api/app/providers"
)

type resetPasswordMail struct {
	baseMail
}

func NewResetPasswordMail(data map[string]any) domains.MailInterface {
	return &resetPasswordMail{
		baseMail: baseMail{
			Template: "reset-password.html",
			Data:     data,
			Subject:  "Reset Your Password",
		},
	}
}

func init() {
	providers.MailProvider.RegisterMail(NewResetPasswordMail)
}
