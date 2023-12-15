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

func (s *adminRepository) Create(c context.Context, data domains.Model, repo ...*gorm.DB) (any, error) {
	db := s.db
	if len(repo) == 1 {
		db = repo[0]
	}
	user, ok := data.(domains.AdminModel)
	if !ok {
		return nil, domains.ErrRepositoryInterfaceConversion
	}
	err := basicCreateRepoFunc(c, db, &s.model, &user)
	return user, err
}
func (s *adminRepository) FindByID(c context.Context, id uint) (domains.Model, error) {
	var user domains.AdminModel
	err := basicFindRepoFunc(c, s.db, &s.model, id, &user, "User")
	return user, err
}
func (s *adminRepository) Update(c context.Context, id uint, data domains.Model) (int64, any, error) {
	dataModel, ok := data.(domains.AdminModel)
	if !ok {
		return 0, nil, domains.ErrRepositoryInterfaceConversion
	}
	aff, err := basicUpdateRepoFunc(c, s.db, &s.model, id, &dataModel)
	return aff, dataModel, err
}
func (s *adminRepository) Delete(c context.Context, id uint) (int64, int64, error) {
	aff, err := basicDeleteRepoFunc(c, s.db, &s.model, id)
	return int64(id), aff, err
}
func (s *adminRepository) Get(c context.Context, id uint, q string, page uint, orderBy string, desc bool) (any, uint, error) {
	var users []domains.AdminModel
	var count int64
	query := s.db.Scopes(usingContextScope(c), usingModelScope(&s.model), orderScope(&s.model, orderBy, desc))
	if id != 0 {
		result := query.Scopes(whereIdEqualScope(id)).Find(&users)
		return users, 1, convertRepoError(result)
	}
	searchQuery := query.Scopes(paginateScope(page)).
		Where("fullname LIKE ?", "%"+q+"%").
		Preload("User")
	_ = *searchQuery.Count(&count)
	maxPage := getMaxPage(uint(count))
	result := searchQuery.Find(&users)
	return users, maxPage, convertRepoError(result)
}
