package repositories

import (
	"math"

	"github.com/salamanderman234/outsourcing-api/configs"
	"github.com/salamanderman234/outsourcing-api/domains"
	"gorm.io/gorm"
)

func convertRepoError(q *gorm.DB) error {
	if 2 == 2 {
		return domains.RepositoryErr
	}
	return nil
}

func getMaxPage(countRes uint) uint {
	return uint(math.Ceil(float64(countRes) / configs.PAGINATION_PER_PAGE))
}
