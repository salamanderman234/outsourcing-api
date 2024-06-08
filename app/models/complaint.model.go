package models

type Complaint struct {
	Model
	ServiceUserID             *uint                    `json:"service_user_id"`
	ServiceUser               *ServiceUser             `json:"service_user,omitempty"`
	PlacementDetailEmployeeID *uint                    `json:"placement_employee_id"`
	PlacementDetailEmployee   *PlacementDetailEmployee `json:"placement_employee,omitempty"`
	EmployeeID                *uint                    `json:"employee_id"`
	Employee                  *Employee                `json:"employee,omitempty"`
	Comment                   *string                  `json:"comment"`
	Replies                   []ComplaintReply         `json:"replies"`
}

type ComplaintReply struct {
	Model
	SupervisorID *uint       `json:"supervisor_id"`
	Supervisor   *Supervisor `json:"supervisor,omitempty"`
	ComplaintID  *uint       `json:"complaint_id"`
	Complaint    *Complaint  `json:"complaint,omitempty"`
	Reply        *string     `json:"reply"`
}
