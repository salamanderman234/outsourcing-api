package services

import (
	"context"

	"github.com/salamanderman234/outsourcing-api/domains"
	"github.com/salamanderman234/outsourcing-api/helpers"
)

// ----- USER SERVICE -----
type userService struct{}

func NewUserService() domains.UserService {
	return &userService{}
}
func (userService) Find(c context.Context, id uint) (domains.UserEntity, error) {
	var userEntity domains.UserEntity
	result, err := domains.RepoRegistry.UserRepo.Find(c, id)
	if err != nil {
		return userEntity, err
	}
	if err := helpers.Convert(result, &userEntity); err != nil {
		return userEntity, err
	}
	return userEntity, nil
}
func (userService) Update(c context.Context, id uint, data domains.UserEditForm) (int64, domains.UserEntity, error) {
	var userEntity domains.UserEntity
	var userModel domains.UserModel
	if ok, err := helpers.Validate(data); !ok {
		return 0, userEntity, err
	}
	if err := helpers.Convert(data, &userModel); err != nil {
		return 0, userEntity, err
	}
	aff, updated, err := domains.RepoRegistry.UserRepo.Update(c, id, userModel)
	if err != nil {
		return aff, userEntity, err
	}
	if err := helpers.Convert(updated, &userEntity); err != nil {
		return 0, userEntity, err
	}
	return aff, userEntity, nil
}
func (userService) Delete(c context.Context, id uint) (uint, int64, error) {
	idResult, aff, err := domains.RepoRegistry.UserRepo.Delete(c, id)
	if err != nil {
		return uint(idResult), aff, err
	}
	return uint(idResult), aff, nil
}

// ----- END OF USER SERVICE -----
// ----- SERVICE USER SERVICE -----
// ----- END OF SERVICE USER SERVICE -----
