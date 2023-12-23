package repositories

import (
	"context"

	"github.com/salamanderman234/outsourcing-api/domains"
	"gorm.io/gorm"
)

type serviceItemRepository struct {
	db    *gorm.DB
	model domains.ServiceItemModel
}

func NewServiceItemRepository(db *gorm.DB) domains.ServiceItemRepository {
	return &serviceItemRepository{
		db:    db,
		model: domains.ServiceItemModel{},
	}
}

func (cr *serviceItemRepository) Create(c context.Context, data domains.ServiceItemModel, repo ...*gorm.DB) (domains.ServiceItemModel, error) {
	db := cr.db
	if len(repo) == 1 {
		db = repo[0]
	}
	err := basicCreateRepoFunc(c, db, &cr.model, &data)
	return data, err
}
func (cr *serviceItemRepository) Find(c context.Context, id uint) (domains.ServiceItemModel, error) {
	var result domains.ServiceItemModel
	err := basicFindRepoFunc(c, cr.db, &cr.model, id, &result)
	return result, err
}
func (cr *serviceItemRepository) Read(c context.Context, q string, page uint, orderBy string, desc bool, withPagination bool) ([]domains.ServiceItemModel, uint, error) {
	var results []domains.ServiceItemModel
	callFunc := func(db *gorm.DB) *gorm.DB {
		return db.Where("service_items.item_name LIKE ?", "%"+q+"%").
			Or("service_items.description LIKE ?", "%"+q+"%").
			Preload("Service")
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
func (cr *serviceItemRepository) Update(c context.Context, id uint, data domains.ServiceItemModel, repo ...*gorm.DB) (int64, domains.ServiceItemModel, error) {
	db := cr.db
	if len(repo) == 1 {
		db = repo[0]
	}
	aff, err := basicUpdateRepoFunc(c, db, &cr.model, id, &data)
	return aff, data, err
}
func (cr *serviceItemRepository) Delete(c context.Context, id uint, repo ...*gorm.DB) (int64, int64, error) {
	db := cr.db
	if len(repo) == 1 {
		db = repo[0]
	}
	aff, err := basicDeleteRepoFunc(c, db, &cr.model, id)
	return int64(id), aff, err
}
