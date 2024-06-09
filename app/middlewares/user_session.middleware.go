package custom_middlewares

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/outsourcing-api/app/helpers"
	"github.com/salamanderman234/outsourcing-api/app/types"
	"github.com/salamanderman234/outsourcing-api/app/types/enums"
	"github.com/salamanderman234/outsourcing-api/configs"
)

func RetrieveUserSession(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie(configs.VarConfig.AuthCookieName)
		if err != nil {
			return next(c)
		}
		token := cookie.Value
		payload, err := helpers.JWT.VerifyToken(token)
		if err != nil {
			status, resp := helpers.Response.CreateResponse(types.ResponseParams{
				Error: err,
			})
			return c.JSON(status, resp)
		}
		email := payload.Email
		role := payload.Role
		tokenType := payload.Subject
		if tokenType != string(enums.AuthenticationTokenType) {
			return next(c)
		}
		if email != "" {
			uri := c.Request().RequestURI
			helpers.Logger.Info(
				fmt.Sprintf("(Session) User (%s) %s is attempting to access %s", role, email, uri),
			)
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
