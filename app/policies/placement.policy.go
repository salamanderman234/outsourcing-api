package policies

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/salamanderman234/outsourcing-api/app/types"
	"github.com/salamanderman234/outsourcing-api/app/types/enums"
)

type PlacementPolicy struct{}

func (PlacementPolicy) Create(claims jwt.Claims) bool {
	c, ok := claims.(types.JWTCLaims)
	if !ok {
		return false
	}
	role := c.Role
	return role == string(enums.SuperAdminUserRole) ||
		role == string(enums.AdminUserRole)
}

func (PlacementPolicy) ReadAll(claims jwt.Claims) bool {
	return true
}

func (PlacementPolicy) Find(data any, claims jwt.Claims) bool {
	return true
}

func (PlacementPolicy) Update(data any, claims jwt.Claims) bool {
	c, ok := claims.(types.JWTCLaims)
	if !ok {
		return false
	}
	role := c.Role
	return role == string(enums.SuperAdminUserRole) ||
		role == string(enums.AdminUserRole)
}

func (PlacementPolicy) Delete(data any, claims jwt.Claims) bool {
	c, ok := claims.(types.JWTCLaims)
	if !ok {
		return false
	}
	role := c.Role
	return role == string(enums.SuperAdminUserRole) ||
		role == string(enums.AdminUserRole)
}

func (PlacementPolicy) UploadFile(data any, claims jwt.Claims) bool {
	c, ok := claims.(types.JWTCLaims)
	if !ok {
		return false
	}
	role := c.Role
	return role == string(enums.SuperAdminUserRole) ||
		role == string(enums.AdminUserRole)
}
func (PlacementPolicy) ViewFile(data any, claims jwt.Claims) bool {
	c, ok := claims.(types.JWTCLaims)
	if !ok {
		return false
	}
	role := c.Role
	return role == string(enums.SuperAdminUserRole) ||
		role == string(enums.AdminUserRole)
}
