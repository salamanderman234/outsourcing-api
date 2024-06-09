package services

import (
	"context"
	"fmt"
	"strconv"
	"time"

	service_domains "github.com/salamanderman234/outsourcing-api/app/domains/services"
	"github.com/salamanderman234/outsourcing-api/app/forms"
	"github.com/salamanderman234/outsourcing-api/app/helpers"
	"github.com/salamanderman234/outsourcing-api/app/jobs"
	"github.com/salamanderman234/outsourcing-api/app/mails"
	"github.com/salamanderman234/outsourcing-api/app/models"
	"github.com/salamanderman234/outsourcing-api/app/policies"
	"github.com/salamanderman234/outsourcing-api/app/providers"
	"github.com/salamanderman234/outsourcing-api/app/types"
	"github.com/salamanderman234/outsourcing-api/app/types/enums"
	"github.com/salamanderman234/outsourcing-api/configs"
	"golang.org/x/crypto/bcrypt"
)

type authService struct{}

func NewAuthService() service_domains.AuthServiceInterface {
	return &authService{}
}
func (authService) Login(ctx context.Context, creds forms.LoginForm) (models.User, string, error) {
	// validate form
	if err := helpers.Validator.Validate(creds); err != nil {
		return models.User{}, "", err
	}
	conds := map[string]any{
		"email": creds.Email,
	}
	// calling repo
	var user models.User
	err := providers.RepoProvider.BaseRepo.FindWhere(ctx,
		conds,
		&user,
		"AdminProfile",
		"SupervisorProfile",
		"EmployeeProfile",
		"ServiceUserProfile",
	)
	if err != nil {
		return models.User{}, "", types.ErrNotMatched
	}
	// check password validity
	err = bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(creds.Password))
	if err != nil {
		return models.User{}, "", err
	}
	// create auth token
	tkn, err := helpers.JWT.CreateToken(user, enums.AuthenticationTokenType)
	if err != nil {
		return models.User{}, "", err
	}
	helpers.Logger.Info(
		fmt.Sprintf("(Auth) User %s is logged in", *user.Email),
	)
	return user, tkn, nil
}
func (authService) RegisterUser(
	ctx context.Context,
	creds forms.UserRegisterForm,
	role enums.UserRolesEnum,
) (models.User, string, error) {
	claims, _ := ctx.Value(configs.VarConfig.UserContextName).(types.JWTCLaims)
	policy := policies.UserPolicy{}.Create(claims)
	if !policy {
		return models.User{}, "", types.ErrForbiden
	}

	if role == enums.SuperAdminUserRole && (claims.Role != string(enums.SuperAdminRole)) {
		return models.User{}, "", types.ErrForbiden
	} else if role == enums.AdminUserRole && (claims.Role != string(enums.SuperAdminRole)) {
		return models.User{}, "", types.ErrForbiden
	} else if (role == enums.EmployeeUserRole || role == enums.SupervisorUserRole) && ((claims.Role != string(enums.SuperAdminRole)) && (claims.Role != string(enums.AdminUserRole))) {
		return models.User{}, "", types.ErrForbiden
	} else if role == enums.ApplicationUserRole && (claims.Role != string(enums.SuperAdminRole)) {
		return models.User{}, "", types.ErrForbiden
	}

	if err := helpers.Validator.Validate(creds); err != nil {
		return models.User{}, "", err
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(creds.Password), 1)
	hashedPasswordStr := string(hashedPassword)
	creds.Password = hashedPasswordStr

	User := models.User{}
	err := helpers.Translator.TranslateStruct(creds, &User)
	if err != nil {
		return models.User{}, "", err
	}
	roleString := string(role)
	User.Role = &roleString
	switch roleString {
	case string(enums.AdminUserRole):
		if User.AdminProfile == nil {
			return User, "", types.ErrBadRequest
		}
		User.SuperAdminProfile = nil
		User.EmployeeProfile = nil
		User.ServiceUserProfile = nil
		User.SupervisorProfile = nil
	case string(enums.EmployeeUserRole):
		if User.EmployeeProfile == nil {
			return User, "", types.ErrBadRequest.SetCustomMsg(
				"missing employee_profile field",
			)
		}
		User.AdminProfile = nil
		User.SuperAdminProfile = nil
		User.ServiceUserProfile = nil
		User.SupervisorProfile = nil
	case string(enums.SupervisorUserRole):
		if User.SupervisorProfile == nil {
			return User, "", types.ErrBadRequest.SetCustomMsg(
				"missing supervisor_profile field",
			)
		}
		User.AdminProfile = nil
		User.SuperAdminProfile = nil
		User.EmployeeProfile = nil
		User.ServiceUserProfile = nil
	case string(enums.ServiceUserRole):
		if User.ServiceUserProfile == nil {
			return User, "", types.ErrBadRequest.SetCustomMsg(
				"missing service_user_profile field",
			)
		}
		User.AdminProfile = nil
		User.SuperAdminProfile = nil
		User.EmployeeProfile = nil
		User.SupervisorProfile = nil
	case string(enums.SuperAdminUserRole):
		if User.SuperAdminProfile == nil {
			return User, "", types.ErrBadRequest.SetCustomMsg(
				"missing super_admin_profile field",
			)
		}
		User.AdminProfile = nil
		User.EmployeeProfile = nil
		User.ServiceUserProfile = nil
		User.SupervisorProfile = nil
	}
	data, err := providers.RepoProvider.UserRepo.RegisterUser(ctx, User)
	if err != nil {
		return models.User{}, "", err
	}
	token, err := helpers.JWT.CreateToken(data, enums.AuthenticationTokenType)
	if err != nil {
		return models.User{}, "", err
	}
	mail := mails.NewVerifyAccountMail(map[string]any{
		"id": data.ID,
	})
	job := jobs.NewSendMailJob(mail, []string{*data.Email})
	err = helpers.JobManager.DispatchNow(job)
	if err != nil {
		return models.User{}, "", err
	}
	helpers.Logger.Info(
		fmt.Sprintf("(Auth) User %s is successfully registered", *data.Email),
	)
	return data, token, nil
}
func (authService) ForgotPassword(ctx context.Context, creds forms.ChangePasswordForm) error {
	if err := helpers.Validator.Validate(creds); err != nil {
		return err
	}
	to := []string{creds.Email}
	var user models.User
	providers.RepoProvider.BaseRepo.FindWhere(
		ctx,
		map[string]any{
			"email": creds.Email,
		},
		&user,
		"AdminProfile",
		"SupervisorProfile",
		"EmployeeProfile",
		"ServiceUserProfile",
	)
	id := user.ID
	token, err := helpers.JWT.CreateToken(user, enums.ResetPasswordTokenType, 12)
	if err != nil {
		return err
	}
	mail := mails.NewResetPasswordMail(map[string]any{
		"user_id": id,
		"token":   token,
	})
	job := jobs.NewSendMailJob(mail, to)
	helpers.JobManager.DispatchNow(job)
	return nil
}
func (authService) ResetPassword(ctx context.Context, creds forms.ResetPasswordForm) error {
	if err := helpers.Validator.Validate(creds); err != nil {
		return err
	}
	claims, err := helpers.JWT.VerifyToken(creds.ResetToken)
	if err != nil {
		return err
	}
	if claims.Subject != string(enums.ResetPasswordTokenType) {
		return types.ErrForbiden.SetCustomMsg(
			"invalid reset token",
		)
	}
	usrIDStr := strconv.Itoa(int(creds.UserID))
	if claims.ID != usrIDStr {
		return types.ErrForbiden.SetCustomMsg(
			"invalid reset token",
		)
	}
	id, _ := strconv.Atoi(claims.ID)
	byteHashedNewPassword, _ := bcrypt.GenerateFromPassword([]byte(creds.NewPassword), 1)
	hashedNewPassword := string(byteHashedNewPassword)
	err = providers.RepoProvider.BaseRepo.Update(
		ctx,
		[]uint{uint(id)},
		&models.User{Password: &hashedNewPassword},
	)
	return err
}
func (authService) VerifyEmail(ctx context.Context, creds forms.VerifyUserForm) error {
	if err := helpers.Validator.Validate(creds); err != nil {
		return err
	}
	now := time.Now()
	providers.RepoProvider.BaseRepo.Update(
		ctx,
		[]uint{creds.UserID},
		&models.User{VerifiedAt: &now},
	)
	return nil
}

func (authService) SendVerifyEmail(ctx context.Context, id uint) error {
	var user models.User
	err := providers.RepoProvider.BaseRepo.Find(ctx, id, &user)
	if err != nil {
		return nil
	}
	email := *user.Email
	mail := mails.NewVerifyAccountMail(map[string]any{
		"id": user.ID,
	})
	job := jobs.NewSendMailJob(mail, []string{email})
	helpers.JobManager.DispatchNow(job)
	return nil
}
