package services

import (
	"context"
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/salamanderman234/outsourcing-api/configs"
	"github.com/salamanderman234/outsourcing-api/domains"
	"github.com/salamanderman234/outsourcing-api/helpers"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
}

func NewAuthService() domains.BasicAuthService {
	return &authService{}
}

func (authService) Login(c context.Context, loginForm domains.BasicLoginForm, remember bool) (domains.TokenPair, domains.UserEntity, error) {
	tokenPair := domains.TokenPair{}
	var userWithProfile domains.UserEntity
	if ok, err := helpers.Validate(loginForm); !ok {
		return tokenPair, userWithProfile, err
	}
	password := loginForm.Password
	user, err := domains.RepoRegistry.UserRepo.GetUserWithCreds(c, loginForm.Email)
	if err != nil {
		errConv := domains.ErrInvalidCreds
		errConv.ValidationErrors = govalidator.Errors{
			govalidator.Error{
				Name:                     "email",
				Validator:                "invalid creds",
				CustomErrorMessageExists: true,
				Err:                      errors.New("incorrect email or password"),
			},
		}
		return tokenPair, userWithProfile, errConv
	}
	err = bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(password))
	if err != nil {
		errConv := domains.ErrInvalidCreds
		errConv.ValidationErrors = govalidator.Errors{
			govalidator.Error{
				Name:                     "email",
				Validator:                "invalid creds",
				CustomErrorMessageExists: true,
				Err:                      errors.New("incorrect email or password"),
			},
		}
		return tokenPair, userWithProfile, errConv
	}
	tokenPair, err = generatePairToken(user.ID, *user.Email, user.Role, user.Profile, remember)
	if err != nil {
		return tokenPair, userWithProfile, err
	}
	err = helpers.Convert(user, &userWithProfile)
	if err != nil {
		return tokenPair, userWithProfile, err
	}
	return tokenPair, userWithProfile, nil
}
func (authService) Register(c context.Context, authData domains.BasicRegisterForm, profileData any, role domains.RoleEnum, remember bool) (domains.TokenPair, domains.UserWithProfileEntity, error) {
	tokenPair := domains.TokenPair{}
	userContext, ok := c.Value(configs.UserKey).(domains.UserEntity)
	if (role == domains.EmployeeRole || role == domains.AdminRole || role == domains.SupervisorRole) && !ok {
		return tokenPair, domains.UserWithProfileEntity{}, domains.ErrInvalidAccess
	}
	if (role == domains.EmployeeRole || role == domains.AdminRole || role == domains.SupervisorRole) && userContext.Role != string(domains.AdminRole) {
		return tokenPair, domains.UserWithProfileEntity{}, domains.ErrInvalidAccess
	}
	var userWithProfile domains.UserWithProfileEntity
	if ok, err := helpers.Validate(profileData); !ok {
		return tokenPair, userWithProfile, err
	}
	var user domains.UserModel
	var profile domains.Model
	err := helpers.Convert(authData, &user)
	if err != nil {
		return tokenPair, userWithProfile, err
	}
	if role == domains.AdminRole {
		var adminData domains.AdminModel
		err := helpers.Convert(profileData, &adminData)
		if err != nil {
			return tokenPair, userWithProfile, err
		}
		profile = adminData
	} else if role == domains.EmployeeRole {
		var employeeData domains.EmployeeModel
		err := helpers.Convert(profileData, &employeeData)
		if err != nil {
			return tokenPair, userWithProfile, err
		}
		profile = employeeData
	} else if role == domains.SupervisorRole {
		var supervisorData domains.SupervisorModel
		err := helpers.Convert(profileData, &supervisorData)
		if err != nil {
			return tokenPair, userWithProfile, err
		}
		profile = supervisorData
	} else if role == domains.ServiceUserRole {
		var serviceUser domains.ServiceUserModel
		err := helpers.Convert(profileData, &serviceUser)
		if err != nil {
			return tokenPair, userWithProfile, err
		}
		profile = serviceUser
	} else {
		return tokenPair, userWithProfile, domains.ErrInvalidRole
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(authData.Password), 1)
	hashedString := string(hashed)
	user.Password = &hashedString
	user.Role = string(role)
	if err != nil {
		return tokenPair, userWithProfile, domains.ErrHashingPassword
	}
	user.Profile = helpers.GenerateAvatar(helpers.GenerateRandomString(5))
	user.JoinedDate = time.Now()
	_, result, err := domains.RepoRegistry.UserRepo.RegisterUser(c, user, profile)
	if err != nil {
		return tokenPair, userWithProfile, err
	}
	userResult := result.User
	profileResult := result.Profile
	tokenPair, err = generatePairToken(userResult.ID, *userResult.Email, userResult.Role, userResult.Profile, remember)
	if err != nil {
		return tokenPair, userWithProfile, err
	}
	err = helpers.Convert(userResult, &userWithProfile.User)
	if err != nil {
		return tokenPair, userWithProfile, err
	}
	userWithProfile.Profile = profileResult
	return tokenPair, userWithProfile, nil
}
func (authService) Check(c context.Context, token string) (domains.JWTPayload, error) {
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
func (authService) Refresh(c context.Context, refreshToken string) (domains.TokenPair, error) {
	tokenPair := domains.TokenPair{}
	claims, err := helpers.VerifyToken(refreshToken)
	if err != nil {
		return tokenPair, err
	}
	id := claims.JWTPayload.ID
	if id == 0 {
		return tokenPair, domains.ErrInvalidToken
	}
	user, err := domains.RepoRegistry.UserRepo.Find(c, uint(id))
	if err != nil {
		if errors.Is(err, domains.ErrRecordNotFound) {
			return tokenPair, domains.ErrInvalidToken
		}
		return tokenPair, err
	}
	expire := helpers.GenerateExpireTime(configs.ACCESS_TOKEN_EXPIRE_TIME)
	accessClaims := helpers.
		CreateJWTClaims(uint(user.ID), user.Email, &user.Role, &user.Profile, expire)
	access, err := helpers.GenerateToken(accessClaims)
	if err != nil {
		return domains.TokenPair{}, err
	}
	tokenPair.Access = access
	tokenPair.Refresh = refreshToken
	return tokenPair, nil
}

func (a authService) UpdateAuthProfile(c context.Context, id uint, password string, profilePic ...domains.EntityFileMap) (bool, error) {
	var dataModel domains.UserModel
	var userEntity domains.UserEntity
	fun := func(id uint) (int, domains.Model, error) {
		userModel, err := domains.RepoRegistry.UserRepo.Find(c, id)
		if err != nil {
			return 0, nil, err
		}
		oldProfile := userModel.Profile
		_, err = a.storeProfile(&dataModel, profilePic...)
		if err != nil {
			return 0, nil, err
		}
		aff, updated, err := domains.RepoRegistry.UserRepo.Update(c, id, dataModel)
		if err != nil {
			go domains.ServiceRegistry.FileServ.Destroy(dataModel.Profile)
		} else {
			go domains.ServiceRegistry.FileServ.Destroy(oldProfile)
		}
		return int(aff), updated, nil
	}
	aff, err := basicUpdateService(true, c, id, &domains.UserEntity{}, &dataModel, &userEntity, fun)
	return aff == 1, err
}

func (authService) storeProfile(model *domains.UserModel, files ...domains.EntityFileMap) (bool, error) {
	if len(files) == 1 {
		file := files[0]
		if file.Field == "profile" && file.File != nil {
			zippedFile := map[string]domains.FileWrapper{
				"profile": {
					Config: configs.IMAGE_FILE_CONFIG,
					File:   file.File,
					Field:  file.Field,
					Dest:   configs.FILE_DESTS["user/profile"],
				},
			}
			savedPaths, _, err := domains.ServiceRegistry.FileServ.BatchStore(zippedFile)
			if err != nil {
				return false, err
			}
			model.Profile = savedPaths["profile"]
			return true, nil
		}
	}
	model.Profile = ""
	return false, nil
}
