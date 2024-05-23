package configs

import (
	"github.com/spf13/viper"
)

type appConfig struct {
	Name              string
	Url               string
	IsDebug           bool
	PaginationPerPage uint
	Env               string
}

var AppConfig appConfig

func setConfigFile(url string) {
	viper.SetConfigFile(url)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func (a *appConfig) SetConfig(path string) {
	setConfigFile(path)
	// app config
	a.Name = viper.GetString("APP_NAME")
	a.Url = viper.GetString("APP_URL")
	a.IsDebug = viper.GetBool("APP_DEBUG")
	a.PaginationPerPage = 10
	a.Env = viper.GetString("APP_ENV")
	// database config
	DatabaseConfig.setDatabaseConfig()
	// jwt config
	JWTConfig.setJWTConfig()
	// mailer config
	MailerConfig.setMailerConfig()
	// router config
	RouterConfig.setRouterConfig()
	// set var config
	VarConfig.setVarConfig()
	// set resource config
	ResourceConfig.setResourceConfig()
	// midtrans
	MidtransConfig.setMidtransConfig()

}
