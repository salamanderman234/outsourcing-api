package domains

var (
	ServiceRegistry = struct {
		AuthServ BasicAuthService
	}{}

	RepoRegistry = struct {
		UserRepo        UserRepository
		ServiceUserRepo ServiceUserRepository
		SupervisorRepo  SupervisorRepository
		EmployeeRepo    EmployeeRepository
		AdminRepo       AdminRepository
	}{}

	ViewRegistry = struct {
		AuthView BasicAuthView
	}{}
)
