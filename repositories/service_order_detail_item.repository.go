package repositories

import (
	"context"

	"github.com/salamanderman234/outsourcing-api/domains"
	"gorm.io/gorm"
)

type serviceOrderDetailItemRepository struct {
	db    *gorm.DB
	model domains.ServiceOrderDetailItemModel
}

func NewServiceOrderDetailItemRepository(db *gorm.DB) domains.ServiceOrderDetailItemRepository {
	return &serviceOrderDetailItemRepository{
		db:    db,
		model: domains.ServiceOrderDetailItemModel{},
	}
}
func (s serviceOrderDetailItemRepository) Create(c context.Context, data domains.ServiceOrderDetailItemModel, repo ...*gorm.DB) (domains.ServiceOrderDetailItemModel, error) {
	db := s.db
	if len(repo) == 1 {
		db = repo[0]
	}
	err := basicCreateRepoFunc(c, db, &s.model, &data)
	return data, err
}
func (s serviceOrderDetailItemRepository) Read(c context.Context, serviceOrderDetailId uint) ([]domains.ServiceOrderDetailItemModel, error) {
	var results []domains.ServiceOrderDetailItemModel
	result := s.db.Scopes(
		usingContextScope(c),
		usingModelScope(&s.model),
		orderScope(&s.model, "", true),
	).Where("service_order_detail_id = ?", serviceOrderDetailId).
		Preload("ServiceItem").
		Find(&results)
	return results, convertRepoError(result)
}
func (s serviceOrderDetailItemRepository) Find(c context.Context, id uint) (domains.ServiceOrderDetailItemModel, error) {
	var result domains.ServiceOrderDetailItemModel
	err := basicFindRepoFunc(c, s.db, &s.model, id, &result, "ServiceItem")
	return result, err
}
func (s serviceOrderDetailItemRepository) Update(c context.Context, id uint, data domains.ServiceOrderDetailItemModel, repo ...*gorm.DB) (int, domains.ServiceOrderDetailItemModel, error) {
	db := s.db
	if len(repo) == 1 {
		db = repo[0]
	}
	aff, err := basicUpdateRepoFunc(c, db, &s.model, id, &data)
	return int(aff), data, err
}
func (s serviceOrderDetailItemRepository) Delete(c context.Context, id uint, repo ...*gorm.DB) (uint, int64, error) {
	db := s.db
	if len(repo) == 1 {
		db = repo[0]
	}
	aff, err := basicDeleteRepoFunc(c, db, &s.model, id)
	return id, aff, err
}
