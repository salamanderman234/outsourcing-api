package policies

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/salamanderman234/outsourcing-api/app/models"
	"github.com/salamanderman234/outsourcing-api/app/types"
	"github.com/salamanderman234/outsourcing-api/app/types/enums"
)

type TransactionPolicy struct{}

func (TransactionPolicy) Create(claims jwt.Claims) bool {
	c, ok := claims.(types.JWTCLaims)
	if !ok {
		return false
	}
	role := c.Role
	return role == string(enums.SuperAdminUserRole) ||
		role == string(enums.AdminUserRole) ||
		role == string(enums.ServiceUserRole)
}

func (TransactionPolicy) ReadAll(claims jwt.Claims) bool {
	c, ok := claims.(types.JWTCLaims)
	if !ok {
		return false
	}
	role := c.Role
	return role == string(enums.SuperAdminUserRole) ||
		role == string(enums.AdminUserRole) ||
		role == string(enums.ServiceUserRole)
}

func (TransactionPolicy) Find(data any, claims jwt.Claims) bool {
	c, ok := claims.(types.JWTCLaims)
	if !ok {
		return false
	}
	role := c.Role
	trans, ok := data.(*models.Transaction)
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

func (TransactionPolicy) Update(data any, claims jwt.Claims) bool {
	c, ok := claims.(types.JWTCLaims)
	if !ok {
		return false
	}
	role := c.Role
	trans, ok := data.(*models.Transaction)
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

func (TransactionPolicy) Delete(data any, claims jwt.Claims) bool {
	c, ok := claims.(types.JWTCLaims)
	if !ok {
		return false
	}
	role := c.Role
	trans, ok := data.(*models.Transaction)
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

func (TransactionPolicy) UploadFile(data any, claims jwt.Claims) bool {
	c, ok := claims.(types.JWTCLaims)
	if !ok {
		return false
	}
	role := c.Role
	trans, ok := data.(*models.Transaction)
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
func (TransactionPolicy) ViewFile(data any, claims jwt.Claims) bool {
	c, ok := claims.(types.JWTCLaims)
	if !ok {
		return false
	}
	role := c.Role
	trans, ok := data.(*models.Transaction)
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
