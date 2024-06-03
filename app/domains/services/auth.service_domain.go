package service_domains

import (
	"context"

	"github.com/salamanderman234/outsourcing-api/app/forms"
	"github.com/salamanderman234/outsourcing-api/app/models"
	"github.com/salamanderman234/outsourcing-api/app/types/enums"
)

type AuthServiceInterface interface {
	Login(ctx context.Context, creds forms.LoginForm) (models.User, string, error)
	RegisterUser(ctx context.Context,
		creds forms.UserRegisterForm,
		role enums.UserRolesEnum,
	) (models.User, string, error)
	SendVerifyEmail(ctx context.Context, id uint) error
	ForgotPassword(ctx context.Context, forgotForm forms.ChangePasswordForm) error
	ResetPassword(ctx context.Context, resetForm forms.ResetPasswordForm) error
	VerifyEmail(ctx context.Context, verifyForm forms.VerifyUserForm) error
}
