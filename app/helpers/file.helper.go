package helpers

// import (
// 	"encoding/base64"
// 	"fmt"
// 	"mime"
// 	"os"

// 	"github.com/salamanderman234/outsourcing-api/configs"
// )

// type fileHelper struct{}

// func (fileHelper) ConvertBase64ToFile(str string) ([]byte, error) {
// 	return base64.StdEncoding.DecodeString(str)
// }

// // path relative to storage directory
// func (fileHelper) SaveFile(file []byte, mime string, resource configs.Resource) (string, error) {
// 	name := String.GenerateRandomString(7)
// 	dirPath := fmt.Sprintf("%s%s", resource.Config.BasePath, resource.Path)
// 	fullpath := fmt.Sprintf("%s/%s.%s", dirPath, name, mime)
// 	f, err := os.Create(fullpath)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer f.Close()
// 	if _, err := f.Write(file); err != nil {
// 		return "", err
// 	}
// 	if err := f.Sync(); err != nil {
// 		return "", err
// 	}
// 	return fullpath, nil
// }

// func (f fileHelper) SaveFromBase64(base64 string, resource configs.Resource) (string, error) {

// 	file, err := f.ConvertBase64ToFile(base64)
// 	if err != nil {
// 		return "", err
// 	}

// }

// func (f fileHelper) CheckMimeCompability(file string, resource configs.Resource) error {
// 	acceptedMimes := resource.Config.AcceptedMimes
// 	mime.ExtensionsByType()
// }
// var File = fileHelper{}
