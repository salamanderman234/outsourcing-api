package domains

import "context"

// ----- AUTH SERVICE -----
type BasicAuthService interface {
	// return access and refresh token or error if there's any
	Login(c context.Context, loginForm BasicLoginForm, remember bool) (TokenPair, error)
	Register(c context.Context, authData BasicRegisterForm, profileData any, role RoleEnum, remember bool) (TokenPair, error)
	Check(c context.Context, token string) (JWTPayload, error)
	Refresh(c context.Context, refreshToken string) (TokenPair, error)
}

// ----- END OF AUTH SERVICE -----
// --> Basic
type BasicCrudService interface {
	Create(c context.Context, data any) (any, error)
	Read(c context.Context, id uint, q string, page uint, orderBy string, isDesc bool) (any, Pagination, error)
	Update(c context.Context, id uint, data any) (int, any, error)
	Delete(c context.Context, id uint) (int, int, error)
}

// ----- MASTER DATA SERVICE -----
type ServiceCategoryService interface {
	BasicCrudService
}
type DistrictService interface {
	BasicCrudService
}
type SubDistrictService interface {
	BasicCrudService
}
type VillageService interface {
	BasicCrudService
}

//----- END OF MASTER DATA SERVICE -----
