package configs

import "github.com/spf13/viper"

type mailerConfigs struct {
	Email    string
	Password string
	Name     string
	Host     string
	Port     int
}

var MailerConfig mailerConfigs

func (m *mailerConfigs) setMailerConfig() {
	m.Email = viper.GetString("SMTP_EMAIL")
	m.Password = viper.GetString("SMTP_PASSWORD")
	m.Host = viper.GetString("SMTP_HOST")
	m.Port = viper.GetInt("SMTP_PORT")
	m.Name = viper.GetString("APP_NAME")
}
