package mails

import (
	"bytes"
	"html/template"
	"path"
)

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
	var filepath = path.Join("storage", "mails", m.Template)
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
