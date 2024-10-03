package forms

type FileUploadForm struct {
	File string `json:"file" valid:"required"`
}
