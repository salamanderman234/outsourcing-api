package main

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/salamanderman234/outsourcing-api/app/helpers"
	"github.com/salamanderman234/outsourcing-api/app/jobs"
	custom_middlewares "github.com/salamanderman234/outsourcing-api/app/middlewares"
	"github.com/salamanderman234/outsourcing-api/app/providers"
	"github.com/salamanderman234/outsourcing-api/app/repositories"
	"github.com/salamanderman234/outsourcing-api/app/services"
	"github.com/salamanderman234/outsourcing-api/app/views"
	"github.com/salamanderman234/outsourcing-api/configs"
	"github.com/salamanderman234/outsourcing-api/routes"
	templates "github.com/salamanderman234/outsourcing-api/views"
	"gorm.io/gorm"
)

func init() {
	configs.AppConfig.SetConfig("./.env")
	providers.SetMidtransClient()
	providers.MailProvider.SetMailClient()
}

func main() {
	server := echo.New()
	server.Static("/public", "./views/public")
	server.Renderer = &templates.DefaultTemplate
	// set up middleware
	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowOrigins:     []string{"http://localhost:3000", "http://web.salamanderman.my.id"},
	}))
	server.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:    true,
		LogURI:       true,
		LogLatency:   true,
		LogUserAgent: true,
		LogMethod:    true,
		LogRemoteIP:  true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			str := fmt.Sprintf("(Request) URI: %s, METHOD: %s, IP: %s, STATUS: %d (%.3fs)",
				v.URI, v.Method, v.RemoteIP, v.Status, v.Latency.Seconds(),
			)
			helpers.Logger.Info(str)
			return nil
		},
	}))
	// server.Use(middleware.Logger())
	server.Use(custom_middlewares.RetrieveUserSession)
	server.Use(middleware.BodyLimit(configs.RouterConfig.MaxBodyLength))
	// server.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
	// 	Skipper:      middleware.DefaultSkipper,
	// 	ErrorMessage: types.ErrTimeOut.Msg,
	// 	Timeout:      time.Duration(configs.RouterConfig.MaxTimeOut) * time.Second,
	// }))

	// set up database connection
	connectDB := func() (*gorm.DB, error) {
		return configs.DatabaseConfig.ConnectDatabase()
	}
	helpers.Logger.Info("(Database) Attempt to connect database...")
	connection, err := connectDB()
	for err != nil {
		helpers.Logger.Fatal(fmt.Sprintf("(Database) Failed to connect database, msg: %s", err.Error()))
		helpers.Logger.Info("(Database) Retrying to connect in 3s...")
		time.Sleep(3 * time.Second)
		connection, err = connectDB()
	}
	providers.SetConnection(connection)
	helpers.Logger.Info("(Database) Successfully connect into the database !")

	// set up app
	repositories.RegisterAllRepos()
	services.RegisterAllServices()
	views.RegisterAllViews()
	apiRoute := server.Group(fmt.Sprintf("/api/v%s", configs.AppConfig.ApiVersion))
	routes.RegisterAllApiRoutes(apiRoute)
	routes.RegisterAllRoutes(server)

	// scheduler
	jobs.RunCron()

	// init first user
	services.Init()

	// start
	helpers.Logger.Info("(Server) Starting the server...")
	server.Logger.Fatal(server.Start(":8080"))
	helpers.Logger.Fatal("(Server) Server shutdown")
}
