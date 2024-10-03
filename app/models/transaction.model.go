package models

import "time"

type Transaction struct {
	Model
	ServiceUserID       *uint               `json:"service_user_id" visible:"true"`
	ServiceUser         *ServiceUser        `json:"service_user" gorm:"foreignKey:ServiceUserID" visible:"true"`
	PackageID           *uint               `json:"package_id" visible:"true"`
	Package             *Package            `json:"package" gorm:"foreignKey:PackageID" visible:"true"`
	OrderDate           *time.Time          `json:"order_date" visible:"true"`
	ContractDuration    *uint               `json:"contract_duration" visible:"true"`
	StartDate           *time.Time          `json:"start_date" visible:"true"`
	RegencyID           *uint               `json:"regency_id" visible:"true"`
	Regency             *Regency            `json:"regency" gorm:"foreignKey:RegencyID" visible:"true"`
	Address             *string             `json:"address" visible:"true"`
	CompanyName         *string             `json:"company_name" visible:"true"`
	BillingName         *string             `json:"billing_name" visible:"true"`
	BillingAddress      *string             `json:"billing_address" visible:"true"`
	PaymentMethod       *string             `json:"payment_method" visible:"true"`
	TotalPrice          *uint64             `json:"total_price" visible:"true"`
	TotalPaid           *uint64             `json:"total_paid" visible:"true"`
	Status              *string             `json:"status" visible:"true"`
	DPStatus            *string             `json:"dp_status"`
	TerminStatus        *string             `json:"termin_status"`
	NextPaymentDeadline *time.Time          `json:"next_payment_deadline" visible:"true"`
	Details             []TransactionDetail `json:"details" visible:"true"`
	MOU                 *string             `json:"mou"`
	Placements          []Placement         `json:"placements" visible:"true"`
	Payments            []Payment           `json:"payments"`
}

type TransactionDetail struct {
	Model
	TransactionID *uint                  `json:"transaction_id" visible:"true"`
	Transaction   *Transaction           `json:"transaction" gorm:"foreignKey:TransactionID" visible:"true"`
	ServiceID     *uint                  `json:"service_id" visible:"true"`
	Service       *Service               `json:"service" visible:"true"`
	TotalEmployee *uint                  `json:"total_employee" visible:"true"`
	ServicePrice  *uint64                `json:"service_price" visible:"true"`
	EmployeePrice *uint64                `json:"employee_price" visible:"true"`
	EtcPrice      *uint64                `json:"etc_price" visible:"true"`
	SubTotalPrice *uint64                `json:"sub_total_price" visible:"true"`
	Etcs          []TransactionDetailEtc `json:"etcs" visible:"true"`
}

type TransactionDetailEtc struct {
	Model
	TransactionDetailID     *uint                  `json:"transaction_detail_id" visible:"true"`
	TransactionDetail       *TransactionDetail     `json:"transaction_detail" gorm:"foreignKey:TransactionDetailID" visible:"true"`
	AdditionalItemServiceID *uint                  `json:"additional_item_service_id" visible:"true"`
	AdditionalItemService   *AdditionalItemService `json:"additional_item_service" gorm:"foreignKey:AdditionalItemServiceID" visible:"true"`
	Qty                     *uint                  `json:"qty" visible:"true"`
	Price                   *uint64                `json:"price" visible:"true"`
	SubTotalPrice           *uint64                `json:"sub_total_price" visible:"true"`
}
