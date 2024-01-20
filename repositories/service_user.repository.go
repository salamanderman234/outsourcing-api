package repositories

import (
	"context"

	"github.com/salamanderman234/outsourcing-api/domains"
	"gorm.io/gorm"
)

type serviceUserRepository struct {
	db    *gorm.DB
	model domains.ServiceUserModel
}

func NewServiceUserRepository(db *gorm.DB) domains.ServiceUserRepository {
	return &serviceUserRepository{
		db:    db,
		model: domains.ServiceUserModel{},
	}
}

func (s *serviceUserRepository) Create(c context.Context, data domains.ServiceUserModel, repo ...*gorm.DB) (domains.ServiceUserModel, error) {
	db := s.db
	if len(repo) == 1 {
		db = repo[0]
	}
	err := basicCreateRepoFunc(c, db, &s.model, &data)
	return data, err
}
func (s *serviceUserRepository) Find(c context.Context, id uint) (domains.ServiceUserModel, error) {
	var result domains.ServiceUserModel
	err := basicFindRepoFunc(c, s.db, &s.model, id, &result)
	return result, err
}
func (s *serviceUserRepository) Update(c context.Context, id uint, data domains.ServiceUserModel, repo ...*gorm.DB) (int64, domains.ServiceUserModel, error) {
	db := s.db
	if len(repo) == 1 {
		db = repo[0]
	}
	aff, err := basicUpdateRepoFunc(c, db, &s.model, id, &data)
	return aff, data, err
}
func (s *serviceUserRepository) Delete(c context.Context, id uint, repo ...*gorm.DB) (int64, int64, error) {
	db := s.db
	if len(repo) == 1 {
		db = repo[0]
	}
	aff, err := basicDeleteRepoFunc(c, db, &s.model, id)
	return int64(id), aff, err
}
func (s *serviceUserRepository) Read(c context.Context, q string, page uint, orderBy string, desc bool, withPagination bool) ([]domains.ServiceUserModel, uint, error) {
	var results []domains.ServiceUserModel
	callFunc := func(db *gorm.DB) *gorm.DB {
		return db.Where("users.email LIKE ?", "%"+q+"%").
			Or("service_users.fullname LIKE ?", "%"+q+"%").
			Joins("JOIN users ON users.id = service_users.user_id").
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
