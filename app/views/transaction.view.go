package views

import (
	"context"

	"github.com/labstack/echo/v4"
	view_domains "github.com/salamanderman234/outsourcing-api/app/domains/views"
	"github.com/salamanderman234/outsourcing-api/app/forms"
	"github.com/salamanderman234/outsourcing-api/app/providers"
	"github.com/salamanderman234/outsourcing-api/app/types"
)

type transactionView struct{}

func NewTransactionView() view_domains.TransactionViewInterface {
	return &transactionView{}
}

func (transactionView) Create(c echo.Context) error {
	var form forms.TransactionCreateForm
	callback := func(ctx context.Context) (any, error) {
		return providers.ServiceProvider.TransactionService.Create(ctx, form)
	}
	return baseCreateFunc(c, &form, callback)
}
func (transactionView) Read(c echo.Context) error {
	callback := func(ctx context.Context, q string, page uint) (any, *types.Pagination, error) {
		return providers.ServiceProvider.TransactionService.Read(ctx, q, page)
	}
	return baseReadFunc(c, callback)
}
func (transactionView) Find(c echo.Context) error {
	callback := func(ctx context.Context, id uint) (any, error) {
		return providers.ServiceProvider.TransactionService.Find(ctx, id)
	}
	return baseFindFunc(c, callback)
}
func (transactionView) Update(c echo.Context) error {
	var form forms.TransactionUpdateForm
	callback := func(ctx context.Context, id uint) (uint, any, error) {
		return providers.ServiceProvider.TransactionService.Update(ctx, id, form)
	}
	return baseUpdateFunc(c, &form, callback)
}
func (transactionView) Delete(c echo.Context) error {
	callback := func(ctx context.Context, id uint) (uint, error) {
		return providers.ServiceProvider.TransactionService.Delete(ctx, id)
	}
	return baseDeleteFunc(c, callback)
}

func (transactionView) UploadMOU(c echo.Context) error {
	return nil
}
