package services

import (
	"context"
	"fmt"

	"github.com/salamanderman234/outsourcing-api/app/domains"
	service_domains "github.com/salamanderman234/outsourcing-api/app/domains/services"
	"github.com/salamanderman234/outsourcing-api/app/helpers"
	"github.com/salamanderman234/outsourcing-api/app/policies"
	"github.com/salamanderman234/outsourcing-api/app/providers"
	"github.com/salamanderman234/outsourcing-api/app/types"
	"github.com/salamanderman234/outsourcing-api/configs"
)

type fileService struct{}

func NewFileService() service_domains.FileServiceInterface {
	return &fileService{}
}

func (fileService) UploadFile(
	ctx context.Context,
	file string,
	resource domains.ResourceInterface,
) (string, error) {
	pol, ok := resource.GetPolicy().(policies.Policy)
	if !ok {
		return "", types.ErrForbiden
	}
	claims, _ := ctx.Value(configs.VarConfig.UserContextName).(types.JWTCLaims)
	if !pol.UploadFile(resource.GetData(), claims) {
		helpers.Logger.Warning(
			fmt.Sprintf("(Forbidden) User: %s", claims.Email),
		)
		return "", types.ErrForbiden
	}
	return helpers.File.SaveFromBase64(file, resource)
}
func (fileService) DeleteFile(ctx context.Context, path string) error {
	return helpers.File.DeleteFile(path)
}
func (fileService) GetFile(
	ctx context.Context,
	id uint,
	resource domains.ResourceInterface,
) ([]byte, string, error) {
	err := providers.RepoProvider.BaseRepo.Find(ctx, id, resource.GetData())
	if err != nil {
		return nil, "", err
	}
	pol, ok := resource.GetPolicy().(policies.Policy)
	if !ok {
		return nil, "", types.ErrForbiden
	}
	claims, _ := ctx.Value(configs.VarConfig.UserContextName).(types.JWTCLaims)
	if !pol.UploadFile(resource.GetData(), claims) {
		helpers.Logger.Warning(
			fmt.Sprintf("(Forbidden) User: %s", claims.Email),
		)
		return nil, "", types.ErrForbiden
	}
	path := resource.GetFieldValue()
	if path == "" {
		return nil, "", types.ErrRecordNotFound
	}
	return helpers.File.GetFile(path)
}
