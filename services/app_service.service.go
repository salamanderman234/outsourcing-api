package services

import (
	"context"

	"github.com/salamanderman234/outsourcing-api/domains"
	"github.com/salamanderman234/outsourcing-api/helpers"
)

// ----- SERVICE ITEM SERVICE -----
type serviceItemService struct{}

func NewServiceItemService() domains.ServiceItemService {
	return &serviceItemService{}
}
func (cs serviceItemService) Create(c context.Context, data domains.ServiceItemCreateForm, files ...domains.EntityFileMap) (domains.ServiceItemEntity, error) {
	var dataModel domains.ServiceItemModel
	var dataEntity domains.ServiceItemEntity
	fun := func() (domains.Model, error) {
		return domains.RepoRegistry.ServiceItemRepo.Create(c, dataModel)
	}
	_, err := basicCreateService(data, &dataModel, &dataEntity, fun)
	return dataEntity, err
}
func (serviceItemService) Read(c context.Context, id uint, q string, page uint, orderBy string, isDesc bool, withPagination bool) (any, *domains.Pagination, error) {
	var pagination domains.Pagination
	if id != 0 {
		result, err := domains.RepoRegistry.ServiceItemRepo.Find(c, id)
		return result, nil, err
	}
	datas, maxPage, err := domains.RepoRegistry.ServiceItemRepo.Read(c, q, page, orderBy, isDesc, withPagination)
	if err != nil {
		return nil, nil, err
	}
	var datasEntity []domains.ServiceItemEntity
	for _, data := range datas {
		var dataEntity domains.ServiceItemEntity
		err := helpers.Convert(data, &dataEntity)
		if err != nil {
			return nil, nil, domains.ErrConversionType
		}
		datasEntity = append(datasEntity, dataEntity)
	}
	queries := helpers.MakeDefaultGetPaginationQueries(q, id, page, orderBy, isDesc, withPagination)
	pagination = helpers.MakePagination(maxPage, uint(page), queries)
	return datasEntity, &pagination, nil
}
func (cs serviceItemService) Update(c context.Context, id uint, data domains.ServiceItemUpdateForm, files ...domains.EntityFileMap) (int, domains.ServiceItemEntity, error) {
	var dataModel domains.ServiceItemModel
	var dataEntity domains.ServiceItemEntity
	fun := func(id uint) (int, domains.Model, error) {
		aff, updated, err := domains.RepoRegistry.ServiceItemRepo.Update(c, id, dataModel)
		return int(aff), updated, err
	}
	aff, _, err := basicUpdateService(id, data, &dataModel, &dataEntity, fun)
	return aff, dataEntity, err
}
func (serviceItemService) Delete(c context.Context, id uint) (int, int, error) {
	idResult, aff, err := domains.RepoRegistry.ServiceItemRepo.Delete(c, id)
	return int(idResult), int(aff), err
}

// ----- END OF SERVICE ITEM SERVICE -----
// ----- PARTIAL SERVICE SERVICE -----
type partialServiceService struct{}

func NewPartialServiceService() domains.PartialServiceService {
	return &partialServiceService{}
}
func (partialServiceService) storeIconImage(model *domains.ServiceModel, files ...domains.EntityFileMap) (bool, error) {
	filesLen := len(files)
	if filesLen <= 2 && filesLen > 0 {
		for _, file := range files {
			if file.Field == "icon" && file.File != nil {
				savedPath, err := domains.
					ServiceRegistry.
					FileServ.
					Store(file.File, "service/icon")
				if err != nil {
					return false, err
				}
				model.Icon = savedPath
			} else if file.Field == "image" && file.File != nil {
				savedPath, err := domains.
					ServiceRegistry.
					FileServ.
					Store(file.File, "service/image")
				if err != nil {
					return false, err
				}
				model.Image = savedPath
			}
		}
		return true, nil
	}
	model.Icon = ""
	model.Image = ""
	return false, nil
}
func (ps partialServiceService) Create(c context.Context, data domains.PartialServiceCreateForm, files ...domains.EntityFileMap) (domains.ServiceEntity, error) {
	var dataModel domains.ServiceModel
	var dataEntity domains.ServiceEntity
	fun := func() (domains.Model, error) {
		_, err := ps.storeIconImage(&dataModel, files...)
		if err != nil {
			return nil, err
		}
		return domains.RepoRegistry.ServiceRepo.Create(c, dataModel)
	}
	_, err := basicCreateService(data, &dataModel, &dataEntity, fun)
	if err != nil {
		return domains.ServiceEntity{}, err
	}
	return dataEntity, nil
}
func (partialServiceService) Read(c context.Context, id uint, q string, page uint, orderBy string, isDesc bool, withPagination bool) (any, *domains.Pagination, error) {
	var pagination domains.Pagination
	if id == 0 {
		result, err := domains.RepoRegistry.ServiceRepo.Find(c, id)
		return result, nil, err
	}
	datas, maxPage, err := domains.RepoRegistry.ServiceRepo.Read(c, q, page, orderBy, isDesc, withPagination)
	if err != nil {
		return nil, nil, err
	}
	var datasEntity []domains.ServiceEntity
	for _, data := range datas {
		var dataEntity domains.ServiceEntity
		err := helpers.Convert(data, &dataEntity)
		if err != nil {
			return nil, nil, domains.ErrConversionType
		}
		datasEntity = append(datasEntity, dataEntity)
	}
	if withPagination {
		queries := helpers.MakeDefaultGetPaginationQueries(q, id, page, orderBy, isDesc, withPagination)
		pagination = helpers.MakePagination(maxPage, uint(page), queries)
		return datasEntity, &pagination, nil
	}
	return datasEntity, nil, nil
}
func (ps partialServiceService) Update(c context.Context, id uint, data domains.PartialServiceUpdateForm, files ...domains.EntityFileMap) (int, domains.ServiceEntity, error) {
	var dataModel domains.ServiceModel
	var dataEntity domains.ServiceEntity
	fun := func(id uint) (int, domains.Model, error) {
		categoryModel, err := domains.RepoRegistry.ServiceRepo.Find(c, id)
		if err != nil {
			return 0, nil, err
		}
		_, err = ps.storeIconImage(&dataModel, files...)
		if err != nil {
			return 0, nil, err
		}
		if dataModel.Icon != "" {
			icon := categoryModel.Icon
			go domains.ServiceRegistry.FileServ.Destroy(icon)
		}
		if dataModel.Image != "" {
			image := categoryModel.Image
			go domains.ServiceRegistry.FileServ.Destroy(image)
		}
		aff, updated, err := domains.RepoRegistry.ServiceRepo.Update(c, id, dataModel)
		return int(aff), updated, err
	}
	aff, result, err := basicUpdateService(id, data, &dataModel, &dataEntity, fun)
	if err != nil {
		return 0, domains.ServiceEntity{}, err
	}
	return aff, result.(domains.ServiceEntity), nil
}
func (partialServiceService) Delete(c context.Context, id uint) (int, int, error) {
	idResult, aff, err := domains.RepoRegistry.ServiceRepo.Delete(c, id)
	return int(idResult), int(aff), err
}

// ----- END OF PARTIAL SERVICE SERVICE -----
// ----- SERVICE PACKAGE -----
// ----- END OF SERVICE PACKAGE
