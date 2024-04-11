package service_domains

import (
	"context"

	"github.com/salamanderman234/outsourcing-api/app/models"
)

type ProfileServiceInterface interface {
	GetProfile(ctx context.Context, id uint, email string) (models.User, error)
	UpdateProfile(ctx context.Context, id uint, changes any) (models.User, error)
}

type UserServiceInterface interface {
	GetUsers(ctx context.Context, q string)
}
