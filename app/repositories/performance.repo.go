package repositories

import (
	"context"
	"time"

	repository_domains "github.com/salamanderman234/outsourcing-api/app/domains/repositories"
	"github.com/salamanderman234/outsourcing-api/app/models"
	"github.com/salamanderman234/outsourcing-api/app/providers"
	"gorm.io/gorm"
)

type performanceRepo struct{}

func NewPerformanceRepo() repository_domains.PerformanceRepositoryInterface {
	return &performanceRepo{}
}

func (performanceRepo) GetEmployeePerformances(ctx context.Context,
	employeeID uint,
	from *time.Time,
	to *time.Time,
) ([]models.Performance, error) {
	var results []models.Performance
	db := providers.GetConnection().
		WithContext(ctx).
		Model(&models.Performance{}).
		Where("employee_id = ?", employeeID)
	if from != nil && to != nil {
		db.Where("performance_forms.date BETWEEN ? AND ?", from, to).
			Joins("JOIN performance_forms ON performance_forms.id=performances.performance_form_id")
	}
	finalResult := db.Find(&results)
	if finalResult.RowsAffected == 0 {
		return results, gorm.ErrRecordNotFound
	}
	return results, nil
}
