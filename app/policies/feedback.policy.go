package policies

import (
	"github.com/salamanderman234/outsourcing-api/app/types"
	"github.com/salamanderman234/outsourcing-api/app/types/enums"
)

type FeedbackPolicy struct{}

func (FeedbackPolicy) Create(claims types.JWTCLaims) bool {
	return claims.Role == string(enums.AdminUserRole) || true
}

func (FeedbackPolicy) ReadAll(claims types.JWTCLaims) bool {
	return claims.Role == string(enums.AdminUserRole) || true
}

func (FeedbackPolicy) Find(data any, claims types.JWTCLaims) bool {
	return true
}

func (FeedbackPolicy) Update(data any, claims types.JWTCLaims) bool {
	return claims.Role == string(enums.AdminUserRole) || true
}

func (FeedbackPolicy) Delete(data any, claims types.JWTCLaims) bool {
	return claims.Role == string(enums.AdminUserRole) || true
}

func (FeedbackPolicy) UploadFile(data any, claims types.JWTCLaims) bool {
	return false
}
func (FeedbackPolicy) ViewFile(data any, claims types.JWTCLaims) bool {
	return false
}
