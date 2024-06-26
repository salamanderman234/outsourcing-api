package views

import (
	"context"
	"strconv"

	"github.com/labstack/echo/v4"
	view_domains "github.com/salamanderman234/outsourcing-api/app/domains/views"
	"github.com/salamanderman234/outsourcing-api/app/forms"
	"github.com/salamanderman234/outsourcing-api/app/providers"
	"github.com/salamanderman234/outsourcing-api/app/types"
)

// province master
type masterProvinceView struct{}

func NewMasterProvinceView() view_domains.MasterProvinceViewInterface {
	return &masterProvinceView{}
}

func (masterProvinceView) Create(c echo.Context) error {
	var form forms.MasterProviceCreateForm
	callback := func(ctx context.Context) (any, error) {
		return providers.ServiceProvider.MasterProvinceService.Create(ctx, form)
	}
	return baseCreateFunc(c, &form, callback)
}
func (masterProvinceView) Read(c echo.Context) error {
	callback := func(ctx context.Context, q string, page uint) (any, *types.Pagination, error) {
		return providers.ServiceProvider.MasterProvinceService.Read(ctx, q, page)
	}
	return baseReadFunc(c, callback)
}
func (masterProvinceView) Find(c echo.Context) error {
	callback := func(ctx context.Context, id uint) (any, error) {
		return providers.ServiceProvider.MasterProvinceService.Find(ctx, id)
	}
	return baseFindFunc(c, callback)
}
func (masterProvinceView) Update(c echo.Context) error {
	var form forms.MasterProviceUpdateForm
	callback := func(ctx context.Context, id uint) (uint, any, error) {
		return providers.ServiceProvider.MasterProvinceService.Update(ctx, id, form)
	}
	return baseUpdateFunc(c, &form, callback)
}
func (masterProvinceView) Delete(c echo.Context) error {
	callback := func(ctx context.Context, id uint) (uint, error) {
		return providers.ServiceProvider.MasterProvinceService.Delete(ctx, id)
	}
	return baseDeleteFunc(c, callback)
}

// end of province master

// regency master
type regencyMasterView struct{}

func NewRegencyMasterView() view_domains.MasterRegencyViewInterface {
	return &regencyMasterView{}
}
func (regencyMasterView) Create(c echo.Context) error {
	var form forms.MasterRegencyCreateForm
	callback := func(ctx context.Context) (any, error) {
		return providers.ServiceProvider.MasterRegencyService.Create(ctx, form)
	}
	return baseCreateFunc(c, &form, callback)
}
func (regencyMasterView) Read(c echo.Context) error {
	callback := func(ctx context.Context, q string, page uint) (any, *types.Pagination, error) {
		return providers.ServiceProvider.MasterRegencyService.Read(ctx, q, page)
	}
	return baseReadFunc(c, callback)
}
func (regencyMasterView) Find(c echo.Context) error {
	callback := func(ctx context.Context, id uint) (any, error) {
		return providers.ServiceProvider.MasterRegencyService.Find(ctx, id)
	}
	return baseFindFunc(c, callback)
}
func (regencyMasterView) Update(c echo.Context) error {
	var form forms.MasterRegencyUpdateForm
	callback := func(ctx context.Context, id uint) (uint, any, error) {
		return providers.ServiceProvider.MasterRegencyService.Update(ctx, id, form)
	}
	return baseUpdateFunc(c, &form, callback)
}
func (regencyMasterView) Delete(c echo.Context) error {
	callback := func(ctx context.Context, id uint) (uint, error) {
		return providers.ServiceProvider.MasterRegencyService.Delete(ctx, id)
	}
	return baseDeleteFunc(c, callback)
}

// end of regency master
// category master
type categoryMasterView struct{}

func NewCategoryMasterView() view_domains.MasterCategoryViewInterface {
	return &categoryMasterView{}
}
func (categoryMasterView) Create(c echo.Context) error {
	var form forms.MasterCategoryCreateForm
	callback := func(ctx context.Context) (any, error) {
		return providers.ServiceProvider.MasterCategoryService.Create(ctx, form)
	}
	return baseCreateFunc(c, &form, callback)
}
func (categoryMasterView) Read(c echo.Context) error {
	callback := func(ctx context.Context, q string, page uint) (any, *types.Pagination, error) {
		return providers.ServiceProvider.MasterCategoryService.Read(ctx, q, page)
	}
	return baseReadFunc(c, callback)
}
func (categoryMasterView) Find(c echo.Context) error {
	callback := func(ctx context.Context, id uint) (any, error) {
		return providers.ServiceProvider.MasterCategoryService.Find(ctx, id)
	}
	return baseFindFunc(c, callback)
}
func (categoryMasterView) Update(c echo.Context) error {
	var form forms.MasterCategoryUpdateForm
	callback := func(ctx context.Context, id uint) (uint, any, error) {
		return providers.ServiceProvider.MasterCategoryService.Update(ctx, id, form)
	}
	return baseUpdateFunc(c, &form, callback)
}
func (categoryMasterView) Delete(c echo.Context) error {
	callback := func(ctx context.Context, id uint) (uint, error) {
		return providers.ServiceProvider.MasterCategoryService.Delete(ctx, id)
	}
	return baseDeleteFunc(c, callback)
}

// end of category master
// question master
type questionMasterView struct{}

func NewQuestionMasterView() view_domains.MasterQuestionViewInterface {
	return &questionMasterView{}
}
func (questionMasterView) Create(c echo.Context) error {
	var form forms.MasterQuestionCreateForm
	callback := func(ctx context.Context) (any, error) {
		return providers.ServiceProvider.MasterQuestionService.Create(ctx, form)
	}
	return baseCreateFunc(c, &form, callback)
}
func (questionMasterView) Read(c echo.Context) error {
	categoryIDStr := c.QueryParam("category_id")
	categoryID, _ := strconv.Atoi(categoryIDStr)
	callback := func(ctx context.Context, q string, page uint) (any, *types.Pagination, error) {
		return providers.ServiceProvider.MasterQuestionService.Read(ctx, uint(categoryID), q, page)
	}
	return baseReadFunc(c, callback)
}
func (questionMasterView) Find(c echo.Context) error {
	callback := func(ctx context.Context, id uint) (any, error) {
		return providers.ServiceProvider.MasterQuestionService.Find(ctx, id)
	}
	return baseFindFunc(c, callback)
}
func (questionMasterView) Update(c echo.Context) error {
	var form forms.MasterQuestionUpdateForm
	callback := func(ctx context.Context, id uint) (uint, any, error) {
		return providers.ServiceProvider.MasterQuestionService.Update(ctx, id, form)
	}
	return baseUpdateFunc(c, &form, callback)
}
func (questionMasterView) Delete(c echo.Context) error {
	callback := func(ctx context.Context, id uint) (uint, error) {
		return providers.ServiceProvider.MasterQuestionService.Delete(ctx, id)
	}
	return baseDeleteFunc(c, callback)
}
func (questionMasterView) AssignQuestion(c echo.Context) error {
	var form forms.MasterQuestionAssignForm
	callback := func(ctx context.Context) (any, error) {
		err := providers.ServiceProvider.MasterQuestionService.AssignQuestion(ctx, form)
		return nil, err
	}
	return baseCreateFunc(c, &form, callback)
}
func (questionMasterView) UnassignQuestion(c echo.Context) error {
	var form forms.MasterQuestionUnassignForm
	callback := func(ctx context.Context, id uint) (uint, any, error) {
		err := providers.ServiceProvider.MasterQuestionService.UnassignQuestion(ctx, form)
		return 0, nil, err
	}
	return baseUpdateFunc(c, &form, callback)
}

// end of question master
// payment config
type paymentConfigMasterView struct{}

func NewPaymentConfigMasterView() view_domains.MasterPaymentConfigViewInterface {
	return &paymentConfigMasterView{}
}
func (paymentConfigMasterView) SetDPPercentage(c echo.Context) error {
	var form forms.MaterPaymentConfigUpdateForm
	callback := func(ctx context.Context, id uint) (uint, any, error) {
		err := providers.ServiceProvider.MasterPaymentConfigService.SetDPPercentage(ctx, form.Amount)
		return 0, nil, err
	}
	return baseUpdateFunc(c, &form, callback)
}
func (paymentConfigMasterView) Set3TerminFirst(c echo.Context) error {
	var form forms.MaterPaymentConfigUpdateForm
	callback := func(ctx context.Context, id uint) (uint, any, error) {
		err := providers.ServiceProvider.MasterPaymentConfigService.Set3TerminFirst(ctx, form.Amount)
		return 0, nil, err
	}
	return baseUpdateFunc(c, &form, callback)
}

func (paymentConfigMasterView) Set3TerminSecond(c echo.Context) error {
	var form forms.MaterPaymentConfigUpdateForm
	callback := func(ctx context.Context, id uint) (uint, any, error) {
		err := providers.ServiceProvider.MasterPaymentConfigService.Set3TerminSecond(ctx, form.Amount)
		return 0, nil, err
	}
	return baseUpdateFunc(c, &form, callback)
}

// end of payment config
