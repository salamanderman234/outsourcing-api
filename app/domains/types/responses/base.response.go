package responses

import "github.com/salamanderman234/outsourcing-api/app/domains/types/enums"

type ResponseInterface interface {
	GetMsg() string
}

type BaseResponse struct {
	Msg string `json:"message"`
}

func (b BaseResponse) GetMsg() string {
	return b.Msg
}

type ResponseConfig struct {
	Action     enums.ActionEnum
	Error      error
	Datas      any
	Data       any
	Pagination *Pagination
}
