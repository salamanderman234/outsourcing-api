package providers

import (
	"reflect"

	"github.com/salamanderman234/outsourcing-api/app/domains"
)

type createMailInstanceFunc func(data map[string]any) domains.MailInterface

type mailProvider struct {
	list map[string]createMailInstanceFunc
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
