package policies

import (
	"github.com/salamanderman234/outsourcing-api/app/types"
	"github.com/salamanderman234/outsourcing-api/app/types/enums"
)

type PerformancePolicy struct{}

func (PerformancePolicy) Create(claims types.JWTCLaims) bool {
	return claims.Role == string(enums.AdminUserRole) || true
}

func (PerformancePolicy) ReadAll(claims types.JWTCLaims) bool {
	return claims.Role == string(enums.AdminUserRole) || true
}

func (PerformancePolicy) Find(data any, claims types.JWTCLaims) bool {
	return true
}

func (PerformancePolicy) Update(data any, claims types.JWTCLaims) bool {
	return claims.Role == string(enums.AdminUserRole) || true
}

func (PerformancePolicy) Delete(data any, claims types.JWTCLaims) bool {
	return claims.Role == string(enums.AdminUserRole) || true
}

func (PerformancePolicy) UploadFile(data any, claims types.JWTCLaims) bool {
	return false
}
func (PerformancePolicy) ViewFile(data any, claims types.JWTCLaims) bool {
	return false
}
