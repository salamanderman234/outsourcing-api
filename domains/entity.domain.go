package domains

import (
	"mime/multipart"

	"github.com/golang-jwt/jwt/v5"
)

// basic
type BasicEntity struct {
	ID uint `json:"id"`
}

// response entity
type BasicResponse struct {
	Message string `json:"message"`
	Body    any    `json:"payload"`
}

type DataBodyResponse struct {
	Data       any         `json:"data,omitempty"`
	Datas      any         `json:"datas,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
	Updated    any         `json:"updated,omitempty"`
	ID         *uint       `json:"id,omitempty"`
	Affected   *int        `json:"affected_row,omitempty"`
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

// file
type EntityFileMap struct {
	Field string
	File  *multipart.FileHeader
}

// pagination entity
type Pagination struct {
	Next     uint           `json:"next"`
	Current  uint           `json:"current"`
	Previous uint           `json:"previous"`
	MaxPage  uint           `json:"max_page"`
	Queries  map[string]any `json:"queries"`
}

type PaginationQuery struct {
	Query string `json:"query,omitempty"`
	Value any    `json:"value,omitempty"`
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

// ----- AUTH ------
type UserEntity struct {
	BasicEntity
	Email       string             `json:"email"`
	Role        string             `json:"role"`
	ServiceUser *ServiceUserEntity `json:"service_user_profile,omitempty"`
	Supervisor  *SupervisorEntity  `json:"supervisor_profile,omitempty"`
	Admin       *AdminEntity       `json:"admin_profile,omitempty"`
	Employee    *EmployeeEntity    `json:"employee_profile,omitempty"`
}
type ServiceUserEntity struct {
	BasicEntity
	User               *UserEntity `json:"user,omitempty"`
	Address            string      `json:"address"`
	Fullname           string      `json:"Fullname"`
	IdentityCardNumber string      `json:"identity_card_number"`
	Phone              string      `json:"phone"`
}
type SupervisorEntity struct {
	BasicEntity
	User               *UserEntity `json:"user,omitempty"`
	Address            string      `json:"address"`
	Fullname           string      `json:"Fullname"`
	IdentityCardNumber string      `json:"identity_card_number"`
	Phone              string      `json:"phone"`
}
type AdminEntity struct {
	BasicEntity
	User     *UserEntity `json:"user,omitempty"`
	Address  string      `json:"address"`
	Fullname string      `json:"Fullname"`
	Phone    string      `json:"phone"`
}
type EmployeeEntity struct {
	BasicEntity
	User               *UserEntity `json:"user,omitempty"`
	Address            string      `json:"address"`
	Fullname           string      `json:"fullname"`
	IdentityCardNumber string      `json:"identity_card_number"`
	Phone              string      `json:"phone"`
}
type UserWithProfileEntity struct {
	User    UserEntity `json:"created_user"`
	Profile any        `json:"created_profile"`
}

// ----- END OF AUTH ----
// ----- MASTER DATA -----
type CategoryEntity struct {
	BasicEntity
	CategoryName *string `json:"category_name"`
	Icon         string  `json:"icon"`
	Description  string  `json:"description"`
}
type DistrictEntity struct {
	BasicEntity
	DisctrictName string              `json:"district_name"`
	Description   string              `json:"description"`
	SubDistricts  []SubDistrictEntity `json:"sub_districts"`
}
type SubDistrictEntity struct {
	BasicEntity
	SubDisctrictName string         `json:"subdistrict_name"`
	Description      string         `json:"description"`
	District         DistrictEntity `json:"district"`
	SubDistricts     []VillageModel `json:"villages"`
}
type VillageEntity struct {
	BasicEntity
	VillageName   *string           `json:"village_name"`
	Description   string            `json:"description"`
	SubDistrictID *uint             `json:"subdistrict_id"`
	SubDistrict   *SubDistrictModel `json:"subdistrict"`
}

// ----- END OF MASTER DATA -----
