package models

type Regency struct {
	Model
	ProvinceID *uint     `json:"province_id,omitempty" visible:"true"`
	Province   *Province `json:"province,omitempty" visible:"true" gorm:"foreignKey:ProvinceID"`
	Regency    *string   `json:"regency,omitempty" visible:"true"`
}

type Province struct {
	Model
	Province *string `json:"province,omitempty" visible:"true"`
}

type Category struct {
	Model
	CategoryName string  `json:"category_name" visible:"true"`
	Icon         *string `json:"icon" visible:"true"`
	Description  string  `json:"description" visible:"true"`
}

type Question struct {
	Model
	Question string `json:"question"`
	Hint     string `json:"hint"`
}

type CategoryQuestion struct {
	Model
	CategoryID *uint     `json:"category_id"`
	Category   *Category `json:"category"`
	QuestionID *uint     `json:"question_id"`
	Question   *Question `json:"question"`
}

type PaymentConfig struct {
	Model
	Type    *string  `json:"type"`
	SubType *string  `json:"sub_type"`
	Amount  *float32 `json:"amount"`
}
