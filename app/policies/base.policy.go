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
	UploadFile(data any, claims types.JWTCLaims) bool
	ViewFile(data any, claims types.JWTCLaims) bool
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

func (baseAdminOnlyPolicy) UploadFile(data any, claims types.JWTCLaims) bool {
	return true
}
func (baseAdminOnlyPolicy) ViewFile(data any, claims types.JWTCLaims) bool {
	return true
}
