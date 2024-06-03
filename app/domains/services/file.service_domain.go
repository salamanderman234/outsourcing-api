package service_domains

import (
	"context"

	"github.com/salamanderman234/outsourcing-api/app/domains"
)

type FileServiceInterface interface {
	UploadFile(ctx context.Context, file string, resource domains.ResourceInterface) (string, error)
	DeleteFile(ctx context.Context, path string) error
	GetFile(ctx context.Context, id uint, resource domains.ResourceInterface) ([]byte, string, error)
}
