package repositories

import (
	"context"

	"github.com/salamanderman234/outsourcing-api/domains"
	"gorm.io/gorm"
)

type serviceOrderDetailRepository struct {
	db    *gorm.DB
	model domains.ServiceOrderDetailModel
}

func NewServiceOrderDetailRepository(db *gorm.DB) domains.ServiceOrderDetailRepository {
	return &serviceOrderDetailRepository{
		db:    db,
		model: domains.ServiceOrderDetailModel{},
	}
}
func (s serviceOrderDetailRepository) Create(c context.Context, data domains.ServiceOrderDetailModel, repo ...*gorm.DB) (domains.ServiceOrderDetailModel, error) {
	db := s.db
	if len(repo) == 1 {
		db = repo[0]
	}
	err := db.Transaction(func(tx *gorm.DB) error {
		result := tx.Scopes(usingContextScope(c), usingModelScope(s.model)).
			Create(&data)
		if result.Error != nil {
			return convertRepoError(result)
		}
		orderDetailId := data.ID
		items := []domains.ServiceOrderDetailItemModel{}
		for _, item := range data.PartialServiceItems {
			item.ServiceOrderDetailID = &orderDetailId
			detailResult, err := domains.RepoRegistry.
				ServiceOrderDetailItemRepo.
				Create(c, item, tx)
			if err != nil {
				return err
			}
			items = append(items, detailResult)
		}
		data.PartialServiceItems = items
		return nil
	})
	return data, err
}
func (s serviceOrderDetailRepository) Read(c context.Context, serviceOrderId uint) ([]domains.ServiceOrderDetailModel, error) {
	var results []domains.ServiceOrderDetailModel
	result := s.db.Scopes(
		usingContextScope(c),
		usingModelScope(&s.model),
		orderScope(&s.model, "", true),
	).Where("service_order_id = ?", serviceOrderId).
		Preload("PartialServiceItems").
		Preload("PartialService").
		Find(&results)
	return results, convertRepoError(result)
}
func (s serviceOrderDetailRepository) Find(c context.Context, id uint) (domains.ServiceOrderDetailModel, error) {
	var result domains.ServiceOrderDetailModel
	err := basicFindRepoFunc(c, s.db, &s.model, id, &result, "PartialServiceItems", "PartialService")
	return result, err
}
func (s serviceOrderDetailRepository) Update(c context.Context, id uint, data domains.ServiceOrderDetailModel, repo ...*gorm.DB) (int, domains.ServiceOrderDetailModel, error) {
	db := s.db
	if len(repo) == 1 {
		db = repo[0]
	}
	aff, err := basicUpdateRepoFunc(c, db, &s.model, id, &data)
	return int(aff), data, err
}
func (s serviceOrderDetailRepository) Delete(c context.Context, id uint, repo ...*gorm.DB) (uint, int64, error) {
	db := s.db
	if len(repo) == 1 {
		db = repo[0]
	}
	aff, err := basicDeleteRepoFunc(c, db, &s.model, id)
	return id, aff, err
}
