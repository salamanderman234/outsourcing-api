package services

import (
	"context"

	"github.com/salamanderman234/outsourcing-api/domains"
	"github.com/salamanderman234/outsourcing-api/helpers"
)

// --> Category Service
type categoryService struct{}

func NewCategoryService() domains.ServiceCategoryService {
	return &categoryService{}
}
func (categoryService) storeIcon(entity *domains.CategoryModel, files ...domains.EntityFileMap) (bool, error) {
	if len(files) == 1 {
		file := files[0]
		if file.Field == "icon" && file.File != nil {
			savedPath, err := domains.
				ServiceRegistry.
				FileServ.
				Store(file.File, "master/category")
			if err != nil {
				return false, err
			}
			entity.Icon = savedPath
			return true, nil
		}
	}
	entity.Icon = ""
	return false, nil
}
func (cs categoryService) Create(c context.Context, data any, files ...domains.EntityFileMap) (any, error) {
	var dataModel domains.CategoryModel
	var categoryEntity domains.CategoryEntity
	fun := func() (any, error) {
		_, err := cs.storeIcon(&dataModel, files...)
		if err != nil {
			return nil, err
		}
		return domains.RepoRegistry.CategoryRepo.Create(c, dataModel)
	}
	return basicCreateService(data, &dataModel, &categoryEntity, fun)
}
func (categoryService) Read(c context.Context, id uint, q string, page uint, orderBy string, isDesc bool) (any, *domains.Pagination, error) {
	var pagination domains.Pagination
	datas, maxPage, err := domains.RepoRegistry.CategoryRepo.Get(c, id, q, page, orderBy, isDesc)
	if err != nil {
		return nil, nil, err
	}
	datasModel, ok := datas.([]domains.CategoryModel)
	if !ok {
		dataModel, ok := datas.(domains.CategoryModel)
		if !ok {
			return nil, nil, domains.ErrConversionType
		}
		var dataEntity domains.CategoryEntity
		err := helpers.Convert(dataModel, &dataEntity)
		if err != nil {
			return nil, nil, domains.ErrConversionType
		}
		return dataEntity, nil, nil
	}
	var datasEntity []domains.CategoryEntity
	for _, data := range datasModel {
		var dataEntity domains.CategoryEntity
		err := helpers.Convert(data, &dataEntity)
		if err != nil {
			return nil, nil, domains.ErrConversionType
		}
		datasEntity = append(datasEntity, dataEntity)
	}
	queries := helpers.MakeDefaultGetPaginationQueries(q, id, page, orderBy, isDesc)
	pagination = helpers.MakePagination(maxPage, uint(page), queries)
	return datasEntity, &pagination, nil
}
func (cs categoryService) Update(c context.Context, id uint, data any, files ...domains.EntityFileMap) (int, any, error) {
	var dataModel domains.CategoryModel
	var categoryEntity domains.CategoryEntity
	fun := func(id uint) (int, any, error) {
		categoryModel, err := domains.RepoRegistry.CategoryRepo.FindByID(c, id)
		if err != nil {
			return 0, nil, err
		}
		icon := categoryModel.(domains.CategoryModel).Icon
		ok, err := cs.storeIcon(&dataModel, files...)
		if err != nil {
			return 0, nil, err
		}
		if ok {
			go domains.ServiceRegistry.FileServ.Destroy(icon)
		}
		aff, updated, err := domains.RepoRegistry.CategoryRepo.Update(c, id, dataModel)
		return int(aff), updated, err
	}

	return basicUpdateService(id, data, &dataModel, &categoryEntity, fun)
}
func (categoryService) Delete(c context.Context, id uint) (int, int, error) {
	idResult, aff, err := domains.RepoRegistry.CategoryRepo.Delete(c, id)
	return int(idResult), int(aff), err
}

// // --> District Service
// type districtService struct{}

// func NewDistrictService() domains.DistrictService {
// 	return &districtService{}
// }

// // --> Subdistrict Service
// type subDistrictService struct{}

// func NewSubDistrictService() domains.SubDistrictService {
// 	return &subDistrictService{}
// }

// // --> Village Service
// type villageService struct{}

// func NewVillageService() domains.VillageService {
// 	return &villageService{}
// }
