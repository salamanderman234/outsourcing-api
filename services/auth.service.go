package services

import (
	"context"
	"strconv"

	"github.com/salamanderman234/outsourcing-api/configs"
	"github.com/salamanderman234/outsourcing-api/domains"
	"github.com/salamanderman234/outsourcing-api/helpers"
	"golang.org/x/crypto/bcrypt"
)

type serviceUserAuthService struct {
}

func NewUserAuthService() domains.ServiceUserAuthService {
	return &serviceUserAuthService{}
}

func (s serviceUserAuthService) Login(c context.Context, loginForm domains.BasicLoginForm, remember bool) (domains.TokenPair, error) {
	tokenPair := domains.TokenPair{}
	if ok, err := helpers.Validate(loginForm); !ok {
		return tokenPair, err
	}
	password, err := bcrypt.GenerateFromPassword([]byte(loginForm.Password), 1)
	if err != nil {
		return tokenPair, domains.ErrHashingPassword
	}
	user, err := domains.RepoRegistry.UserRepo.GetUserWithCreds(c, loginForm.Email)
	if err != nil {
		return tokenPair, err
	}
	userModel, ok := user.(domains.UserModel)
	if !ok {
		return tokenPair, domains.ErrConversionType
	}
	err = bcrypt.CompareHashAndPassword(password, []byte(*userModel.Password))
	if err != nil {
		return tokenPair, domains.ErrInvalidCreds
	}
	tokenPair, err = generatePairToken(userModel.ID, *userModel.Email, userModel.Role, userModel.Profile, remember)
	if err != nil {
		return tokenPair, err
	}
	return tokenPair, nil
}
func (s serviceUserAuthService) Register(c context.Context, authData domains.BasicRegisterForm, profileData any, remember bool) (domains.TokenPair, error) {
	tokenPair := domains.TokenPair{}
	if ok, err := helpers.Validate(authData); !ok {
		return tokenPair, err
	}
	if ok, err := helpers.Validate(profileData); !ok {
		return tokenPair, err
	}
	var user domains.ServiceUserModel
	err := helpers.Convert(authData, &user)
	if err != nil {
		return tokenPair, err
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(authData.Password), 1)
	authData.Password = string(hashed)
	if err != nil {
		return tokenPair, domains.ErrHashingPassword
	}
	_, result, err := domains.RepoRegistry.UserRepo.RegisterUser(c, authData, profileData)
	if err != nil {
		return tokenPair, err
	}
	resultModel, ok := result.(domains.UserModel)
	if !ok {
		return tokenPair, domains.ErrConversionType
	}

	tokenPair, err = generatePairToken(resultModel.ID, *resultModel.Email, resultModel.Role, resultModel.Profile, remember)
	if err != nil {
		return tokenPair, err
	}
	return tokenPair, nil
}
func (serviceUserAuthService) Check(c context.Context, token string) (domains.JWTPayload, error) {
	claims, err := helpers.VerifyToken(token)
	if err != nil {
		return domains.JWTPayload{}, err
	}
	payload := claims.JWTPayload
	return payload, nil
}
func (s serviceUserAuthService) Refresh(c context.Context, refreshToken string) (domains.TokenPair, error) {
	tokenPair := domains.TokenPair{}
	claims, err := helpers.VerifyToken(refreshToken)
	if err != nil {
		return tokenPair, err
	}
	id, _ := strconv.Atoi(claims.ID)
	if id == 0 {
		return tokenPair, domains.ErrInvalidToken
	}
	user, err := domains.RepoRegistry.UserRepo.FindByID(c, uint(id))
	if err != nil {
		return tokenPair, err
	}
	userModel, ok := user.(domains.UserModel)
	if !ok {
		return tokenPair, domains.ErrInvalidToken
	}
	expire := helpers.GenerateExpireTime(configs.ACCESS_TOKEN_EXPIRE_TIME)
	accessClaims := helpers.
		CreateJWTClaims(uint(userModel.ID), userModel.Email, &userModel.Role, &userModel.Profile, expire)
	access, err := helpers.GenerateToken(accessClaims)
	if err != nil {
		return domains.TokenPair{}, err
	}
	tokenPair.Access = access
	tokenPair.Refresh = refreshToken
	return tokenPair, nil
}
