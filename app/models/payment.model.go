package models

type Payment struct {
	Model
	TransactionID  *uint        `json:"transaction_id,omitempty" visible:"true"`
	Transaction    *Transaction `json:"transaction,omitempty" visible:"true" gorm:"foreignKey:TransactionID"`
	BillingName    *string      `json:"billing_name,omitempty" visible:"true"`
	BillingContact *string      `json:"billing_contact,omitempty" visible:"true"`
	BillingAddress *string      `json:"billing_address,omitempty" visible:"true"`
	TotalAmount    *int64       `json:"total_amount,omitempty" visible:"true"`
	Status         *string      `json:"status,omitempty" visible:"true"`
}
