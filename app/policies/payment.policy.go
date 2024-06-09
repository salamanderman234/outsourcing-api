package policies

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/salamanderman234/outsourcing-api/app/types"
	"github.com/salamanderman234/outsourcing-api/app/types/enums"
)

type PaymentPolicy struct {
}

func (PaymentPolicy) Create(claims jwt.Claims) bool {
	c, ok := claims.(types.JWTCLaims)
	if !ok {
		return false
	}
	return c.Role == string(enums.AdminUserRole) ||
		c.Role == string(enums.SuperAdminUserRole) ||
		c.Role == string(enums.ServiceUserRole)
}

func (PaymentPolicy) ReadAll(claims jwt.Claims) bool {
	c, ok := claims.(types.JWTCLaims)
	if !ok {
		return false
	}
	return c.Role == string(enums.AdminUserRole) ||
		c.Role == string(enums.SuperAdminUserRole)
}

func (PaymentPolicy) Find(data any, claims jwt.Claims) bool {
	c, ok := claims.(types.JWTCLaims)
	if !ok {
		return false
	}
	return c.Role == string(enums.AdminUserRole) ||
		c.Role == string(enums.SuperAdminUserRole) ||
		c.Role == string(enums.ServiceUserRole)
}

func (PaymentPolicy) Update(data any, claims jwt.Claims) bool {
	c, ok := claims.(types.JWTCLaims)
	if !ok {
		return false
	}
	return c.Role == string(enums.AdminUserRole) ||
		c.Role == string(enums.SuperAdminUserRole) ||
		c.Role == string(enums.ServiceUserRole)
}

func (PaymentPolicy) Delete(data any, claims jwt.Claims) bool {
	c, ok := claims.(types.JWTCLaims)
	if !ok {
		return false
	}

	return c.Role == string(enums.AdminUserRole) ||
		c.Role == string(enums.SuperAdminUserRole) ||
		c.Role == string(enums.ServiceUserRole)
}

func (PaymentPolicy) UploadFile(data any, claims jwt.Claims) bool {
	return false
}
func (PaymentPolicy) ViewFile(data any, claims jwt.Claims) bool {
	return false
}
