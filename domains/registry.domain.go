package domains

var (
	ServiceRegistry = struct {
		AuthServ        BasicAuthService
		CategoryServ    CategoryService
		FileServ        FileService
		ServiceItemServ ServiceItemService
		ServiceServ     PartialServiceService
		OrderServ       OrderService
		UserServ        UserService
		ServiceUserServ ServiceUserService
		EmployeeServ    EmployeeService
		AdminServ       AdminService
		SupervisorServ  SupervisorService
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
		UserView        UserView
		ServiceUserView UserServiceView
		EmployeeView    EmployeeView
		SupervisorView  SupervisorView
		AdminView       AdminView
		CategoryView    CategoryView
		ServiceItemView ServiceItemView
		ServiceView     PartialServiceView
		OrderView       ServiceOrderView
	}{}
)
