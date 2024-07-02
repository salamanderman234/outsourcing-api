package models

import "time"

type Placement struct {
	Model
	TransactionID        *uint             `json:"transaction_id,omitempty"`
	Transaction          *Transaction      `json:"transaction,omitempty"`
	SupervisorID         *uint             `json:"supervisor_id,omitempty"`
	Supervisor           *Supervisor       `json:"supervisor,omitempty"`
	RegencyID            *uint             `json:"regency_id,omitempty"`
	Regency              *Regency          `json:"regency,omitempty"`
	StartDate            *time.Time        `json:"start_date,omitempty"`
	EndDate              *time.Time        `json:"end_date,omitempty"`
	Status               *string           `json:"status,omitempty"`
	Note                 *string           `json:"note,omitempty"`
	Duration             *uint             `json:"duration,omitempty"`
	TotalEmployee        *uint             `json:"total_employee,omitempty"`
	Name                 *string           `json:"name,omitempty"`
	CompanyName          *string           `json:"company_name,omitempty"`
	Address              *string           `json:"address,omitempty"`
	Details              []PlacementDetail `json:"details,omitempty"`
	FormGenerateSchedule *string           `json:"form_generate_schedule"`
	LastFormDate         *time.Time        `json:"last_form_date"`
	NextFormDate         *time.Time        `json:"next_form_date"`
	Forms                []PerformanceForm `json:"forms"`
}

type PlacementDetail struct {
	Model
	PlacementID   *uint                     `json:"placement_id,omitempty"`
	Placement     *Placement                `json:"placement,omitempty"`
	ServiceID     *uint                     `json:"service_id,omitempty"`
	Service       *Service                  `json:"service,omitempty"`
	TotalEmployee *uint                     `json:"total_employee,omitempty"`
	Filled        *uint                     `json:"filled,omitempty"`
	Employees     []PlacementDetailEmployee `json:"employees,omitempty"`
	Salary        *uint                     `json:"salary,omitempty"`
}

type PlacementDetailEmployee struct {
	Model
	PlacementDetailID        *uint                     `json:"placement_detail_id,omitempty"`
	PlacementDetail          *PlacementDetail          `json:"placement,omitempty"`
	EmployeeID               *uint                     `json:"employee_id,omitempty"`
	Employee                 *Employee                 `json:"employee,omitempty"`
	Status                   *string                   `json:"status,omitempty"`
	PlacementDate            *time.Time                `json:"placement_date,omitempty"`
	StartDate                *time.Time                `json:"start_date,omitempty"`
	EndDate                  *time.Time                `json:"end_date,omitempty"`
	ExitDate                 *time.Time                `json:"exit_date,omitempty"`
	Duration                 *uint                     `json:"duration,omitempty"`
	ExpectedSalary           *uint                     `json:"expected_salary,omitempty"`
	ExpectedSalaryTotal      *uint                     `json:"expected_salary_total,omitempty"`
	ActualSalary             *uint                     `json:"actual_salary,omitempty"`
	Complaints               []Complaint               `json:"complaints"`
	PerformanceFormFeedbacks []PerformanceFormFeedback `json:"performance_form_feedbacks"`
}
