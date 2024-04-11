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
	CategoryName string `json:"category_name" visible:"true"`
	Icon         string `json:"icon" visible:"true"`
	Description  string `json:"description" visible:"true"`
}
