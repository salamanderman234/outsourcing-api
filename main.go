package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/salamanderman234/outsourcing-api/app/helpers"
	custom_middlewares "github.com/salamanderman234/outsourcing-api/app/middlewares"
	"github.com/salamanderman234/outsourcing-api/app/providers"
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
	// server.Use(middleware.BodyLimit(configs.RouterConfig.MaxBodyLength))
	// server.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
	// 	Skipper:      middleware.DefaultSkipper,
	// 	ErrorMessage: custom_errors.ErrTimeOut.Msg,
	// 	Timeout:      time.Duration(configs.RouterConfig.MaxTimeOut) * time.Second,
	// }))

	// set up database connection
	connection, err := configs.DatabaseConfig.ConnectDatabase()
	helpers.Logger.Info("(Database) Attempt to connect database...")
	providers.SetConnection(connection)
	if err != nil {
		helpers.Logger.Fatal(fmt.Sprintf("(Database) Failed to connect database, msg: %s", err.Error()))
		panic(err)
	}
	helpers.Logger.Info("(Database) Successfully connect into the database !")

	// set up app
	repositories.RegisterAllRepos()
	services.RegisterAllServices()
	views.RegisterAllViews()
	routes.RegisterAllRoutes(server)

	// scheduler
	helpers.Cron.SetupCron()
	helpers.Cron.AddDurationJob(300, helpers.JobManager.ExecuteQueue)
	helpers.Cron.StartCron()

	// start
	helpers.Logger.Info("(Server) Starting the server...")
	server.Logger.Fatal(server.Start(":8080"))
	helpers.Logger.Fatal("(Server) Server shutdown")
}
