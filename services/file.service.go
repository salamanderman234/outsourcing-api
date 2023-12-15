package services

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strings"

	"github.com/salamanderman234/outsourcing-api/configs"
	"github.com/salamanderman234/outsourcing-api/domains"
	"github.com/salamanderman234/outsourcing-api/helpers"
)

type fileService struct{}

func NewFileService() domains.FileService {
	return &fileService{}
}
func (f fileService) Store(file *multipart.FileHeader, dest string) (string, error) {
	vaultPath := configs.FILE_VAULT_PATH
	src, err := file.Open()
	if err != nil {
		return "", domains.ErrFileOpen
	}
	defer src.Close()
	fileName := helpers.GenerateRandomString(10)
	splittedFileName := strings.Split(file.Filename, ".")
	fileType := splittedFileName[len(splittedFileName)-1]
	nameType := fileName + "." + fileType
	finalDest := "." + vaultPath + "/" + dest + "/"
	// Destination
	if _, err := os.Stat(finalDest); errors.Is(err, os.ErrNotExist) {
		os.MkdirAll(finalDest, os.ModePerm)
	}
	finalDest += nameType
	dst, err := os.Create(finalDest)
	if err != nil {
		fmt.Println(err)
		return "", domains.ErrFileCreate
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return "", domains.ErrFileCopy
	}
	return finalDest, nil
}
func (f fileService) BatchStore(files []*multipart.FileHeader, dest string) ([]string, error) {
	filePaths := []string{}
	for _, file := range files {
		filePath, err := f.Store(file, dest)
		if err != nil {
			return nil, err
		}
		filePaths = append(filePaths, filePath)
	}
	return filePaths, nil
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
