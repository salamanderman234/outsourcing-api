package main

import (
	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/outsourcing-api/configs"
	"github.com/salamanderman234/outsourcing-api/domains"
	"github.com/salamanderman234/outsourcing-api/repositories"
	"github.com/salamanderman234/outsourcing-api/services"
	"github.com/salamanderman234/outsourcing-api/views"
)

func init() {
	configs.SetConfig("./.env")
}

func main() {
	server := echo.New()
	// database
	connection, err := configs.ConnectDatabase()
	if err != nil {
		panic(err)
	}
	// repositories
	domains.RepoRegistry.UserRepo = repositories.NewUserRepository(connection)
	// services
	domains.AuthServiceRegistry.AuthServ = services.NewUserAuthService()

	authView := views.NewAuthView()
	server.GET("/login", authView.Login)
	server.GET("/verify", authView.Verify)
	server.Logger.Fatal(server.Start(":1323"))
}
