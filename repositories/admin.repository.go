package repositories

import (
	"context"

	"github.com/salamanderman234/outsourcing-api/domains"
	"gorm.io/gorm"
)

type adminRepository struct {
	db    *gorm.DB
	model domains.AdminModel
}

func NewAdminRepository(db *gorm.DB) domains.AdminRepository {
	return &adminRepository{
		db:    db,
		model: domains.AdminModel{},
	}
}

func (s *adminRepository) Create(c context.Context, data domains.AdminModel, repo ...*gorm.DB) (domains.AdminModel, error) {
	db := s.db
	if len(repo) == 1 {
		db = repo[0]
	}
	err := basicCreateRepoFunc(c, db, &s.model, &data)
	return data, err
}
func (s *adminRepository) Find(c context.Context, id uint) (domains.AdminModel, error) {
	var result domains.AdminModel
	err := basicFindRepoFunc(c, s.db, &s.model, id, &result, "User")
	return result, err
}
func (s *adminRepository) Update(c context.Context, id uint, data domains.AdminModel, repo ...*gorm.DB) (int64, domains.AdminModel, error) {
	db := s.db
	if len(repo) == 1 {
		db = repo[0]
	}
	aff, err := basicUpdateRepoFunc(c, db, &s.model, id, &data)
	return aff, data, err
}
func (s *adminRepository) Delete(c context.Context, id uint, repo ...*gorm.DB) (int64, int64, error) {
	db := s.db
	if len(repo) == 1 {
		db = repo[0]
	}
	aff, err := basicDeleteRepoFunc(c, db, &s.model, id)
	return int64(id), aff, err
}
func (s *adminRepository) Read(c context.Context, q string, page uint, orderBy string, desc bool, withPagination bool) ([]domains.AdminModel, uint, error) {
	var admins []domains.AdminModel
	callFunc := func(db *gorm.DB) *gorm.DB {
		return db.
			Where("fullname LIKE ?", "%"+q+"%").
			Preload("User")
	}
	maxPage, err := basicReadFunc(
		c,
		&admins,
		s.db,
		callFunc,
		page,
		orderBy,
		desc,
		withPagination,
		&s.model,
	)
	if err != nil {
		return admins, 0, err
	}
	if len(admins) == 0 {
		return nil, 0, domains.ErrRecordNotFound
	}
	return admins, maxPage, nil
}
