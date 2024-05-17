package types

import "github.com/golang-jwt/jwt/v5"

type JWTCLaims struct {
	jwt.RegisteredClaims
	Email string
	Role  string
}

type AuthSessionKey string

var (
	UserContextKey AuthSessionKey = "user"
)
