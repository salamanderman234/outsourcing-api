package service_domains

import (
	"context"

	"github.com/salamanderman234/outsourcing-api/app/domains/types/responses"
	"github.com/salamanderman234/outsourcing-api/app/forms"
	"github.com/salamanderman234/outsourcing-api/app/models"
)

type MasterProvinceServiceInterface interface {
	Create(ctx context.Context, data forms.MasterProviceCreateForm) (models.Province, error)
	Read(ctx context.Context, q string, page uint) ([]models.Province, *responses.Pagination, error)
	Find(ctx context.Context, id uint) (models.Province, error)
	Update(ctx context.Context, id uint, data forms.MasterProviceUpdateForm) (uint, models.Province, error)
	Delete(ctx context.Context, id uint) (uint, error)
}

type MasterRegencyServiceInterface interface {
	Create(ctx context.Context, data forms.MasterRegencyCreateForm) (models.Regency, error)
	Read(ctx context.Context, q string, page uint) ([]models.Regency, *responses.Pagination, error)
	Find(ctx context.Context, id uint) (models.Regency, error)
	Update(ctx context.Context, id uint, data forms.MasterRegencyUpdateForm) (uint, models.Regency, error)
	Delete(ctx context.Context, id uint) (uint, error)
}
