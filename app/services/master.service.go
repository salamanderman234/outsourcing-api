package services

import (
	"context"
	"fmt"
	"strconv"

	service_domains "github.com/salamanderman234/outsourcing-api/app/domains/services"
	"github.com/salamanderman234/outsourcing-api/app/forms"
	"github.com/salamanderman234/outsourcing-api/app/helpers"
	"github.com/salamanderman234/outsourcing-api/app/models"
	"github.com/salamanderman234/outsourcing-api/app/policies"
	"github.com/salamanderman234/outsourcing-api/app/providers"
	"github.com/salamanderman234/outsourcing-api/app/types"
	"github.com/salamanderman234/outsourcing-api/app/types/enums"
	"github.com/salamanderman234/outsourcing-api/configs"
)

// province master
type masterProvinceService struct{}

func NewMasterProvinceService() service_domains.MasterProvinceServiceInterface {
	return &masterProvinceService{}
}

func (masterProvinceService) Create(ctx context.Context,
	data forms.MasterProviceCreateForm,
) (models.Province, error) {
	var province models.Province
	err := baseCreateFunc(ctx, policies.MasterPolicy{}, &province, data)
	return province, err
}
func (masterProvinceService) Read(ctx context.Context,
	q string, page uint,
) ([]models.Province, *types.Pagination, error) {
	var results []models.Province
	params := types.DBSearchParams{
		Page:           page,
		WithPagination: page > 0,
		Params: []types.WhereQuery{
			{Field: "province", Operator: "LIKE", Str: q},
		},
		Query: q,
		Model: &models.Province{},
	}

	pagination, err := baseReadFunc(
		ctx,
		policies.MasterPolicy{},
		params,
		&results,
	)
	return results, pagination, err
}
func (masterProvinceService) Find(ctx context.Context, id uint) (models.Province, error) {
	var province models.Province
	err := baseFindFunc(ctx, policies.MasterPolicy{}, id, &province)
	return province, err
}
func (masterProvinceService) Update(ctx context.Context,
	id uint, data forms.MasterProviceUpdateForm,
) (uint, models.Province, error) {
	var province models.Province
	err := baseUpdateFunc(ctx, policies.MasterPolicy{}, id, &province, data)
	return id, province, err
}
func (masterProvinceService) Delete(ctx context.Context, id uint) (uint, error) {
	err := baseDeleteFunc(ctx, policies.MasterPolicy{}, id, &models.Province{})
	return id, err
}

// end of province master

// regency master
type regencyMasterService struct{}

func NewRegencyMasterService() service_domains.MasterRegencyServiceInterface {
	return &regencyMasterService{}
}
func (regencyMasterService) Create(
	ctx context.Context,
	data forms.MasterRegencyCreateForm,
) (models.Regency, error) {
	var regency models.Regency
	err := baseCreateFunc(ctx, policies.MasterPolicy{}, &regency, data)
	return regency, err
}
func (regencyMasterService) Read(ctx context.Context,
	q string,
	page uint,
) ([]models.Regency, *types.Pagination, error) {
	var results []models.Regency
	params := types.DBSearchParams{
		Page:           page,
		WithPagination: page > 0,
		Params: []types.WhereQuery{
			{Field: "regency", Operator: "LIKE", Str: q},
		},
		Query: q,
		Model: &models.Regency{},
	}

	pagination, err := baseReadFunc(
		ctx,
		policies.MasterPolicy{},
		params,
		&results,
	)
	return results, pagination, err
}
func (regencyMasterService) Find(ctx context.Context, id uint) (models.Regency, error) {
	var regency models.Regency
	err := baseFindFunc(ctx, policies.MasterPolicy{}, id, &regency)
	return regency, err
}
func (regencyMasterService) Update(ctx context.Context,
	id uint,
	data forms.MasterRegencyUpdateForm,
) (uint, models.Regency, error) {
	var regency models.Regency
	err := baseUpdateFunc(ctx, policies.MasterPolicy{}, id, &regency, data)
	return id, regency, err
}
func (regencyMasterService) Delete(ctx context.Context, id uint) (uint, error) {
	err := baseDeleteFunc(ctx, policies.MasterPolicy{}, id, &models.Regency{})
	return id, err
}

// end of regency master

// category master
type categoryMasterService struct{}

func NewCategoryMasterService() service_domains.MasterCategoryServiceInterface {
	return &categoryMasterService{}
}
func (categoryMasterService) Create(
	ctx context.Context,
	data forms.MasterCategoryCreateForm,
) (models.Category, error) {

	var category models.Category
	err := baseCreateFunc(ctx, policies.MasterPolicy{}, &category, data)
	return category, err
}
func (categoryMasterService) Read(ctx context.Context, q string, page uint) ([]models.Category, *types.Pagination, error) {
	var results []models.Category
	params := types.DBSearchParams{
		Page:           page,
		WithPagination: page > 0,
		Params: []types.WhereQuery{
			{Field: "category_name", Operator: "LIKE", Str: q},
			{Field: "description", Operator: "LIKE", Str: q, IsOr: true},
		},
		Query: q,
		Model: &models.Category{},
	}

	pagination, err := baseReadFunc(
		ctx,
		policies.MasterPolicy{},
		params,
		&results,
	)
	return results, pagination, err
}
func (categoryMasterService) Find(ctx context.Context, id uint) (models.Category, error) {
	var category models.Category
	err := baseFindFunc(ctx, policies.MasterPolicy{}, id, &category)
	return category, err
}
func (categoryMasterService) Update(ctx context.Context, id uint, data forms.MasterCategoryUpdateForm) (uint, models.Category, error) {
	var category models.Category
	err := baseUpdateFunc(ctx, policies.MasterPolicy{}, id, &category, data)
	return id, category, err
}
func (categoryMasterService) Delete(ctx context.Context, id uint) (uint, error) {
	err := baseDeleteFunc(ctx, policies.MasterPolicy{}, id, &models.Category{})
	return id, err
}

// end of category master

// question master
type questionMasterService struct{}

func NewQuestionMasterService() service_domains.MasterQuestionServiceInterface {
	return &questionMasterService{}
}
func (questionMasterService) Create(
	ctx context.Context,
	data forms.MasterQuestionCreateForm,
) (models.Question, error) {
	var question models.Question
	err := baseCreateFunc(ctx, policies.MasterPolicy{}, &question, data)
	return question, err
}

func (questionMasterService) AssignQuestion(ctx context.Context, data forms.MasterQuestionAssignForm) error {
	var question models.CategoryQuestion
	before := func() error {
		var exists []models.CategoryQuestion
		categoryIDStr := strconv.Itoa(int(data.CategoryID))
		questionIDStr := strconv.Itoa(int(data.QuestionID))
		params := types.DBSearchParams{
			Page:           0,
			WithPagination: false,
			Params: []types.WhereQuery{
				{Field: "category_id", Operator: "=", Str: categoryIDStr},
				{Field: "question_id", Operator: "=", Str: questionIDStr},
			},
			Model: &models.CategoryQuestion{},
		}
		providers.RepoProvider.BaseRepo.ReadAll(ctx, &exists, params)
		if len(exists) != 0 {
			return types.ErrUnprocessableEntity.SetCustomMsg(
				"this question has been linked to this category",
			)
		}
		return nil
	}
	err := baseCreateFunc(ctx, policies.MasterPolicy{}, &question, data, before)
	return err
}
func (questionMasterService) UnassignQuestion(ctx context.Context, data forms.MasterQuestionUnassignForm) error {
	id := uint(0)
	before := func() error {
		var exists models.CategoryQuestion
		err := providers.RepoProvider.BaseRepo.FindWhere(ctx, map[string]any{
			"category_id": data.CategoryID,
			"question_id": data.QuestionID,
		}, &exists)
		return err
	}
	err := baseDeleteFunc(ctx, policies.MasterPolicy{}, id, &models.CategoryQuestion{}, before)
	return err
}

func (questionMasterService) Read(
	ctx context.Context,
	categoryID uint,
	q string,
	page uint,
) ([]models.Question, *types.Pagination, error) {
	var results []models.Question
	params := types.DBSearchParams{
		Page:           page,
		WithPagination: page > 0,
		Params: []types.WhereQuery{
			{Field: "question", Operator: "LIKE", Str: q},
			{Field: "hint", Operator: "LIKE", Str: q, IsOr: true},
		},
		Query: q,
		Model: &models.Question{},
	}

	if categoryID != 0 {
		var categoryQuestions []models.CategoryQuestion
		categoryIDStr := strconv.Itoa(int(categoryID))
		params.Params = []types.WhereQuery{
			{Field: "category_id", Operator: "=", Str: categoryIDStr},
		}
		params.Preloads = []string{
			"Question",
		}
		params.Model = &models.CategoryQuestion{}
		pagination, err := baseReadFunc(
			ctx,
			policies.MasterPolicy{},
			params,
			&categoryQuestions,
		)
		if err != nil {
			return results, pagination, err
		}
		for _, categoryQuestion := range categoryQuestions {
			question := categoryQuestion.Question
			if question != nil {
				results = append(results, *question)
			}
		}
		return results, pagination, nil
	} else {
		pagination, err := baseReadFunc(
			ctx,
			policies.MasterPolicy{},
			params,
			&results,
		)
		return results, pagination, err
	}
}

func (questionMasterService) Find(ctx context.Context, id uint) (models.Question, error) {
	var question models.Question
	err := baseFindFunc(ctx, policies.MasterPolicy{}, id, &question)
	return question, err
}
func (questionMasterService) Update(ctx context.Context, id uint, data forms.MasterQuestionUpdateForm) (uint, models.Question, error) {
	var question models.Question
	err := baseUpdateFunc(ctx, policies.MasterPolicy{}, id, &question, data)
	return id, question, err
}
func (questionMasterService) Delete(ctx context.Context, id uint) (uint, error) {
	err := baseDeleteFunc(ctx, policies.MasterPolicy{}, id, &models.Question{})
	return id, err
}

// end of question master

// payment config
type paymentConfigMasterService struct{}

func NewPaymentConfigMasterService() service_domains.MasterPaymentConfigServiceInterface {
	return &paymentConfigMasterService{}
}
func (paymentConfigMasterService) SetDPPercentage(ctx context.Context, amount uint) error {
	var data models.PaymentConfig
	claims, _ := ctx.Value(configs.VarConfig.UserContextName).(types.JWTCLaims)
	policyResult := policies.MasterPolicy{}.Update(data, claims)
	if !policyResult {
		helpers.Logger.Warning(
			fmt.Sprintf("(Forbidden) User: %s", claims.Email),
		)
		return types.ErrForbiden
	}
	params := types.DBSearchParams{
		Params: []types.WhereQuery{
			{Field: "type", Operator: "=", Str: string(enums.DpPayment)},
			{Field: "sub_type", Operator: "=", Str: string(enums.DPSubType)},
		},
	}
	var exists []models.PaymentConfig
	providers.RepoProvider.BaseRepo.ReadAll(ctx, &exists, params)
	if len(exists) == 0 {
		typ := string(enums.DpPayment)
		data.Type = &typ
		subtyp := string(enums.DPSubType)
		data.SubType = &subtyp
		return providers.RepoProvider.BaseRepo.Create(ctx, &data)
	} else {
		err := providers.RepoProvider.BaseRepo.FindWhere(ctx, map[string]any{
			"type":     string(enums.DpPayment),
			"sub_type": string(enums.DPSubType),
		}, &data)
		if err != nil {
			return err
		}
		floatAmount := float32(amount)
		data.Amount = &floatAmount
		return providers.RepoProvider.BaseRepo.Update(ctx, []uint{data.ID}, &data)
	}
}
func (paymentConfigMasterService) Set3TerminFirst(ctx context.Context, amount uint) error {
	var data models.PaymentConfig
	claims, _ := ctx.Value(configs.VarConfig.UserContextName).(types.JWTCLaims)
	policyResult := policies.MasterPolicy{}.Update(data, claims)
	if !policyResult {
		helpers.Logger.Warning(
			fmt.Sprintf("(Forbidden) User: %s", claims.Email),
		)
		return types.ErrForbiden
	}
	params := types.DBSearchParams{
		Params: []types.WhereQuery{
			{Field: "type", Operator: "=", Str: string(enums.ThreeTermin)},
			{Field: "sub_type", Operator: "=", Str: string(enums.ThreeTerminFirstSubType)},
		},
	}
	var exists []models.PaymentConfig
	providers.RepoProvider.BaseRepo.ReadAll(ctx, &exists, params)
	if len(exists) == 0 {
		typ := string(enums.ThreeTermin)
		data.Type = &typ
		subtyp := string(enums.ThreeTerminFirstSubType)
		data.SubType = &subtyp
		return providers.RepoProvider.BaseRepo.Create(ctx, &data)
	} else {
		err := providers.RepoProvider.BaseRepo.FindWhere(ctx, map[string]any{
			"type":     string(enums.ThreeTermin),
			"sub_type": string(enums.ThreeTerminFirstSubType),
		}, &data)
		if err != nil {
			return err
		}
		floatAmount := float32(amount)
		data.Amount = &floatAmount
		return providers.RepoProvider.BaseRepo.Update(ctx, []uint{data.ID}, &data)
	}
}
func (paymentConfigMasterService) Set3TerminSecond(ctx context.Context, amount uint) error {
	var data models.PaymentConfig
	claims, _ := ctx.Value(configs.VarConfig.UserContextName).(types.JWTCLaims)
	policyResult := policies.MasterPolicy{}.Update(data, claims)
	if !policyResult {
		helpers.Logger.Warning(
			fmt.Sprintf("(Forbidden) User: %s", claims.Email),
		)
		return types.ErrForbiden
	}
	params := types.DBSearchParams{
		Params: []types.WhereQuery{
			{Field: "type", Operator: "=", Str: string(enums.ThreeTermin)},
			{Field: "sub_type", Operator: "=", Str: string(enums.ThreeTerminSecondSubType)},
		},
	}
	var exists []models.PaymentConfig
	providers.RepoProvider.BaseRepo.ReadAll(ctx, &exists, params)
	if len(exists) == 0 {
		typ := string(enums.ThreeTermin)
		data.Type = &typ
		subtyp := string(enums.ThreeTerminSecondSubType)
		data.SubType = &subtyp
		return providers.RepoProvider.BaseRepo.Create(ctx, &data)
	} else {
		err := providers.RepoProvider.BaseRepo.FindWhere(ctx, map[string]any{
			"type":     string(enums.ThreeTermin),
			"sub_type": string(enums.ThreeTerminSecondSubType),
		}, &data)
		if err != nil {
			return err
		}
		floatAmount := float32(amount)
		data.Amount = &floatAmount
		return providers.RepoProvider.BaseRepo.Update(ctx, []uint{data.ID}, &data)
	}
}

// end of payment config
