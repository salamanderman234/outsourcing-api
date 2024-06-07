package policies

import (
	"github.com/salamanderman234/outsourcing-api/app/types"
)

type PlacementPolicy struct{}

func (PlacementPolicy) Create(claims types.JWTCLaims) bool {
	return true
}

func (PlacementPolicy) ReadAll(claims types.JWTCLaims) bool {
	// role := claims.Role
	// return (role == string(enums.AdminUserRole) || role == string(enums.SuperAdminRole))
	return true
}

func (PlacementPolicy) Find(data any, claims types.JWTCLaims) bool {
	// user, ok := data.(models.User)
	// if !ok {
	// 	return false
	// }
	// role := *user.Role
	// id, _ := strconv.Atoi(claims.ID)
	// return (role == string(enums.AdminUserRole) || role == string(enums.SuperAdminRole)) || user.ID == uint(id)
	return true
}

func (PlacementPolicy) Update(data any, claims types.JWTCLaims) bool {
	// user, ok := data.(models.User)
	// if !ok {
	// 	return false
	// }
	// role := *user.Role
	// id, _ := strconv.Atoi(claims.ID)
	// return (role == string(enums.AdminUserRole) || role == string(enums.SuperAdminRole)) || user.ID == uint(id)
	return true
}

func (PlacementPolicy) Delete(data any, claims types.JWTCLaims) bool {
	// role := claims.Role
	// return (role == string(enums.AdminUserRole) || role == string(enums.SuperAdminRole))
	return true
}

func (PlacementPolicy) UploadFile(data any, claims types.JWTCLaims) bool {
	return true
}
func (PlacementPolicy) ViewFile(data any, claims types.JWTCLaims) bool {
	return true
}
