package policies

import "github.com/salamanderman234/outsourcing-api/app/types"

type userPolicy struct{}

func (userPolicy) RegisterUser(role string, claims types.JWTCLaims) bool {
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
