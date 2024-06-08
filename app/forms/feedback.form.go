package forms

type FeedbackCreateForm struct {
	TransactionID uint    `json:"transaction_id" valid:"required"`
	Review        float64 `json:"review" valid:"required,range(1|5)"`
	Comment       *string `json:"comment" valid:"optional,stringlength(1|5000)"`
}

type FeedbackUpdateForm struct {
	Review  *float64 `json:"review" valid:"optional,range(1|5)"`
	Comment *string  `json:"comment" valid:"optional,stringlength(1|5000)"`
}

type PerformanceCreateForm struct{}

type PerformanceUpdateForm struct{}

type UserPerformanceUploadForm struct{}

type UserPerformanceUpdateForm struct{}

type PerformanceUploadInputForm struct{}

type PerformanceUpdateInputForm struct{}
