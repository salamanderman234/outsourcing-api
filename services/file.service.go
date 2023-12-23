package services

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
	"slices"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/salamanderman234/outsourcing-api/configs"
	"github.com/salamanderman234/outsourcing-api/domains"
	"github.com/salamanderman234/outsourcing-api/helpers"
)

type fileService struct{}

func NewFileService() domains.FileService {
	return &fileService{}
}
func (f fileService) Store(file *multipart.FileHeader, dest string, fileConfig configs.FileConfig) (string, error) {
	vaultPath := configs.FILE_VAULT_PATH
	if file.Size > fileConfig.MaximumFileSize {
		return "", domains.ErrFileSize
	}
	src, err := file.Open()
	if err != nil {
		return "", domains.ErrFileOpen
	}
	defer src.Close()
	fileName := helpers.GenerateRandomString(10)
	splittedFileName := strings.Split(file.Filename, ".")
	fileType := splittedFileName[len(splittedFileName)-1]
	valid := slices.Contains(fileConfig.AcceptedFileTypes, fileType)
	if !valid {
		return "", domains.ErrInvalidFileType
	}
	nameType := fileName + "." + fileType
	finalDest := "." + vaultPath + "/" + dest + "/"
	// Destination
	if _, err := os.Stat(finalDest); errors.Is(err, os.ErrNotExist) {
		os.MkdirAll(finalDest, os.ModePerm)
	}
	finalDest += nameType
	dst, err := os.Create(finalDest)
	if err != nil {
		return "", domains.ErrFileCreate
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return "", domains.ErrFileCopy
	}
	return finalDest, nil
}
func (f fileService) cancelOperations(filePaths map[string]string) {
	for _, filePath := range filePaths {
		f.Destroy(filePath)
	}
}
func (f fileService) BatchStore(files map[string]domains.FileWrapper) (map[string]string, string, error) {
	filePaths := map[string]string{}
	for key, file := range files {
		content := file.File
		dest := file.Dest
		config := file.Config
		filePath, err := f.Store(content, dest, config)
		if err != nil {
			go f.cancelOperations(filePaths)
			if errors.Is(err, domains.ErrInvalidFileType) {
				conv := err.(domains.GeneralError)
				conv.ValidationErrors = govalidator.Errors{
					govalidator.Error{
						Name:                     key,
						Validator:                "file type",
						CustomErrorMessageExists: true,
						Err:                      errors.New(files[key].Config.AcceptedErrMsg),
					},
				}
				return filePaths, key, conv
			} else if errors.Is(err, domains.ErrFileSize) {
				conv := err.(domains.GeneralError)
				conv.ValidationErrors = govalidator.Errors{
					govalidator.Error{
						Name:                     key,
						Validator:                "file size",
						CustomErrorMessageExists: true,
						Err:                      errors.New(files[key].Config.MaximumErrMsg),
					},
				}
				return filePaths, key, conv
			}
			return filePaths, key, err
		}
		filePaths[key] = filePath
	}
	return filePaths, "", nil
}

func (f fileService) Destroy(target string) error {
	err := os.Remove(target)
	if err != nil {
		return domains.ErrDeleteFile
	}
	return nil
}

// func(f fileService) Read(target string) (*multipart.FileHeader, error){

// }
