package helpers

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gabriel-vasile/mimetype"
	"github.com/salamanderman234/outsourcing-api/app/domains"
	"github.com/salamanderman234/outsourcing-api/app/types"
)

type fileHelper struct{}

func (fileHelper) ConvertBase64ToFile(str string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(str)
}

// path relative to storage directory
func (fi fileHelper) SaveFile(file []byte, resource domains.ResourceInterface) (string, error) {
	ext, err := fi.CheckMimeCompability(file, resource)
	if err != nil {
		return "", err
	}
	name := String.GenerateRandomString(10)
	dirPath := resource.GetFullPath()
	if _, err := os.Stat(dirPath); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			return "", err
		}
	}

	fullpath := fmt.Sprintf("%s/%s%s", dirPath, name, ext)
	f, err := os.Create(fullpath)
	if err != nil {
		return "", err
	}
	defer f.Close()
	if _, err := f.Write(file); err != nil {
		return "", err
	}
	if err := f.Sync(); err != nil {
		return "", err
	}
	return fullpath, nil
}

func (f fileHelper) SaveFromBase64(base64 string, resource domains.ResourceInterface) (string, error) {
	file, err := f.ConvertBase64ToFile(base64)
	if err != nil {
		return "", err
	}
	return f.SaveFile(file, resource)
}

func (f fileHelper) CheckMimeCompability(file []byte, resource domains.ResourceInterface) (string, error) {
	fileConfig, ok := resource.GetFileConfig().(types.FileConfig)
	if !ok {
		return "", types.ErrInternalServer
	}
	acceptedMimes := fileConfig.AcceptedMimes
	mime := mimetype.Detect(file)

	fmt.Println(acceptedMimes, mime.String())
	if !mimetype.EqualsAny(mime.String(), acceptedMimes...) {
		return "", types.ErrBadRequest
	}

	extension := mime.Extension()
	return extension, nil
}

func (f fileHelper) DeleteFile(path string) error {
	err := os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}

func (f fileHelper) GetFile(path string) ([]byte, string, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, "", err
	}
	mime := http.DetectContentType(file)
	return file, mime, nil
}

var File = fileHelper{}
