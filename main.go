package main

import (
	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/outsourcing-api/views"
)

func main() {
	server := echo.New()

	authView := views.NewAuthView()
	server.GET("/test", authView.Login)
	server.Logger.Fatal(server.Start(":1323"))
}
