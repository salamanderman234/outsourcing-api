package repositories

import (
	"context"

	"github.com/salamanderman234/outsourcing-api/domains"
	"gorm.io/gorm"
)

type categoryRepository struct {
	db    *gorm.DB
	model domains.CategoryModel
}

func NewCategoryRepository(db *gorm.DB) domains.CategoryRepository {
	return &categoryRepository{
		db:    db,
		model: domains.CategoryModel{},
	}
}

func (cr *categoryRepository) Create(c context.Context, data domains.CategoryModel, repo ...*gorm.DB) (domains.CategoryModel, error) {
	db := cr.db
	if len(repo) == 1 {
		db = repo[0]
	}
	err := basicCreateRepoFunc(c, db, cr.model, &data)
	return data, err
}
func (cr *categoryRepository) Find(c context.Context, id uint) (domains.CategoryModel, error) {
	var result domains.CategoryModel
	err := basicFindRepoFunc(c, cr.db, &cr.model, id, &result, "Employees", "PartialServices")
	return result, err
}
func (cr *categoryRepository) Read(c context.Context, q string, page uint, orderBy string, desc bool, withPagination bool) ([]domains.CategoryModel, uint, error) {
	var results []domains.CategoryModel
	callFunc := func(db *gorm.DB) *gorm.DB {
		return db.Where("category_name LIKE ?", "%"+q+"%").
			Or("description LIKE ?", "%"+q+"%")
	}
	maxPage, err := basicReadFunc(
		c,
		&results,
		cr.db,
		callFunc,
		page,
		orderBy,
		desc,
		withPagination,
		&cr.model,
	)
	if err != nil {
		return results, 0, err
	}
	if len(results) == 0 {
		return nil, 0, domains.ErrRecordNotFound
	}
	return results, maxPage, nil
}
func (cr *categoryRepository) Update(c context.Context, id uint, data domains.CategoryModel, repo ...*gorm.DB) (int64, domains.CategoryModel, error) {
	db := cr.db
	if len(repo) == 1 {
		db = repo[0]
	}
	aff, err := basicUpdateRepoFunc(c, db, &cr.model, id, &data)
	return aff, data, err
}
func (cr *categoryRepository) Delete(c context.Context, id uint, repo ...*gorm.DB) (int64, int64, error) {
	db := cr.db
	if len(repo) == 1 {
		db = repo[0]
	}
	aff, err := basicDeleteRepoFunc(c, db, &cr.model, id)
	return int64(id), aff, err
}
