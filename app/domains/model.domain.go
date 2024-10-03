package domains

import "time"

type ModelInterface interface {
	SetUpdatedEmail(email string)
	SetID(id uint)
	GetID() uint
	GetTimeStamps() (time.Time, time.Time)
}
