package models

type Feedback struct {
	Model
	ServiceUserID *uint        `json:"service_user_id"`
	ServiceUser   *ServiceUser `json:"service_user,omitempty"`
	TransactionID *uint        `json:"transaction_id"`
	Transaction   *Transaction `json:"transaction,omitempty"`
	Review        *float64     `json:"review"`
	Comment       *string      `json:"comment"`
}
