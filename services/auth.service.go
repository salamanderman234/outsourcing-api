package services

import (
	"context"
	"errors"
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
	password := loginForm.Password
	user, err := domains.RepoRegistry.UserRepo.GetUserWithCreds(c, loginForm.Email)
	if err != nil {
		return tokenPair, err
	}
	userModel, ok := user.(domains.UserModel)
	if !ok {
		return tokenPair, domains.ErrConversionType
	}
	err = bcrypt.CompareHashAndPassword([]byte(*userModel.Password), []byte(password))

	if err != nil {
		return tokenPair, domains.ErrInvalidCreds
	}
	tokenPair, err = generatePairToken(userModel.ID, *userModel.Email, userModel.Role, userModel.Profile, remember)
	if err != nil {
		return tokenPair, err
	}
	return tokenPair, nil
}
func (s serviceUserAuthService) Register(c context.Context, authData domains.BasicRegisterForm, profileData any, role domains.RoleEnum, remember bool) (domains.TokenPair, error) {
	tokenPair := domains.TokenPair{}
	if ok, err := helpers.Validate(authData); !ok {
		return tokenPair, err
	}
	if ok, err := helpers.Validate(profileData); !ok {
		return tokenPair, err
	}
	var user domains.UserModel
	var profile any
	err := helpers.Convert(authData, &user)
	if err != nil {
		return tokenPair, err
	}
	if role == domains.AdminRole {
		var adminData domains.AdminModel
		err := helpers.Convert(profileData, &adminData)
		if err != nil {
			return tokenPair, err
		}
		profile = adminData
	} else if role == domains.EmployeeRole {
		var employeeData domains.EmployeeModel
		err := helpers.Convert(profileData, &employeeData)
		if err != nil {
			return tokenPair, err
		}
		profile = employeeData
	} else if role == domains.SupervisorRole {
		var supervisorData domains.SupervisorModel
		err := helpers.Convert(profileData, &supervisorData)
		if err != nil {
			return tokenPair, err
		}
		profile = supervisorData
	} else if role == domains.ServiceUserRole {
		var serviceUser domains.ServiceUserModel
		err := helpers.Convert(profileData, &serviceUser)
		if err != nil {
			return tokenPair, err
		}
		profile = serviceUser
	} else {
		return tokenPair, domains.ErrInvalidRole
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(authData.Password), 1)
	hashedString := string(hashed)
	user.Password = &hashedString
	user.Role = string(role)
	if err != nil {
		return tokenPair, domains.ErrHashingPassword
	}
	_, result, err := domains.RepoRegistry.UserRepo.RegisterUser(c, user, profile)
	if err != nil {
		return tokenPair, err
	}
	resultModel, ok := result.(domains.UserWithProfileModel)
	if !ok {
		return tokenPair, domains.ErrConversionType
	}
	userResult := resultModel.User
	tokenPair, err = generatePairToken(userResult.ID, *userResult.Email, userResult.Role, userResult.Profile, remember)
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
	if payload.Username == nil {
		return domains.JWTPayload{}, domains.ErrInvalidToken
	}
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
		if errors.Is(err, domains.ErrRecordNotFound) {
			return tokenPair, domains.ErrInvalidToken
		}
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
