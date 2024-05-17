package jobs

import (
	"reflect"

	"github.com/salamanderman234/outsourcing-api/app/domains"
	"github.com/salamanderman234/outsourcing-api/app/helpers"
	"github.com/salamanderman234/outsourcing-api/app/providers"
)

type sendMailJob struct {
	MailName string
	Data     map[string]any
	To       []string
}

func NewSendMailJob(mail domains.MailInterface, to []string) domains.JobInterface {
	mailName := reflect.TypeOf(mail).Name()
	data := mail.GetData()

	return &sendMailJob{
		MailName: mailName,
		Data:     data,
		To:       to,
	}
}

func (s sendMailJob) Handle(err chan<- error) {
	mail := providers.MailProvider.NewMailInstance(s.MailName, s.Data)
	tmpl, errs := mail.GetTemplate()
	if errs != nil {
		err <- errs
	}
	errs = helpers.Mailer.SendEmail(s.To, mail.GetSubject(), tmpl)
	if errs != nil {
		err <- errs
	}
}

func (s sendMailJob) GetData() map[string]any {
	return map[string]any{
		"to":        s.To,
		"mail_name": s.MailName,
		"data":      s.Data,
	}
}
func (s *sendMailJob) SetData(data map[string]any) {
	if to, ok := data["to"]; ok {
		s.To, _ = to.([]string)
	}
	if mail, ok := data["mail"]; ok {
		s.MailName, _ = mail.(string)
	}
	if data, ok := data["data"]; ok {
		s.Data, _ = data.(map[string]any)
	}
}

func init() {
	providers.JobProvider.RegisterJob(&sendMailJob{})
}
