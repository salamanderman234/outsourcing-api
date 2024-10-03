package policies

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/salamanderman234/outsourcing-api/app/types"
	"github.com/salamanderman234/outsourcing-api/app/types/enums"
)

type baseAdminOnlyPolicy struct{}

func (baseAdminOnlyPolicy) Create(claims jwt.Claims) bool {
	c, ok := claims.(types.JWTCLaims)
	if !ok {
		return false
	}
	return c.Role == string(enums.AdminUserRole) || c.Role == string(enums.SuperAdminRole)
}

func (baseAdminOnlyPolicy) ReadAll(claims jwt.Claims) bool {
	return true
}

func (baseAdminOnlyPolicy) Find(data any, claims jwt.Claims) bool {
	return true
}

func (baseAdminOnlyPolicy) Update(data any, claims jwt.Claims) bool {
	c, ok := claims.(types.JWTCLaims)
	if !ok {
		return false
	}
	return c.Role == string(enums.AdminUserRole) || c.Role == string(enums.SuperAdminRole)
}

func (baseAdminOnlyPolicy) Delete(data any, claims jwt.Claims) bool {
	c, ok := claims.(types.JWTCLaims)
	if !ok {
		return false
	}
	return c.Role == string(enums.AdminUserRole) || c.Role == string(enums.SuperAdminRole)
}

func (baseAdminOnlyPolicy) UploadFile(data any, claims jwt.Claims) bool {
	return true
}
func (baseAdminOnlyPolicy) ViewFile(data any, claims jwt.Claims) bool {
	return true
}
