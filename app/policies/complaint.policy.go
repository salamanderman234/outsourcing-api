package policies

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/salamanderman234/outsourcing-api/app/models"
	"github.com/salamanderman234/outsourcing-api/app/types"
	"github.com/salamanderman234/outsourcing-api/app/types/enums"
)

type ComplaintPolicy struct {
}

func (ComplaintPolicy) Create(claims jwt.Claims) bool {
	c, ok := claims.(types.JWTCLaims)
	if !ok {
		return false
	}
	return c.Role == string(enums.AdminUserRole) ||
		c.Role == string(enums.SuperAdminUserRole) ||
		c.Role == string(enums.ServiceUserRole)
}

func (ComplaintPolicy) ReadAll(claims jwt.Claims) bool {
	c, ok := claims.(types.JWTCLaims)
	if !ok {
		return false
	}
	return c.Role == string(enums.AdminUserRole) ||
		c.Role == string(enums.SuperAdminUserRole)
}

func (ComplaintPolicy) Find(data any, claims jwt.Claims) bool {
	c, ok := claims.(types.JWTCLaims)
	if !ok {
		return false
	}
	role := c.Role
	trans, ok := data.(*models.Complaint)
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

func (ComplaintPolicy) Update(data any, claims jwt.Claims) bool {
	c, ok := claims.(types.JWTCLaims)
	if !ok {
		return false
	}
	role := c.Role
	trans, ok := data.(*models.Complaint)
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

func (ComplaintPolicy) Delete(data any, claims jwt.Claims) bool {
	c, ok := claims.(types.JWTCLaims)
	if !ok {
		return false
	}
	role := c.Role
	trans, ok := data.(*models.Complaint)
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

func (ComplaintPolicy) UploadFile(data any, claims jwt.Claims) bool {
	return false
}
func (ComplaintPolicy) ViewFile(data any, claims jwt.Claims) bool {
	return false
}

type ComplaintReplyPolicy struct {
}

func (ComplaintReplyPolicy) Create(claims jwt.Claims) bool {
	c, ok := claims.(types.JWTCLaims)
	if !ok {
		return false
	}
	return c.Role == string(enums.AdminUserRole) ||
		c.Role == string(enums.SuperAdminUserRole) ||
		c.Role == string(enums.SupervisorUserRole)
}

func (ComplaintReplyPolicy) ReadAll(claims jwt.Claims) bool {
	c, ok := claims.(types.JWTCLaims)
	if !ok {
		return false
	}
	return c.Role == string(enums.AdminUserRole) ||
		c.Role == string(enums.SuperAdminUserRole) ||
		c.Role == string(enums.SupervisorUserRole)
}

func (ComplaintReplyPolicy) Find(data any, claims jwt.Claims) bool {
	c, ok := claims.(types.JWTCLaims)
	if !ok {
		return false
	}
	role := c.Role
	trans, ok := data.(models.ComplaintReply)
	if !ok {
		return false
	}
	transUserID := uint(0)
	if trans.Supervisor != nil {
		transUserID = *trans.SupervisorID
	}
	return role == string(enums.SuperAdminUserRole) ||
		role == string(enums.AdminUserRole) ||
		(role == string(enums.SupervisorUserRole) && transUserID == c.ProfileID) ||
		role == string(enums.ServiceUserRole)
}

func (ComplaintReplyPolicy) Update(data any, claims jwt.Claims) bool {
	c, ok := claims.(types.JWTCLaims)
	if !ok {
		return false
	}
	role := c.Role
	trans, ok := data.(models.ComplaintReply)
	if !ok {
		return false
	}
	transUserID := uint(0)
	if trans.Supervisor != nil {
		transUserID = *trans.SupervisorID
	}
	return role == string(enums.SuperAdminUserRole) ||
		role == string(enums.AdminUserRole) ||
		(role == string(enums.SupervisorUserRole) && transUserID == c.ProfileID)
}

func (ComplaintReplyPolicy) Delete(data any, claims jwt.Claims) bool {
	c, ok := claims.(types.JWTCLaims)
	if !ok {
		return false
	}
	role := c.Role
	trans, ok := data.(models.ComplaintReply)
	if !ok {
		return false
	}
	transUserID := uint(0)
	if trans.Supervisor != nil {
		transUserID = *trans.SupervisorID
	}
	return role == string(enums.SuperAdminUserRole) ||
		role == string(enums.AdminUserRole) ||
		(role == string(enums.SupervisorUserRole) && transUserID == c.ProfileID)
}

func (ComplaintReplyPolicy) UploadFile(data any, claims jwt.Claims) bool {
	return false
}
func (ComplaintReplyPolicy) ViewFile(data any, claims jwt.Claims) bool {
	return false
}
