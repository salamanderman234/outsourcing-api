package mails

type verifyAccountMail struct {
	baseMail
}

func NewVerifyAccountMail(data map[string]any) MailInterface {
	return &verifyAccountMail{
		baseMail: baseMail{
			Template: "verify-account.html",
			Data:     data,
			Subject:  "Verify Your Account",
		},
	}
}

func init() {
	registerMail(NewVerifyAccountMail)
}
