package helpers

import (
	"math"

	database_types "github.com/salamanderman234/outsourcing-api/app/domains/types/databases"
	custom_errors "github.com/salamanderman234/outsourcing-api/app/domains/types/errors"
	"github.com/salamanderman234/outsourcing-api/app/domains/types/responses"
	"github.com/salamanderman234/outsourcing-api/configs"
)

type responseHelper struct{}

func (r responseHelper) CreateResponse(con responses.ResponseConfig) (int, responses.ResponseInterface) {
	var response responses.ResponseInterface
	if con.Error != nil {
		newResponse := responses.FailResponse{}
		debugMsg := con.Error.Error()
		if configs.AppConfig.IsDebug {
			newResponse.DebugMsg = &debugMsg
		}

		err := Translator.TranslateError(con.Error)
		generalMsg, msg, details := r.ExtractError(err)

		newResponse.Msg = generalMsg
		newResponse.Detail = msg
		newResponse.Errors = details

		response = newResponse
		return err.Status, response
	}

	newResponse := responses.DefaultSuccessResponse{
		BaseResponse: responses.BaseResponse{
			Msg: con.Action.Msg,
		},
		Datas:      Struct.LoopGetVisibleStruct(con.Datas),
		Data:       Struct.GetVisibleAttributes(con.Data),
		Pagination: con.Pagination,
	}
	response = newResponse
	return con.Action.Status, response
}

func (responseHelper) ExtractError(err custom_errors.GeneralError) (string, string, []responses.FieldError) {
	return err.GeneralMessage, err.Msg, err.ValidationErrors
}

func (responseHelper) CreatePagination(max int64, config database_types.DBSearchConfig) *responses.Pagination {
	next := uint(math.Min(float64(config.Page+1), float64(max)))
	prev := uint(math.Max(float64(config.Page-1), 1))
	return &responses.Pagination{
		PreviousPage: prev,
		CurrentPage:  config.Page,
		NextPage:     next,
		Query:        config.Query,
		MaxPage:      max,
	}
}

var Response = responseHelper{}
