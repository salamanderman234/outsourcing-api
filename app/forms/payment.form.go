package forms

type PaymentNotifForm struct {
	// TransactionTime   time.Time `json:"transaction_time" valid:"required"`
	OrderID           string `json:"order_id" valid:"required"`
	PaymentType       string `json:"payment_type" valid:"required"`
	GrossAmount       string `json:"gross_amount" valid:"required"`
	SignatureKey      string `json:"signature_key" valid:"required"`
	StatusCode        string `json:"status_code" valid:"required"`
	FraudStatus       string `json:"fraud_status" valid:"required"`
	TransactionStatus string `json:"transaction_status" valid:"required"`
}
