package policies

import (
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/salamanderman234/outsourcing-api/app/models"
	"github.com/salamanderman234/outsourcing-api/app/types"
	"github.com/salamanderman234/outsourcing-api/app/types/enums"
)

type UserPolicy struct {
}

func (UserPolicy) Create(claims jwt.Claims) bool {
	return true
}

func (UserPolicy) ReadAll(claims jwt.Claims) bool {
	c, ok := claims.(types.JWTCLaims)
	if !ok {
		return false
	}
	role := c.Role
	return (role == string(enums.AdminUserRole) || role == string(enums.SuperAdminRole))
}

func (UserPolicy) Find(data any, claims jwt.Claims) bool {
	// c, ok := claims.(types.JWTCLaims)
	// if !ok {
	// 	return false
	// }
	// user, ok := data.(*models.User)
	// if !ok {
	// 	return false
	// }
	// role := c.Role
	// id, _ := strconv.Atoi(c.ID)
	// return (role == string(enums.AdminUserRole) || role == string(enums.SuperAdminRole)) || user.ID == uint(id)
	return true
}

func (UserPolicy) Update(data any, claims jwt.Claims) bool {
	c, ok := claims.(types.JWTCLaims)
	if !ok {
		return false
	}
	user, ok := data.(*models.User)
	if !ok {
		return false
	}
	role := c.Role
	id, _ := strconv.Atoi(c.ID)
	return (role == string(enums.AdminUserRole) || role == string(enums.SuperAdminRole)) || user.ID == uint(id)
}

func (UserPolicy) Delete(data any, claims jwt.Claims) bool {
	c, ok := claims.(types.JWTCLaims)
	if !ok {
		return false
	}
	role := c.Role
	return (role == string(enums.AdminUserRole) || role == string(enums.SuperAdminRole))
}

func (UserPolicy) UploadFile(data any, claims jwt.Claims) bool {
	c, ok := claims.(types.JWTCLaims)
	if !ok {
		return false
	}
	user, ok := data.(*models.User)
	if !ok {
		return false
	}
	role := c.Role
	id, _ := strconv.Atoi(c.ID)
	return (role == string(enums.AdminUserRole) || role == string(enums.SuperAdminRole)) || user.ID == uint(id)
}
func (UserPolicy) ViewFile(data any, claims jwt.Claims) bool {
	c, ok := claims.(types.JWTCLaims)
	if !ok {
		return false
	}
	user, ok := data.(*models.User)
	if !ok {
		return false
	}
	role := c.Role
	id, _ := strconv.Atoi(c.ID)
	return (role == string(enums.AdminUserRole) || role == string(enums.SuperAdminRole)) || user.ID == uint(id) || true
}
