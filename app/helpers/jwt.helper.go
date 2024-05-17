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
	secret     string
	exp        int
	signMethod jwt.SigningMethod
}

func (j jwtHelper) CreateToken(user models.User, subject enums.TokenType) (string, error) {
	idStr := strconv.Itoa(int(user.ID))
	claims := types.JWTCLaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ID:      idStr,
			Issuer:  configs.AppConfig.Name,
			Subject: string(subject),
		},
		Email: *user.Email,
		Role:  *user.Role,
	}
	if j.exp > 1 {
		exp := jwt.NewNumericDate(time.Now().Add(time.Duration(j.exp) * time.Hour))
		claims.RegisteredClaims.ExpiresAt = exp
	}
	token := jwt.NewWithClaims(
		configs.JWTConfig.SigningMethod,
		claims,
	)
	return token.SignedString([]byte(j.secret))
}

func (j jwtHelper) VerifyToken(token string) (types.JWTCLaims, error) {
	tkn, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrTokenSignatureInvalid
		} else if method != j.signMethod {
			return nil, jwt.ErrTokenSignatureInvalid
		}
		return []byte(j.secret), nil
	})
	if err != nil {
		return types.JWTCLaims{}, err
	}
	if !tkn.Valid {
		return types.JWTCLaims{}, jwt.ErrTokenInvalidClaims
	}
	claims, _ := tkn.Claims.(types.JWTCLaims)
	return claims, nil
}

var JWT = jwtHelper{
	secret:     configs.JWTConfig.Secret,
	exp:        configs.JWTConfig.Exp,
	signMethod: configs.JWTConfig.SigningMethod,
}
