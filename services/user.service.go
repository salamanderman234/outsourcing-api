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
type serviceUserService struct{}

func NewServiceUserService() domains.ServiceUserService {
	return &serviceUserService{}
}
func (serviceUserService) Read(c context.Context, id uint, q string, page uint, orderBy string, isDesc bool, withPagination bool) (any, *domains.Pagination, error) {
	var pagination domains.Pagination
	if id != 0 {
		var resultEntity domains.ServiceUserEntity
		result, err := domains.RepoRegistry.ServiceUserRepo.Find(c, id)
		if err != nil {
			return nil, nil, err
		}
		err = helpers.Convert(result, &resultEntity)
		if err != nil {
			return nil, nil, err
		}
		return resultEntity, nil, nil
	}
	datas, maxPage, err := domains.RepoRegistry.ServiceUserRepo.Read(c, q, page, orderBy, isDesc, withPagination)
	if err != nil {
		return nil, nil, err
	}
	var datasEntity []domains.ServiceUserEntity
	for _, data := range datas {
		var dataEntity domains.ServiceUserEntity
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
func (serviceUserService) Update(c context.Context, id uint, data domains.ServiceUserUpdateForm, files ...domains.EntityFileMap) (int, domains.ServiceUserEntity, error) {
	var dataModel domains.ServiceUserModel
	var serviceUserEntity domains.ServiceUserEntity
	fun := func(id uint) (int, domains.Model, error) {
		aff, updated, err := domains.RepoRegistry.ServiceUserRepo.Update(c, id, dataModel)
		return int(aff), updated, err
	}
	aff, _, err := basicUpdateService(id, data, &dataModel, &serviceUserEntity, fun)
	return aff, serviceUserEntity, err
}
func (serviceUserService) Delete(c context.Context, id uint) (int, int, error) {
	idResult, aff, err := domains.RepoRegistry.ServiceUserRepo.Delete(c, id)
	return int(idResult), int(aff), err
}

// ----- END OF SERVICE USER SERVICE -----
// ----- EMPLOYEE SERVICE -----
type employeeService struct{}

func NewEmployeeService() domains.EmployeeService {
	return &employeeService{}
}
func (employeeService) Read(c context.Context, id uint, q string, page uint, orderBy string, isDesc bool, withPagination bool) (any, *domains.Pagination, error) {
	var pagination domains.Pagination
	if id != 0 {
		var resultEntity domains.EmployeeEntity
		result, err := domains.RepoRegistry.EmployeeRepo.Find(c, id)
		if err != nil {
			return nil, nil, err
		}
		err = helpers.Convert(result, &resultEntity)
		if err != nil {
			return nil, nil, err
		}
		return resultEntity, nil, nil
	}
	datas, maxPage, err := domains.RepoRegistry.EmployeeRepo.Read(c, q, page, orderBy, isDesc, withPagination)
	if err != nil {
		return nil, nil, err
	}
	var datasEntity []domains.EmployeeEntity
	for _, data := range datas {
		var dataEntity domains.EmployeeEntity
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
func (employeeService) Update(c context.Context, id uint, data domains.EmployeeUpdateForm, files ...domains.EntityFileMap) (int, domains.EmployeeEntity, error) {
	var dataModel domains.EmployeeModel
	var employeeEntity domains.EmployeeEntity
	fun := func(id uint) (int, domains.Model, error) {
		aff, updated, err := domains.RepoRegistry.EmployeeRepo.Update(c, id, dataModel)
		return int(aff), updated, err
	}
	aff, _, err := basicUpdateService(id, data, &dataModel, &employeeEntity, fun)
	return aff, employeeEntity, err
}
func (employeeService) Delete(c context.Context, id uint) (int, int, error) {
	idResult, aff, err := domains.RepoRegistry.EmployeeRepo.Delete(c, id)
	return int(idResult), int(aff), err
}

// ----- END OF EMPLOYEE SERVICE -----
// ----- SUPERVISOR SERVICE -----
type supervisorService struct{}

func NewSupervisorService() domains.SupervisorService {
	return &supervisorService{}
}
func (supervisorService) Read(c context.Context, id uint, q string, page uint, orderBy string, isDesc bool, withPagination bool) (any, *domains.Pagination, error) {
	var pagination domains.Pagination
	if id != 0 {
		var resultEntity domains.SupervisorEntity
		result, err := domains.RepoRegistry.SupervisorRepo.Find(c, id)
		if err != nil {
			return nil, nil, err
		}
		err = helpers.Convert(result, &resultEntity)
		if err != nil {
			return nil, nil, err
		}
		return resultEntity, nil, nil
	}
	datas, maxPage, err := domains.RepoRegistry.SupervisorRepo.Read(c, q, page, orderBy, isDesc, withPagination)
	if err != nil {
		return nil, nil, err
	}
	var datasEntity []domains.SupervisorEntity
	for _, data := range datas {
		var dataEntity domains.SupervisorEntity
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
func (supervisorService) Update(c context.Context, id uint, data domains.SupervisorUpdateForm, files ...domains.EntityFileMap) (int, domains.SupervisorEntity, error) {
	var dataModel domains.SupervisorModel
	var supervisorEntity domains.SupervisorEntity
	fun := func(id uint) (int, domains.Model, error) {
		aff, updated, err := domains.RepoRegistry.SupervisorRepo.Update(c, id, dataModel)
		return int(aff), updated, err
	}
	aff, _, err := basicUpdateService(id, data, &dataModel, &supervisorEntity, fun)
	return aff, supervisorEntity, err
}
func (supervisorService) Delete(c context.Context, id uint) (int, int, error) {
	idResult, aff, err := domains.RepoRegistry.SupervisorRepo.Delete(c, id)
	return int(idResult), int(aff), err
}

// ----- END OF SUPERVISOR SERVICE -----
// ----- ADMIN SERVICE -----
type adminService struct{}

func NewAdminService() domains.AdminService {
	return &adminService{}
}
func (adminService) Read(c context.Context, id uint, q string, page uint, orderBy string, isDesc bool, withPagination bool) (any, *domains.Pagination, error) {
	var pagination domains.Pagination
	if id != 0 {
		var resultEntity domains.AdminEntity
		result, err := domains.RepoRegistry.AdminRepo.Find(c, id)
		if err != nil {
			return nil, nil, err
		}
		err = helpers.Convert(result, &resultEntity)
		if err != nil {
			return nil, nil, err
		}
		return resultEntity, nil, nil
	}
	datas, maxPage, err := domains.RepoRegistry.AdminRepo.Read(c, q, page, orderBy, isDesc, withPagination)
	if err != nil {
		return nil, nil, err
	}
	var datasEntity []domains.AdminEntity
	for _, data := range datas {
		var dataEntity domains.AdminEntity
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
func (adminService) Update(c context.Context, id uint, data domains.AdminUpdateForm, files ...domains.EntityFileMap) (int, domains.AdminEntity, error) {
	var dataModel domains.AdminModel
	var adminEntity domains.AdminEntity
	fun := func(id uint) (int, domains.Model, error) {
		aff, updated, err := domains.RepoRegistry.AdminRepo.Update(c, id, dataModel)
		return int(aff), updated, err
	}
	aff, _, err := basicUpdateService(id, data, &dataModel, &adminEntity, fun)
	return aff, adminEntity, err
}
func (adminService) Delete(c context.Context, id uint) (int, int, error) {
	idResult, aff, err := domains.RepoRegistry.AdminRepo.Delete(c, id)
	return int(idResult), int(aff), err
}

// ----- END OF ADMIN SERVICE -----
