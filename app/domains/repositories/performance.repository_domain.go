package repository_domains

import (
	"context"
	"time"

	"github.com/salamanderman234/outsourcing-api/app/models"
)

type PerformanceRepositoryInterface interface {
	GetEmployeePerformances(ctx context.Context, employeeID uint, from *time.Time, to *time.Time) ([]models.Performance, error)
}
