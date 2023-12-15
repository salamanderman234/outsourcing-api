package domains

var (
	ServiceRegistry = struct {
		AuthServ     BasicAuthService
		CategoryServ ServiceCategoryService
		FileServ     FileService
	}{}

	RepoRegistry = struct {
		UserRepo        UserRepository
		ServiceUserRepo ServiceUserRepository
		SupervisorRepo  SupervisorRepository
		EmployeeRepo    EmployeeRepository
		AdminRepo       AdminRepository
		CategoryRepo    ServiceCategoryRepository
	}{}

	ViewRegistry = struct {
		AuthView     BasicAuthView
		CategoryView ServiceCategoryView
	}{}
)
