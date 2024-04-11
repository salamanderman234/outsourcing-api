package mails

import (
	"bytes"
	"html/template"
	"path"
)

type MailInterface interface {
	GetTemplate() (string, error)
}

type baseMail struct {
	Template string
	Data     map[string]any
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
