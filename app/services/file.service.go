package services

import (
	"context"

	service_domains "github.com/salamanderman234/outsourcing-api/app/domains/services"
)

type fileService struct{}

func NewFileService() service_domains.FileServiceInterface {
	return &fileService{}
}

func (fileService) UploadFile(ctx context.Context, file string, path string) (string, error) {
	return "", nil
}
func (fileService) DeleteFile(ctx context.Context, path string) error {
	return nil
}
func (fileService) GetFile(ctx context.Context, path string) (string, error) {
	return "", nil
}
