package repositories

import (
	"context"

	"github.com/salamanderman234/outsourcing-api/domains"
	"gorm.io/gorm"
)

type userRepository struct {
	db    *gorm.DB
	model domains.UserModel
}

func NewUserRepository(db *gorm.DB) domains.UserRepository {
	return &userRepository{
		db:    db,
		model: domains.UserModel{},
	}
}

func (s *userRepository) GetUserWithCreds(c context.Context, username string) (any, error) {
	var user domains.UserModel
	usernameField := "email"
	return user, basicCredsSearch(c, s.db, usernameField, username, &user)
}
func (s *userRepository) RegisterUser(c context.Context, authData any, profileData any) (int64, any, error) {
	user, ok := authData.(domains.UserModel)
	if !ok {
		return 0, nil, domains.ErrRepositoryInterfaceConversion
	}
	result, err := s.Create(c, user)
	_, errProfile := s.CreateProfile(c, user.Role, profileData)
	if errProfile != nil {
		return 0, nil, err
	}
	return 1, result, err
}
func (s *userRepository) CreateProfile(c context.Context, role string, data any) (any, error) {
	if role == string(domains.AdminRole) {
		return domains.RepoRegistry.AdminRepo.Create(c, data)
	}
	if role == string(domains.EmployeeRole) {
		return domains.RepoRegistry.EmployeeRepo.Create(c, data)
	}
	if role == string(domains.SupervisorRole) {
		return domains.RepoRegistry.SupervisorRepo.Create(c, data)
	}
	return domains.RepoRegistry.ServiceUserRepo.Create(c, data)
}
func (s *userRepository) Create(c context.Context, data any) (any, error) {
	user, ok := data.(domains.UserModel)
	if !ok {
		return nil, domains.ErrRepositoryInterfaceConversion
	}
	err := basicCreateRepoFunc(c, s.db, &s.model, &user)
	return user, err
}
func (s *userRepository) FindByID(c context.Context, id uint) (any, error) {
	var user domains.UserModel
	err := basicFindRepoFunc(c, s.db, &s.model, id, &user)
	return user, err
}
func (s *userRepository) Update(c context.Context, id uint, data any) (int64, any, error) {
	dataModel, ok := data.(domains.UserModel)
	if !ok {
		return 0, nil, domains.ErrRepositoryInterfaceConversion
	}
	aff, err := basicUpdateRepoFunc(c, s.db, &s.model, id, &dataModel)
	return aff, dataModel, err
}
func (s *userRepository) Delete(c context.Context, id uint) (int64, int64, error) {
	aff, err := basicDeleteRepoFunc(c, s.db, &s.model, id)
	return int64(id), aff, err
}
func (s *userRepository) Get(c context.Context, id uint, q string, page uint, orderBy string, desc bool) (any, uint, error) {
	var users []domains.UserModel
	var count int64
	query := s.db.Scopes(usingContextScope(c), usingModelScope(&s.model), orderScope(&s.model, orderBy, desc))
	if id != 0 {
		result := query.Scopes(whereIdEqualScope(id)).Find(&users)
		return users, 1, convertRepoError(result)
	}
	searchQuery := query.Scopes(paginateScope(page)).
		Where("email LIKE ?", "%"+q+"%")
	_ = *searchQuery.Count(&count)
	maxPage := getMaxPage(uint(count))
	result := searchQuery.Find(&users)
	return users, maxPage, convertRepoError(result)
}
