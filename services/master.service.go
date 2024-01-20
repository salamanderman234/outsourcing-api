package services

import (
	"context"

	"github.com/salamanderman234/outsourcing-api/configs"
	"github.com/salamanderman234/outsourcing-api/domains"
	"github.com/salamanderman234/outsourcing-api/helpers"
)

// --> Category Service
type categoryService struct{}

func NewCategoryService() domains.CategoryService {
	return &categoryService{}
}
func (categoryService) storeIcon(entity *domains.CategoryModel, files ...domains.EntityFileMap) (bool, error) {
	if len(files) == 1 {
		file := files[0]
		if file.Field == "icon" && file.File != nil {
			zippedFile := map[string]domains.FileWrapper{
				"icon": {
					Config: configs.IMAGE_FILE_CONFIG,
					File:   file.File,
					Field:  file.Field,
					Dest:   configs.FILE_DESTS["category/"+file.Field],
				},
			}
			savedPaths, _, err := domains.ServiceRegistry.FileServ.BatchStore(zippedFile)
			if err != nil {
				return false, err
			}
			entity.Icon = savedPaths["icon"]
			return true, nil
		}
	}
	entity.Icon = ""
	return false, nil
}
func (cs categoryService) Create(c context.Context, data domains.CategoryCreateForm, files ...domains.EntityFileMap) (domains.CategoryEntity, error) {
	var dataModel domains.CategoryModel
	var categoryEntity domains.CategoryEntity
	fun := func() (domains.Model, error) {
		_, err := cs.storeIcon(&dataModel, files...)
		if err != nil {
			return nil, err
		}
		result, err := domains.RepoRegistry.CategoryRepo.Create(c, dataModel)
		if err != nil {
			go domains.ServiceRegistry.FileServ.Destroy(dataModel.Icon)
		}
		return result, nil
	}
	err := basicCreateService(true, c, data, &dataModel, &categoryEntity, fun)
	return categoryEntity, err
}
func (categoryService) Read(c context.Context, id uint, q string, page uint, orderBy string, isDesc bool, withPagination bool) (any, *domains.Pagination, error) {
	var resultEntity domains.CategoryEntity
	var datasEntity []domains.CategoryEntity
	findFun := func() (any, error) {
		return domains.RepoRegistry.CategoryRepo.Find(c, id)
	}
	readFun := func() (any, uint, error) {
		return domains.RepoRegistry.CategoryRepo.Read(c, q, page, orderBy, isDesc, withPagination)
	}
	conFun := func(datas any) error {
		datasModel := datas.([]domains.CategoryModel)
		for _, data := range datasModel {
			var dataEntity domains.CategoryEntity
			err := helpers.Convert(data, &dataEntity)
			if err != nil {
				return domains.ErrConversionType
			}
			datasEntity = append(datasEntity, dataEntity)
		}
		return nil
	}
	pagination, err := basicReadService(true,
		c, id, q, page, orderBy, isDesc, withPagination,
		&resultEntity, findFun, readFun, conFun, domains.CategoryModel{},
	)
	if id != 0 {
		return resultEntity, nil, err
	}
	return datasEntity, pagination, err
}
func (cs categoryService) Update(c context.Context, id uint, data domains.CategoryUpdateForm, files ...domains.EntityFileMap) (int, domains.CategoryEntity, error) {
	var dataModel domains.CategoryModel
	var categoryEntity domains.CategoryEntity
	fun := func(id uint) (int, domains.Model, error) {
		categoryModel, err := domains.RepoRegistry.CategoryRepo.Find(c, id)
		if err != nil {
			return 0, nil, err
		}
		icon := categoryModel.Icon
		ok, err := cs.storeIcon(&dataModel, files...)
		if err != nil {
			return 0, nil, err
		}
		// destroy after success
		if ok {
			go domains.ServiceRegistry.FileServ.Destroy(icon)
		}
		aff, updated, err := domains.RepoRegistry.CategoryRepo.Update(c, id, dataModel)
		// destroy if failed update
		if err != nil {
			go domains.ServiceRegistry.FileServ.Destroy(dataModel.Icon)
		}
		return int(aff), updated, nil
	}
	aff, err := basicUpdateService(true, c, id, data, &dataModel, &categoryEntity, fun)
	return aff, categoryEntity, err
}
func (categoryService) Delete(c context.Context, id uint) (int, int, error) {
	delFun := func() (int, int, error) {
		id, aff, err := domains.RepoRegistry.CategoryRepo.Delete(c, id)
		return int(id), int(aff), err
	}
	idResult, aff, err := basicDeleteService(true, c, id, delFun, domains.CategoryModel{})
	return int(idResult), int(aff), err
}
