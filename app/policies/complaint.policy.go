package policies

import (
	"github.com/salamanderman234/outsourcing-api/app/types"
	"github.com/salamanderman234/outsourcing-api/app/types/enums"
)

type ComplaintPolicy struct {
}

func (ComplaintPolicy) Create(claims types.JWTCLaims) bool {
	return claims.Role == string(enums.AdminUserRole) || true
}

func (ComplaintPolicy) ReadAll(claims types.JWTCLaims) bool {
	return claims.Role == string(enums.AdminUserRole) || true
}

func (ComplaintPolicy) Find(data any, claims types.JWTCLaims) bool {
	return true
}

func (ComplaintPolicy) Update(data any, claims types.JWTCLaims) bool {
	return claims.Role == string(enums.AdminUserRole) || true
}

func (ComplaintPolicy) Delete(data any, claims types.JWTCLaims) bool {
	return claims.Role == string(enums.AdminUserRole) || true
}

func (ComplaintPolicy) UploadFile(data any, claims types.JWTCLaims) bool {
	return false
}
func (ComplaintPolicy) ViewFile(data any, claims types.JWTCLaims) bool {
	return false
}

type ComplaintReplyPolicy struct {
}

func (ComplaintReplyPolicy) Create(claims types.JWTCLaims) bool {
	return claims.Role == string(enums.AdminUserRole) || true
}

func (ComplaintReplyPolicy) ReadAll(claims types.JWTCLaims) bool {
	return claims.Role == string(enums.AdminUserRole) || true
}

func (ComplaintReplyPolicy) Find(data any, claims types.JWTCLaims) bool {
	return true
}

func (ComplaintReplyPolicy) Update(data any, claims types.JWTCLaims) bool {
	return claims.Role == string(enums.AdminUserRole) || true
}

func (ComplaintReplyPolicy) Delete(data any, claims types.JWTCLaims) bool {
	return claims.Role == string(enums.AdminUserRole) || true
}

func (ComplaintReplyPolicy) UploadFile(data any, claims types.JWTCLaims) bool {
	return false
}
func (ComplaintReplyPolicy) ViewFile(data any, claims types.JWTCLaims) bool {
	return false
}
