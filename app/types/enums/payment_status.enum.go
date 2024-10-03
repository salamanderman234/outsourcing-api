package enums

type PaymentStatusEnum string

var (
	PendingPayment PaymentStatusEnum = "pending"
	SuccessPayment PaymentStatusEnum = "success"
	FailedPayment  PaymentStatusEnum = "failed"
)
