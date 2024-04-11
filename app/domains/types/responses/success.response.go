package responses

type DefaultSuccessResponse struct {
	BaseResponse
	Datas      any         `json:"datas,omitempty"`
	Data       any         `json:"data,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}
