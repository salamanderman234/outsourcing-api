package service_domains

import (
	"context"

	"github.com/salamanderman234/outsourcing-api/app/forms"
	"github.com/salamanderman234/outsourcing-api/app/models"
	"github.com/salamanderman234/outsourcing-api/app/types"
	"github.com/salamanderman234/outsourcing-api/app/types/enums"
)

type UserServiceInterface interface {
	GetUsers(ctx context.Context, q string)
	Read(ctx context.Context, q string, role enums.UserRolesEnum, page uint) ([]models.User, *types.Pagination, error)
	Find(ctx context.Context, id uint) (models.User, error)
	Update(ctx context.Context, id uint, data forms.MasterProviceUpdateForm) (uint, models.User, error)
	Delete(ctx context.Context, id uint) (uint, error)
}
