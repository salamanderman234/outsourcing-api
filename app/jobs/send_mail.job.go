package jobs

import (
	"reflect"

	"github.com/salamanderman234/outsourcing-api/app/helpers"
	"github.com/salamanderman234/outsourcing-api/app/mails"
)

type sendMailJob struct {
	MailName string
	Data     map[string]any
	To       []string
}

func NewSendMailJob(mail mails.MailInterface, to []string) JobInterface {
	mailName := reflect.TypeOf(mail).Name()
	data := mail.GetData()

	return &sendMailJob{
		MailName: mailName,
		Data:     data,
		To:       to,
	}
}

func (s sendMailJob) Handle() error {
	mail := mails.NewMailInstance(s.MailName, s.Data)
	tmpl, err := mail.GetTemplate()
	if err != nil {
		return err
	}
	go helpers.Mailer.SendEmail(s.To, mail.GetSubject(), tmpl)
	return nil
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
	registerJob(&sendMailJob{})
}
