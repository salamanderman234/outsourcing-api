package policies

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/salamanderman234/outsourcing-api/app/types"
	"github.com/salamanderman234/outsourcing-api/app/types/enums"
)

type PerformancePolicy struct{}

func (PerformancePolicy) Create(claims jwt.Claims) bool {
	c, ok := claims.(types.JWTCLaims)
	if !ok {
		return false
	}
	return c.Role == string(enums.AdminUserRole) || true
}

func (PerformancePolicy) ReadAll(claims jwt.Claims) bool {
	c, ok := claims.(types.JWTCLaims)
	if !ok {
		return false
	}
	return c.Role == string(enums.AdminUserRole) || true
}

func (PerformancePolicy) Find(data any, claims jwt.Claims) bool {
	return true
}

func (PerformancePolicy) Update(data any, claims jwt.Claims) bool {
	c, ok := claims.(types.JWTCLaims)
	if !ok {
		return false
	}
	return c.Role == string(enums.AdminUserRole) || true
}

func (PerformancePolicy) Delete(data any, claims jwt.Claims) bool {
	c, ok := claims.(types.JWTCLaims)
	if !ok {
		return false
	}
	return c.Role == string(enums.AdminUserRole) || true
}

func (PerformancePolicy) UploadFile(data any, claims jwt.Claims) bool {
	return false
}
func (PerformancePolicy) ViewFile(data any, claims jwt.Claims) bool {
	return false
}
