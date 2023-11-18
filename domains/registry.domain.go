package domains

var (
	AuthServiceRegistry = struct {
		ServiceUserAuthServ ServiceUserAuthService
		SupervisorAuthServ  SupervisorAuthService
		EmployeeSAuthServ   EmployeeAuthService
		AdminAuthServ       AdminAuthService
	}{}

	RepoRegistry = struct {
		ServiceUserRepo ServiceUserRepository
		SupervisorRepo  SupervisorRepository
		EmployeeSRepo   EmployeeRepository
		AdminRepo       AdminRepository
	}{}
)
