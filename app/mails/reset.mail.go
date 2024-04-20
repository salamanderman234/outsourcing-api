package mails

type resetPasswordMail struct {
	baseMail
}

func NewResetPasswordMail(data map[string]any) MailInterface {
	return &resetPasswordMail{
		baseMail: baseMail{
			Template: "reset-password.html",
			Data:     data,
			Subject:  "Reset Your Password",
		},
	}
}

func init() {
	registerMail(NewResetPasswordMail)
}
