package service_domains

import (
	"context"

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
	CreateForm(ctx context.Context, placementID uint) (models.PerformanceForm, error)
	DeleteForm(ctx context.Context, formID uint) error
	SubmitAnswer(ctx context.Context, data forms.PerformanceSubmitForm) error
	GetEmployeePerformances(ctx context.Context, employeeID uint, month uint, year uint) ([]models.Performance, error)
}
