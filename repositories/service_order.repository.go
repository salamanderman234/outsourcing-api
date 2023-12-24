package repositories

import (
	"context"

	"github.com/salamanderman234/outsourcing-api/domains"
	"gorm.io/gorm"
)

type serviceOrderRepository struct {
	db    *gorm.DB
	model domains.ServiceOrderModel
}

func NewServiceOrderRepository(db *gorm.DB) domains.ServiceOrderRepository {
	return &serviceOrderRepository{
		db:    db,
		model: domains.ServiceOrderModel{},
	}
}

func (s serviceOrderRepository) Create(c context.Context, data domains.ServiceOrderModel, repo ...*gorm.DB) (domains.ServiceOrderModel, error) {
	err := s.db.Transaction(func(tx *gorm.DB) error {
		result := tx.Scopes(usingContextScope(c), usingModelScope(s.model)).
			Create(&data)
		if result != nil {
			tx.Rollback()
			return convertRepoError(result)
		}
		details := []domains.ServiceOrderDetailModel{}
		for _, detail := range data.ServiceOrderDetails {
			detailResult, err := domains.RepoRegistry.
				ServiceOrderDetailRepo.
				Create(c, detail, tx)
			if err != nil {
				tx.Rollback()
				return err
			}
			details = append(details, detailResult)
		}
		data.ServiceOrderDetails = details
		return nil
	})
	return data, err
}
func (s serviceOrderRepository) Read(c context.Context, status string, service_user_id uint, page uint, orderBy string, desc bool, withPagination bool) ([]domains.ServiceOrderModel, uint, error) {
	var results []domains.ServiceOrderModel
	var count int64
	q := s.db.Scopes(
		usingContextScope(c),
		usingModelScope(&s.model),
		orderScope(&s.model, orderBy, desc),
	)
	if status != "" {
		q = q.Where("status = ?", status)
	}
	if service_user_id != 0 {
		q = q.Where("service_user_id")
	}
	if withPagination {
		q = q.Scopes(paginateScope(page))
	}
	result := q.
		Preload("ServicePackage").
		Preload("ServiceUser").
		Preload("ServiceOrderDetails").
		Find(results)
	if withPagination {
		_ = q.Count(&count)
	}
	maxPage := getMaxPage(uint(count))
	return results, maxPage, convertRepoError(result)
}
func (s serviceOrderRepository) Find(c context.Context, id uint) (domains.ServiceOrderModel, error) {
	var result domains.ServiceOrderModel
	err := basicFindRepoFunc(c, s.db, &s.model, id, &result, "ServicePackage", "ServiceUser", "ServiceOrderDetails")
	return result, err
}
func (s serviceOrderRepository) Update(c context.Context, id uint, data domains.ServiceOrderModel, repo ...*gorm.DB) (int64, domains.ServiceOrderModel, error) {
	db := s.db
	if len(repo) == 1 {
		db = repo[0]
	}
	aff, err := basicUpdateRepoFunc(c, db, &s.model, id, &data)
	return aff, data, err
}
func (s serviceOrderRepository) Delete(c context.Context, id uint, repo ...*gorm.DB) (uint, int64, error) {
	db := s.db
	if len(repo) == 1 {
		db = repo[0]
	}
	aff, err := basicDeleteRepoFunc(c, db, &s.model, id)
	return id, aff, err
}
