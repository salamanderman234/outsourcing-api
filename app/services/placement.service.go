package services

import (
	"context"
	"strconv"
	"time"

	service_domains "github.com/salamanderman234/outsourcing-api/app/domains/services"
	"github.com/salamanderman234/outsourcing-api/app/forms"
	"github.com/salamanderman234/outsourcing-api/app/models"
	"github.com/salamanderman234/outsourcing-api/app/policies"
	"github.com/salamanderman234/outsourcing-api/app/providers"
	"github.com/salamanderman234/outsourcing-api/app/types"
	"github.com/salamanderman234/outsourcing-api/app/types/enums"
)

type placementService struct{}

func NewPlacementService() service_domains.PlacementServiceInterface {
	return &placementService{}
}

func (placementService) CreatePlacement(ctx context.Context, data forms.PlacementCreateForm) (models.Placement, error) {
	var placement models.Placement
	before := func() error {
		var transaction models.Transaction
		err := providers.RepoProvider.BaseRepo.Find(ctx, *placement.TransactionID, &transaction,
			"Placements", "Details", "ServiceUser", "Details.Service",
		)
		if err != nil {
			return err
		}
		if len(transaction.Placements) >= 1 {
			return types.ErrUnprocessableEntity
		}

		var supervisor models.Supervisor
		err = providers.RepoProvider.BaseRepo.Find(ctx, *placement.SupervisorID, &supervisor)
		if err != nil {
			return err
		}

		if *supervisor.RegencyID != *transaction.RegencyID {
			return types.ErrUnprocessableEntity
		}

		details := transaction.Details
		totalEmployee := uint(0)
		for _, detail := range details {
			zero := uint(0)
			salary := uint(*detail.EmployeePrice)
			placementDetail := models.PlacementDetail{
				ServiceID:     detail.ServiceID,
				TotalEmployee: detail.TotalEmployee,
				Filled:        &zero,
				Salary:        &salary,
			}
			placement.Details = append(placement.Details, placementDetail)
			totalEmployee += *detail.TotalEmployee
		}

		if transaction.ServiceUser != nil {
			placement.Name = transaction.ServiceUser.Fullname
		}

		placement.Address = transaction.Address
		placement.RegencyID = transaction.RegencyID
		transactionStartDate := *transaction.StartDate
		duration := *transaction.ContractDuration
		endDate := transactionStartDate.Add(time.Duration(duration) * (time.Hour * 24))
		placement.StartDate = &transactionStartDate
		placement.EndDate = &endDate
		placement.Duration = &duration
		placement.TotalEmployee = &totalEmployee
		status := string(enums.PlacementOngoingStatus)
		placement.Status = &status

		return nil
	}
	err := baseCreateFunc(ctx, policies.PlacementPolicy{}, &placement, data, before)
	return placement, err
}
func (placementService) GetPlacementOrder(ctx context.Context, id uint) (models.Placement, error) {
	var results []models.Placement
	transactionID := strconv.Itoa(int(id))
	params := types.DBSearchParams{
		WithPagination: false,
		Params: []types.WhereQuery{
			{Field: "transaction_id", Operator: "=", Str: transactionID},
		},
		Model: &models.Placement{},
	}

	_, err := baseReadFunc(
		ctx,
		policies.PlacementPolicy{},
		params,
		&results,
	)
	if len(results) != 1 {
		return models.Placement{}, types.ErrRecordNotFound
	}
	return results[0], err
}
func (placementService) Find(ctx context.Context, id uint) (models.Placement, error) {
	var placement models.Placement
	err := baseFindFunc(ctx, policies.PlacementPolicy{}, id, &placement)
	return placement, err
}
func (placementService) Read(ctx context.Context, q string, page uint) ([]models.Placement, *types.Pagination, error) {
	var results []models.Placement
	params := types.DBSearchParams{
		Page:           page,
		WithPagination: page > 0,
		Params: []types.WhereQuery{
			{Field: "status", Operator: "LIKE", Str: q},
		},
		Query: q,
		Model: &models.Placement{},
	}

	pagination, err := baseReadFunc(
		ctx,
		policies.PlacementPolicy{},
		params,
		&results,
	)
	return results, pagination, err
}
func (placementService) PlaceNewEmployee(ctx context.Context, data forms.PlacementDetailEmployeeCreateForm) (models.PlacementDetailEmployee, error) {
	var placementDetailEmployee models.PlacementDetailEmployee
	before := func() error {
		var detail models.PlacementDetail
		var employee models.Employee
		placementDetailID := data.PlacementDetailID

		err := providers.RepoProvider.BaseRepo.Find(ctx, placementDetailID, &detail, "Placement")
		if err != nil {
			return err
		}
		err = providers.RepoProvider.BaseRepo.Find(ctx, data.EmployeeID, &employee)
		if err != nil {
			return err
		}
		if *employee.RegencyID != *detail.Placement.RegencyID {
			return types.ErrUnprocessableEntity
		}
		ongoing := string(enums.PlacementEmployeeOngoingStatus)
		placementDetailEmployee.Status = &ongoing
		now := time.Now()
		if data.StartDate == nil {
			placementDetailEmployee.StartDate = &now
		}
		placementDetailEmployee.PlacementDate = &now
		placementDetailEmployee.ExpectedSalary = detail.Salary
		if detail.Placement != nil {
			placementDetailEmployee.Duration = detail.Placement.Duration
			placementDetailEmployee.EndDate = detail.Placement.EndDate
			diff := detail.Placement.EndDate.Sub(*placementDetailEmployee.StartDate).Hours()
			expectedWorkDay := uint(diff / 24)
			expectedTotalSalary := uint(expectedWorkDay * (*placementDetailEmployee.ExpectedSalary))
			placementDetailEmployee.ExpectedSalaryTotal = &expectedTotalSalary
		} else {
			return types.ErrUnprocessableEntity
		}
		filled := *detail.Filled + 1
		detail.Filled = &filled
		return nil
	}
	err := baseCreateFunc(ctx, policies.PlacementPolicy{}, &placementDetailEmployee, data, before)
	return placementDetailEmployee, err
}

func (placementService) CutoffEmployeePlacement(ctx context.Context, placementDetailEmployeeID uint, data forms.PlacementDetailEmployeeUpdateForm) (uint, error) {
	var placement models.PlacementDetailEmployee
	before := func() error {
		var pl models.PlacementDetailEmployee
		err := providers.RepoProvider.BaseRepo.Find(ctx, placementDetailEmployeeID, &pl)
		if err != nil {
			return err
		}
		if placement.ExitDate == nil {
			now := time.Now()
			placement.ExitDate = &now
		}
		salary := *pl.ExpectedSalary
		diff := placement.ExitDate.Sub(*placement.StartDate).Hours()
		workDiff := uint(diff / 24)
		actualSalary := workDiff * salary
		placement.ActualSalary = &actualSalary

		return nil
	}
	err := baseUpdateFunc(ctx, policies.PlacementPolicy{}, placementDetailEmployeeID, &placement, data, before)
	return placementDetailEmployeeID, err
}

func (placementService) RemoveEmployeePlacement(ctx context.Context, placementDetailEmployeeID uint) (uint, error) {
	err := baseDeleteFunc(ctx, policies.PlacementPolicy{}, placementDetailEmployeeID, &models.PlacementDetailEmployee{})
	return placementDetailEmployeeID, err
}
func (placementService) Delete(ctx context.Context, id uint) (uint, error) {
	err := baseDeleteFunc(ctx, policies.PlacementPolicy{}, id, &models.Placement{})
	return id, err
}
