package domains

var (
	ServiceRegistry = struct {
		AuthServ        BasicAuthService
		CategoryServ    CategoryService
		FileServ        FileService
		ServiceItemServ ServiceItemService
		ServiceServ     PartialServiceService
		OrderServ       OrderService
	}{}

	RepoRegistry = struct {
		UserRepo                   UserRepository
		ServiceUserRepo            ServiceUserRepository
		SupervisorRepo             SupervisorRepository
		EmployeeRepo               EmployeeRepository
		AdminRepo                  AdminRepository
		CategoryRepo               CategoryRepository
		ServiceItemRepo            ServiceItemRepository
		ServiceRepo                PartialServiceRepository
		ServiceOrderRepo           ServiceOrderRepository
		ServiceOrderDetailRepo     ServiceOrderDetailRepository
		ServiceOrderDetailItemRepo ServiceOrderDetailItemRepository
	}{}

	ViewRegistry = struct {
		AuthView        BasicAuthView
		CategoryView    CategoryView
		ServiceItemView ServiceItemView
		ServiceView     PartialServiceView
	}{}
)
