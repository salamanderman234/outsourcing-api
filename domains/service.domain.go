package domains

import (
	"context"
	"mime/multipart"

	"github.com/salamanderman234/outsourcing-api/configs"
)

// ----- FILE SERVICE -----
type FileService interface {
	Store(file *multipart.FileHeader, dest string, fileConfig configs.FileConfig) (string, error)
	BatchStore(files map[string]FileWrapper) (map[string]string, string, error)
	Destroy(target string) error
	// Read(target string) (*multipart.aff,FileHeader, error)
}

// ----- END OF FILE SERVICE -----

// ----- AUTH SERVICE -----
type BasicAuthService interface {
	// return access and refresh token or error if there's any
	Login(c context.Context, loginForm BasicLoginForm, remember bool) (TokenPair, UserEntity, error)
	Register(c context.Context, authData BasicRegisterForm, profileData any, role RoleEnum, remember bool) (TokenPair, UserWithProfileEntity, error)
	Check(c context.Context, token string) (JWTPayload, error)
	Refresh(c context.Context, refreshToken string) (TokenPair, error)
}

// ----- END OF AUTH SERVICE -----
// --> Basic
type BasicCrudService interface {
	Create(c context.Context, data any, ffiles ...EntityFileMap) (Entity, error)
	Read(c context.Context, id uint, q string, page uint, orderBy string, isDesc bool, withPagination bool) (any, *Pagination, error)
	Update(c context.Context, id uint, data any, files ...EntityFileMap) (int, any, error)
	Delete(c context.Context, id uint) (int, int, error)
}

// ----- USER SERVICE ------
type ServiceUserService interface {
	Read(c context.Context, id uint, q string, page uint, orderBy string, isDesc bool, withPagination bool) (any, *Pagination, error)
	Update(c context.Context, id uint, data ServiceUserUpdateForm, files ...EntityFileMap) (int, ServiceUserEntity, error)
	Delete(c context.Context, id uint) (int, int, error)
}
type EmployeeService interface {
	Read(c context.Context, id uint, q string, page uint, orderBy string, isDesc bool, withPagination bool) (any, *Pagination, error)
	Update(c context.Context, id uint, data EmployeeUpdateForm, files ...EntityFileMap) (int, EmployeeEntity, error)
	Delete(c context.Context, id uint) (int, int, error)
}
type AdminService interface {
	Read(c context.Context, id uint, q string, page uint, orderBy string, isDesc bool, withPagination bool) (any, *Pagination, error)
	Update(c context.Context, id uint, data AdminUpdateForm, files ...EntityFileMap) (int, AdminEntity, error)
	Delete(c context.Context, id uint) (int, int, error)
}
type SupervisorService interface {
	Read(c context.Context, id uint, q string, page uint, orderBy string, isDesc bool, withPagination bool) (any, *Pagination, error)
	Update(c context.Context, id uint, data SupervisorUpdateForm, files ...EntityFileMap) (int, SupervisorEntity, error)
	Delete(c context.Context, id uint) (int, int, error)
}

// ----- END OF USER SERVICE -----

// ----- MASTER DATA SERVICE -----
type CategoryService interface {
	Create(c context.Context, data CategoryCreateForm, files ...EntityFileMap) (CategoryEntity, error)
	Read(c context.Context, id uint, q string, page uint, orderBy string, isDesc bool, withPagination bool) (any, *Pagination, error)
	Update(c context.Context, id uint, data CategoryUpdateForm, files ...EntityFileMap) (int, CategoryEntity, error)
	Delete(c context.Context, id uint) (int, int, error)
}

// ----- END OF MASTER DATA SERVICE -----
// ---- APP SERVICE SERVICE -----
type ServiceItemService interface {
	Create(c context.Context, data ServiceItemCreateForm, files ...EntityFileMap) (ServiceItemEntity, error)
	Read(c context.Context, serviceId uint, id uint, q string, page uint, orderBy string, isDesc bool, withPagination bool) (any, *Pagination, error)
	Update(c context.Context, id uint, data ServiceItemUpdateForm, files ...EntityFileMap) (int, ServiceItemEntity, error)
	Delete(c context.Context, id uint) (int, int, error)
}
type PartialServiceService interface {
	Create(c context.Context, data PartialServiceCreateForm, files ...EntityFileMap) (ServiceEntity, error)
	Read(c context.Context, categoryId uint, id uint, q string, page uint, orderBy string, isDesc bool, withPagination bool) (any, *Pagination, error)
	Update(c context.Context, id uint, data PartialServiceUpdateForm, files ...EntityFileMap) (int, ServiceEntity, error)
	Delete(c context.Context, id uint) (int, int, error)
}

// ---- END OF APP SERVICE SERVICE -----
// ---- ORDER SERVICE ----
type OrderService interface {
	MakeOrder(c context.Context, user UserModel, orderData ServiceOrderForm) (ServiceOrderEntity, error)
	CancelOrder(c context.Context, user UserModel, orderId uint) (bool, error)
	ListOrder(c context.Context, user UserModel, serviceUserId uint, status string, page uint, orderBy string, desc bool, withPagination bool) ([]ServiceOrderEntity, *Pagination, error)
	DetailOrder(c context.Context, user UserModel, id uint) (ServiceOrderEntity, error)
	UpdateOrderStatus(c context.Context, user UserModel, id uint, updateForm ServiceOrderUpdateStatusForm) (int, ServiceOrderEntity, error)
}

// ---- END OF ORDER SERVICE -----
