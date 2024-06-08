package providers

import (
	"reflect"

	"github.com/salamanderman234/outsourcing-api/app/domains"
	"github.com/salamanderman234/outsourcing-api/configs"
	"gopkg.in/gomail.v2"
)

type createMailInstanceFunc func(data map[string]any) domains.MailInterface

type mailProvider struct {
	list   map[string]createMailInstanceFunc
	Dialer *gomail.Dialer
}

func (m *mailProvider) SetMailClient() {
	m.Dialer = gomail.NewDialer(
		configs.MailerConfig.Host,
		configs.MailerConfig.Port,
		configs.MailerConfig.Email,
		configs.MailerConfig.Password,
	)
}

func (m *mailProvider) RegisterMail(fun createMailInstanceFunc) {
	mailName := reflect.TypeOf(fun(map[string]any{})).String()
	m.list[mailName] = fun
}

func (m *mailProvider) NewMailInstance(name string, data map[string]any) domains.MailInterface {
	fun, ok := m.list[name]
	if !ok {
		return nil
	}
	return fun(data)
}

var MailProvider = mailProvider{
	list: make(map[string]createMailInstanceFunc),
}
