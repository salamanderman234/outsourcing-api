package policies

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/salamanderman234/outsourcing-api/app/models"
	"github.com/salamanderman234/outsourcing-api/app/types"
	"github.com/salamanderman234/outsourcing-api/app/types/enums"
)

type FeedbackPolicy struct{}

func (FeedbackPolicy) Create(claims jwt.Claims) bool {
	c, ok := claims.(types.JWTCLaims)
	if !ok {
		return false
	}
	return c.Role == string(enums.AdminUserRole) ||
		c.Role == string(enums.SuperAdminUserRole) ||
		c.Role == string(enums.ServiceUserRole)
}

func (FeedbackPolicy) ReadAll(claims jwt.Claims) bool {
	c, ok := claims.(types.JWTCLaims)
	if !ok {
		return false
	}
	return c.Role == string(enums.AdminUserRole) ||
		c.Role == string(enums.SuperAdminUserRole) ||
		c.Role == string(enums.ServiceUserRole)
}

func (FeedbackPolicy) Find(data any, claims jwt.Claims) bool {
	c, ok := claims.(types.JWTCLaims)
	if !ok {
		return false
	}
	role := c.Role
	trans, ok := data.(*models.Feedback)
	if !ok {
		return false
	}
	transUserID := uint(0)
	if trans.ServiceUserID != nil {
		transUserID = *trans.ServiceUserID
	}
	return role == string(enums.SuperAdminUserRole) ||
		role == string(enums.AdminUserRole) ||
		(role == string(enums.ServiceUserRole) && transUserID == c.ProfileID)
}

func (FeedbackPolicy) Update(data any, claims jwt.Claims) bool {
	c, ok := claims.(types.JWTCLaims)
	if !ok {
		return false
	}
	role := c.Role
	trans, ok := data.(*models.Feedback)
	if !ok {
		return false
	}
	transUserID := uint(0)
	if trans.ServiceUserID != nil {
		transUserID = *trans.ServiceUserID
	}
	return role == string(enums.SuperAdminUserRole) ||
		role == string(enums.AdminUserRole) ||
		(role == string(enums.ServiceUserRole) && transUserID == c.ProfileID)
}

func (FeedbackPolicy) Delete(data any, claims jwt.Claims) bool {
	c, ok := claims.(types.JWTCLaims)
	if !ok {
		return false
	}
	role := c.Role
	trans, ok := data.(*models.Feedback)
	if !ok {
		return false
	}
	transUserID := uint(0)
	if trans.ServiceUserID != nil {
		transUserID = *trans.ServiceUserID
	}
	return role == string(enums.SuperAdminUserRole) ||
		role == string(enums.AdminUserRole) ||
		(role == string(enums.ServiceUserRole) && transUserID == c.ProfileID)
}

func (FeedbackPolicy) UploadFile(data any, claims jwt.Claims) bool {
	return false
}
func (FeedbackPolicy) ViewFile(data any, claims jwt.Claims) bool {
	return false
}
