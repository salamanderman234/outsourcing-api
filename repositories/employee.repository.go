package repositories

import (
	"context"
	"errors"
	"net/http"

	"github.com/salamanderman234/outsourcing-api/domains"
	"gorm.io/gorm"
)

type employeeRepository struct {
	db    *gorm.DB
	model domains.EmployeeModel
}

func NewEmployeeRepository(db *gorm.DB) domains.EmployeeRepository {
	return &employeeRepository{
		db:    db,
		model: domains.EmployeeModel{},
	}
}

func (s *employeeRepository) Create(c context.Context, data domains.EmployeeModel, repo ...*gorm.DB) (domains.EmployeeModel, error) {
	db := s.db
	if len(repo) == 1 {
		db = repo[0]
	}
	err := basicCreateRepoFunc(c, db, &s.model, &data)
	if errors.Is(err, domains.ErrDuplicateEntries) {
		err = domains.DatabaseKeyError{
			Field:  "identity_card_number",
			Msg:    "this identity card number already exists",
			Status: http.StatusConflict,
		}
		return data, err
	}
	return data, err
}
func (s *employeeRepository) Find(c context.Context, id uint) (domains.EmployeeModel, error) {
	var user domains.EmployeeModel
	err := basicFindRepoFunc(c, s.db, &s.model, id, &user, "User")
	return user, err
}
func (s *employeeRepository) Update(c context.Context, id uint, data domains.EmployeeModel, repo ...*gorm.DB) (int64, domains.EmployeeModel, error) {
	db := s.db
	if len(repo) == 1 {
		db = repo[0]
	}
	aff, err := basicUpdateRepoFunc(c, db, &s.model, id, &data)
	if errors.Is(err, domains.ErrDuplicateEntries) {
		err = domains.DatabaseKeyError{
			Field:  "identity_card_number",
			Msg:    "this identity card number already exists",
			Status: http.StatusConflict,
		}
		return aff, data, err
	}
	return aff, data, err
}
func (s *employeeRepository) Delete(c context.Context, id uint, repo ...*gorm.DB) (int64, int64, error) {
	db := s.db
	if len(repo) == 1 {
		db = repo[0]
	}
	aff, err := basicDeleteRepoFunc(c, db, &s.model, id)
	return int64(id), aff, err
}
func (s *employeeRepository) Read(c context.Context, q string, page uint, orderBy string, desc bool, withPagination bool) ([]domains.EmployeeModel, uint, error) {
	var results []domains.EmployeeModel
	callFunc := func(db *gorm.DB) *gorm.DB {
		return db.Where("fullname LIKE ?", "%"+q+"%").
			Preload("User")
	}
	maxPage, err := basicReadFunc(
		c,
		&results,
		s.db,
		callFunc,
		page,
		orderBy,
		desc,
		withPagination,
		&s.model,
	)
	if err != nil {
		return results, 0, err
	}
	if len(results) == 0 {
		return nil, 0, domains.ErrRecordNotFound
	}
	return results, maxPage, nil
}
