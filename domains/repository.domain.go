package domains

import "context"

// ----- AUTH REPOSITORY -----
type AuthRepository interface {
	// return user if creds is valid
	GetUserWithCreds(c context.Context, username string, password string) (any, error)
	RegisterUser(c context.Context, data any) (int64, any, error)
}

// ----- END OF AUTH REPOSITORY -----
