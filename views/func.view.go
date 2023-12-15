package views

import (
	"context"
	"errors"
	"mime/multipart"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/outsourcing-api/domains"
	"github.com/salamanderman234/outsourcing-api/helpers"
)

// func readMultipartForm(c echo.Context, formName string) ([]*multipart.FileHeader, error) {
// 	form, err := c.MultipartForm()
// 	if err != nil {
// 		return nil, err
// 	}
// 	files, ok := form.File[formName]
// 	if !ok {
// 		return files, err
// 	}
// 	for _, file := range files {
// 		file.Open()
// 	}
// 	return files, nil
// }

func readFile(c echo.Context, formName string) (*multipart.FileHeader, error) {
	file, err := c.FormFile(formName)
	if errors.Is(err, http.ErrMissingFile) {
		return nil, nil
	}
	if err != nil {
		return nil, domains.ErrGetMultipartFormData
	}
	return file, nil
}

type callCreateServFunc func(ctx context.Context) (any, error)

func basicCreateView(c echo.Context, data any, callFunc callCreateServFunc) error {
	ctx := c.Request().Context()
	resp := domains.BasicResponse{
		Message: "created",
	}
	if err := c.Bind(data); err != nil {
		msg := "invalid request"
		payload := domains.ErrorBodyResponse{
			Error: &msg,
		}
		resp.Message = domains.ErrBadRequest.Error()
		resp.Body = payload
		return c.JSON(http.StatusBadRequest, resp)
	}
	created, err := callFunc(ctx)
	if err != nil {
		status, msg, errBody := helpers.HandleError(err)
		resp.Message = msg
		resp.Body = *errBody
		return c.JSON(status, resp)
	}
	resp.Body = domains.DataBodyResponse{
		Data: created,
	}
	return c.JSON(http.StatusCreated, resp)
}

type callUpdateServFunc func(ctx context.Context, id uint) (int, any, error)

func basicUpdateView(c echo.Context, data any, callFunc callUpdateServFunc) error {
	ctx := c.Request().Context()
	resp := domains.BasicResponse{
		Message: "ok",
	}
	var id uint
	if err := echo.QueryParamsBinder(c).Uint("id", &id).BindError(); err != nil || id == 0 {
		msg := "id is required, must be an unsigned integer and cannot be 0"
		payload := domains.ErrorBodyResponse{
			Error: &msg,
		}
		resp.Message = domains.ErrBadRequest.Error()
		resp.Body = payload
		return c.JSON(http.StatusBadRequest, resp)
	}
	if err := c.Bind(data); err != nil {
		msg := "invalid request"
		payload := domains.ErrorBodyResponse{
			Error: &msg,
		}
		resp.Message = domains.ErrBadRequest.Error()
		resp.Body = payload
		return c.JSON(http.StatusBadRequest, resp)
	}
	aff, _, err := callFunc(ctx, id)
	if err != nil {
		status, msg, errBody := helpers.HandleError(err)
		resp.Message = msg
		resp.Body = *errBody
		return c.JSON(status, resp)
	}
	resp.Body = domains.DataBodyResponse{
		ID:       &id,
		Affected: &aff,
	}
	return c.JSON(http.StatusOK, resp)
}
