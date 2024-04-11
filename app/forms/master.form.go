package forms

type MasterProviceCreateForm struct {
	Province string `json:"province" form:"province" valid:"required,stringlength(1|255)"`
}

type MasterProviceUpdateForm struct {
	Province string `json:"province" form:"province" valid:"required,stringlength(1|255)"`
}

type MasterRegencyCreateForm struct {
	ProvinceID uint   `json:"province_id" form:"province_id" valid:"required"`
	Regency    string `json:"regency" form:"regency" valid:"required,stringlength(1|255)"`
}
type MasterRegencyUpdateForm struct {
	Regency string `json:"regency" form:"regency" valid:"required,stringlength(1|255)"`
}
