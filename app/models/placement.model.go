package models

import "time"

type Placement struct {
	Model
	TransactionID        *uint             `json:"transaction_id"`
	Transaction          *Transaction      `json:"transaction"`
	SupervisorID         *uint             `json:"supervisor_id"`
	Supervisor           *Supervisor       `json:"supervisor"`
	RegencyID            *uint             `json:"regency_id"`
	Regency              *Regency          `json:"regency"`
	StartDate            *time.Time        `json:"start_date"`
	EndDate              *time.Time        `json:"end_date"`
	Status               *string           `json:"status"`
	Note                 *string           `json:"note"`
	Duration             *uint             `json:"duration"`
	TotalEmployee        *uint             `json:"total_employee"`
	Name                 *string           `json:"name"`
	CompanyName          *string           `json:"company_name"`
	Address              *string           `json:"address"`
	Details              []PlacementDetail `json:"details"`
	FormGenerateSchedule *string           `json:"form_generate_schedule"`
	LastFormDate         *time.Time        `json:"last_form_date"`
	NextFormDate         *time.Time        `json:"next_form_date"`
	Forms                []PerformanceForm `json:"forms"`
}

type PlacementDetail struct {
	Model
	PlacementID   *uint                     `json:"placement_id"`
	Placement     *Placement                `json:"placement"`
	ServiceID     *uint                     `json:"service_id"`
	Service       *Service                  `json:"service"`
	TotalEmployee *uint                     `json:"total_employee"`
	Filled        *uint                     `json:"filled"`
	Employees     []PlacementDetailEmployee `json:"employees"`
	Salary        *uint                     `json:"salary"`
}

type PlacementDetailEmployee struct {
	Model
	PlacementDetailID        *uint                     `json:"placement_detail_id"`
	PlacementDetail          *PlacementDetail          `json:"placement"`
	EmployeeID               *uint                     `json:"employee_id"`
	Employee                 *Employee                 `json:"employee"`
	Status                   *string                   `json:"status"`
	PlacementDate            *time.Time                `json:"placement_date"`
	StartDate                *time.Time                `json:"start_date"`
	EndDate                  *time.Time                `json:"end_date"`
	ExitDate                 *time.Time                `json:"exit_date"`
	Duration                 *uint                     `json:"duration"`
	ExpectedSalary           *uint                     `json:"expected_salary"`
	ExpectedSalaryTotal      *uint                     `json:"expected_salary_total"`
	ActualSalary             *uint                     `json:"actual_salary"`
	Complaints               []Complaint               `json:"complaints"`
	PerformanceFormFeedbacks []PerformanceFormFeedback `json:"performance_form_feedbacks"`
}
