package custom_middlewares

import (
	"context"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/outsourcing-api/app/domains/types/responses"
	"github.com/salamanderman234/outsourcing-api/app/helpers"
	"github.com/salamanderman234/outsourcing-api/configs"
)

func RetrieveUserSession(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		if token == "" {
			return next(c)
		}
		token = strings.ReplaceAll(c.Request().Header.Get("Authorization"), "Bearer ", "")
		if token == "" {
			return next(c)
		}
		payload, err := helpers.JWT.VerifyToken(token)
		if err != nil {
			status, resp := helpers.Response.CreateResponse(responses.ResponseConfig{
				Error: err,
			})
			return c.JSON(status, resp)
		}
		c.Set(string(configs.VarConfig.UserContextName), payload)
		c.SetRequest(c.Request().WithContext(
			context.WithValue(
				c.Request().Context(),
				configs.VarConfig.UserContextName,
				payload),
		),
		)
		return next(c)
	}
}
