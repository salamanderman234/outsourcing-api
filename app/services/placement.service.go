package services

import (
	"context"
	"fmt"
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
		tranStatus := transaction.Status
		if len(transaction.Placements) >= 1 || *tranStatus != string(enums.WaitingForPlacement) {
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

		stat := string(enums.Ongoing)
		transaction.Status = &stat
		err = providers.RepoProvider.BaseRepo.Update(ctx, []uint{transaction.ID}, &transaction)
		if err != nil {
			return err
		}

		return nil
	}
	err := baseCreateFunc(ctx, policies.PlacementPolicy{}, &placement, data, before)
	return placement, err
}
func (placementService) GetPlacementOrder(ctx context.Context, id uint) (models.Placement, error) {
	var results []models.Placement
	transactionID := strconv.Itoa(int(id))
	params := types.DBSearchParams{
		Params: []types.WhereQuery{
			{Field: "transaction_id", Operator: "=", Str: transactionID},
		},
		Model: &models.Placement{},
		Preloads: []string{
			"Regency",
			"Supervisor",
			"Details",
			"Details.Service",
			"Details.Employees",
			"Details.Employees.Employee",
		},
	}

	_, err := baseReadFunc(
		ctx,
		policies.PlacementPolicy{},
		params,
		&results,
	)
	if err != nil {
		return models.Placement{}, err
	}
	if len(results) != 1 {
		return models.Placement{}, types.ErrRecordNotFound
	}
	return results[0], err
}
func (placementService) Find(ctx context.Context, id uint) (models.Placement, error) {
	var placement models.Placement
	err := baseFindFunc(ctx, policies.PlacementPolicy{}, id, &placement,
		"Regency",
		"Supervisor",
		"Details",
		"Details.Service",
		"Details.Employees",
		"Details.Employees.Employee")
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
		Preloads: []string{
			"Regency",
			"Supervisor",
			"Details",
			"Details.Service",
			"Details.Employees",
			"Details.Employees.Employee",
		},
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

		var exists models.PlacementDetailEmployee
		err := providers.RepoProvider.BaseRepo.FindWhere(ctx, map[string]any{
			"employee_id":         data.EmployeeID,
			"placement_detail_id": data.PlacementDetailID,
		}, &exists)

		if err == nil {
			return types.ErrDuplicateEntries
		}
		err = providers.RepoProvider.BaseRepo.Find(ctx, placementDetailID, &detail, "Placement")
		if err != nil {
			return err
		}
		err = providers.RepoProvider.BaseRepo.Find(ctx, data.EmployeeID, &employee)
		if err != nil {
			return err
		}

		fill := *detail.Filled
		max := *detail.TotalEmployee

		if fill == max {
			return types.ErrUnprocessableEntity
		}

		employeeRegency := *employee.RegencyID
		placementRegency := *detail.Placement.RegencyID
		if employeeRegency != placementRegency {
			return types.ErrUnprocessableEntity
		}
		ongoing := string(enums.PlacementEmployeeOngoingStatus)
		placementDetailEmployee.Status = &ongoing
		now := time.Now()
		if data.StartDate == nil {
			placementDetailEmployee.StartDate = &now
		}
		placementDetailEmployee.PlacementDate = &now
		tranEndDate := detail.Placement.EndDate
		tranStartDate := detail.Placement.StartDate

		if tranStartDate.Compare(*placementDetailEmployee.StartDate) < 0 {
			placementDetailEmployee.StartDate = tranStartDate
		}
		if tranEndDate.Compare(*placementDetailEmployee.PlacementDate) < 0 {
			return types.ErrUnprocessableEntity
		}
		placementDetailEmployee.ExpectedSalary = detail.Salary
		if detail.Placement != nil {
			placementDetailEmployee.Duration = detail.Placement.Duration
			placementDetailEmployee.EndDate = detail.Placement.EndDate
			diff := detail.Placement.EndDate.Sub(*placementDetailEmployee.StartDate).Hours()
			realDuration := *detail.Placement.Duration
			expectedWorkDay := uint(min(int(realDuration), int(diff/24)))
			expectedTotalSalary := uint(expectedWorkDay * (*placementDetailEmployee.ExpectedSalary))
			fmt.Println(expectedWorkDay, *placementDetailEmployee.ExpectedSalary, expectedTotalSalary)
			placementDetailEmployee.ExpectedSalaryTotal = &expectedTotalSalary
		} else {
			return types.ErrUnprocessableEntity
		}
		filled := *detail.Filled + 1
		detail.Filled = &filled
		err = providers.RepoProvider.BaseRepo.Update(ctx, []uint{detail.ID}, &detail)
		if err != nil {
			return err
		}
		return nil
	}
	err := baseCreateFunc(ctx, policies.PlacementPolicy{}, &placementDetailEmployee, data, before)
	return placementDetailEmployee, err
}

func (placementService) CutoffEmployeePlacement(ctx context.Context, placementDetailEmployeeID uint, data forms.PlacementDetailEmployeeUpdateForm) (uint, error) {
	var placement models.PlacementDetailEmployee
	before := func() error {
		var pl models.PlacementDetailEmployee
		err := providers.RepoProvider.BaseRepo.Find(ctx, placementDetailEmployeeID, &pl, "PlacementDetail")
		if err != nil {
			return err
		}
		if *pl.Status != string(enums.PlacementEmployeeSuspendStatus) && *pl.Status != string(enums.PlacementEmployeeOngoingStatus) {
			return types.ErrUnprocessableEntity
		}
		if placement.ExitDate == nil {
			now := time.Now()
			placement.ExitDate = &now
		}
		salary := *pl.ExpectedSalary
		diff := placement.ExitDate.Sub(*placement.StartDate).Hours()
		limitWork := *pl.Duration
		workDiff := uint(min(limitWork, uint(diff/24)))
		actualSalary := uint(min(workDiff*salary, *pl.ExpectedSalaryTotal))
		placement.ActualSalary = &actualSalary
		stat := enums.PlacementEmployeeDismissedStatus
		placement.Status = &stat

		detail := pl.PlacementDetail
		if detail != nil {
			filled := *detail.Filled
			updatefilled := filled - 1
			detail.Filled = &updatefilled
			err = providers.RepoProvider.BaseRepo.Update(ctx, []uint{
				detail.ID,
			}, detail)
			if err != nil {
				return err
			}
		}

		return nil
	}
	err := baseUpdateFunc(ctx, policies.PlacementPolicy{}, placementDetailEmployeeID, &placement, data, before)
	return placementDetailEmployeeID, err
}

func (placementService) Delete(ctx context.Context, id uint) (uint, error) {
	err := baseDeleteFunc(ctx, policies.PlacementPolicy{}, id, &models.Placement{})
	return id, err
}

func (placementService) RemoveEmployeePlacement(ctx context.Context, placementDetailEmployeeID uint) (uint, error) {
	before := func() error {
		var emp models.PlacementDetailEmployee
		err := providers.RepoProvider.BaseRepo.Find(ctx, placementDetailEmployeeID, &emp, "PlacementDetail")
		if err != nil {
			return err
		}
		status := emp.Status
		detail := emp.PlacementDetail
		if detail != nil && status != nil {
			if *status != string(enums.PlacementEmployeeDismissedStatus) {
				fill := *detail.Filled
				fill = max(0, fill-1)
				detail.Filled = &fill

				err := providers.RepoProvider.BaseRepo.Update(ctx, []uint{detail.ID}, detail)
				if err != nil {
					return err
				}
			}
		}
		return nil
	}
	err := baseDeleteFunc(ctx, policies.PlacementPolicy{}, placementDetailEmployeeID, &models.PlacementDetailEmployee{}, before)
	return placementDetailEmployeeID, err
}

func (placementService) Update(ctx context.Context, id uint, data forms.PlacementUpdateForm) (uint, models.Placement, error) {
	var regency models.Placement
	before := func() error {

		var supervisor models.Supervisor
		if data.SupervisorID != nil {
			var placement models.Placement
			err := providers.RepoProvider.BaseRepo.Find(ctx, id, &placement)
			if err != nil {
				return err
			}
			err = providers.RepoProvider.BaseRepo.Find(ctx, *data.SupervisorID, &supervisor)
			if err != nil {
				return err
			}

			if *supervisor.RegencyID != *placement.RegencyID {
				return types.ErrUnprocessableEntity
			}
		}
		return nil
	}
	err := baseUpdateFunc(ctx, policies.PlacementPolicy{}, id, &regency, data, before)
	return id, regency, err
}

func (placementService) EmployeePlacementDetail(ctx context.Context, id uint) (models.PlacementDetailEmployee, error) {
	var placementEmployee models.PlacementDetailEmployee
	err := baseFindFunc(ctx, policies.PlacementPolicy{}, id, &placementEmployee)
	return placementEmployee, err
}

func (placementService) PlacementDetails(ctx context.Context, id uint) ([]models.PlacementDetail, error) {
	var results []models.PlacementDetail
	idStr := strconv.Itoa(int(id))
	params := types.DBSearchParams{
		Page:           0,
		WithPagination: false,
		Params: []types.WhereQuery{
			{Field: "placement_id", Operator: "=", Str: idStr},
		},
		Query: idStr,
		Model: &models.PlacementDetail{},
		Preloads: []string{
			"Service",
			"Employees",
		},
	}

	_, err := baseReadFunc(
		ctx,
		policies.PlacementPolicy{},
		params,
		&results,
	)
	return results, err
}
