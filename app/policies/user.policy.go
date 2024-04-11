package policies

import (
	auth_types "github.com/salamanderman234/outsourcing-api/app/domains/types/auth"
)

type userPolicy struct{}

func (userPolicy) RegisterUser(role string, claims auth_types.JWTCLaims) bool {
	// if role == string(enums.AdminUserRole) ||
	// 	role == string(enums.EmployeeUserRole) ||
	// 	role == string(enums.SupervisorUserRole) {

	// 	return claims.Role == string(enums.AdminUserRole)
	// } else if role == string(enums.ServiceUserRole) {
	// 	return true
	// }
	// return false
	return true
}

var UserPolicy = userPolicy{}
