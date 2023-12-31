package domains

import "context"

// ----- AUTH SERVICE -----
type BasicAuthService interface {
	// return access and refresh token or error if there's any
	Login(c context.Context, loginForm BasicLoginForm) (TokenPair, error)
	Register(c context.Context, data any) (TokenPair, error)
	Check(c context.Context, token string) (JWTPayload, error)
	Refresh(c context.Context, refreshToken string) (TokenPair, error)
}

type ServiceUserAuthService interface {
	BasicAuthService
}
type SupervisorAuthService interface {
	BasicAuthService
}
type EmployeeAuthService interface {
	BasicAuthService
}
type AdminAuthService interface {
	BasicAuthService
}

//----- END OF AUTH SERVICE -----
