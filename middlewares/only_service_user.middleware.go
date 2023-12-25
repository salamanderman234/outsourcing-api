package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/outsourcing-api/domains"
	"github.com/salamanderman234/outsourcing-api/helpers"
)

func OnlyServiceUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		resp := domains.BasicResponse{}
		user := c.Get("user")
		if user == nil {
			status, msg, errBody := helpers.HandleError(domains.ErrInvalidAccess)
			resp.Message = msg
			resp.Body = *errBody
			return c.JSON(status, resp)
		}
		userEnt, ok := user.(domains.UserEntity)
		if !ok {
			status, msg, errBody := helpers.HandleError(domains.ErrInvalidAccess)
			resp.Message = msg
			resp.Body = *errBody
			return c.JSON(status, resp)
		}
		if userEnt.Role != string(domains.ServiceUserRole) {
			status, msg, errBody := helpers.HandleError(domains.ErrInvalidAccess)
			resp.Message = msg
			resp.Body = *errBody
			return c.JSON(status, resp)
		}
		return next(c)
	}
}
