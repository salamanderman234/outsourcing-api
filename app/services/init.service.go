package services

import (
	"context"
	"time"

	"github.com/salamanderman234/outsourcing-api/app/models"
	"github.com/salamanderman234/outsourcing-api/app/providers"
	"github.com/salamanderman234/outsourcing-api/app/types"
	"github.com/salamanderman234/outsourcing-api/app/types/enums"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

func Init() {
	sups := []models.SuperAdmin{}
	providers.RepoProvider.BaseRepo.ReadAll(context.Background(), &sups, types.DBSearchParams{})
	if len(sups) == 0 {
		email := viper.GetString("INIT_USER")
		password := viper.GetString("INIT_PASSWORD")
		pass, err := bcrypt.GenerateFromPassword([]byte(password), 1)
		if err != nil {
			return
		}
		strPass := string(pass)
		now := time.Now()
		role := string(enums.SuperAdminRole)
		user := models.User{
			Email:      &email,
			Password:   &strPass,
			Role:       &role,
			VerifiedAt: &now,
		}
		err = providers.RepoProvider.BaseRepo.Create(context.Background(), &user)
		if err != nil {
			return
		}
		fullname := "Super Admin"
		profile := models.SuperAdmin{
			UserID:   &user.ID,
			Fullname: &fullname,
		}
		providers.RepoProvider.BaseRepo.Create(context.Background(), &profile)
	}
}
