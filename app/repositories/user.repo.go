package repositories

import (
	"context"

	repository_domains "github.com/salamanderman234/outsourcing-api/app/domains/repositories"
	"github.com/salamanderman234/outsourcing-api/app/models"
	"github.com/salamanderman234/outsourcing-api/app/providers"
	"github.com/salamanderman234/outsourcing-api/app/types/enums"
	"gorm.io/gorm"
)

type userRepo struct{}

func NewUserRepo() repository_domains.UserRepositoryInterface {
	return &userRepo{}
}

func (userRepo) UpdateUser(ctx context.Context, id uint, data models.User) (models.User, error) {
	db := providers.GetConnection()
	var user models.User
	err := db.Transaction(func(tx *gorm.DB) error {
		err := providers.RepoProvider.BaseRepo.Find(ctx, id, &user,
			"AdminProfile",
			"SupervisorProfile",
			"EmployeeProfile",
			"ServiceUserProfile",
			"SuperAdminProfile",
		)
		if err != nil {
			return err
		}
		role := *user.Role
		switch role {
		case string(enums.AdminUserRole):
			profile := data.AdminProfile
			if profile != nil {
				err = providers.RepoProvider.BaseRepo.Update(ctx, []uint{user.AdminProfile.ID}, profile, tx)
				if err != nil {
					tx.Rollback()
					return err
				}
			}
		case string(enums.EmployeeUserRole):
			profile := data.EmployeeProfile
			if profile != nil {
				err = providers.RepoProvider.BaseRepo.Update(ctx, []uint{user.EmployeeProfile.ID}, profile, tx)
				if err != nil {
					tx.Rollback()
					return err
				}
			}
		case string(enums.SupervisorUserRole):
			profile := data.SupervisorProfile
			if profile != nil {
				err = providers.RepoProvider.BaseRepo.Update(ctx, []uint{user.SuperAdminProfile.ID}, profile, tx)
				if err != nil {
					tx.Rollback()
					return err
				}
			}
		case string(enums.ServiceUserRole):
			profile := data.ServiceUserProfile
			if profile != nil {
				err = providers.RepoProvider.BaseRepo.Update(ctx, []uint{user.ServiceUserProfile.ID}, profile, tx)
				if err != nil {
					tx.Rollback()
					return err
				}
			}
		case string(enums.SuperAdminUserRole):
			profile := data.SuperAdminProfile
			if profile != nil {
				err = providers.RepoProvider.BaseRepo.Update(ctx, []uint{user.SuperAdminProfile.ID}, profile, tx)
				if err != nil {
					tx.Rollback()
					return err
				}
			}
		}

		data.AdminProfile = nil
		data.SuperAdminProfile = nil
		data.EmployeeProfile = nil
		data.ServiceUserProfile = nil
		data.SupervisorProfile = nil

		err = providers.RepoProvider.BaseRepo.Update(ctx, []uint{id}, &data, tx)
		if err != nil {
			tx.Rollback()
			return err
		}
		return nil
	})
	return user, err
}

func (userRepo) RegisterUser(ctx context.Context, data models.User) (models.User, error) {
	db := providers.GetConnection()
	var user models.User
	err := db.Transaction(func(tx *gorm.DB) error {
		users := []models.User{
			data,
		}
		err := providers.RepoProvider.BaseRepo.Create(ctx, users, tx)
		if err != nil {
			tx.Rollback()
			return err
		}
		user = users[0]
		return nil
	})
	return user, err
}
