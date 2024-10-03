package enums

type UserRolesEnum string

const (
	AdminUserRole              UserRolesEnum = "admin"
	SuperAdminUserRole         UserRolesEnum = "super_admin"
	ServiceUserRole            UserRolesEnum = "service_user"
	SupervisorUserRole         UserRolesEnum = "supervisor"
	EmployeeSupervisorUserRole UserRolesEnum = "employee_supervisor"
	EmployeeUserRole           UserRolesEnum = "employee"
	ApplicationUserRole        UserRolesEnum = "application"
)
