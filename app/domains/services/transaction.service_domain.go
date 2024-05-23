package service_domains

import (
	"context"

	"github.com/salamanderman234/outsourcing-api/app/forms"
	"github.com/salamanderman234/outsourcing-api/app/models"
	"github.com/salamanderman234/outsourcing-api/app/types"
)

type TransactionServiceInterface interface {
	Create(ctx context.Context, data forms.TransactionCreateForm) (models.Transaction, error)
	Update(ctx context.Context, id uint, data forms.TransactionUpdateForm) (uint, models.Transaction, error)
	Read(ctx context.Context, q string, page uint) ([]models.Transaction, *types.Pagination, error)
	Find(ctx context.Context, id uint) (models.Transaction, error)
	UploadMOU(ctx context.Context, id uint, file string) error
	Delete(ctx context.Context, id uint) (uint, error)
}
