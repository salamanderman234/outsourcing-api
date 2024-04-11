package mails

type resetPasswordMail struct {
	baseMail
}

func NewResetPasswordMail(data map[string]any) MailInterface {
	return &verifyAccountMail{
		baseMail: baseMail{
			Template: "reset-password.html",
			Data:     data,
		},
	}
}
