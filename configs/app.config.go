package configs

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

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

// Application
type ContextKey string

var (
	UserKey ContextKey = "user"
)

type ContextValue struct {
	echo.Context
}

func (ctx ContextValue) Get(key string) interface{} {
	// get old context value
	val := ctx.Context.Get(key)
	if val != nil {
		return val
	}
	return ctx.Request().Context().Value(key)
}

func (ctx ContextValue) Set(key string, val interface{}) {
	contextKey := ContextKey(key)
	ctx.SetRequest(ctx.Request().WithContext(context.WithValue(ctx.Request().Context(), contextKey, val)))
}
func MiddlewareContextValue(fn echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return fn(ContextValue{ctx})
	}
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
		"user/profile":         "user/profile",
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
