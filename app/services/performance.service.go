package services

import (
	"context"
	"fmt"
	"time"

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

type performanceService struct{}

func NewPerformanceService() service_domains.PerformanceServiceInterface {
	return &performanceService{}
}

func (performanceService) CreateForm(ctx context.Context, placementID uint) (models.PerformanceForm, error) {
	data := models.PerformanceForm{}
	claims, _ := ctx.Value(configs.VarConfig.UserContextName).(types.JWTCLaims)
	policyResult := policies.PerformancePolicy{}.Create(claims)
	if !policyResult {
		helpers.Logger.Warning(
			fmt.Sprintf("(Forbidden) User: %s", claims.Email),
		)
		return data, types.ErrForbiden
	}
	var placement models.Placement
	err := providers.RepoProvider.BaseRepo.Find(ctx, placementID, &placement)
	if err != nil {
		return data, err
	}
	if claims.Role == string(enums.SupervisorUserRole) {
		if placement.SupervisorID != &claims.ProfileID {
			return data, types.ErrForbiden
		}
	}
	data.SetUpdatedEmail(claims.Email)
	data.PlacementID = &placementID
	now := time.Now()
	data.Date = &now
	err = providers.RepoProvider.BaseRepo.Create(ctx, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}
func (performanceService) DeleteForm(ctx context.Context, formID uint) error {
	return baseDeleteFunc(ctx, policies.PerformancePolicy{}, formID, &models.PerformanceForm{})
}
func (performanceService) SubmitAnswer(ctx context.Context, data forms.PerformanceSubmitForm) error {
	datas := []models.PerformanceFormFeedback{}

	claims, _ := ctx.Value(configs.VarConfig.UserContextName).(types.JWTCLaims)
	if err := helpers.Validator.Validate(data); err != nil {
		return err
	}
	var form models.PerformanceForm
	err := providers.RepoProvider.BaseRepo.Find(ctx, data.PerformanceFormID, &form,
		"Placement", "Placement.Transaction", "Placement.Details", "Placement.Details.Employees",
	)
	if err != nil {
		return err
	}
	if claims.Role != string(enums.ServiceUserRole) {
		return types.ErrForbiden
	}
	if *form.Placement.Transaction.ServiceUserID != claims.ProfileID {
		return types.ErrForbiden
	}
	if form.FilledDate != nil {
		return types.ErrUnprocessableEntity.SetCustomMsg("you already filled this form")
	}
	answers := data.Details
	performances := []models.Performance{}
	for _, answer := range answers {
		cont := models.PerformanceFormFeedback{}
		performance := models.Performance{}
		placementEmployee := models.PlacementDetailEmployee{}
		if err := helpers.Translator.TranslateStruct(answer, &cont); err != nil {
			return err
		}
		placementEmployeeID := answer.PlacementDetailEmployeeID
		err := providers.RepoProvider.BaseRepo.Find(ctx,
			placementEmployeeID,
			&placementEmployee,
		)
		if err != nil {
			return err
		}
		performance.EmployeeID = placementEmployee.EmployeeID
		cont.PerformanceFormID = &form.ID
		performance.PerformanceFormID = &form.ID
		cont.ServiceUserID = &claims.ProfileID
		performance.ServiceUserID = &claims.ProfileID
		now := time.Now()
		cont.Date = &now
		performance.Date = &now
		rate, ok := enums.PerformanceRateMap[answer.Answer]
		if !ok {
			rate = 1
		}
		cont.Rate = &rate
		performance.Rate = &rate
		datas = append(datas, cont)
		performances = append(performances, performance)
	}
	err = providers.RepoProvider.BaseRepo.Create(ctx, datas)
	if err != nil {
		return err
	}
	err = providers.RepoProvider.BaseRepo.Create(ctx, performances)
	if err != nil {
		return err
	}
	now := time.Now()
	form.FilledDate = &now
	err = providers.RepoProvider.BaseRepo.Update(ctx, []uint{
		form.ID,
	}, &form)
	if err != nil {
		return err
	}
	return nil
}
func (performanceService) GetEmployeePerformances(ctx context.Context, employeeID uint, from *time.Time, to *time.Time) ([]models.Performance, error) {
	results := []models.Performance{}
	claims, _ := ctx.Value(configs.VarConfig.UserContextName).(types.JWTCLaims)
	policyResult := policies.PerformancePolicy{}.ReadAll(claims)
	if !policyResult {
		helpers.Logger.Warning(
			fmt.Sprintf("(Forbidden) User: %s", claims.Email),
		)
		return results, types.ErrForbiden
	}
	results, err := providers.RepoProvider.PerformanceRepo.GetEmployeePerformances(
		ctx,
		employeeID,
		from,
		to,
	)
	if err != nil {
		return results, err
	}
	return results, nil
}

func (performanceService) GetForm(ctx context.Context, formID uint) (models.PerformanceForm, error) {
	var data models.PerformanceForm
	err := baseFindFunc(ctx, policies.PerformancePolicy{}, formID, &data,
		"Placement",
		"Placement.Details",
		"Placement.Details.Service",
		"Placement.Details.Employees",
		"Placement.Details.Service.Category",
		"Feedbacks",
		"Feedbacks.PlacementDetailEmployee",
		"Feedbacks.PlacementDetailEmployee.Employee",
	)
	return data, err
}
