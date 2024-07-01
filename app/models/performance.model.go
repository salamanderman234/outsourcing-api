package models

import "time"

type PerformanceForm struct {
	Model
	PlacementID *uint                     `json:"placement_id"`
	Placement   *Placement                `json:"placement"`
	Date        *time.Time                `json:"date"`
	FilledDate  *time.Time                `json:"fiiled_date"`
	Feedbacks   []PerformanceFormFeedback `json:"feedbacks"`
}

type PerformanceFormFeedback struct {
	Model
	PerformanceFormID         *uint                    `json:"performance_form_id"`
	PerformanceForm           *PerformanceForm         `json:"performance_form_"`
	PlacementDetailEmployeeID *uint                    `json:"placement_detail_employee_id"`
	PlacementDetailEmployee   *PlacementDetailEmployee `json:"placement_detail_employee"`
	ServiceUserID             *uint                    `json:"service_user_id"`
	ServiceUser               *ServiceUser             `json:"service_user"`
	Question                  *string                  `json:"question"`
	Answer                    string                   `json:"answer"`
	Rate                      *uint                    `json:"rate"`
	Note                      *string                  `json:"note"`
	Date                      *time.Time               `json:"date"`
}

type Performance struct {
	Model
	EmployeeID        *uint            `json:"employee_id"`
	Employee          *Employee        `json:"employee"`
	PerformanceFormID *uint            `json:"performance_form_id"`
	PerformanceForm   *PerformanceForm `json:"performance_form"`
	ServiceUserID     *uint            `json:"service_user_id"`
	ServiceUser       *ServiceUser     `json:"service_user"`
	Rate              *uint            `json:"rate"`
	Date              *time.Time       `json:"date"`
}
