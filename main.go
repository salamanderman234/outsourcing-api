package main

import (
	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/outsourcing-api/configs"
	"github.com/salamanderman234/outsourcing-api/domains"
	"github.com/salamanderman234/outsourcing-api/repositories"
	"github.com/salamanderman234/outsourcing-api/routes"
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
	domains.RepoRegistry.AdminRepo = repositories.NewAdminRepository(connection)
	domains.RepoRegistry.EmployeeRepo = repositories.NewEmployeeRepository(connection)
	domains.RepoRegistry.ServiceUserRepo = repositories.NewServiceUserRepository(connection)
	domains.RepoRegistry.SupervisorRepo = repositories.NewSupervisorRepository(connection)
	// services
	domains.ServiceRegistry.AuthServ = services.NewUserAuthService()
	// views
	domains.ViewRegistry.AuthView = views.NewAuthView()
	// register routes
	routes.RegisterAllRoutes(server)
	server.Logger.Fatal(server.Start(":8080"))
}
