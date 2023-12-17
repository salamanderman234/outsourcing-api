package repositories

import (
	"context"
	"errors"

	"github.com/salamanderman234/outsourcing-api/domains"
	"gorm.io/gorm"
)

type supervisorRepository struct {
	db    *gorm.DB
	model domains.SupervisorModel
}

func NewSupervisorRepository(db *gorm.DB) domains.SupervisorRepository {
	return &supervisorRepository{
		db:    db,
		model: domains.SupervisorModel{},
	}
}

func (s *supervisorRepository) Create(c context.Context, data domains.Model, repo ...*gorm.DB) (any, error) {
	db := s.db
	if len(repo) == 1 {
		db = repo[0]
	}
	user, ok := data.(domains.SupervisorModel)
	if !ok {
		return nil, domains.ErrRepositoryInterfaceConversion
	}
	err := basicCreateRepoFunc(c, db, &s.model, &user)
	if errors.Is(err, domains.ErrDuplicateEntries) {
		return user, domains.ErrCardIdDuplicate
	}
	return user, err
}
func (s *supervisorRepository) FindByID(c context.Context, id uint) (domains.Model, error) {
	var user domains.SupervisorModel
	err := basicFindRepoFunc(c, s.db, &s.model, id, &user, "User")
	return user, err
}
func (s *supervisorRepository) Update(c context.Context, id uint, data domains.Model) (int64, any, error) {
	dataModel, ok := data.(domains.SupervisorModel)
	if !ok {
		return 0, nil, domains.ErrRepositoryInterfaceConversion
	}
	aff, err := basicUpdateRepoFunc(c, s.db, &s.model, id, &dataModel)
	return aff, dataModel, err
}
func (s *supervisorRepository) Delete(c context.Context, id uint) (int64, int64, error) {
	aff, err := basicDeleteRepoFunc(c, s.db, &s.model, id)
	return int64(id), aff, err
}
func (s *supervisorRepository) Get(c context.Context, id uint, q string, page uint, orderBy string, desc bool) (any, uint, error) {
	var users []domains.SupervisorModel
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
