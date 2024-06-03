package mails

import (
	"bytes"
	"html/template"
	"path"

	"github.com/salamanderman234/outsourcing-api/configs"
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
	data := m.Data
	data["app_base_url"] = configs.AppConfig.Url
	data["app_name"] = configs.AppConfig.Name
	if err := tmpl.Execute(&tpl, data); err != nil {
		return "", err
	}
	return tpl.String(), nil
}
