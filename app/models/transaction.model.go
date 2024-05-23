package models

import "time"

type Transaction struct {
	Model
	ServiceUserID       *uint               `json:"service_user_id,omitempty" visible:"true"`
	ServiceUser         *ServiceUser        `json:"service_user,omitempty" gorm:"foreignKey:ServiceUserID" visible:"true"`
	PackageID           *uint               `json:"package_id,omitempty" visible:"true"`
	Package             *Package            `json:"package,omitempty" gorm:"foreignKey:PackageID" visible:"true"`
	OrderDate           *time.Time          `json:"order_date,omitempty" visible:"true"`
	ContractDuration    *uint               `json:"contract_duration,omitempty" visible:"true"`
	StartDate           *time.Time          `json:"start_date,omitempty" visible:"true"`
	RegencyID           *uint               `json:"regency_id,omitempty" visible:"true"`
	Regency             *Regency            `json:"regency,omitempty" gorm:"foreignKey:RegencyID" visible:"true"`
	Address             *string             `json:"address,omitempty" visible:"true"`
	CompanyName         *string             `json:"company_name,omitempty" visible:"true"`
	BillingName         *string             `json:"billing_name,omitempty" visible:"true"`
	BillingAddress      *string             `json:"billing_address,omitempty" visible:"true"`
	PaymentMethod       *string             `json:"payment_method,omitempty" visible:"true"`
	TotalPrice          *uint64             `json:"total_price,omitempty" visible:"true"`
	TotalPaid           *uint64             `json:"total_paid,omitempty" visible:"true"`
	Status              *string             `json:"status,omitempty" visible:"true"`
	NextPaymentDeadline *time.Time          `json:"next_payment_deadline" visible:"true"`
	Details             []TransactionDetail `json:"details,omitempty" visible:"true"`
	MOU                 *string             `json:"mou,omitempty"`
}

type TransactionDetail struct {
	Model
	TransactionID *uint                  `json:"transaction_id,omitempty" visible:"true"`
	Transaction   *Transaction           `json:"transaction,omitempty" gorm:"foreignKey:TransactionID" visible:"true"`
	ServiceID     *uint                  `json:"service_id,omitempty" visible:"true"`
	Service       *Service               `json:"service,omitempty" gorm:"foreignKey:ServiceID" visible:"true"`
	TotalEmployee *uint                  `json:"total_employee,omitempty" visible:"true"`
	ServicePrice  *uint64                `json:"service_price,omitempty" visible:"true"`
	EmployeePrice *uint64                `json:"employee_price,omitempty" visible:"true"`
	EtcPrice      *uint64                `json:"etc_price,omitempty" visible:"true"`
	SubTotalPrice *uint64                `json:"sub_total_price,omitempty" visible:"true"`
	Etcs          []TransactionDetailEtc `json:"etcs,omitempty" visible:"true"`
}

type TransactionDetailEtc struct {
	Model
	TransactionDetailID     *uint                  `json:"transaction_detail_id,omitempty" visible:"true"`
	TransactionDetail       *TransactionDetail     `json:"transaction_detail,omitempty" gorm:"foreignKey:TransactionDetailID" visible:"true"`
	AdditionalItemServiceID *uint                  `json:"additional_item_service_id,omitempty" visible:"true"`
	AdditionalItemService   *AdditionalItemService `json:"additional_item_service,omitempty" gorm:"foreignKey:AdditionalItemServiceID" visible:"true"`
	Qty                     *uint                  `json:"qty,omitempty" visible:"true"`
	Price                   *uint64                `json:"price,omitempty" visible:"true"`
	SubTotalPrice           *uint64                `json:"sub_total_price,omitempty" visible:"true"`
}
