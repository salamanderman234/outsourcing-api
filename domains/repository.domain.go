package domains

import (
	"context"

	"gorm.io/gorm"
)

// ----- AUTH REPOSITORY -----
type BasicAuthRepository interface {
	// return user if creds is valid
	GetUserWithCreds(c context.Context, username string) (any, error)
	RegisterUser(c context.Context, autData any, profileData any) (int64, any, error)
	CreateProfile(c context.Context, role string, data any, userID uint, repo ...*gorm.DB) (any, error)
	BasicCrudRepository
}
type UserRepository interface {
	BasicAuthRepository
}

// ----- END OF AUTH REPOSITORY -----

// ----- CRUD REPOSITORY -----
type BasicCrudRepository interface {
	Create(c context.Context, data any, repo ...*gorm.DB) (any, error)
	FindByID(c context.Context, id uint) (any, error)
	Get(c context.Context, id uint, q string, page uint, orderBy string, desc bool) (any, uint, error)
	Update(c context.Context, id uint, data any) (int64, any, error)
	Delete(c context.Context, id uint) (int64, int64, error)
}
type ServiceUserRepository interface {
	BasicCrudRepository
}
type SupervisorRepository interface {
	BasicCrudRepository
}
type AdminRepository interface {
	BasicCrudRepository
}
type EmployeeRepository interface {
	BasicCrudRepository
}

type ServiceCategoryRepository interface {
	BasicCrudRepository
}
type DistrictRepository interface {
	BasicCrudRepository
}
type SubDistrictRepository interface {
	BasicCrudRepository
}
type VillageRepository interface {
	BasicCrudRepository
}

// ----- END OF CRUD REPOSITORY -----
