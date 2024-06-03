package types

type FileConfig struct {
	MaxSize       uint
	BasePath      string
	AcceptedMimes []string
}

type FileConfigKey string

var (
	ImageConfig FileConfigKey = "image"
	PDFConfig   FileConfigKey = "pdf"
)
