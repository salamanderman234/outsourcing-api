package domains

var (
	AuthServiceRegistry = struct {
		AuthServ BasicAuthService
	}{}

	ServiceRegistry = struct{}{}

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
