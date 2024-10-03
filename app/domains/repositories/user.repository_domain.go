package repository_domains

import (
	"context"

	"github.com/salamanderman234/outsourcing-api/app/models"
)

type UserRepositoryInterface interface {
	RegisterUser(ctx context.Context, data models.User) (models.User, error)
	UpdateUser(ctx context.Context, id uint, data models.User) (models.User, error)
}
