package domains

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
	JWTPayload
}
type JWTPayload struct {
	ID         uint   `json:"id"`
	Username   string `json:"username"`
	Role       string `json:"role"`
	ProfilePic string `json:"profile_pic"`
}
type TokenPair struct {
	Refresh string `json:"refresh_token"`
	Access  string `json:"access_token"`
}
