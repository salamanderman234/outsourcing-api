package domains

type RoleEnum string

const (
	AdminRole       RoleEnum = "admin"
	EmployeeRole    RoleEnum = "employee"
	ServiceUserRole RoleEnum = "service_user"
	SupervisorRole  RoleEnum = "supervisor"
)
