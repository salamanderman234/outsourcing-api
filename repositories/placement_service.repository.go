package repositories

import (
	"context"
	"errors"

	"github.com/asaskevich/govalidator"
	"github.com/salamanderman234/outsourcing-api/domains"
	"gorm.io/gorm"
)

type serviceOrderPlacementServiceRepository struct {
	db    *gorm.DB
	model domains.ServiceOrderPlacementServiceModel
}

func NewServiceOrderPlacementServiceRepository(db *gorm.DB) domains.ServiceOrderPlacementServiceRepository {
	return &serviceOrderPlacementServiceRepository{
		db:    db,
		model: domains.ServiceOrderPlacementServiceModel{},
	}
}

func (pr *serviceOrderPlacementServiceRepository) BatchCreate(c context.Context, placementId uint, datas []domains.ServiceOrderPlacementServiceModel, repo ...*gorm.DB) ([]domains.ServiceOrderPlacementServiceModel, error) {
	db := pr.db
	if len(repo) == 1 {
		db = repo[0]
	}
	result := db.Scopes(usingContextScope(c), usingModelScope(&pr.model)).Create(&datas)
	if result.Error != nil {
		return datas, convertRepoError(result)
	}
	return datas, nil
}
func (pr *serviceOrderPlacementServiceRepository) Create(c context.Context, data domains.ServiceOrderPlacementServiceModel, repo ...*gorm.DB) (domains.ServiceOrderPlacementServiceModel, error) {
	db := pr.db
	if len(repo) == 1 {
		db = repo[0]
	}
	err := basicCreateRepoFunc(c, db, &pr.model, &data)
	if errors.Is(err, domains.ErrForeignKeyViolated) {
		conv := err.(domains.GeneralError)
		conv.ValidationErrors = govalidator.Errors{
			domains.DatabaseKeyError{
				Field: "service_order_placement_id",
				Msg:   "placement does not exists",
			},
			domains.DatabaseKeyError{
				Field: "partial_service_id",
				Msg:   "partial service does not exists",
			},
		}
		return data, conv
	}
	return data, err
}
func (pr *serviceOrderPlacementServiceRepository) Find(c context.Context, id uint) (domains.ServiceOrderPlacementServiceModel, error) {
	var result domains.ServiceOrderPlacementServiceModel
	err := basicFindRepoFunc(c, pr.db, &pr.model, id, &result, "Employees", "ServiceOrderPlacement", "PartialService")
	return result, err
}
func (pr *serviceOrderPlacementServiceRepository) Update(c context.Context, id uint, data domains.ServiceOrderPlacementServiceModel, repo ...*gorm.DB) (int, domains.ServiceOrderPlacementServiceModel, error) {
	db := pr.db
	if len(repo) == 1 {
		db = repo[0]
	}
	aff, err := basicUpdateRepoFunc(c, db, &pr.model, id, &data)
	if errors.Is(err, domains.ErrForeignKeyViolated) {
		conv := err.(domains.GeneralError)
		conv.ValidationErrors = govalidator.Errors{
			domains.DatabaseKeyError{
				Field: "service_order_placement_id",
				Msg:   "placement does not exists",
			},
			domains.DatabaseKeyError{
				Field: "partial_service_id",
				Msg:   "partial service does not exists",
			},
		}
		return 0, data, conv
	}
	return int(aff), data, err
}
func (pr *serviceOrderPlacementServiceRepository) Delete(c context.Context, id uint, repo ...*gorm.DB) (int, int, error) {
	db := pr.db
	if len(repo) == 1 {
		db = repo[0]
	}
	aff, err := basicDeleteRepoFunc(c, db, &pr.model, id)
	return int(id), int(aff), err
}
