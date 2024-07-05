package main

import (
	"fmt"

	"github.com/salamanderman234/outsourcing-api/app/helpers"
	"github.com/salamanderman234/outsourcing-api/app/models"
	"github.com/salamanderman234/outsourcing-api/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	configs.AppConfig.SetConfig("./.env")
}

func CreateDatabase() error {
	host := configs.DatabaseConfig.Host
	port := configs.DatabaseConfig.Port
	user := configs.DatabaseConfig.UserName
	password := configs.DatabaseConfig.Password
	name := configs.DatabaseConfig.Name
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local",
		user,
		password,
		host,
		port,
	)
	helpers.Logger.Info(dsn)
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		TranslateError:                           true,
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		conn.Exec(fmt.Sprintf("CREATE DATABASE %s", name))
		return nil
	}
	return err
}

func main() {
	models := []any{
		// master
		models.Province{},
		models.Regency{},
		models.Category{},
		models.Question{},
		models.CategoryQuestion{},
		models.PaymentConfig{},
		// user
		models.User{},
		models.Admin{},
		models.Supervisor{},
		models.Employee{},
		models.ServiceUser{},
		models.Job{},
		models.SuperAdmin{},
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
		// placement
		models.Placement{},
		models.PlacementDetail{},
		models.PlacementDetailEmployee{},
		// feedback
		models.Feedback{},
		models.Complaint{},
		models.ComplaintReply{},
		// performance
		models.PerformanceForm{},
		models.PerformanceFormFeedback{},
		models.Performance{},
	}
	connection, err := configs.DatabaseConfig.ConnectDatabase()
	if err != nil {
		err = CreateDatabase()
		if err != nil {
			panic(err)
		}
	}
	connection.AutoMigrate(models...)
}
