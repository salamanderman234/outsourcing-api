package domains

var (
	AuthServiceRegistry = struct {
		AuthServ BasicAuthService
	}{}

	RepoRegistry = struct {
		UserRepo        UserRepository
		ServiceUserRepo ServiceUserRepository
		SupervisorRepo  SupervisorRepository
		EmployeeRepo    EmployeeRepository
		AdminRepo       AdminRepository
	}{}
)
