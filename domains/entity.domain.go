package domains

import "github.com/golang-jwt/jwt/v5"

// response entity
type BasicResponse struct {
	Message string `json:"message"`
	Body    any    `json:"Body"`
}

type DataBodyResponse struct {
	Data       any         `json:"data,omitempty"`
	Datas      any         `json:"datas,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

type ErrorBodyResponse struct {
	Error  *string               `json:"error,omitempty"`
	Errors []ErrorDetailResponse `json:"errors,omitempty"`
}

type ErrorDetailResponse struct {
	Field  *string `json:"field"`
	Rule   *string `json:"rule"`
	Detail *string `json:"detail"`
}

// pagination entity
type Pagination struct {
	Next     uint             `json:"next"`
	Current  uint             `json:"current"`
	Previous uint             `json:"previous"`
	MaxPage  uint             `json:"max_page"`
	Queries  []map[string]any `json:"queries"`
}

// authentication entity
type JWTClaims struct {
	jwt.RegisteredClaims
	JWTPayload
}
type JWTPayload struct {
	Username   *string `json:"username"`
	Role       *string `json:"role"`
	ProfilePic *string `json:"profile_pic"`
}
type TokenPair struct {
	Refresh string `json:"refresh_token"`
	Access  string `json:"access_token"`
}
