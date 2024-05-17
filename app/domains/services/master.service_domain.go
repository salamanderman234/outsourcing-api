package service_domains

import (
	"context"

	"github.com/salamanderman234/outsourcing-api/app/forms"
	"github.com/salamanderman234/outsourcing-api/app/models"
	"github.com/salamanderman234/outsourcing-api/app/types"
)

type MasterProvinceServiceInterface interface {
	Create(ctx context.Context, data forms.MasterProviceCreateForm) (models.Province, error)
	Read(ctx context.Context, q string, page uint) ([]models.Province, *types.Pagination, error)
	Find(ctx context.Context, id uint) (models.Province, error)
	Update(ctx context.Context, id uint, data forms.MasterProviceUpdateForm) (uint, models.Province, error)
	Delete(ctx context.Context, id uint) (uint, error)
}

type MasterRegencyServiceInterface interface {
	Create(ctx context.Context, data forms.MasterRegencyCreateForm) (models.Regency, error)
	Read(ctx context.Context, q string, page uint) ([]models.Regency, *types.Pagination, error)
	Find(ctx context.Context, id uint) (models.Regency, error)
	Update(ctx context.Context, id uint, data forms.MasterRegencyUpdateForm) (uint, models.Regency, error)
	Delete(ctx context.Context, id uint) (uint, error)
}

type MasterCategoryServiceInterface interface {
	Create(ctx context.Context, data forms.MasterCategoryCreateForm) (models.Category, error)
	Read(ctx context.Context, q string, page uint) ([]models.Category, *types.Pagination, error)
	Find(ctx context.Context, id uint) (models.Category, error)
	Update(ctx context.Context, id uint, data forms.MasterCategoryUpdateForm) (uint, models.Category, error)
	Delete(ctx context.Context, id uint) (uint, error)
}
