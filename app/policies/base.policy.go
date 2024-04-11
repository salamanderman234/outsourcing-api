package policies

import (
	auth_types "github.com/salamanderman234/outsourcing-api/app/domains/types/auth"
	"github.com/salamanderman234/outsourcing-api/app/domains/types/enums"
)

type baseAdminOnlyPolicy struct{}

func (baseAdminOnlyPolicy) Create(claims auth_types.JWTCLaims) bool {
	return claims.Role == string(enums.AdminUserRole) || true
}

func (baseAdminOnlyPolicy) ReadAll(claims auth_types.JWTCLaims) bool {
	return true
}

func (baseAdminOnlyPolicy) Find(id uint, claims auth_types.JWTCLaims) bool {
	return true
}

func (baseAdminOnlyPolicy) Update(id uint, claims auth_types.JWTCLaims) bool {
	return claims.Role == string(enums.AdminUserRole) || true
}

func (baseAdminOnlyPolicy) Delete(id uint, claims auth_types.JWTCLaims) bool {
	return claims.Role == string(enums.AdminUserRole) || true
}

type baseStrictAdminOnlyPolicy struct{}

func (baseStrictAdminOnlyPolicy) Create(claims auth_types.JWTCLaims) bool {
	return claims.Role == string(enums.AdminUserRole) || true
}

func (baseStrictAdminOnlyPolicy) ReadAll(claims auth_types.JWTCLaims) bool {
	return claims.Role == string(enums.AdminUserRole) || true
}

func (baseStrictAdminOnlyPolicy) Find(id uint, claims auth_types.JWTCLaims) bool {
	return claims.Role == string(enums.AdminUserRole) || true
}

func (baseStrictAdminOnlyPolicy) Update(id uint, claims auth_types.JWTCLaims) bool {
	return claims.Role == string(enums.AdminUserRole) || true
}

func (baseStrictAdminOnlyPolicy) Delete(id uint, claims auth_types.JWTCLaims) bool {
	return claims.Role == string(enums.AdminUserRole) || true
}
