package models

type Access struct {
	Model
	AccessName  *string        `json:"access_name,omitempty"`
	Description *string        `json:"description,omitempty"`
	Actions     []AccessAction `json:"actions"`
}

type AccessAction struct {
	Model
	AccessID *uint   `json:"access_id,omitempty"`
	Access   *Access `json:"access,omitempty" gorm:"foreignKey:AccessID"`
	UrlID    *uint   `json:"url_id,omitempty"`
	Url      *Url    `json:"url,omitempty" gorm:"foreignKey:UrlID"`
}

type Url struct {
	Model
	Url    *string `json:"url,omitempty"`
	Name   *string `json:"name,omitempty"`
	Method *string `json:"method,omitempty"`
}
