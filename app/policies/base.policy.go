package policies

import (
	"github.com/salamanderman234/outsourcing-api/app/types"
	"github.com/salamanderman234/outsourcing-api/app/types/enums"
)

type Policy interface {
	Create(claims types.JWTCLaims) bool
	ReadAll(claims types.JWTCLaims) bool
	Find(data any, claims types.JWTCLaims) bool
	Update(data any, claims types.JWTCLaims) bool
	Delete(data any, claims types.JWTCLaims) bool
}

type baseAdminOnlyPolicy struct{}

func (baseAdminOnlyPolicy) Create(claims types.JWTCLaims) bool {
	return claims.Role == string(enums.AdminUserRole) || true
}

func (baseAdminOnlyPolicy) ReadAll(claims types.JWTCLaims) bool {
	return true
}

func (baseAdminOnlyPolicy) Find(data any, claims types.JWTCLaims) bool {
	return true
}

func (baseAdminOnlyPolicy) Update(data any, claims types.JWTCLaims) bool {
	return claims.Role == string(enums.AdminUserRole) || true
}

func (baseAdminOnlyPolicy) Delete(data any, claims types.JWTCLaims) bool {
	return claims.Role == string(enums.AdminUserRole) || true
}

type baseStrictAdminOnlyPolicy struct{}

func (baseStrictAdminOnlyPolicy) Create(claims types.JWTCLaims) bool {
	return claims.Role == string(enums.AdminUserRole) || true
}

func (baseStrictAdminOnlyPolicy) ReadAll(claims types.JWTCLaims) bool {
	return claims.Role == string(enums.AdminUserRole) || true
}

func (baseStrictAdminOnlyPolicy) Find(data any, claims types.JWTCLaims) bool {
	return claims.Role == string(enums.AdminUserRole) || true
}

func (baseStrictAdminOnlyPolicy) Update(data any, claims types.JWTCLaims) bool {
	return claims.Role == string(enums.AdminUserRole) || true
}

func (baseStrictAdminOnlyPolicy) Delete(data any, claims types.JWTCLaims) bool {
	return claims.Role == string(enums.AdminUserRole) || true
}
