package repositories

import (
	"context"

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

func (s *employeeRepository) GetUserWithCreds(c context.Context, username string) (any, error) {
	var user domains.EmployeeModel
	usernameField := "email"
	return user, basicCredsSearch(c, s.db, usernameField, username, &user)
}
func (s *employeeRepository) RegisterUser(c context.Context, data any) (int64, any, error) {
	user, ok := data.(domains.EmployeeModel)
	if !ok {
		return 0, nil, domains.ErrRepositoryInterfaceConversion
	}
	result, err := s.Create(c, user)
	return 1, result, err
}
func (s *employeeRepository) Create(c context.Context, data any) (any, error) {
	user, ok := data.(domains.EmployeeModel)
	if !ok {
		return nil, domains.ErrRepositoryInterfaceConversion
	}
	err := basicCreateRepoFunc(c, s.db, &s.model, &user)
	return user, err
}
func (s *employeeRepository) FindByID(c context.Context, id uint) (any, error) {
	var user domains.EmployeeModel
	err := basicFindRepoFunc(c, s.db, &s.model, id, &user)
	return user, err
}
func (s *employeeRepository) Update(c context.Context, id uint, data any) (int64, any, error) {
	dataModel, ok := data.(domains.EmployeeModel)
	if !ok {
		return 0, nil, domains.ErrRepositoryInterfaceConversion
	}
	aff, err := basicUpdateRepoFunc(c, s.db, &s.model, id, &dataModel)
	return aff, dataModel, err
}
func (s *employeeRepository) Delete(c context.Context, id uint) (int64, int64, error) {
	aff, err := basicDeleteRepoFunc(c, s.db, &s.model, id)
	return int64(id), aff, err
}
func (s *employeeRepository) Get(c context.Context, id uint, q string, page uint, orderBy string, desc bool) ([]domains.EmployeeModel, uint, error) {
	var users []domains.EmployeeModel
	var count int64
	query := s.db.Scopes(usingContextScope(c), usingModelScope(&s.model), orderScope(&s.model, orderBy, desc))
	if id != 0 {
		result := query.Scopes(whereIdEqualScope(id)).Find(&users)
		return users, 1, convertRepoError(result)
	}
	searchQuery := query.Scopes(paginateScope(page)).
		Where("email LIKE ?", "%"+q+"%").
		Where("name LIKE ?", "%"+q+"%")
	_ = *searchQuery.Count(&count)
	maxPage := getMaxPage(uint(count))
	result := searchQuery.Find(&users)
	return users, maxPage, convertRepoError(result)
}
