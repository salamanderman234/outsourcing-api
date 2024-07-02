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

type MasterAccessServiceInterface interface {
	// Create(ctx context.Context, data forms.MasterAccessCreateForm) (models.Access)

}

type MasterQuestionServiceInterface interface {
	Create(ctx context.Context, data forms.MasterQuestionCreateForm) (models.Question, error)
	AssignQuestion(ctx context.Context, data forms.MasterQuestionAssignForm) error
	UnassignQuestion(ctx context.Context, data forms.MasterQuestionUnassignForm) error
	Read(ctx context.Context, categoryID uint, q string, page uint) ([]models.Question, *types.Pagination, error)
	Find(ctx context.Context, id uint) (models.Question, error)
	Update(ctx context.Context, id uint, data forms.MasterQuestionUpdateForm) (uint, models.Question, error)
	Delete(ctx context.Context, id uint) (uint, error)
}

type MasterPaymentConfigServiceInterface interface {
	SetDPPercentage(ctx context.Context, amount uint) error
	Set3TerminFirst(ctx context.Context, amount uint) error
	Set3TerminSecond(ctx context.Context, amount uint) error
	GetConfigs(ctx context.Context) ([]models.PaymentConfig, error)
}
