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
	polResult := policies.PaymentPolicy{}.Update(&order, claims)
	if !polResult {
		helpers.Logger.Warning(
			fmt.Sprintf("(Forbidden) User: %s", claims.Email),
		)
		return "", "", types.ErrForbiden
	}
	status := *order.Status
	if !slices.Contains([]string{
		string(enums.WaitingForInitialPayment), string(enums.WaitingForFurtherPayments),
	}, status) {
		return "", "", types.ErrForbiden
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
	paymentMethod := *order.PaymentMethod
	userClient := order.ServiceUser

	if paymentMethod == string(enums.DpPayment) {
		if paid == 0 && status == string(enums.WaitingForInitialPayment) {
			percentageAmount := float64(20)
			percentage := int64(percentageAmount / float64(100))
			totalAmount *= percentage
		} else {
			totalAmount -= int64(paid)
		}
	} else if paymentMethod == string(enums.ThreeTermin) {
		float20 := float64(20)
		float50 := float64(50)
		float70 := float64(70)
		if paid == 0 && status == string(enums.WaitingForInitialPayment) {
			percentage := int64(float20 / float64(100))
			totalAmount *= percentage
		} else if int64(paid) == (totalAmount*(int64(float20)/100)) && status == string(enums.WaitingForFurtherPayments) {
			percentage := int64(float50 / float64(100))
			totalAmount *= percentage
		} else if int64(paid) == (totalAmount*(int64(float70)/100)) && status == string(enums.WaitingForFurtherPayments) {
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
	transaction := &models.Transaction{}

	if slices.Contains([]enums.MidtransStatusEnum{
		enums.AuthorizeMidtransStatus,
		enums.SettlementMidtransStatus,
		enums.CaptureMidtranstatus,
	}, enums.MidtransStatusEnum(status)) {

		totalPaid := *payment.Transaction.TotalPaid
		totalPaid += uint64(*payment.TotalAmount)
		payment.Transaction.TotalPaid = &totalPaid

		paymentStatus := string(enums.SuccessPayment)
		transactionStatus := string(enums.WaitingForPlacement)
		tranStatus := payment.Transaction.Status
		if tranStatus != nil {
			if *tranStatus == string(enums.WaitingForFurtherPayments) {
				transactionStatus = string(enums.Ongoing)
			}
		}

		payment.Status = &paymentStatus
		payment.Transaction.Status = &transactionStatus

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
