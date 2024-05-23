package service_domains

import "context"

type FileServiceInterface interface {
	UploadFile(ctx context.Context, file string, path string) (string, error)
	DeleteFile(ctx context.Context, path string) error
	GetFile(ctx context.Context, path string) (string, error)
}
