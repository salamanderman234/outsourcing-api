package helpers

import (
	"fmt"
	"math"
	"reflect"

	"github.com/salamanderman234/outsourcing-api/app/types"
	"github.com/salamanderman234/outsourcing-api/configs"
)

type responseHelper struct{}

func (r responseHelper) CreateResponse(con types.ResponseParams) (int, types.ResponseInterface) {
	var response types.ResponseInterface
	if con.Error != nil {
		newResponse := types.FailResponse{}
		debugMsg := con.Error.Error()
		if configs.AppConfig.IsDebug {
			newResponse.DebugMsg = &debugMsg
			errorType := fmt.Sprint(reflect.TypeOf(con.Error))
			newResponse.ErrorType = &errorType
		}

		err := Translator.TranslateError(con.Error)
		generalMsg, msg, details := r.ExtractError(err)

		newResponse.Msg = generalMsg
		newResponse.Detail = msg
		newResponse.Errors = details

		Logger.Error(fmt.Sprintf("(%s) %s", generalMsg, con.Error.Error()))

		response = newResponse
		return err.Status, response
	}

	newResponse := types.DefaultSuccessResponse{
		BaseResponse: types.BaseResponse{
			Msg: con.Action.Msg,
		},
		// Datas:      Struct.LoopGetVisibleStruct(con.Datas),
		// Data:       Struct.GetVisibleAttributes(con.Data),
		Datas:      con.Datas,
		Data:       con.Data,
		Pagination: con.Pagination,
	}
	response = newResponse
	return con.Action.Status, response
}

func (responseHelper) ExtractError(err types.GeneralError) (string, string, []types.FieldError) {
	return err.GeneralMessage, err.Msg, err.ValidationErrors
}

func (responseHelper) CreatePagination(max int64, config types.DBSearchParams) *types.Pagination {
	next := uint(math.Min(float64(config.Page+1), float64(max)))
	prev := uint(math.Max(float64(config.Page-1), 1))
	return &types.Pagination{
		PreviousPage: prev,
		CurrentPage:  config.Page,
		NextPage:     next,
		Query:        config.Query,
		MaxPage:      max,
	}
}

var Response = responseHelper{}
