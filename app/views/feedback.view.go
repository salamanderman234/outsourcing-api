package views

import (
	"context"

	"github.com/labstack/echo/v4"
	view_domains "github.com/salamanderman234/outsourcing-api/app/domains/views"
	"github.com/salamanderman234/outsourcing-api/app/forms"
	"github.com/salamanderman234/outsourcing-api/app/providers"
	"github.com/salamanderman234/outsourcing-api/app/types"
)

type feedbackView struct{}

func NewFeedbackView() view_domains.FeedbackViewInterface {
	return &feedbackView{}
}

func (feedbackView) Create(c echo.Context) error {
	var form forms.FeedbackCreateForm
	callback := func(ctx context.Context) (any, error) {
		return providers.ServiceProvider.FeedbackService.Create(ctx, form)
	}
	return baseCreateFunc(c, &form, callback)
}
func (feedbackView) Read(c echo.Context) error {
	callback := func(ctx context.Context, q string, page uint) (any, *types.Pagination, error) {
		return providers.ServiceProvider.FeedbackService.Read(ctx, q, page)
	}
	return baseReadFunc(c, callback)
}
func (feedbackView) Find(c echo.Context) error {
	callback := func(ctx context.Context, id uint) (any, error) {
		return providers.ServiceProvider.FeedbackService.Find(ctx, id)
	}
	return baseFindFunc(c, callback)
}
func (feedbackView) Update(c echo.Context) error {
	var form forms.FeedbackUpdateForm
	callback := func(ctx context.Context, id uint) (uint, any, error) {
		return providers.ServiceProvider.FeedbackService.Update(ctx, id, form)
	}
	return baseUpdateFunc(c, &form, callback)
}
func (feedbackView) Delete(c echo.Context) error {
	callback := func(ctx context.Context, id uint) (uint, error) {
		return providers.ServiceProvider.FeedbackService.Delete(ctx, id)
	}
	return baseDeleteFunc(c, callback)
}
