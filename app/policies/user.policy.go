package policies

import (
	"github.com/salamanderman234/outsourcing-api/app/types"
	"github.com/salamanderman234/outsourcing-api/app/types/enums"
)

type UserPolicy struct {
}

func (UserPolicy) Create(claims types.JWTCLaims) bool {
	return true
}

func (UserPolicy) ReadAll(claims types.JWTCLaims) bool {
	// role := claims.Role
	// return (role == string(enums.AdminUserRole) || role == string(enums.SuperAdminRole))
	return true
}

func (UserPolicy) Find(data any, claims types.JWTCLaims) bool {
	// user, ok := data.(models.User)
	// if !ok {
	// 	return false
	// }
	// role := *user.Role
	// id, _ := strconv.Atoi(claims.ID)
	// return (role == string(enums.AdminUserRole) || role == string(enums.SuperAdminRole)) || user.ID == uint(id)
	return true
}

func (UserPolicy) Update(data any, claims types.JWTCLaims) bool {
	// user, ok := data.(models.User)
	// if !ok {
	// 	return false
	// }
	// role := *user.Role
	// id, _ := strconv.Atoi(claims.ID)
	// return (role == string(enums.AdminUserRole) || role == string(enums.SuperAdminRole)) || user.ID == uint(id)
	return true
}

func (UserPolicy) Delete(data any, claims types.JWTCLaims) bool {
	role := claims.Role
	return (role == string(enums.AdminUserRole) || role == string(enums.SuperAdminRole))
}

func (UserPolicy) UploadFile(data any, claims types.JWTCLaims) bool {
	return true
}
func (UserPolicy) ViewFile(data any, claims types.JWTCLaims) bool {
	return true
}
