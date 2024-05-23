package main

import (
	"github.com/salamanderman234/outsourcing-api/app/models"
	"github.com/salamanderman234/outsourcing-api/configs"
)

func init() {
	configs.AppConfig.SetConfig("./.env")
}

func main() {
	models := []any{
		// master
		models.Province{},
		models.Regency{},
		models.Category{},
		// user
		models.User{},
		models.Admin{},
		models.Supervisor{},
		models.Employee{},
		models.ServiceUser{},
		models.Job{},
		// service
		models.Service{},
		models.RequiredItemService{},
		models.AdditionalItemService{},
		models.Package{},
		models.PackageService{},
		models.PackageServiceAdditionalItem{},
		// transaction
		models.Transaction{},
		models.TransactionDetail{},
		models.TransactionDetailEtc{},
		// payments
		models.Payment{},
	}
	connection, err := configs.DatabaseConfig.ConnectDatabase()
	if err != nil {
		panic(err)
	}
	connection.AutoMigrate(models...)
}
