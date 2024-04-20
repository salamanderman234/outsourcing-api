package models

import "time"

type ModelInterface interface {
	GetID() uint
	GetTimeStamps() (time.Time, time.Time)
}

type Model struct {
	ID              uint       `json:"id" visible:"true"`
	CreatedAt       *time.Time `json:"created_at,omitempty"`
	UpdatedAt       *time.Time `json:"updated_at,omitempty"`
	UpdateEmailUser *string    `json:"update_email_user,omitempty"`
}

func (m Model) GetID() uint {
	return m.ID
}

func (m Model) GetTimeStamps() (time.Time, time.Time) {
	return *m.CreatedAt, *m.UpdatedAt
}

type Job struct {
	Model
	Payload    string    `json:"payload"`
	JobName    string    `json:"job_name"`
	ReservedAt time.Time `json:"reserved_at"`
	Attempts   uint      `json:"attemps" gorm:"default:0"`
}
