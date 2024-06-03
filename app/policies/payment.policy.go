package policies

import (
	"github.com/salamanderman234/outsourcing-api/app/types"
	"github.com/salamanderman234/outsourcing-api/app/types/enums"
)

type PaymentPolicy struct {
}

func (PaymentPolicy) Create(claims types.JWTCLaims) bool {
	return claims.Role == string(enums.AdminUserRole) || true
}

func (PaymentPolicy) ReadAll(claims types.JWTCLaims) bool {
	return claims.Role == string(enums.AdminUserRole)
}

func (PaymentPolicy) Find(data any, claims types.JWTCLaims) bool {
	return true
}

func (PaymentPolicy) Update(data any, claims types.JWTCLaims) bool {
	return claims.Role == string(enums.AdminUserRole) || true
}

func (PaymentPolicy) Delete(data any, claims types.JWTCLaims) bool {
	return claims.Role == string(enums.AdminUserRole) || true
}

func (PaymentPolicy) UploadFile(data any, claims types.JWTCLaims) bool {
	return true
}
func (PaymentPolicy) ViewFile(data any, claims types.JWTCLaims) bool {
	return true
}
