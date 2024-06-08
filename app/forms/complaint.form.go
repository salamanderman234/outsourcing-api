package forms

type ComplaintCreateForm struct {
	PlacementDetailEmployeeID uint   `json:"placement_employee_id" valid:"required"`
	Comment                   string `json:"comment" valid:"required,stringlength(1|10000)"`
}

type ComplaintUpdateForm struct {
	Comment *string `json:"comment" valid:"optional,stringlength(1|10000)"`
}

type ReplyComplaintCreateForm struct {
	ComplaintID uint   `json:"complaint_id" valid:"required"`
	Reply       string `json:"reply" valid:"required,stringlength(1|10000)"`
}

type ReplyComplaintUpdateForm struct {
	Reply *string `json:"reply" valid:"optional,stringlength(1|10000)"`
}
