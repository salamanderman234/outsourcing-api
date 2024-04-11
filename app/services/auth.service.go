package services

import (
	"context"
	"strconv"
	"time"

	"github.com/salamanderman234/outsourcing-api/app/domains"
	service_domains "github.com/salamanderman234/outsourcing-api/app/domains/services"
	auth_types "github.com/salamanderman234/outsourcing-api/app/domains/types/auth"
	"github.com/salamanderman234/outsourcing-api/app/domains/types/enums"
	custom_errors "github.com/salamanderman234/outsourcing-api/app/domains/types/errors"
	"github.com/salamanderman234/outsourcing-api/app/forms"
	"github.com/salamanderman234/outsourcing-api/app/helpers"
	"github.com/salamanderman234/outsourcing-api/app/mails"
	"github.com/salamanderman234/outsourcing-api/app/models"
	"github.com/salamanderman234/outsourcing-api/app/policies"
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
	err := domains.RepoRegistry.BaseRepo.FindWhere(ctx,
		conds,
		user,
		"AdminProfile",
		"SupervisorProfile",
		"EmployeeProfile",
		"ServiceUserProfile",
	)
	if err != nil {
		return models.User{}, "", custom_errors.ErrNotMatched
	}
	// check password validity
	err = bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(creds.Password))
	if err != nil {
		return models.User{}, "", err
	}
	// check verification status
	if user.VerifiedAt == nil {
		return models.User{}, "", custom_errors.ErrNotVerifiedUser
	}
	// create auth token
	tkn, err := helpers.JWT.CreateToken(user, enums.AuthenticationTokenType)
	if err != nil {
		return models.User{}, "", err
	}
	return user, tkn, nil
}
func (authService) RegisterUser(
	ctx context.Context,
	creds forms.UserRegisterForm,
	role enums.UserRolesEnum,
) (models.User, string, error) {
	claims, _ := ctx.Value(configs.VarConfig.UserContextName).(auth_types.JWTCLaims)
	if !policies.UserPolicy.RegisterUser(string(role), claims) {
		return models.User{}, "", custom_errors.ErrForbiden
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
	data, err := domains.RepoRegistry.UserRepo.RegisterUser(ctx, User)
	if err != nil {
		return models.User{}, "", err
	}
	token, err := helpers.JWT.CreateToken(data, enums.AuthenticationTokenType)
	if err != nil {
		return models.User{}, "", err
	}
	mail, err := mails.NewVerifyAccountMail(map[string]any{
		"id": data.ID,
	}).GetTemplate()
	if err != nil {
		return models.User{}, "", err
	}
	go helpers.Mailer.SendEmail([]string{*data.Email}, "Verify your account", mail)
	return data, token, nil
}
func (authService) ForgotPassword(ctx context.Context, creds forms.ChangePasswordForm) error {
	if err := helpers.Validator.Validate(creds); err != nil {
		return err
	}
	to := []string{creds.Email}
	token, err := helpers.JWT.CreateToken(models.User{
		Model: models.Model{
			ID: 0,
		},
		Email: &creds.Email,
	}, enums.ResetPasswordTokenType)
	if err != nil {
		return err
	}
	mail, err := mails.NewResetPasswordMail(map[string]any{
		"email": creds.Email,
		"token": token,
	}).GetTemplate()
	if err != nil {
		return err
	}
	go helpers.Mailer.SendEmail(to, "Change your password", mail)
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
		return custom_errors.ErrForbiden
	}
	if claims.Email != creds.Email {
		return custom_errors.ErrForbiden
	}
	id, _ := strconv.Atoi(claims.ID)
	byteHashedNewPassword, _ := bcrypt.GenerateFromPassword([]byte(creds.NewPassword), 1)
	hashedNewPassword := string(byteHashedNewPassword)
	err = domains.RepoRegistry.BaseRepo.Update(
		ctx,
		[]uint{uint(id)},
		models.User{Password: &hashedNewPassword},
	)
	return err
}
func (authService) VerifyEmail(ctx context.Context, creds forms.VerifyUserForm) error {
	if err := helpers.Validator.Validate(creds); err != nil {
		return err
	}
	now := time.Now()
	err := domains.RepoRegistry.BaseRepo.Update(
		ctx,
		[]uint{creds.UserID},
		models.User{VerifiedAt: &now},
	)
	return err
}
