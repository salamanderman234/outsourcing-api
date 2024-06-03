package policies

import (
	"strconv"

	"github.com/salamanderman234/outsourcing-api/app/models"
	"github.com/salamanderman234/outsourcing-api/app/types"
	"github.com/salamanderman234/outsourcing-api/app/types/enums"
)

type PlacementPolicy struct{}

func (PlacementPolicy) Create(claims types.JWTCLaims) bool {
	return true
}

func (PlacementPolicy) ReadAll(claims types.JWTCLaims) bool {
	role := claims.Role
	return (role == string(enums.AdminUserRole) || role == string(enums.SuperAdminRole))
}

func (PlacementPolicy) Find(data any, claims types.JWTCLaims) bool {
	user, ok := data.(models.User)
	if !ok {
		return false
	}
	role := *user.Role
	id, _ := strconv.Atoi(claims.ID)
	return (role == string(enums.AdminUserRole) || role == string(enums.SuperAdminRole)) || user.ID == uint(id)
}

func (PlacementPolicy) Update(data any, claims types.JWTCLaims) bool {
	user, ok := data.(models.User)
	if !ok {
		return false
	}
	role := *user.Role
	id, _ := strconv.Atoi(claims.ID)
	return (role == string(enums.AdminUserRole) || role == string(enums.SuperAdminRole)) || user.ID == uint(id)
}

func (PlacementPolicy) Delete(data any, claims types.JWTCLaims) bool {
	role := claims.Role
	return (role == string(enums.AdminUserRole) || role == string(enums.SuperAdminRole))
}

func (PlacementPolicy) UploadFile(data any, claims types.JWTCLaims) bool {
	return true
}
func (PlacementPolicy) ViewFile(data any, claims types.JWTCLaims) bool {
	return true
}
