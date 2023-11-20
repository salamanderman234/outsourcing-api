package configs

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDSN() string {
	host := viper.GetString("DB_HOST")
	port := viper.GetString("DB_PORT")
	user := viper.GetString("DB_USER")
	password := viper.GetString("DB_PASS")
	name := viper.GetString("DB_NAME")

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user,
		password,
		host,
		port,
		name,
	)
}

func ConnectDatabase() (*gorm.DB, error) {
	dsn := GetDSN()
	return gorm.Open(mysql.Open(dsn), &gorm.Config{
		TranslateError: true,
	})
}
