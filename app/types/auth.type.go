package types

import (
	"github.com/golang-jwt/jwt/v5"
)

type JWTCLaims struct {
	jwt.RegisteredClaims
	ProfileID uint
	Email     string
	Role      string
	V         string
}

type AuthSessionKey string

var (
	UserContextKey   AuthSessionKey = "user"
	AccessContextKey AuthSessionKey = "access"
)
