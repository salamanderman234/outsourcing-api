package repositories

import (
	"context"
	"errors"

	"github.com/salamanderman234/outsourcing-api/domains"
	"gorm.io/gorm"
)

type partialServiceRepository struct {
	db    *gorm.DB
	model domains.ServiceModel
}

func NewPartialServiceRepository(db *gorm.DB) domains.PartialServiceRepository {
	return &partialServiceRepository{
		db:    db,
		model: domains.ServiceModel{},
	}
}
func (ps *partialServiceRepository) Create(ctx context.Context, data domains.ServiceModel, repo ...*gorm.DB) (domains.ServiceModel, error) {
	db := ps.db
	if len(repo) == 1 {
		db = repo[0]
	}
	err := basicCreateRepoFunc(ctx, db, &ps.model, &data)
	if errors.Is(err, domains.ErrForeignKeyViolated) {
		conv := err.(domains.GeneralError)
		conv.DatabaseError = domains.DatabaseKeyError{
			Field: "category_id",
			Msg:   "this category id does not exists",
		}
		return data, conv
	}
	return data, err
}
func (ps *partialServiceRepository) Read(ctx context.Context, q string, page uint, orderBy string, desc bool, withPagination bool) ([]domains.ServiceModel, uint, error) {
	var results []domains.ServiceModel
	callFunc := func(db *gorm.DB) *gorm.DB {
		return db.Where("service_name LIKE ?", "%"+q+"%").
			Or("description LIKE ?", "%"+q+"%").
			Preload("Category").
			Preload("ServiceItems")
	}
	maxPage, err := basicReadFunc(
		ctx,
		&results,
		ps.db,
		callFunc,
		page,
		orderBy,
		desc,
		withPagination,
		&ps.model,
	)
	if err != nil {
		return results, 0, err
	}
	if len(results) == 0 {
		return nil, 0, domains.ErrRecordNotFound
	}
	return results, maxPage, nil
}
func (ps *partialServiceRepository) Find(ctx context.Context, id uint) (domains.ServiceModel, error) {
	var result domains.ServiceModel
	err := basicFindRepoFunc(ctx, ps.db, &ps.model, id, &result)
	return result, err
}
func (ps *partialServiceRepository) Update(ctx context.Context, id uint, data domains.ServiceModel, repo ...*gorm.DB) (int, domains.ServiceModel, error) {
	db := ps.db
	if len(repo) == 1 {
		db = repo[0]
	}
	aff, err := basicUpdateRepoFunc(ctx, db, &ps.model, id, &data)
	if errors.Is(err, domains.ErrForeignKeyViolated) {
		conv := err.(domains.GeneralError)
		conv.DatabaseError = domains.DatabaseKeyError{
			Field: "category_id",
			Msg:   "this category id does not exists",
		}
		return 0, data, conv
	}
	return int(aff), data, err
}
func (ps *partialServiceRepository) Delete(ctx context.Context, id uint, repo ...*gorm.DB) (int, int, error) {
	db := ps.db
	if len(repo) == 1 {
		db = repo[0]
	}
	aff, err := basicDeleteRepoFunc(ctx, db, &ps.model, id)
	return int(id), int(aff), err
}
