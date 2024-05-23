package providers

import (
	"reflect"

	"github.com/salamanderman234/outsourcing-api/app/domains"
	"github.com/salamanderman234/outsourcing-api/configs"
)

type createMailInstanceFunc func(data map[string]any) domains.MailInterface

type mailProvider struct {
	list     map[string]createMailInstanceFunc
	Host     string
	Port     int
	Email    string
	Password string
}

func (m *mailProvider) SetMailClient() {
	m.Host = configs.MailerConfig.Host
	m.Port = configs.MailerConfig.Port
	m.Email = configs.MailerConfig.Email
	m.Password = configs.MailerConfig.Password
}

func (m *mailProvider) RegisterMail(fun createMailInstanceFunc) {
	mailName := reflect.TypeOf(fun(map[string]any{})).Name()
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
