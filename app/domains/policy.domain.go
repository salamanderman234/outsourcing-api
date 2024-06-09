package domains

import (
	"github.com/golang-jwt/jwt/v5"
)

type Policy interface {
	Create(claims jwt.Claims) bool
	ReadAll(claims jwt.Claims) bool
	Find(data any, claims jwt.Claims) bool
	Update(data any, claims jwt.Claims) bool
	Delete(data any, claims jwt.Claims) bool
	UploadFile(data any, claims jwt.Claims) bool
	ViewFile(data any, claims jwt.Claims) bool
}
