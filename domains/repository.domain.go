package domains

import (
	"context"

	"gorm.io/gorm"
)

// ----- AUTH REPOSITORY -----
type UserRepository interface {
	GetUserWithCreds(c context.Context, username string) (UserModel, error)
	RegisterUser(c context.Context, autData UserModel, profileData Model) (int64, UserWithProfileModel, error)
	CreateProfile(c context.Context, role string, data Model, userID uint, repo ...*gorm.DB) (Model, error)
	Find(c context.Context, id uint) (UserModel, error)
	Read(c context.Context, q string, page uint, orderBy string, desc bool, withPagination bool) ([]UserModel, uint, error)
	Update(c context.Context, id uint, data UserModel, repo ...*gorm.DB) (int64, UserModel, error)
	Delete(c context.Context, id uint, repo ...*gorm.DB) (int64, int64, error)
}

// ----- END OF AUTH REPOSITORY -----

// ----- CRUD REPOSITORY -----

// --> Profile
type ServiceUserRepository interface {
	Create(c context.Context, data ServiceUserModel, repo ...*gorm.DB) (ServiceUserModel, error)
	Find(c context.Context, id uint) (ServiceUserModel, error)
	Read(c context.Context, q string, page uint, orderBy string, desc bool, withPagination bool) ([]ServiceUserModel, uint, error)
	Update(c context.Context, id uint, data ServiceUserModel, repo ...*gorm.DB) (int64, ServiceUserModel, error)
	Delete(c context.Context, id uint, repo ...*gorm.DB) (int64, int64, error)
}
type SupervisorRepository interface {
	Create(c context.Context, data SupervisorModel, repo ...*gorm.DB) (SupervisorModel, error)
	Find(c context.Context, id uint) (SupervisorModel, error)
	Read(c context.Context, employeeStatus EmployeeStatusEnum, q string, page uint, orderBy string, desc bool, withPagination bool) ([]SupervisorModel, uint, error)
	Update(c context.Context, id uint, data SupervisorModel, repo ...*gorm.DB) (int64, SupervisorModel, error)
	Delete(c context.Context, id uint, repo ...*gorm.DB) (int64, int64, error)
}
type AdminRepository interface {
	Create(c context.Context, data AdminModel, repo ...*gorm.DB) (AdminModel, error)
	Find(c context.Context, id uint) (AdminModel, error)
	Read(c context.Context, q string, page uint, orderBy string, desc bool, withPagination bool) ([]AdminModel, uint, error)
	Update(c context.Context, id uint, data AdminModel, repo ...*gorm.DB) (int64, AdminModel, error)
	Delete(c context.Context, id uint, repo ...*gorm.DB) (int64, int64, error)
}
type EmployeeRepository interface {
	Create(c context.Context, data EmployeeModel, repo ...*gorm.DB) (EmployeeModel, error)
	Find(c context.Context, id uint) (EmployeeModel, error)
	Read(c context.Context, category string, employeeStatus EmployeeStatusEnum, q string, page uint, orderBy string, desc bool, withPagination bool) ([]EmployeeModel, uint, error)
	Update(c context.Context, id uint, data EmployeeModel, repo ...*gorm.DB) (int64, EmployeeModel, error)
	Delete(c context.Context, id uint, repo ...*gorm.DB) (int64, int64, error)
}

// --> Master data
type CategoryRepository interface {
	Create(c context.Context, data CategoryModel, repo ...*gorm.DB) (CategoryModel, error)
	Find(c context.Context, id uint) (CategoryModel, error)
	Read(c context.Context, q string, page uint, orderBy string, desc bool, withPagination bool) ([]CategoryModel, uint, error)
	Update(c context.Context, id uint, data CategoryModel, repo ...*gorm.DB) (int64, CategoryModel, error)
	Delete(c context.Context, id uint, repo ...*gorm.DB) (int64, int64, error)
}
type DistrictRepository interface {
	Create(c context.Context, data DistrictModel, repo ...*gorm.DB) (DistrictModel, error)
	Find(c context.Context, id uint) (DistrictModel, error)
	Read(c context.Context, q string, page uint, orderBy string, desc bool, withPagination bool) ([]DistrictModel, uint, error)
	Update(c context.Context, id uint, data DistrictModel, repo ...*gorm.DB) (int64, DistrictModel, error)
	Delete(c context.Context, id uint, repo ...*gorm.DB) (int64, int64, error)
}
type SubDistrictRepository interface {
	Create(c context.Context, data SubDistrictModel, repo ...*gorm.DB) (SubDistrictModel, error)
	Find(c context.Context, id uint) (SubDistrictModel, error)
	Read(c context.Context, q string, page uint, orderBy string, desc bool, withPagination bool) ([]SubDistrictModel, uint, error)
	Update(c context.Context, id uint, data SubDistrictModel, repo ...*gorm.DB) (int64, SubDistrictModel, error)
	Delete(c context.Context, id uint, repo ...*gorm.DB) (int64, int64, error)
}
type VillageRepository interface {
	Create(c context.Context, data VillageModel, repo ...*gorm.DB) (VillageModel, error)
	Find(c context.Context, id uint) (VillageModel, error)
	Read(c context.Context, q string, page uint, orderBy string, desc bool, withPagination bool) ([]VillageModel, uint, error)
	Update(c context.Context, id uint, data VillageModel, repo ...*gorm.DB) (int64, VillageModel, error)
	Delete(c context.Context, id uint, repo ...*gorm.DB) (int64, int64, error)
}

// ----- END OF CRUD REPOSITORY -----
// ----- APP SERVICE REPOSITORY -----
type ServiceItemRepository interface {
	Create(c context.Context, data ServiceItemModel, repo ...*gorm.DB) (ServiceItemModel, error)
	Find(c context.Context, id uint) (ServiceItemModel, error)
	ReadIn(c context.Context, field string, conds []string) ([]ServiceItemModel, error)
	Read(c context.Context, serviceId uint, q string, page uint, orderBy string, desc bool, withPagination bool) ([]ServiceItemModel, uint, error)
	Update(c context.Context, id uint, data ServiceItemModel, repo ...*gorm.DB) (int64, ServiceItemModel, error)
	Delete(c context.Context, id uint, repo ...*gorm.DB) (int64, int64, error)
}
type PartialServiceRepository interface {
	Create(ctx context.Context, partialServiceData ServiceModel, repo ...*gorm.DB) (ServiceModel, error)
	Read(ctx context.Context, categoryId uint, q string, page uint, orderBy string, desc bool, withPagination bool) ([]ServiceModel, uint, error)
	ReadIn(c context.Context, field string, conds []string) ([]ServiceModel, error)
	Find(ctx context.Context, id uint) (ServiceModel, error)
	Update(ctx context.Context, id uint, data ServiceModel, repo ...*gorm.DB) (int, ServiceModel, error)
	Delete(ctx context.Context, id uint, repo ...*gorm.DB) (int, int, error)
}

// ----- END OF APP SERVICE REPOSITORY -----
// ----- SERVICE ORDER REPOSITORY -----
type ServiceOrderRepository interface {
	Create(c context.Context, data ServiceOrderModel, repo ...*gorm.DB) (ServiceOrderModel, error)
	Read(c context.Context, status string, service_user_id uint, page uint, orderBy string, desc bool, withPagination bool) ([]ServiceOrderModel, uint, error)
	Find(c context.Context, id uint) (ServiceOrderModel, error)
	Update(c context.Context, id uint, data ServiceOrderModel, repo ...*gorm.DB) (int64, ServiceOrderModel, error)
	Delete(c context.Context, id uint, repo ...*gorm.DB) (uint, int64, error)
}
type ServiceOrderDetailRepository interface {
	Create(c context.Context, data ServiceOrderDetailModel, repo ...*gorm.DB) (ServiceOrderDetailModel, error)
	Read(c context.Context, serviceOrderId uint) ([]ServiceOrderDetailModel, error)
	Find(c context.Context, id uint) (ServiceOrderDetailModel, error)
	Update(c context.Context, id uint, data ServiceOrderDetailModel, repo ...*gorm.DB) (int, ServiceOrderDetailModel, error)
	Delete(c context.Context, id uint, repo ...*gorm.DB) (uint, int64, error)
}
type ServiceOrderDetailItemRepository interface {
	Create(c context.Context, data ServiceOrderDetailItemModel, repo ...*gorm.DB) (ServiceOrderDetailItemModel, error)
	Read(c context.Context, serviceOrderDetailId uint) ([]ServiceOrderDetailItemModel, error)
	Find(c context.Context, id uint) (ServiceOrderDetailItemModel, error)
	Update(c context.Context, id uint, data ServiceOrderDetailItemModel, repo ...*gorm.DB) (int, ServiceOrderDetailItemModel, error)
	Delete(c context.Context, id uint, repo ...*gorm.DB) (uint, int64, error)
}

// ----- END OF SERVICE ORDER REPOSITORY -----
// ----- PLACEMENT REPOSITORY -----
type ServiceOrderPlacementRepository interface {
	Create(c context.Context, data ServiceOrderPlacementModel, repo ...*gorm.DB) (ServiceOrderPlacementModel, error)
	Read(c context.Context, orderId uint, page uint, orderBy string, desc bool, withPagination bool) ([]ServiceOrderPlacementModel, uint, error)
	Find(c context.Context, id uint) (ServiceOrderPlacementModel, error)
	Update(c context.Context, id uint, data ServiceOrderPlacementModel, repo ...*gorm.DB) (int, ServiceOrderPlacementModel, error)
	Delete(c context.Context, id uint, repo ...*gorm.DB) (int, int, error)
}
type ServiceOrderPlacementDailyReportRepository interface {
	Create(c context.Context, data ServiceOrderPlacementDailyReportModel, repo ...*gorm.DB) (ServiceOrderPlacementDailyReportModel, error)
	Find(c context.Context, id uint) (ServiceOrderPlacementDailyReportModel, error)
	Update(c context.Context, id uint, data ServiceOrderPlacementDailyReportModel, repo ...*gorm.DB) (int, ServiceOrderPlacementDailyReportModel, error)
	Delete(c context.Context, id uint, repo ...*gorm.DB) (int, int, error)
}
type ServiceOrderPlacementServiceRepository interface {
	BatchCreate(c context.Context, placementId uint, datas []ServiceOrderPlacementServiceModel, repo ...*gorm.DB) ([]ServiceOrderPlacementServiceModel, error)
	Create(c context.Context, data ServiceOrderPlacementServiceModel, repo ...*gorm.DB) (ServiceOrderPlacementServiceModel, error)
	Find(c context.Context, id uint) (ServiceOrderPlacementServiceModel, error)
	Update(c context.Context, id uint, data ServiceOrderPlacementServiceModel, repo ...*gorm.DB) (int, ServiceOrderPlacementServiceModel, error)
	Delete(c context.Context, id uint, repo ...*gorm.DB) (int, int, error)
}
type ServiceOrderPlacementServiceEmployeeRepository interface {
	Create(c context.Context, data ServiceOrderPlacementServiceEmployeeModel, repo ...*gorm.DB) (ServiceOrderPlacementServiceEmployeeModel, error)
	// Read(c context.Context, placementId uint, page uint, orderBy string, desc bool, withPagination bool) ([]ServiceOrderPlacementServiceEmployeeModel, uint, error)
	Find(c context.Context, id uint) (ServiceOrderPlacementServiceEmployeeModel, error)
	Update(c context.Context, id uint, data ServiceOrderPlacementServiceEmployeeModel, repo ...*gorm.DB) (int, ServiceOrderPlacementServiceEmployeeModel, error)
	Delete(c context.Context, id uint, repo ...*gorm.DB) (int, int, error)
}
type ServiceOrderPlacementServiceEmployeeScheduleRepository interface {
	Create(c context.Context, data ServiceOrderPlacementServiceEmployeeScheduleModel, repo ...*gorm.DB) (ServiceOrderPlacementServiceEmployeeScheduleModel, error)
	Find(c context.Context, id uint) (ServiceOrderPlacementServiceEmployeeScheduleModel, error)
	Update(c context.Context, id uint, data ServiceOrderPlacementServiceEmployeeScheduleModel, repo ...*gorm.DB) (int, ServiceOrderPlacementServiceEmployeeModel, error)
	Delete(c context.Context, id uint, repo ...*gorm.DB) (int, int, error)
}

// ----- END OF PLACEMENT REPOSITORY -----
