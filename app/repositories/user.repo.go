package repositories

import (
	"context"

	"github.com/salamanderman234/outsourcing-api/app/domains"
	repository_domains "github.com/salamanderman234/outsourcing-api/app/domains/repositories"
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
		users := []models.User{
			data,
		}
		err := domains.RepoRegistry.BaseRepo.Create(ctx, users, tx)
		if err != nil {
			tx.Rollback()
			return err
		}
		user = users[0]
		return nil
	})
	return user, err
}
