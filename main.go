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
	// server.Use(middleware.Logger())
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
	domains.RepoRegistry.CategoryRepo = repositories.NewCategoryRepository(connection)
	// services
	domains.ServiceRegistry.AuthServ = services.NewUserAuthService()
	domains.ServiceRegistry.CategoryServ = services.NewCategoryService()
	domains.ServiceRegistry.FileServ = services.NewFileService()
	// views
	domains.ViewRegistry.AuthView = views.NewAuthView()
	domains.ViewRegistry.CategoryView = views.NewCategoryView()
	// register routes
	routes.RegisterAllRoutes(server)
	server.Logger.Fatal(server.Start(":8080"))
}
