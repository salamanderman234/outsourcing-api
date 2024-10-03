package repository_domains

import (
	"context"

	"github.com/salamanderman234/outsourcing-api/app/domains"
	"github.com/salamanderman234/outsourcing-api/app/types"
	"gorm.io/gorm"
)

type BaseRepositoryInterface interface {
	Create(ctx context.Context, datas any, conn ...*gorm.DB) error
	ReadAll(
		ctx context.Context,
		result any,
		config types.DBSearchParams,
	) (*types.Pagination, error)
	Find(ctx context.Context, id uint, cont domains.ModelInterface, preloads ...string) error
	FindWhere(ctx context.Context, conds map[string]any, cont domains.ModelInterface, preloads ...string) error
	Update(ctx context.Context, ids []uint, data domains.ModelInterface, conn ...*gorm.DB) error
	Delete(ctx context.Context, ids []uint, model domains.ModelInterface, conn ...*gorm.DB) error
}

type DoEachFunc func([]domains.ModelInterface, error) error
type DoFunc func(conn *gorm.DB) ([]domains.ModelInterface, error)
