package service_domains

import (
	"context"

	"github.com/salamanderman234/outsourcing-api/app/forms"
	"github.com/salamanderman234/outsourcing-api/app/models"
	"github.com/salamanderman234/outsourcing-api/app/types"
)

type ApplicationServiceServiceInterface interface {
	Create(ctx context.Context, data forms.ServiceCreateForm) (models.Service, error)
	Read(ctx context.Context, q string, page uint) ([]models.Service, *types.Pagination, error)
	Find(ctx context.Context, id uint) (models.Service, error)
	Update(ctx context.Context, id uint, data forms.ServiceUpdateForm) (uint, models.Service, error)
	Delete(ctx context.Context, id uint) (uint, error)
}

type ApplicationPackageInterface interface {
	Create(ctx context.Context, data forms.PackageServiceCreateForm) (models.Package, error)
	Read(ctx context.Context, q string, page uint) ([]models.Package, *types.Pagination, error)
	Find(ctx context.Context, id uint) (models.Package, error)
	Update(ctx context.Context, id uint, data forms.PackageUpdateForm) (uint, models.Package, error)
	Delete(ctx context.Context, id uint) (uint, error)
}
