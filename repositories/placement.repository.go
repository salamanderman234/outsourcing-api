package repositories

import (
	"context"
	"errors"

	"github.com/asaskevich/govalidator"
	"github.com/salamanderman234/outsourcing-api/domains"
	"gorm.io/gorm"
)

type serviceOrderPlacementRepository struct {
	db    *gorm.DB
	model domains.ServiceOrderPlacementModel
}

func NewServiceOrderPlacementRepository(db *gorm.DB) domains.ServiceOrderPlacementRepository {
	return &serviceOrderPlacementRepository{
		db:    db,
		model: domains.ServiceOrderPlacementModel{},
	}
}
func (p *serviceOrderPlacementRepository) Create(c context.Context, data domains.ServiceOrderPlacementModel, repo ...*gorm.DB) (domains.ServiceOrderPlacementModel, error) {
	db := p.db
	if len(repo) == 1 {
		db = repo[0]
	}
	err := db.Transaction(func(tx *gorm.DB) error {
		result := tx.Scopes(usingContextScope(c)).Model(&p.model).Create(&data)
		if result.Error != nil {
			tx.Rollback()
			return convertRepoError(result)
		}
		placementServices := data.ServicePlacements
		createdPlacementServices, err := domains.RepoRegistry.
			PlacementServiceRepo.
			BatchCreate(c, data.ID, placementServices)
		if err != nil {
			return err
		}
		data.ServicePlacements = createdPlacementServices
		return nil
	})
	return data, err
}
func (p *serviceOrderPlacementRepository) Read(c context.Context, orderId uint, page uint, orderBy string, desc bool, withPagination bool) ([]domains.ServiceOrderPlacementModel, uint, error) {
	var results []domains.ServiceOrderPlacementModel
	callFunc := func(db *gorm.DB) *gorm.DB {
		query := db
		if orderId != 0 {
			query = query.Where("service_order_id = ?", orderId)
		}
		return query.Preload("Supervisor").
			Preload("ServiceOrder").
			Preload("ServicePlacements")
	}
	maxPage, err := basicReadFunc(
		c,
		&results,
		p.db,
		callFunc,
		page,
		orderBy,
		desc,
		withPagination,
		&p.model,
	)
	if err != nil {
		return results, 0, err
	}
	if len(results) == 0 {
		return nil, 0, domains.ErrRecordNotFound
	}
	return results, maxPage, nil
}
func (p *serviceOrderPlacementRepository) Find(c context.Context, id uint) (domains.ServiceOrderPlacementModel, error) {
	var result domains.ServiceOrderPlacementModel
	err := basicFindRepoFunc(c, p.db, &p.model, id, &result, "Supervisor", "ServiceOrder", "ServicePlacements", "Reports")
	return result, err
}
func (p *serviceOrderPlacementRepository) Update(c context.Context, id uint, data domains.ServiceOrderPlacementModel, repo ...*gorm.DB) (int, domains.ServiceOrderPlacementModel, error) {
	db := p.db
	if len(repo) == 1 {
		db = repo[0]
	}
	aff, err := basicUpdateRepoFunc(c, db, &p.model, id, &data)
	if errors.Is(err, domains.ErrForeignKeyViolated) {
		conv := err.(domains.GeneralError)
		conv.ValidationErrors = govalidator.Errors{
			domains.DatabaseKeyError{
				Field: "supervisor_id",
				Msg:   "supervisor does not exists",
			},
			domains.DatabaseKeyError{
				Field: "service_order_id",
				Msg:   "service order does not exists",
			},
		}
		return 0, data, conv
	}
	return int(aff), data, err
}
func (p *serviceOrderPlacementRepository) Delete(c context.Context, id uint, repo ...*gorm.DB) (int, int, error) {
	db := p.db
	if len(repo) == 1 {
		db = repo[0]
	}
	aff, err := basicDeleteRepoFunc(c, db, &p.model, id)
	return int(id), int(aff), err
}
