package configs

import "github.com/spf13/viper"

const (
	PAGINATION_PER_PAGE         = 10
	PAGINATION_ORDER_BY_DEFAULT = "updated_at"
	PAGINATION_ORDER            = "DESC"
	// token
	ACCESS_TOKEN_EXPIRE_TIME  = 4320
	REFRESH_TOKEN_EXPIRE_TIME = 4320
	REFRESH_TOKEN_COOKIE_NAME = "refresh"
	// file
	FILENAME_LENGTH      = 10
	MAXIMUM_CONTENT_SIZE = 15
	FILE_VAULT_PATH      = "/storage"
)

type FileConfig struct {
	AcceptedFileTypes []string
	MaximumFileSize   int64
	AcceptedErrMsg    string
	MaximumErrMsg     string
}

// file configuration
var (
	IMAGE_FILE_CONFIG = FileConfig{
		AcceptedFileTypes: []string{"webp", "png", "jpg", "jpeg", "svg"},
		MaximumFileSize:   524288,
		AcceptedErrMsg:    "accepted file type : png, webp, svg, jpg, and jpeg",
		MaximumErrMsg:     "maximum file size : 500kb",
	}
	PDF_FILE_CONFIG = FileConfig{
		AcceptedFileTypes: []string{"pdf"},
		MaximumFileSize:   10485760,
		AcceptedErrMsg:    "accepted file type : pdf",
		MaximumErrMsg:     "maximum file size : 10MB",
	}
)
var (
	FILE_DESTS = map[string]string{
		"partial-service/icon": "service/icon",
		"patial-service/image": "service/image",
		"category/icon":        "master/category/icon",
		"order/mou":            "order/mou",
	}
)

func GetApplicationSecret() string {
	return viper.GetString("APP_SECRET")
}

func SetConfig(url string) {
	viper.SetConfigFile(url)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
