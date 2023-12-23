package repositories

import (
	"context"
	"errors"

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

func (s *userRepository) GetUserWithCreds(c context.Context, username string) (domains.UserModel, error) {
	var user domains.UserModel
	result := s.db.WithContext(c).Model(&s.model).
		Preload("ServiceUser").
		Preload("Admin").
		Preload("Employee").
		Preload("Supervisor").
		First(&user, "email = ?", username)
	return user, convertRepoError(result)
}
func (s *userRepository) RegisterUser(c context.Context, authData domains.UserModel, profileData domains.Model) (int64, domains.UserWithProfileModel, error) {
	aff := 0
	data := domains.UserWithProfileModel{}
	err := s.db.Transaction(func(tx *gorm.DB) error {
		user := authData
		created, err := s.Create(c, user, tx)
		if errors.Is(err, domains.ErrDuplicateEntries) {
			tx.Rollback()
			return err
		}
		if err != nil {
			tx.Rollback()
			return err
		}
		profile, errProfile := s.CreateProfile(c, user.Role, profileData, created.ID, tx)
		if errProfile != nil {
			tx.Rollback()
			return errProfile
		}
		data.User = created
		data.Profile = profile
		aff = 2
		return nil
	})
	return int64(aff), data, err
}
func (s *userRepository) CreateProfile(c context.Context, role string, data domains.Model, userID uint, repo ...*gorm.DB) (domains.Model, error) {
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
func (s *userRepository) Create(c context.Context, data domains.UserModel, repo ...*gorm.DB) (domains.UserModel, error) {
	db := s.db
	if len(repo) == 1 {
		db = repo[0]
	}
	err := basicCreateRepoFunc(c, db, &s.model, &data)
	if errors.Is(err, domains.ErrDuplicateEntries) {
		conv := err.(domains.GeneralError)
		conv.DatabaseError = domains.DatabaseKeyError{
			Field: "email",
			Msg:   "this email already exists",
		}
		return data, conv
	}
	return data, err
}
func (s *userRepository) Find(c context.Context, id uint) (domains.UserModel, error) {
	var user domains.UserModel
	err := basicFindRepoFunc(c, s.db, &s.model, id, &user, "Admin", "ServiceUser", "Employee", "Supervisor")
	return user, err
}
func (s *userRepository) Update(c context.Context, id uint, data domains.UserModel, repo ...*gorm.DB) (int64, domains.UserModel, error) {
	db := s.db
	if len(repo) == 1 {
		db = repo[0]
	}
	aff, err := basicUpdateRepoFunc(c, db, &s.model, id, &data)
	if errors.Is(err, domains.ErrDuplicateEntries) {
		conv := err.(domains.GeneralError)
		conv.DatabaseError = domains.DatabaseKeyError{
			Field: "email",
			Msg:   "this email already exists",
		}
		return aff, data, conv
	}
	return aff, data, err
}
func (s *userRepository) Delete(c context.Context, id uint, repo ...*gorm.DB) (int64, int64, error) {
	db := s.db
	if len(repo) == 1 {
		db = repo[0]
	}
	aff, err := basicDeleteRepoFunc(c, db, &s.model, id)
	return int64(id), aff, err
}
func (s *userRepository) Read(c context.Context, q string, page uint, orderBy string, desc bool, withPagination bool) ([]domains.UserModel, uint, error) {
	var results []domains.UserModel
	callFunc := func(db *gorm.DB) *gorm.DB {
		return db.Where("email LIKE ?", "%"+q+"%").
			Preload("ServiceUser").
			Preload("Admin").
			Preload("Employee").
			Preload("Supervisor")
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
