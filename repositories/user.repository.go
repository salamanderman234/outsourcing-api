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
	aff := 0
	data := domains.UserWithProfileModel{}
	err := s.db.Transaction(func(tx *gorm.DB) error {
		user, ok := authData.(domains.UserModel)
		if !ok {
			return domains.ErrRepositoryInterfaceConversion
		}
		created, err := s.Create(c, user, tx)
		if err != nil {
			return err
		}
		profile, errProfile := s.CreateProfile(c, user.Role, profileData, created.(domains.UserModel).ID, tx)
		if errProfile != nil {
			return err
		}
		data.User = created.(domains.UserModel)
		data.Profile = profile
		aff = 2
		return nil
	})
	return int64(aff), data, err
}
func (s *userRepository) CreateProfile(c context.Context, role string, data any, userID uint, repo ...*gorm.DB) (any, error) {
	if role == string(domains.AdminRole) {
		convertedData, ok := data.(domains.AdminModel)
		if !ok {
			return nil, domains.ErrConversionType
		}
		convertedData.UserID = &userID
		if len(repo) == 1 {
			return domains.RepoRegistry.AdminRepo.Create(c, convertedData, repo[0])
		}
		return domains.RepoRegistry.AdminRepo.Create(c, convertedData)
	}
	if role == string(domains.EmployeeRole) {
		convertedData, ok := data.(domains.EmployeeModel)
		if !ok {
			return nil, domains.ErrConversionType
		}
		convertedData.UserID = &userID
		if len(repo) == 1 {
			return domains.RepoRegistry.EmployeeRepo.Create(c, convertedData, repo[0])
		}
		return domains.RepoRegistry.EmployeeRepo.Create(c, convertedData)
	}
	if role == string(domains.SupervisorRole) {
		convertedData, ok := data.(domains.SupervisorModel)
		if !ok {
			return nil, domains.ErrConversionType
		}
		convertedData.UserID = &userID
		if len(repo) == 1 {
			return domains.RepoRegistry.SupervisorRepo.Create(c, convertedData, repo[0])
		}
		return domains.RepoRegistry.SupervisorRepo.Create(c, convertedData)
	}
	convertedData, ok := data.(domains.ServiceUserModel)
	if !ok {
		return nil, domains.ErrConversionType
	}
	convertedData.UserID = &userID
	if len(repo) == 1 {
		return domains.RepoRegistry.ServiceUserRepo.Create(c, convertedData, repo[0])
	}
	return domains.RepoRegistry.ServiceUserRepo.Create(c, convertedData)
}
func (s *userRepository) Create(c context.Context, data domains.Model, repo ...*gorm.DB) (any, error) {
	db := s.db
	if len(repo) == 1 {
		db = repo[0]
	}
	user, ok := data.(domains.UserModel)
	if !ok {
		return nil, domains.ErrRepositoryInterfaceConversion
	}
	err := basicCreateRepoFunc(c, db, &s.model, &user)
	return user, err
}
func (s *userRepository) FindByID(c context.Context, id uint) (domains.Model, error) {
	var user domains.UserModel
	result := s.db.Model(&user).WithContext(c).Where("id = ?", id).
		Preload("Admin").
		Preload("ServiceUser").
		Preload("Employee").
		Preload("Supervisor").
		First(&user)
	if result.Error != nil {
		return user, convertRepoError(result)
	}
	return user, nil
}
func (s *userRepository) Update(c context.Context, id uint, data domains.Model) (int64, any, error) {
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
		result, err := s.FindByID(c, id)
		users = append(users, result.(domains.UserModel))
		return users, 1, err
	}
	searchQuery := query.Scopes(paginateScope(page)).
		Where("email LIKE ?", "%"+q+"%").
		Where("role = ?", q)
	_ = *searchQuery.Count(&count)
	maxPage := getMaxPage(uint(count))
	result := searchQuery.Find(&users)
	return users, maxPage, convertRepoError(result)
}
