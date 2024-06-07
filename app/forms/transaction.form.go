package forms

import "time"

type TransactionDetailEtcCreateForm struct {
	AdditionalItemServiceID uint `json:"additional_item_service_id" valid:"required,int"`
	Qty                     uint `json:"qty" valid:"required,int"`
}

type TransactionDetailCreateForm struct {
	ServiceID     uint                             `json:"service_id" valid:"required,int"`
	TotalEmployee uint                             `json:"total_employee" valid:"required,int"`
	Etcs          []TransactionDetailEtcCreateForm `json:"etcs"`
}

type TransactionCreateForm struct {
	ServiceUserID    uint                          `json:"service_user_id" valid:"required,int"`
	PackageID        uint                          `json:"package_id" valid:"optional,int"`
	ContractDuration uint                          `json:"contract_duration" valid:"required,int"`
	StartDate        time.Time                     `json:"start_date" valid:"required"`
	RegencyID        uint                          `json:"regency_id" valid:"required,int"`
	Address          string                        `json:"address" valid:"required,stringlength(1|3000)"`
	CompanyName      string                        `json:"company_name" valid:"required,stringlength(1|3000)"`
	BillingName      string                        `json:"billing_name" valid:"required,stringlength(1|3000)"`
	BillingAddress   string                        `json:"billing_address" valid:"required,stringlength(1|3000)"`
	PaymentMethod    string                        `json:"payment_method" valid:"required,in(full|dp|3_termin)"`
	Details          []TransactionDetailCreateForm `json:"details"`
	// UsingMOU         bool                          `json:"using_mou"`
}

type TransactionUpdateForm struct {
	StartDate      *time.Time `json:"start_date" valid:"optional"`
	RegencyID      *uint      `json:"regency_id" valid:"optional,int"`
	Address        *string    `json:"address" valid:"optional,stringlength(1|3000)"`
	CompanyName    *string    `json:"company_name" valid:"optional,stringlength(1|3000)"`
	BillingName    *string    `json:"billing_name" valid:"optional,stringlength(1|3000)"`
	BillingAddress *string    `json:"billing_address" valid:"optional,stringlength(1|3000)"`
}
