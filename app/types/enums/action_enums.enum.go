package enums

import "net/http"

type ActionEnum struct {
	Status int
	Msg    string
}

var (
	CreateAction ActionEnum = ActionEnum{
		Status: http.StatusCreated,
		Msg:    "created",
	}
	ReadAction ActionEnum = ActionEnum{
		Status: http.StatusOK,
		Msg:    "ok",
	}
	UpdateAction ActionEnum = ActionEnum{
		Status: http.StatusOK,
		Msg:    "updated",
	}
	DeleteAction = ActionEnum{
		Status: http.StatusOK,
		Msg:    "deleted",
	}
)
