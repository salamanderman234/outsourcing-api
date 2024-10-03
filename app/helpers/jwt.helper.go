package helpers

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/salamanderman234/outsourcing-api/app/models"
	"github.com/salamanderman234/outsourcing-api/app/types"
	"github.com/salamanderman234/outsourcing-api/app/types/enums"
	"github.com/salamanderman234/outsourcing-api/configs"
)

type jwtHelper struct {
}

func (j jwtHelper) CreateToken(user models.User, subject enums.TokenType, exp ...int) (string, error) {
	email := ""
	role := ""
	v := ""

	if user.Email != nil {
		email = *user.Email
	}
	if user.Role != nil {
		role = *user.Role
	}
	if user.VerifiedAt != nil {
		v = user.VerifiedAt.String()
	}

	idRole := uint(0)

	if role == string(enums.SuperAdminUserRole) && user.SuperAdminProfile != nil {
		idRole = user.SuperAdminProfile.ID
	} else if role == string(enums.AdminUserRole) && user.AdminProfile != nil {
		idRole = user.AdminProfile.ID
	} else if role == string(enums.ServiceUserRole) && user.ServiceUserProfile != nil {
		idRole = user.ServiceUserProfile.ID
	} else if role == string(enums.EmployeeUserRole) && user.EmployeeProfile != nil {
		idRole = user.EmployeeProfile.ID
	} else if role == string(enums.SupervisorUserRole) && user.SupervisorProfile != nil {
		idRole = user.SupervisorProfile.ID
	}
	idStr := strconv.Itoa(int(user.ID))
	claims := types.JWTCLaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ID:      idStr,
			Issuer:  configs.AppConfig.Name,
			Subject: string(subject),
		},
		Email:     email,
		Role:      role,
		ProfileID: idRole,
		V:         v,
	}
	dur := configs.JWTConfig.GetDefaultEXP()
	if len(exp) == 1 {
		dur = exp[0]
	}
	if dur > 1 {
		exp := jwt.NewNumericDate(time.Now().Add(time.Duration(dur) * time.Hour))
		claims.RegisteredClaims.ExpiresAt = exp
	}
	token := jwt.NewWithClaims(
		configs.JWTConfig.SigningMethod,
		claims,
	)
	return token.SignedString([]byte(configs.JWTConfig.GetSecret()))
}

func (j jwtHelper) VerifyToken(token string) (types.JWTCLaims, error) {
	configs.JWTConfig.GetSigningMethod()
	tkn, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		method := t.Method
		if method != configs.JWTConfig.GetSigningMethod() {
			return nil, jwt.ErrTokenSignatureInvalid
		}
		return []byte(configs.JWTConfig.GetSecret()), nil
	})
	if err != nil {
		return types.JWTCLaims{}, err
	}
	if !tkn.Valid {
		return types.JWTCLaims{}, jwt.ErrTokenInvalidClaims
	}
	var result types.JWTCLaims
	claims, _ := tkn.Claims.(jwt.MapClaims)
	err = Translator.TranslateStruct(claims, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

var JWT = jwtHelper{}
