package repositories

import (
	"context"

	"github.com/salamanderman234/outsourcing-api/domains"
	"gorm.io/gorm"
)

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) domains.ServiceCategoryRepository {
	return &categoryRepository{
		db: db,
	}
}

func (cr *categoryRepository) Create(c context.Context, data domains.Model, repo ...*gorm.DB) (any, error) {
	db := cr.db
	if len(repo) == 1 {
		db = repo[0]
	}
	dataModel, ok := data.(domains.CategoryModel)
	if !ok {
		return nil, domains.ErrRepositoryInterfaceConversion
	}
	err := basicCreateRepoFunc(c, db, &dataModel, &dataModel)
	return dataModel, err
}
func (cr *categoryRepository) FindByID(c context.Context, id uint) (domains.Model, error) {
	var result domains.CategoryModel
	err := basicFindRepoFunc(c, cr.db, &result, id, &result)
	return result, err
}
func (cr *categoryRepository) Get(c context.Context, id uint, q string, page uint, orderBy string, desc bool) (any, uint, error) {
	var results []domains.CategoryModel
	var model domains.CategoryModel
	var count int64
	query := cr.db.Scopes(usingContextScope(c), usingModelScope(&model), orderScope(&model, orderBy, desc))
	if id != 0 {
		result, err := cr.FindByID(c, id)
		return result, 0, err
	}
	searchQuery := *query.Where("category_name LIKE ?", "%"+q+"%").
		Or("description LIKE ?", "%"+q+"%")
	result := searchQuery.Scopes(paginateScope(page)).Find(&results)
	if len(results) == 0 {
		return nil, 0, domains.ErrRecordNotFound
	}
	_ = *query.Where("category_name LIKE ?", "%"+q+"%").
		Or("description LIKE ?", "%"+q+"%").Count(&count)
	maxPage := getMaxPage(uint(count))
	return results, maxPage, convertRepoError(result)
}
func (cr *categoryRepository) Update(c context.Context, id uint, data domains.Model) (int64, any, error) {
	dataModel, ok := data.(domains.CategoryModel)
	if !ok {
		return 0, nil, domains.ErrRepositoryInterfaceConversion
	}
	aff, err := basicUpdateRepoFunc(c, cr.db, &dataModel, id, &dataModel)
	return aff, dataModel, err
}
func (cr *categoryRepository) Delete(c context.Context, id uint) (int64, int64, error) {
	aff, err := basicDeleteRepoFunc(c, cr.db, &domains.CategoryModel{}, id)
	return int64(id), aff, err
}
