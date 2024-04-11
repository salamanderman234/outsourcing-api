package repositories

import (
	"context"

	"github.com/salamanderman234/outsourcing-api/app/domains"
	repository_domains "github.com/salamanderman234/outsourcing-api/app/domains/repositories"
	"github.com/salamanderman234/outsourcing-api/app/domains/types/enums"
	"github.com/salamanderman234/outsourcing-api/app/models"
	"gorm.io/gorm"
)

type userRepo struct{}

func NewUserRepo() repository_domains.UserRepositoryInterface {
	return &userRepo{}
}

func (userRepo) RegisterUser(ctx context.Context, data models.User) (models.User, error) {
	db := domains.Connection
	var user models.User
	err := db.Transaction(func(tx *gorm.DB) error {
		var profile models.ModelInterface
		users := []models.User{
			data,
		}
		err := domains.RepoRegistry.BaseRepo.Create(ctx, users, tx)
		if err != nil {
			tx.Rollback()
			return err
		}
		switch *data.Role {
		case string(enums.AdminUserRole):
			profile = data.AdminProfile
			user.AdminProfile = data.AdminProfile
		case string(enums.SupervisorUserRole):
			profile = data.SupervisorProfile
			user.SupervisorProfile = data.SupervisorProfile
		case string(enums.EmployeeUserRole):
			profile = data.EmployeeProfile
			user.EmployeeProfile = data.EmployeeProfile
		default:
			profile = data.ServiceUserProfile
			user.ServiceUserProfile = data.ServiceUserProfile
		}

		err = domains.RepoRegistry.BaseRepo.Create(ctx, []models.ModelInterface{
			profile,
		}, tx)
		if err != nil {
			tx.Rollback()
			return err
		}
		tx.Commit()
		user = users[0]
		return nil
	})
	return user, err
}
