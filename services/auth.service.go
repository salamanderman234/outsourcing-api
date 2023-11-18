package services

import (
	"context"

	"github.com/salamanderman234/outsourcing-api/domains"
	"github.com/salamanderman234/outsourcing-api/helpers"
	"golang.org/x/crypto/bcrypt"
)

type serviceUserAuthService struct{}

func NewServiceUserAuthService() domains.ServiceUserAuthService {
	return &serviceUserAuthService{}
}

func (serviceUserAuthService) Login(c context.Context, loginForm domains.BasicLoginForm, remember bool) (domains.TokenPair, error) {
	tokenPair := domains.TokenPair{}
	if ok, err := helpers.Validate(loginForm); !ok {
		return tokenPair, err
	}
	password, err := bcrypt.GenerateFromPassword([]byte(loginForm.Password), 1)
	if err != nil {
		return tokenPair, domains.ErrHashingPassword
	}
	user, err := domains.RepoRegistry.ServiceUserRepo.GetUserWithCreds(c, loginForm.Username)
	if err != nil {
		return tokenPair, err
	}
	userModel, ok := user.(domains.ServiceUserModel)
	if !ok {
		return tokenPair, domains.ErrConversionType
	}
	err = bcrypt.CompareHashAndPassword(password, []byte(*userModel.Password))
	if err != nil {
		return tokenPair, domains.ErrInvalidCreds
	}
	tokenPair, err = generatePairToken(userModel.ID, *userModel.Email, "service_user", userModel.Profile, remember)
	if err != nil {
		return tokenPair, err
	}
	return tokenPair, nil
}
func (serviceUserAuthService) Register(c context.Context, data any, remember bool) (domains.TokenPair, error) {
	return domains.TokenPair{}, nil
}
func (serviceUserAuthService) Check(c context.Context, token string) (domains.JWTPayload, error) {
	return domains.JWTPayload{}, nil
}
func (serviceUserAuthService) Refresh(c context.Context, refreshToken string) (domains.TokenPair, error) {
	return domains.TokenPair{}, nil
}
