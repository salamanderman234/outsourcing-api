package views

import (
	"context"

	"github.com/labstack/echo/v4"
	view_domains "github.com/salamanderman234/outsourcing-api/app/domains/views"
	"github.com/salamanderman234/outsourcing-api/app/forms"
	"github.com/salamanderman234/outsourcing-api/app/providers"
	"github.com/salamanderman234/outsourcing-api/app/types"
)

type complaintView struct{}

func NewComplaintView() view_domains.ComplaintViewInterface {
	return &complaintView{}
}

func (complaintView) Create(c echo.Context) error {
	var form forms.ComplaintCreateForm
	callback := func(ctx context.Context) (any, error) {
		return providers.ServiceProvider.ComplaintService.Create(ctx, form)
	}
	return baseCreateFunc(c, &form, callback)
}
func (complaintView) Find(c echo.Context) error {
	callback := func(ctx context.Context, id uint) (any, error) {
		return providers.ServiceProvider.ComplaintService.Find(ctx, id)
	}
	return baseFindFunc(c, callback)
}
func (complaintView) Read(c echo.Context) error {
	callback := func(ctx context.Context, q string, page uint) (any, *types.Pagination, error) {
		return providers.ServiceProvider.ComplaintService.Read(ctx, q, page)
	}
	return baseReadFunc(c, callback)
}
func (complaintView) Update(c echo.Context) error {
	var form forms.ComplaintUpdateForm
	callback := func(ctx context.Context, id uint) (uint, any, error) {
		return providers.ServiceProvider.ComplaintService.Update(ctx, id, form)
	}
	return baseUpdateFunc(c, &form, callback)
}
func (complaintView) Delete(c echo.Context) error {
	callback := func(ctx context.Context, id uint) (uint, error) {
		return providers.ServiceProvider.ComplaintService.Delete(ctx, id)
	}
	return baseDeleteFunc(c, callback)
}
func (complaintView) Reply(c echo.Context) error {
	var form forms.ReplyComplaintCreateForm
	callback := func(ctx context.Context) (any, error) {
		return providers.ServiceProvider.ComplaintService.Reply(ctx, form)
	}
	return baseCreateFunc(c, &form, callback)
}
func (complaintView) UpdateReply(c echo.Context) error {
	var form forms.ReplyComplaintUpdateForm
	callback := func(ctx context.Context, id uint) (uint, any, error) {
		return providers.ServiceProvider.ComplaintService.UpdateReply(ctx, id, form)
	}
	return baseUpdateFunc(c, &form, callback)
}
func (complaintView) DeleteReply(c echo.Context) error {
	callback := func(ctx context.Context, id uint) (uint, error) {
		return providers.ServiceProvider.ComplaintService.DeleteReply(ctx, id)
	}
	return baseDeleteFunc(c, callback)
}
