package service_domains

import (
	"context"
	"time"
)

type FeedbackServiceInterface interface {
	Create(ctx context.Context) error
	Find(ctx context.Context) error
	Read(ctx context.Context) error
	Update(ctx context.Context) error
	Delete(ctx context.Context) error
}

type ComplaintServiceInterface interface {
	Create(ctx context.Context) error
	Find(ctx context.Context) error
	Read(ctx context.Context) error
	Update(ctx context.Context) error
	Reply(ctx context.Context) error
	UpdateReply(ctx context.Context) error
	DeleteReply(ctx context.Context) error
	Delete(ctx context.Context) error
}

type PerformanceServiceInterface interface {
	UploadFeedback(ctx context.Context) error
	UpdateFeedback(ctx context.Context) error
	UploadEmployeePerformance(ctx context.Context) error
	FindEmployeePerformance(ctx context.Context, id uint) error
	Read(ctx context.Context, month time.Month, employeeID uint) error
	ReadFromPlacement(ctx context.Context, placementID uint, month time.Month) error
	UpdateEmployeePerformance(ctx context.Context, id uint) error
	DeleteEmployeePerformance(ctx context.Context, id uint) error
}
