package repositories

import (
	"context"
	"errors"

	"github.com/asaskevich/govalidator"
	"github.com/salamanderman234/outsourcing-api/domains"
	"gorm.io/gorm"
)

type serviceOrderPlacementEmployee struct {
	db    *gorm.DB
	model domains.ServiceOrderPlacementServiceEmployeeModel
}

func NewServiceOrderPlacementEmployee(db *gorm.DB) domains.ServiceOrderPlacementServiceEmployeeRepository {
	return &serviceOrderPlacementEmployee{
		db:    db,
		model: domains.ServiceOrderPlacementServiceEmployeeModel{},
	}
}

func (pe *serviceOrderPlacementEmployee) Create(c context.Context, data domains.ServiceOrderPlacementServiceEmployeeModel, repo ...*gorm.DB) (domains.ServiceOrderPlacementServiceEmployeeModel, error) {
	db := pe.db
	if len(repo) == 1 {
		db = repo[0]
	}
	err := basicCreateRepoFunc(c, db, &pe.model, &data)
	if errors.Is(err, domains.ErrForeignKeyViolated) {
		conv := err.(domains.GeneralError)
		conv.ValidationErrors = govalidator.Errors{
			domains.DatabaseKeyError{
				Field: "employee_id",
				Msg:   "employee does not exists",
			},
			domains.DatabaseKeyError{
				Field: "service_order_placement_service_id",
				Msg:   "service order placement service does not exists",
			},
		}
		return data, conv
	}
	return data, err
}

//	func (pe *serviceOrderPlacementEmployee) Read(c context.Context, placementId uint, page uint, orderBy string, desc bool, withPagination bool) ([]domains.ServiceOrderPlacementServiceEmployeeModel, uint, error) {
//		var results []domains.ServiceOrderPlacementServiceEmployeeModel
//		callFunc := func(db *gorm.DB) *gorm.DB {
//			query := db
//			if placementId != 0 {
//				query = query.Where("service_order_placement_id = ?", placementId)
//			}
//			return query.Preload("Employee")
//		}
//		maxPage, err := basicReadFunc(
//			c,
//			&results,
//			pe.db,
//			callFunc,
//			page,
//			orderBy,
//			desc,
//			withPagination,
//			&pe.model,
//		)
//		if err != nil {
//			return results, 0, err
//		}
//		if len(results) == 0 {
//			return nil, 0, domains.ErrRecordNotFound
//		}
//		return results, maxPage, nil
//	}
func (pe *serviceOrderPlacementEmployee) Find(c context.Context, id uint) (domains.ServiceOrderPlacementServiceEmployeeModel, error) {
	var result domains.ServiceOrderPlacementServiceEmployeeModel
	err := basicFindRepoFunc(c, pe.db, &pe.model, id, &result, "Employee", "Schedules")
	return result, err
}
func (pe *serviceOrderPlacementEmployee) Update(c context.Context, id uint, data domains.ServiceOrderPlacementServiceEmployeeModel, repo ...*gorm.DB) (int, domains.ServiceOrderPlacementServiceEmployeeModel, error) {
	db := pe.db
	if len(repo) == 1 {
		db = repo[0]
	}
	aff, err := basicUpdateRepoFunc(c, db, &pe.model, id, &data)
	if errors.Is(err, domains.ErrForeignKeyViolated) {
		conv := err.(domains.GeneralError)
		conv.ValidationErrors = govalidator.Errors{
			domains.DatabaseKeyError{
				Field: "employee_id",
				Msg:   "employee does not exists",
			},
			domains.DatabaseKeyError{
				Field: "service_order_placement_id",
				Msg:   "service order placement does not exists",
			},
		}
		return 0, data, conv
	}
	return int(aff), data, err
}
func (pe *serviceOrderPlacementEmployee) Delete(c context.Context, id uint, repo ...*gorm.DB) (int, int, error) {
	db := pe.db
	if len(repo) == 1 {
		db = repo[0]
	}
	aff, err := basicDeleteRepoFunc(c, db, &pe.model, id)
	return int(id), int(aff), err
}
