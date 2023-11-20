package helpers

import (
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/salamanderman234/outsourcing-api/configs"
	"github.com/salamanderman234/outsourcing-api/domains"
)

func GenerateExpireTime(add float32) *jwt.NumericDate {
	return jwt.NewNumericDate(time.Now().Add(time.Duration(add) * time.Hour))
}

func GenerateToken(claims domains.JWTClaims) (string, error) {
	secret := []byte(configs.GetApplicationSecret())
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(secret)
	if err != nil {
		return "", domains.ErrGenerateToken
	}
	return signed, nil
}

func CreateJWTClaims(id uint, username *string, role *string, profilePic *string, expiresAt *jwt.NumericDate) domains.JWTClaims {
	return domains.JWTClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        strconv.Itoa(int(id)),
			ExpiresAt: expiresAt,
		},
		JWTPayload: domains.JWTPayload{
			Username:   username,
			Role:       role,
			ProfilePic: profilePic,
		},
	}
}

func VerifyToken(token string) (domains.JWTClaims, error) {
	claims := domains.JWTClaims{}
	tkn, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (any, error) {
		return []byte(configs.GetApplicationSecret()), nil
	})
	if errors.Is(err, jwt.ErrTokenExpired) {
		return claims, domains.ErrExpiredToken
	}
	if err != nil || !tkn.Valid {
		return claims, domains.ErrInvalidToken
	}
	return claims, nil
}
