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

type transactionService struct{}

func NewTransactionService() service_domains.TransactionServiceInterface {
	return &transactionService{}
}

func (transactionService) Create(ctx context.Context, data forms.TransactionCreateForm) (models.Transaction, error) {
	var transaction models.Transaction
	before := func() error {
		now := time.Now()
		initialPaymentDeadline := time.Now().Add(72 * time.Hour)
		transaction.OrderDate = &now
		transaction.NextPaymentDeadline = &initialPaymentDeadline
		etcPrice := uint64(0)
		employeePrice := uint64(0)
		servicePrice := uint64(0)

		if transaction.PackageID != nil {
			if *transaction.PackageID != 0 {
				var pack models.Package
				err := providers.RepoProvider.BaseRepo.Find(ctx, *transaction.PackageID, &pack,
					"Services",
					"Services.AdditionalPackageServiceItems",
					"Services.Service",
					"Services.AdditionalPackageServiceItems.AdditionalItemService",
				)
				if err != nil {
					return err
				}
				packageDetails := pack.Services
				details := []models.TransactionDetail{}
				for _, service := range packageDetails {
					detailEtcs := []models.TransactionDetailEtc{}
					etcs := service.AdditionalPackageServiceItems
					for _, etc := range etcs {
						detailEtcs = append(detailEtcs, models.TransactionDetailEtc{
							AdditionalItemServiceID: etc.AdditionalItemServiceID,
							Qty:                     etc.Quantity,
						})
					}
					details = append(details, models.TransactionDetail{
						ServiceID:     service.ServiceID,
						TotalEmployee: service.TotalEmployee,
						Etcs:          detailEtcs,
					})
				}
				transaction.Details = details
			}
		}

		if transaction.Details != nil {
			for index, detail := range transaction.Details {
				var service models.Service
				serviceID := detail.ServiceID
				if serviceID != nil {
					err := providers.RepoProvider.BaseRepo.Find(ctx, *serviceID, &service, "RequiredItems", "AdditionalItems", "Category")
					if err != nil {
						return err
					}
					servPrice := uint64(*service.ServicePrice) * uint64(*transaction.ContractDuration)
					empPrice := uint64(*service.EmployeePrice) * uint64(*detail.TotalEmployee) * uint64(*transaction.ContractDuration)

					transaction.Details[index].ServicePrice = &servPrice
					transaction.Details[index].EmployeePrice = &empPrice
				}
				detailEtcPrice := uint64(0)
				if service.EtcPrice != nil {
					detailEtcPrice = uint64(*service.EtcPrice)
				}
				etcs := detail.Etcs
				for indexY, etc := range etcs {
					var additional models.AdditionalItemService
					additionalItemID := etc.AdditionalItemServiceID
					if additionalItemID != nil {
						err := providers.RepoProvider.BaseRepo.Find(ctx, *additionalItemID, &additional)
						if err != nil {
							return err
						}
						pricePerItem := uint64(*additional.PricePerItem)
						total := pricePerItem * uint64(*etc.Qty)
						detailEtcPrice += total
						transaction.Details[index].Etcs[indexY].Price = &pricePerItem
						transaction.Details[index].Etcs[indexY].SubTotalPrice = &total

					}
				}
				transaction.Details[index].EtcPrice = &detailEtcPrice
				etcPrice += detailEtcPrice
				servicePrice += *transaction.Details[index].ServicePrice
				employeePrice += *transaction.Details[index].EmployeePrice
				subtotal := detailEtcPrice + (*transaction.Details[index].ServicePrice) + (*transaction.Details[index].EmployeePrice)
				transaction.Details[index].SubTotalPrice = &subtotal
			}
			total := etcPrice + servicePrice + employeePrice
			transaction.TotalPrice = &total
			zero := uint64(0)
			transaction.TotalPaid = &zero
			status := string(enums.WaitingForConfirmationStatus)
			transaction.Status = &status
		}
		return nil
	}
	err := baseCreateFunc(ctx, &policies.TransactionPolicy{}, &transaction, data, before)
	return transaction, err
}
func (transactionService) Update(ctx context.Context, id uint, data forms.TransactionUpdateForm) (uint, models.Transaction, error) {
	var transaction models.Transaction
	err := baseUpdateFunc(ctx, &policies.TransactionPolicy{}, id, &transaction, data)
	return id, transaction, err
}
func (transactionService) Read(ctx context.Context, q string, page uint) ([]models.Transaction, *types.Pagination, error) {
	var results []models.Transaction
	params := types.DBSearchParams{
		Page:           page,
		WithPagination: page > 0,
		Query:          q,
		Model:          &models.Transaction{},
		Params: []types.WhereQuery{
			{Field: "status", Operator: "LIKE", Str: q},
		},
		Preloads: []string{
			"Details",
			"ServiceUser",
			"Details.Etcs",
			"Details.Service",
			"Details.Etcs.AdditionalItemService",
		},
	}
	claims, _ := ctx.Value(configs.VarConfig.UserContextName).(types.JWTCLaims)
	if claims.Role == string(enums.ServiceUserRole) {
		idStr := claims.ID
		params.Params = []types.WhereQuery{
			{Field: "service_user_id", Operator: "=", Str: idStr},
		}
	}
	pagination, err := baseReadFunc(
		ctx,
		policies.TransactionPolicy{},
		params,
		&results,
	)
	return results, pagination, err
}
func (transactionService) Find(ctx context.Context, id uint) (models.Transaction, error) {
	var transaction models.Transaction
	err := baseFindFunc(
		ctx,
		&policies.TransactionPolicy{},
		id,
		&transaction,
		"Details",
		"ServiceUser",
		"Details.Etcs",
		"Details.Service",
		"Details.Etcs.AdditionalItemService",
	)
	return transaction, err
}
func (transactionService) UploadMOU(ctx context.Context, id uint, file string) error {
	claims, _ := ctx.Value(configs.VarConfig.UserContextName).(types.JWTCLaims)
	var transaction models.Transaction
	err := providers.RepoProvider.BaseRepo.Find(ctx, id, &transaction)
	if err != nil {
		return err
	}

	policyResult := policies.TransactionPolicy{}.Update(transaction, claims)
	if !policyResult {
		helpers.Logger.Warning(
			fmt.Sprintf("(Forbidden) User: %s", claims.Email),
		)
		return types.ErrForbiden
	}

	if (*transaction.Status) != string(enums.WaitingForMOU) {
		return types.ErrUnprocessableEntity
	}
	resource, err := providers.ResourceProvider.CreateResource("transactions.mou")
	if err != nil {
		return err
	}
	resource.SetData(&transaction)
	result, err := providers.ServiceProvider.FileService.UploadFile(ctx, file, resource)
	if err != nil {
		return err
	}
	status := string(enums.WaitingForMOUConfirmation)
	transaction.MOU = &result
	transaction.Status = &status
	err = providers.RepoProvider.BaseRepo.Update(ctx, []uint{id}, &transaction)
	return err
}
func (transactionService) Delete(ctx context.Context, id uint) (uint, error) {
	err := baseDeleteFunc(ctx, policies.TransactionPolicy{}, id, &models.Transaction{})
	return id, err
}

func (transactionService) AskForMOU(ctx context.Context, id uint) error {
	data := struct {
		Status string `json:"status"`
	}{Status: string(enums.WaitingForMOU)}
	var transaction models.Transaction
	before := func() error {
		var temp models.Transaction
		err := providers.RepoProvider.BaseRepo.Find(ctx, id, &temp)
		if err != nil {
			return err
		}
		status := temp.Status
		if status != nil {
			if *status != string(enums.WaitingForConfirmationStatus) {
				return types.ErrUnprocessableEntity
			}
		}
		return nil
	}
	err := baseUpdateFunc(ctx, policies.MasterPolicy{}, id, &transaction, data, before)
	return err
}

func (transactionService) ConfirmTransaction(ctx context.Context, id uint) error {
	data := struct {
		Status string `json:"status"`
	}{Status: string(enums.WaitingForInitialPayment)}
	var transaction models.Transaction
	before := func() error {
		var temp models.Transaction
		err := providers.RepoProvider.BaseRepo.Find(ctx, id, &temp)
		if err != nil {
			return err
		}
		status := temp.Status
		if status != nil {
			if *status != string(enums.WaitingForConfirmationStatus) && *status != string(enums.WaitingForMOUConfirmation) {
				return types.ErrUnprocessableEntity
			}
		}
		return nil
	}
	err := baseUpdateFunc(ctx, policies.MasterPolicy{}, id, &transaction, data, before)
	return err
}
