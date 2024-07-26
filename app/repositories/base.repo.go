package repositories

import (
	"context"
	"math"

	"github.com/salamanderman234/outsourcing-api/app/domains"
	repository_domains "github.com/salamanderman234/outsourcing-api/app/domains/repositories"
	"github.com/salamanderman234/outsourcing-api/app/helpers"
	"github.com/salamanderman234/outsourcing-api/app/providers"
	"github.com/salamanderman234/outsourcing-api/app/types"
	"github.com/salamanderman234/outsourcing-api/configs"
	"gorm.io/gorm"
)

type baseRepo struct{}

func NewBaseRepo() repository_domains.BaseRepositoryInterface {
	return &baseRepo{}
}

func (baseRepo) Create(ctx context.Context, datas any, conn ...*gorm.DB) error {
	db := providers.GetConnection()
	if len(conn) == 1 {
		db = conn[0]
	}
	result := db.WithContext(ctx).CreateInBatches(datas, 100)
	return result.Error
}
func (baseRepo) ReadAll(
	ctx context.Context,
	result any,
	config types.DBSearchParams,
) (*types.Pagination, error) {
	db := providers.GetConnection().WithContext(ctx).Model(config.Model)
	paginateQuery := *providers.GetConnection().WithContext(ctx)
	var pagination *types.Pagination
	for _, join := range config.Joins {
		db = db.Joins(join)
		paginateQuery = *paginateQuery.Joins(join)
	}
	for _, param := range config.Params {
		val := param.Str
		if param.Operator == "LIKE" {
			val = "%" + val + "%"
		}
		if param.IsOr {
			db = db.Or(param.Field+" "+param.Operator+" ?", val)
			paginateQuery = *paginateQuery.Or(param.Field+" "+param.Operator+" ?", val)
		} else {
			db = db.Where(param.Field+" "+param.Operator+" ?", val)
			paginateQuery = *paginateQuery.Where(param.Field+" "+param.Operator+" ?", val)
		}
	}
	for _, preload := range config.Preloads {
		db = db.Preload(preload)
	}
	var max int64
	if config.Page > 0 {
		paginateQuery = *paginateQuery.Model(config.Model).Count(&max)
		maxPage := math.Ceil(float64(max) / float64(configs.AppConfig.PaginationPerPage))
		pagination = helpers.Response.CreatePagination(int64(maxPage), config)
		db = db.Scopes(paginateScope(config.Page))
	} else if config.Limit > 0 {
		db = db.Limit(config.Limit)
	}
	finalResult := db.Order("created_at DESC").Find(result)
	if finalResult.RowsAffected == 0 {
		return pagination, gorm.ErrRecordNotFound
	}
	return pagination, finalResult.Error
}

func (baseRepo) Find(ctx context.Context, id uint, cont domains.ModelInterface, preloads ...string) error {
	db := providers.GetConnection()
	result := db.WithContext(ctx).Model(cont).Where("id", id)
	for _, preload := range preloads {
		result = result.Preload(preload)
	}
	result = result.First(cont)
	return result.Error
}

func (baseRepo) FindWhere(ctx context.Context, conds map[string]any, cont domains.ModelInterface, preloads ...string) error {
	db := providers.GetConnection()
	result := db.WithContext(ctx).Model(cont).Where(conds)
	for _, preload := range preloads {
		result = result.Preload(preload)
	}
	result = result.First(cont)
	return result.Error
}

func (baseRepo) Update(ctx context.Context, ids []uint, data domains.ModelInterface, conn ...*gorm.DB) error {
	db := providers.GetConnection()
	if len(conn) == 1 {
		db = conn[0]
	}
	result := db.Model(data).
		WithContext(ctx).
		Where("id IN ?", ids).
		Updates(data)
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return result.Error
}
func (baseRepo) Delete(
	ctx context.Context,
	ids []uint, model domains.ModelInterface,
	conn ...*gorm.DB,
) error {
	db := providers.GetConnection()
	if len(conn) == 1 {
		db = conn[0]
	}
	result := db.WithContext(ctx).
		Where("id IN ?", ids).
		Delete(model)
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return result.Error
}
