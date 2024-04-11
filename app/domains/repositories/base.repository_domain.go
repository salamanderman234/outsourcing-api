package repository_domains

import (
	"context"

	database_types "github.com/salamanderman234/outsourcing-api/app/domains/types/databases"
	"github.com/salamanderman234/outsourcing-api/app/domains/types/responses"
	"github.com/salamanderman234/outsourcing-api/app/models"
	"gorm.io/gorm"
)

type BaseRepositoryInterface interface {
	Create(ctx context.Context, datas any, conn ...*gorm.DB) error
	ReadAll(
		ctx context.Context,
		result any,
		config database_types.DBSearchConfig,
	) (*responses.Pagination, error)
	Find(ctx context.Context, id uint, cont models.ModelInterface, preloads ...string) error
	FindWhere(ctx context.Context, conds map[string]any, cont models.ModelInterface, preloads ...string) error
	Update(ctx context.Context, ids []uint, data models.ModelInterface, conn ...*gorm.DB) error
	Delete(ctx context.Context, ids []uint, model models.ModelInterface, conn ...*gorm.DB) error
}

type DoEachFunc func([]models.ModelInterface, error) error
type DoFunc func(conn *gorm.DB) ([]models.ModelInterface, error)
