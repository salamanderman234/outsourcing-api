package service_domains

import (
	"context"

	"github.com/salamanderman234/outsourcing-api/app/forms"
)

type MidtransServiceInterface interface {
	CreatePaymentFromTransaction(ctx context.Context, orderID uint) (string, string, error)
	AfterPaymentAction(ctx context.Context, data forms.PaymentNotifForm) error
}
