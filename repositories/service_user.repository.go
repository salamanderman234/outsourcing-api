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

func (s *serviceUserRepository) Create(c context.Context, data any, repo ...*gorm.DB) (any, error) {
	db := s.db
	if len(repo) == 1 {
		db = repo[0]
	}
	user, ok := data.(domains.ServiceUserModel)
	if !ok {
		return nil, domains.ErrRepositoryInterfaceConversion
	}
	result := db.Scopes(usingContextScope(c), usingModelScope(&s.model)).Create(&user)
	return user, convertRepoError(result)
}
func (s *serviceUserRepository) FindByID(c context.Context, id uint) (any, error) {
	var user domains.ServiceUserModel
	result := s.db.Scopes(usingContextScope(c), usingModelScope(&s.model), whereIdEqualScope(id)).First(&user)
	return user, convertRepoError(result)
}
func (s *serviceUserRepository) Update(c context.Context, id uint, data any) (int64, any, error) {
	result := s.db.Scopes(usingContextScope(c), usingModelScope(&s.model), whereIdEqualScope(id)).Updates(data)
	return result.RowsAffected, data, convertRepoError(result)
}
func (s *serviceUserRepository) Delete(c context.Context, id uint) (int64, int64, error) {
	result := s.db.Scopes(usingContextScope(c), whereIdEqualScope(id)).Delete(&s.model)
	return result.RowsAffected, int64(id), convertRepoError(result)
}
func (s *serviceUserRepository) Get(c context.Context, id uint, q string, page uint, orderBy string, desc bool) (any, uint, error) {
	var users []domains.ServiceUserModel
	var count int64
	query := s.db.Scopes(usingContextScope(c), usingModelScope(&s.model), orderScope(&s.model, orderBy, desc))
	if id != 0 {
		result := query.Scopes(whereIdEqualScope(id)).Find(&users)
		return users, 1, convertRepoError(result)
	}
	searchQuery := query.Scopes(paginateScope(page)).
		Where("fullname LIKE ?", "%"+q+"%")
	_ = *searchQuery.Count(&count)
	maxPage := getMaxPage(uint(count))
	result := searchQuery.Find(&users)
	return users, maxPage, convertRepoError(result)
}
