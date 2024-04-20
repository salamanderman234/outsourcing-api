package repositories

import (
	"context"
	"math"

	"github.com/salamanderman234/outsourcing-api/app/domains"
	repository_domains "github.com/salamanderman234/outsourcing-api/app/domains/repositories"
	database_types "github.com/salamanderman234/outsourcing-api/app/domains/types/databases"
	"github.com/salamanderman234/outsourcing-api/app/domains/types/responses"
	"github.com/salamanderman234/outsourcing-api/app/helpers"
	"github.com/salamanderman234/outsourcing-api/app/models"
	"github.com/salamanderman234/outsourcing-api/configs"
	"gorm.io/gorm"
)

type baseRepo struct{}

func NewBaseRepo() repository_domains.BaseRepositoryInterface {
	return &baseRepo{}
}

func (baseRepo) Create(ctx context.Context, datas any, conn ...*gorm.DB) error {
	db := domains.Connection
	if len(conn) == 1 {
		db = conn[0]
	}
	result := db.WithContext(ctx).CreateInBatches(datas, 100)
	return result.Error
}
func (baseRepo) ReadAll(
	ctx context.Context,
	result any,
	config database_types.DBSearchConfig,
) (*responses.Pagination, error) {
	db := domains.Connection.WithContext(ctx)
	paginateQuery := *domains.Connection.WithContext(ctx)
	var pagination *responses.Pagination
	for _, param := range config.Params {
		val := config.Query
		if param.Operator == "LIKE" {
			val = "%" + val + "%"
		}
		db = db.Where(param.Field+" "+param.Operator+" ?", val)
		paginateQuery = *paginateQuery.Where(param.Field+" "+param.Operator+" ?", val)
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
	finalResult := db.Find(result)
	if finalResult.RowsAffected == 0 {
		return pagination, gorm.ErrRecordNotFound
	}
	return pagination, finalResult.Error
}

func (baseRepo) Find(ctx context.Context, id uint, cont models.ModelInterface, preloads ...string) error {
	db := domains.Connection
	result := db.WithContext(ctx).Model(cont).Where("id", id)
	for _, preload := range preloads {
		result = result.Preload(preload)
	}
	result = result.First(cont)
	return result.Error
}

func (baseRepo) FindWhere(ctx context.Context, conds map[string]any, cont models.ModelInterface, preloads ...string) error {
	db := domains.Connection
	result := db.WithContext(ctx).Model(cont).Where(conds)
	for _, preload := range preloads {
		result = result.Preload(preload)
	}
	result = result.First(cont)
	return result.Error
}

func (baseRepo) Update(ctx context.Context, ids []uint, data models.ModelInterface, conn ...*gorm.DB) error {
	db := domains.Connection
	if len(conn) == 1 {
		db = conn[0]
	}
	result := db.Model(&data).
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
	ids []uint, model models.ModelInterface,
	conn ...*gorm.DB,
) error {
	db := domains.Connection
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
