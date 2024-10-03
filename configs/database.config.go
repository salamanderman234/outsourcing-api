package configs

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type databaseConfig struct {
	Host     string
	Port     string
	UserName string
	Password string
	Name     string
}

var (
	DatabaseConfig databaseConfig
)

func (db databaseConfig) getDSN() string {
	host := db.Host
	port := db.Port
	user := db.UserName
	password := db.Password
	name := db.Name
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user,
		password,
		host,
		port,
		name,
	)
}

func (db databaseConfig) ConnectDatabase() (*gorm.DB, error) {
	dsn := db.getDSN()
	return gorm.Open(mysql.Open(dsn), &gorm.Config{
		TranslateError: true,
		// Logger:                                   logger.Default.LogMode(logger.Info),
		DisableForeignKeyConstraintWhenMigrating: true,
	})
}

func (d *databaseConfig) setDatabaseConfig() {
	d.Host = viper.GetString("DB_HOST")
	d.Port = viper.GetString("DB_PORT")
	d.UserName = viper.GetString("DB_USER")
	d.Password = viper.GetString("DB_PASS")
	d.Name = viper.GetString("DB_NAME")
}
