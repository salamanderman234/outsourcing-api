package service_domains

import (
	"context"

	"github.com/salamanderman234/outsourcing-api/app/forms"
	"github.com/salamanderman234/outsourcing-api/app/models"
	"github.com/salamanderman234/outsourcing-api/app/types"
	"github.com/salamanderman234/outsourcing-api/app/types/enums"
)

type UserServiceInterface interface {
	Read(ctx context.Context, q string, regencyID uint, role enums.UserRolesEnum, page uint) ([]models.User, *types.Pagination, error)
	Find(ctx context.Context, id uint) (models.User, error)
	Update(ctx context.Context, id uint, data forms.UserUpdateForm) (uint, models.User, error)
}
