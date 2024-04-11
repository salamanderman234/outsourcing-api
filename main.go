package main

import (
	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/outsourcing-api/app/domains"
	"github.com/salamanderman234/outsourcing-api/app/repositories"
	"github.com/salamanderman234/outsourcing-api/app/services"
	"github.com/salamanderman234/outsourcing-api/app/views"
	"github.com/salamanderman234/outsourcing-api/configs"
	"github.com/salamanderman234/outsourcing-api/routes"
)

func init() {
	configs.AppConfig.SetConfig("./.env")
}

func main() {
	server := echo.New()
	// set up middleware
	// server.Use(middleware.Logger())
	// server.Use(custom_middlewares.RetrieveUserSession)
	// server.Use(middleware.BodyLimit(configs.RouterConfig.MaxBodyLength))
	// server.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
	// 	Skipper:      middleware.DefaultSkipper,
	// 	ErrorMessage: custom_errors.ErrTimeOut.Msg,
	// 	Timeout:      time.Duration(configs.RouterConfig.MaxTimeOut) * time.Second,
	// }))

	// set up database connection
	connection, err := configs.DatabaseConfig.ConnectDatabase()
	domains.Connection = connection
	if err != nil {
		panic(err)
	}

	// set up app
	repositories.RegisterAllRepos()
	services.RegisterAllServices()
	views.RegisterAllViews()
	routes.RegisterAllRoutes(server)

	// start
	server.Logger.Fatal(server.Start(":8080"))
}
