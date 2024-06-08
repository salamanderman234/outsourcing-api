package service_domains

import (
	"context"
	"time"

	"github.com/salamanderman234/outsourcing-api/app/forms"
	"github.com/salamanderman234/outsourcing-api/app/models"
	"github.com/salamanderman234/outsourcing-api/app/types"
)

type FeedbackServiceInterface interface {
	Create(ctx context.Context, data forms.FeedbackCreateForm) (models.Feedback, error)
	Find(ctx context.Context, id uint) (models.Feedback, error)
	Read(ctx context.Context, q string, page uint) ([]models.Feedback, *types.Pagination, error)
	Update(ctx context.Context, id uint, data forms.FeedbackUpdateForm) (uint, models.Feedback, error)
	Delete(ctx context.Context, id uint) (uint, error)
}

type ComplaintServiceInterface interface {
	Create(ctx context.Context, data forms.ComplaintCreateForm) (models.Complaint, error)
	Find(ctx context.Context, id uint) (models.Complaint, error)
	Read(ctx context.Context, q string, page uint) ([]models.Complaint, *types.Pagination, error)
	Update(ctx context.Context, id uint, data forms.ComplaintUpdateForm) (uint, models.Complaint, error)
	Reply(ctx context.Context, data forms.ReplyComplaintCreateForm) (models.ComplaintReply, error)
	UpdateReply(ctx context.Context, id uint, data forms.ReplyComplaintUpdateForm) (uint, models.ComplaintReply, error)
	DeleteReply(ctx context.Context, id uint) (uint, error)
	Delete(ctx context.Context, id uint) (uint, error)
}

type PerformanceServiceInterface interface {
	CreatePerformanceForm(ctx context.Context, data forms.PerformanceCreateForm) (models.EmployeePerformance, error)
	UpdatePerformanceForm(ctx context.Context, id uint, data forms.PerformanceUpdateForm) (models.EmployeePerformance, error)
	UserUploadPerformanceForm(ctx context.Context, feedbackID uint, data forms.UserPerformanceUploadForm) (models.Feedback, error)
	UserUpdatePerformanceForm(ctx context.Context, feedbackID uint, data forms.UserPerformanceUpdateForm) (uint, models.Feedback, error)
	InputEmployeePerformance(ctx context.Context, data forms.PerformanceUploadInputForm) (models.EmployeePerformance, error)
	UpdateEmployeePerformance(ctx context.Context, id uint, data forms.PerformanceUpdateInputForm) (uint, models.EmployeePerformance, error)
	FindEmployeePerformance(ctx context.Context, id uint) (models.EmployeePerformance, error)
	// Read(ctx context.Context, month time.Month, employeeID uint) error
	ReadFromFeedback(ctx context.Context, feedbackID uint) (models.Feedback, []models.EmployeePerformance, error)
	ReadFromPlacement(ctx context.Context, placementID uint, month time.Month) (models.Placement, []models.EmployeePerformance, error)
	ReadFromEmployee(ctx context.Context, employeeID uint, month time.Month) (models.Employee, []models.EmployeePerformance, error)
	DeleteEmployeePerformance(ctx context.Context, id uint) (uint, error)
}
