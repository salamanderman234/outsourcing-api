package forms

import "time"

type FeedbackCreateForm struct {
	TransactionID uint    `json:"transaction_id" valid:"required"`
	Review        float64 `json:"review" valid:"required,range(1|5)"`
	Comment       *string `json:"comment" valid:"optional,stringlength(1|5000)"`
}

type FeedbackUpdateForm struct {
	Review  *float64 `json:"review" valid:"optional,range(1|5)"`
	Comment *string  `json:"comment" valid:"optional,stringlength(1|5000)"`
}

type PerformanceSubmitAnswerForm struct {
	PlacementDetailEmployeeID uint    `json:"placement_detail_employee_id" valid:"required"`
	Question                  string  `json:"question" valid:"required"`
	Answer                    string  `json:"answer" valid:"required,in(A|B|C|D|E)"`
	Note                      *string `json:"note" valid:"optional,stringlength(1|5000)"`
}

type PerformanceSubmitForm struct {
	PerformanceFormID uint                          `json:"performance_form_id" valid:"required"`
	Details           []PerformanceSubmitAnswerForm `json:"details" valid:"required"`
}

type PerformanceCreateForm struct {
	PlacementID uint `json:"placement_id" valid:"required"`
}

type PerformanceFilter struct {
	EmployeeID uint      `query:"employee_id"`
	From       time.Time `query:"from"`
	To         time.Time `query:"to"`
}
