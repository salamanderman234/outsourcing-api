package forms

import "time"

type PlacementCreateForm struct {
	TransactionID uint    `json:"transaction_id" valid:"required,int"`
	Note          *string `json:"note" valid:"optional,stringlength(1|30000)"`
	SupervisorID  uint    `json:"supervisor_id" valid:"required,int"`
}

type PlacementUpdateForm struct {
	Note         *string `json:"note" valid:"optional,stringlength(1|30000)"`
	SupervisorID *uint   `json:"supervisor_id" valid:"optional,int"`
	Status       *string `json:"status" valid:"optional,in(ongoing|suspend|cancel|end)"`
}

type PlacementDetailEmployeeCreateForm struct {
	PlacementDetailID uint       `json:"placement_detail_id" valid:"required,int"`
	EmployeeID        uint       `json:"employee_id" valid:"required,int"`
	StartDate         *time.Time `json:"start_date" valid:"optional"`
}

type PlacementDetailEmployeeUpdateForm struct {
	ExitDate *time.Time `json:"exit_date" valid:"optional"`
}
