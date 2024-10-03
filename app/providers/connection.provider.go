package providers

import "gorm.io/gorm"

// database connection
var connection *gorm.DB

func SetConnection(conn *gorm.DB) {
	connection = conn
}

func GetConnection() *gorm.DB {
	return connection
}
