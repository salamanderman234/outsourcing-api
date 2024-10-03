package forms

type MasterProviceCreateForm struct {
	Province string `json:"province" form:"province" valid:"required,stringlength(1|255)"`
}

type MasterProviceUpdateForm struct {
	Province *string `json:"province" form:"province" valid:"optional,stringlength(1|255)"`
}

type MasterRegencyCreateForm struct {
	ProvinceID uint   `json:"province_id" form:"province_id" valid:"required,int"`
	Regency    string `json:"regency" form:"regency" valid:"required,stringlength(1|255)"`
}
type MasterRegencyUpdateForm struct {
	Regency *string `json:"regency" form:"regency" valid:"optional,stringlength(1|255)"`
}

type MasterCategoryCreateForm struct {
	CategoryName string `json:"category_name" valid:"required,stringlength(1|255)"`
	Icon         string `json:"icon"`
	Description  string `json:"description" valid:"required,stringlength(1|3000)"`
}

type MasterCategoryUpdateForm struct {
	CategoryName *string `json:"category_name" valid:"optional,stringlength(1|255)"`
	Icon         *string `json:"icon"`
	Description  *string `json:"description" valid:"optional,stringlength(1|3000)"`
}

type MasterQuestionCreateForm struct {
	Question string  `json:"question" valid:"required,stringlength(1|5000)"`
	Hint     *string `json:"hint" valid:"optional,stringlength(1|5000)"`
}
type MasterQuestionUpdateForm struct {
	Question *string `json:"question" valid:"optional,stringlength(1|5000)"`
	Hint     *string `json:"hint" valid:"optional,stringlength(1|5000)"`
}

type MasterQuestionAssignForm struct {
	QuestionID uint `json:"question_id" valid:"required"`
	CategoryID uint `json:"category_id" valid:"required"`
}
type MasterQuestionUnassignForm struct {
	QuestionID uint `json:"question_id" valid:"required"`
	CategoryID uint `json:"category_id" valid:"required"`
}

type MaterPaymentConfigUpdateForm struct {
	Amount uint `json:"amount" valid:"required"`
}
