package enums

type UserRolesEnum string

const (
	AdminUserRole              UserRolesEnum = "admin"
	ServiceUserRole            UserRolesEnum = "service"
	SupervisorUserRole         UserRolesEnum = "supervisor"
	EmployeeSupervisorUserRole UserRolesEnum = "employee_supervisor"
	EmployeeUserRole           UserRolesEnum = "employee"
)
