package services

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/salamanderman234/outsourcing-api/configs"
	"github.com/salamanderman234/outsourcing-api/domains"
	"github.com/salamanderman234/outsourcing-api/helpers"
)

func generatePairToken(id uint, username string, role string, profilePic string, remember bool) (domains.TokenPair, error) {
	var accessExpire *jwt.NumericDate
	var refreshExpire *jwt.NumericDate
	result := domains.TokenPair{}

	accessExpire = helpers.GenerateExpireTime(configs.ACCESS_TOKEN_EXPIRE_TIME)
	if !remember {
		refreshExpire = helpers.GenerateExpireTime(configs.REFRESH_TOKEN_EXPIRE_TIME)
	}

	accessClaims := helpers.CreateJWTClaims(id, &username, &role, &profilePic, accessExpire)
	refreshClaims := helpers.CreateJWTClaims(id, nil, nil, nil, refreshExpire)

	accessToken, err := helpers.GenerateToken(accessClaims)
	if err != nil {
		return result, err
	}
	refreshToken, err := helpers.GenerateToken(refreshClaims)
	if err != nil {
		return result, err
	}
	result.Access = accessToken
	result.Refresh = refreshToken
	return result, nil
}
