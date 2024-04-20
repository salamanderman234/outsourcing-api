package mails

import (
	"bytes"
	"html/template"
	"path"
	"reflect"
)

type createMailInstanceFunc func(data map[string]any) MailInterface

var mailRegistry map[string]createMailInstanceFunc

type MailInterface interface {
	GetTemplate() (string, error)
	GetData() map[string]any
	GetSubject() string
}

type baseMail struct {
	Template string
	Data     map[string]any
	Subject  string
}

func (m baseMail) GetSubject() string {
	return m.Subject
}

func (m baseMail) GetData() map[string]any {
	return m.Data
}

func (m baseMail) GetTemplate() (string, error) {
	var filepath = path.Join("storage", "mail", m.Template)
	tmpl, err := template.ParseFiles(filepath)
	if err != nil {
		return "", err
	}
	var tpl bytes.Buffer
	if err := tmpl.Execute(&tpl, m.Data); err != nil {
		return "", err
	}
	return tpl.String(), nil
}

func registerMail(fun createMailInstanceFunc) {
	mailName := reflect.TypeOf(fun(map[string]any{})).Name()
	mailRegistry[mailName] = fun
}

func NewMailInstance(name string, data map[string]any) MailInterface {
	fun, ok := mailRegistry[name]
	if !ok {
		return nil
	}
	return fun(data)
}
