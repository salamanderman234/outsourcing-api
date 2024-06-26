package services

import (
	"context"
	"crypto"
	"encoding/hex"
	"fmt"
	"slices"
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

type midtransService struct{}

func NewMidtransService() service_domains.MidtransServiceInterface {
	return &midtransService{}
}

func (midtransService) CreatePaymentFromTransaction(ctx context.Context, orderID uint) (string, string, error) {
	var order models.Transaction
	err := providers.RepoProvider.BaseRepo.Find(
		ctx,
		orderID,
		&order,
		"Details",
		"ServiceUser",
		"ServiceUser.User",
		"Details.Etcs",
		"Details.Service",
		"Details.Etcs.AdditionalItemService",
	)
	if err != nil {
		return "", "", err
	}
	claims, _ := ctx.Value(configs.VarConfig.UserContextName).(types.JWTCLaims)
	polResult := policies.TransactionPolicy{}.Update(&order, claims)
	if !polResult {
		helpers.Logger.Warning(
			fmt.Sprintf("(Forbidden) User: %s", claims.Email),
		)
		return "", "", types.ErrForbiden
	}

	if order.Status == nil || order.PaymentMethod == nil {
		return "", "", types.ErrUnprocessableEntity.SetCustomMsg(
			"the transaction does not meet the criteria for using this service",
		)
	}
	paymentMethod := *order.PaymentMethod
	status := *order.Status

	if (paymentMethod == string(enums.DpPayment) && order.DPStatus == nil) ||
		(paymentMethod == string(enums.ThreeTermin) && order.TerminStatus == nil) {

		return "", "", types.ErrUnprocessableEntity.SetCustomMsg(
			"the transaction does not meet the criteria for using this service",
		)
	}

	if !slices.Contains([]string{
		string(enums.WaitingForInitialPayment), string(enums.WaitingForFurtherPayments),
	}, status) {
		return "", "", types.ErrUnprocessableEntity.SetCustomMsg(
			"the transaction does not meet the criteria for using this service",
		)
	}
	items := order.Details
	if order.TotalPrice == nil {
		zero := uint64(0)
		order.TotalPrice = &zero
	}
	if order.TotalPaid == nil {
		zero := uint64(0)
		order.TotalPaid = &zero
	}
	totalAmount := int64(*order.TotalPrice)
	paid := *order.TotalPaid
	userClient := order.ServiceUser
	// buat sesuai dengan status dp dan termin yang baru
	if paymentMethod == string(enums.DpPayment) {
		if order.DPStatus == nil {
			return "", "", types.ErrUnprocessableEntity.SetCustomMsg(
				"the transaction does not meet the criteria for using this service",
			)
		}
		dpStatus := *order.DPStatus
		if dpStatus == string(enums.WaitingForDP) && status == string(enums.WaitingForInitialPayment) {
			percentageAmount := float64(20)
			paymentConfig := models.PaymentConfig{}
			providers.RepoProvider.BaseRepo.FindWhere(ctx, map[string]any{
				"type":     enums.DpPayment,
				"sub_type": string(enums.DPSubType),
			}, &paymentConfig)
			if paymentConfig.Amount != nil {
				percentageAmount = float64(*paymentConfig.Amount)
			}
			percentage := int64(percentageAmount / float64(100))
			totalAmount *= percentage
		} else {
			totalAmount -= int64(paid)
		}
	} else if paymentMethod == string(enums.ThreeTermin) {
		if order.TerminStatus == nil {
			return "", "", types.ErrUnprocessableEntity.SetCustomMsg(
				"the transaction does not meet the criteria for using this service",
			)
		}
		terminStatus := *order.TerminStatus
		float20 := float32(20)
		float50 := float32(50)
		paymentConfig := models.PaymentConfig{}
		providers.RepoProvider.BaseRepo.FindWhere(ctx, map[string]any{
			"type":     enums.ThreeTermin,
			"sub_type": string(enums.ThreeTerminFirstSubType),
		}, &paymentConfig)
		if paymentConfig.Amount != nil {
			float20 = *paymentConfig.Amount
		}
		paymentConfigSecond := models.PaymentConfig{}
		providers.RepoProvider.BaseRepo.FindWhere(ctx, map[string]any{
			"type":     enums.ThreeTermin,
			"sub_type": string(enums.ThreeTerminSecondSubType),
		}, &paymentConfigSecond)
		if paymentConfigSecond.Amount != nil {
			float50 = *paymentConfigSecond.Amount
		}
		if terminStatus == string(enums.WaitingForFirstTermin) && status == string(enums.WaitingForInitialPayment) {
			percentage := int64(float20 / float32(100))
			totalAmount *= percentage
		} else if terminStatus == string(enums.WaitingForSecondTermin) && status == string(enums.WaitingForFurtherPayments) {
			percentage := int64(float50 / float32(100))
			totalAmount *= percentage
		} else if terminStatus == string(enums.WaitingForThirdTermin) && status == string(enums.WaitingForFurtherPayments) {
			totalAmount -= int64(paid)
		}
	}
	if userClient == nil {
		return "", "", types.ErrForbiden
	}
	pending := string(enums.PendingPayment)
	payment := models.Payment{
		TransactionID:  &order.ID,
		BillingName:    userClient.Fullname,
		BillingContact: userClient.Phone,
		BillingAddress: userClient.FullAddress,
		TotalAmount:    &totalAmount,
		Status:         &pending,
	}

	err = providers.RepoProvider.BaseRepo.Create(ctx, &payment)
	if err != nil {
		return "", "", err
	}
	idStr := strconv.Itoa(int(payment.ID))
	resp, errMid := helpers.Midtrans.CreatePayment(idStr, totalAmount, items, *userClient)
	if errMid != nil {
		return "", "", errMid
	}
	token := resp.Token
	redir := resp.RedirectURL
	return token, redir, nil
}
func (midtransService) AfterPaymentAction(ctx context.Context, form forms.PaymentNotifForm) error {
	if err := helpers.Validator.Validate(form); err != nil {
		return types.ErrForbiden
	}
	signature := form.SignatureKey
	validate := []byte(fmt.Sprintf("%s%s%s%s", form.OrderID, form.StatusCode, form.GrossAmount, providers.MidtransClient.ServerKey))
	hasher := crypto.SHA512.New()
	hasher.Write(validate)
	validateStr := hex.EncodeToString(hasher.Sum(nil))

	if signature != validateStr {
		return types.ErrForbiden
	}

	paymentID, _ := strconv.Atoi(form.OrderID)
	var payment models.Payment

	err := providers.RepoProvider.BaseRepo.Find(ctx, uint(paymentID), &payment, "Transaction")
	if err != nil {
		return err
	}

	status := form.TransactionStatus
	if payment.Transaction == nil {
		return types.ErrUnprocessableEntity.SetCustomMsg(
			"invalid payment",
		)
	}
	transaction := payment.Transaction

	if slices.Contains([]enums.MidtransStatusEnum{
		enums.AuthorizeMidtransStatus,
		enums.SettlementMidtransStatus,
		enums.CaptureMidtranstatus,
	}, enums.MidtransStatusEnum(status)) {

		totalPaid := *transaction.TotalPaid
		totalPaid += uint64(*payment.TotalAmount)
		transaction.TotalPaid = &totalPaid

		paymentStatus := string(enums.SuccessPayment)
		transactionStatus := string(enums.WaitingForPlacement)
		tranStatus := transaction.Status
		if tranStatus != nil {
			if *tranStatus == string(enums.WaitingForFurtherPayments) {
				transactionStatus = string(enums.Ongoing)
			}
		}
		if *tranStatus == string(enums.DpPayment) {
			dpStatus := transaction.DPStatus
			if *dpStatus == string(enums.WaitingForDP) {
				stat := string(enums.DPCompleted)
				transaction.DPStatus = &stat
			} else if *dpStatus == string(enums.WaitingForRemainingDP) {
				stat := string(enums.DPRemainingCompleted)
				transaction.DPStatus = &stat
			}
		}
		if *tranStatus == string(enums.ThreeTermin) {
			terminStatus := transaction.TerminStatus
			if *terminStatus == string(enums.WaitingForFirstTermin) {
				stat := string(enums.FirstTerminCompleted)
				transaction.TerminStatus = &stat
			} else if *terminStatus == string(enums.WaitingForSecondTermin) {
				stat := string(enums.SecondTerminCompleted)
				transaction.TerminStatus = &stat
			} else if *terminStatus == string(enums.WaitingForThirdTermin) {
				stat := string(enums.ThirdTerminCompleted)
				transaction.TerminStatus = &stat
			}
		}

		payment.Status = &paymentStatus
		transaction.Status = &transactionStatus

		transaction = payment.Transaction
	} else if enums.MidtransStatusEnum(status) == enums.PendingMidtransStatus {
		paymentStatus := string(enums.PendingMidtransStatus)
		payment.Status = &paymentStatus
	} else if slices.Contains([]enums.MidtransStatusEnum{
		enums.DenyMidtransStatus,
		enums.CancelMidtransStatus,
		enums.ExpireMidtransStatus,
		enums.FailureMidtransStatus,
	}, enums.MidtransStatusEnum(status)) {
		paymentStatus := string(enums.FailedPayment)
		payment.Status = &paymentStatus
	}

	err = providers.RepoProvider.BaseRepo.Update(ctx, []uint{uint(paymentID)}, &payment)

	if err != nil {
		return err
	}

	if transaction != nil {
		transactionID := *payment.TransactionID
		err = providers.RepoProvider.BaseRepo.Update(ctx, []uint{transactionID}, transaction)
		if err != nil {
			return err
		}
	}
	return nil
}
