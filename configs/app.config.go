package configs

import "github.com/spf13/viper"

const (
	PAGINATION_PER_PAGE         = 10
	PAGINATION_ORDER_BY_DEFAULT = "updated_at"
	PAGINATION_ORDER            = "DESC"
	// token
	ACCESS_TOKEN_EXPIRE_TIME  = 0.3
	REFRESH_TOKEN_EXPIRE_TIME = 72
)

func GetApplicationSecret() string {
	return viper.GetString("APP_SECRET")
}

func SetConfig(url string) {
	viper.SetConfigFile(url)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
