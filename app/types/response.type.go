package types

import "github.com/salamanderman234/outsourcing-api/app/types/enums"

type ResponseInterface interface {
	GetMsg() string
}

type BaseResponse struct {
	Msg string `json:"message"`
}

func (b BaseResponse) GetMsg() string {
	return b.Msg
}

type Pagination struct {
	PreviousPage uint   `json:"previous_page"`
	CurrentPage  uint   `json:"current_page"`
	NextPage     uint   `json:"next_page"`
	MaxPage      int64  `json:"max_page"`
	Query        string `json:"query"`
}

type ResponseParams struct {
	Action     enums.ActionEnum
	Error      error
	Datas      any
	Data       any
	Pagination *Pagination
}

type FailResponse struct {
	BaseResponse
	Detail    string       `json:"detail"`
	Errors    []FieldError `json:"errors,omitempty"`
	DebugMsg  *string      `json:"debug,omitempty"`
	ErrorType *string      `json:"error_type,omitempty"`
}

type DefaultSuccessResponse struct {
	BaseResponse
	Datas      any         `json:"datas,omitempty"`
	Data       any         `json:"data,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}
