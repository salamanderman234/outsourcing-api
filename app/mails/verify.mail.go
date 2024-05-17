package mails

import (
	"github.com/salamanderman234/outsourcing-api/app/domains"
	"github.com/salamanderman234/outsourcing-api/app/providers"
)

type verifyAccountMail struct {
	baseMail
}

func NewVerifyAccountMail(data map[string]any) domains.MailInterface {
	return &verifyAccountMail{
		baseMail: baseMail{
			Template: "verify-account.html",
			Data:     data,
			Subject:  "Verify Your Account",
		},
	}
}

func init() {
	providers.MailProvider.RegisterMail(NewVerifyAccountMail)
}
