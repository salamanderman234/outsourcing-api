package domains

import "context"

// ----- AUTH REPOSITORY -----
type BasicAuthRepository interface {
	// return user if creds is valid
	GetUserWithCreds(c context.Context, username string, password string) (any, error)
	RegisterUser(c context.Context, data any) (int64, any, error)
}

// ----- END OF AUTH REPOSITORY -----

// ----- CRUD REPOSITORY -----
type BasicCrudRepository interface {
	Create(c context.Context, data any) (any, error)
	FindByID(c context.Context, id uint) (any, error)
	Update(c context.Context, id uint, data any) (int64, any, error)
	Delete(c context.Context, id uint) (int64, any, error)
}
type ServiceUserRepository interface {
	BasicAuthRepository
	BasicCrudRepository
	Get(c context.Context, id uint, q string, page uint, orderBy string, desc bool) ([]ServiceUserModel, uint, error)
}

// ----- END OF CRUD REPOSITORY -----
