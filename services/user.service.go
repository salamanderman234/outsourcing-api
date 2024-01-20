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
	var entity domains.UserEntity
	var model domains.UserModel
	findFunc := func(id uint) (domains.Model, error) {
		return domains.RepoRegistry.UserRepo.Find(c, id)
	}
	err := basicFindService(false, c, id, &model, &entity, findFunc)
	if err != nil {
		return entity, err
	}
	return entity, nil
}
func (userService) Update(c context.Context, id uint, data domains.UserEditForm) (int64, domains.UserEntity, error) {
	var dataModel domains.UserModel
	var dataEntity domains.UserEntity
	fun := func(id uint) (int, domains.Model, error) {
		aff, updated, err := domains.RepoRegistry.UserRepo.Update(c, id, dataModel)
		return int(aff), updated, err
	}
	aff, err := basicUpdateService(true, c, id, data, &dataModel, &dataEntity, fun)
	return int64(aff), dataEntity, err
}
func (userService) Delete(c context.Context, id uint) (uint, int64, error) {
	delFun := func() (int, int, error) {
		id, aff, err := domains.RepoRegistry.UserRepo.Delete(c, id)
		return int(id), int(aff), err
	}
	idResult, aff, err := basicDeleteService(true, c, id, delFun, domains.UserModel{})
	return uint(idResult), int64(aff), err
}

// ----- END OF USER SERVICE -----
// ----- SERVICE USER SERVICE -----
type serviceUserService struct{}

func NewServiceUserService() domains.ServiceUserService {
	return &serviceUserService{}
}
func (serviceUserService) Read(c context.Context, id uint, q string, page uint, orderBy string, isDesc bool, withPagination bool) (any, *domains.Pagination, error) {
	var resultEntity domains.ServiceUserEntity
	var datasEntity []domains.ServiceUserEntity
	findFun := func() (any, error) {
		return domains.RepoRegistry.ServiceUserRepo.Find(c, id)
	}
	readFun := func() (any, uint, error) {
		return domains.RepoRegistry.ServiceUserRepo.Read(c, q, page, orderBy, isDesc, withPagination)
	}
	conFun := func(datas any) error {
		datasModel, ok := datas.([]domains.ServiceUserModel)
		if !ok {
			return domains.ErrConversionType
		}
		for _, data := range datasModel {
			var dataEntity domains.ServiceUserEntity
			err := helpers.Convert(data, &dataEntity)
			if err != nil {
				return domains.ErrConversionType
			}
			datasEntity = append(datasEntity, dataEntity)
		}
		return nil
	}
	pagination, err := basicReadService(
		true,
		c,
		id,
		q,
		page,
		orderBy,
		isDesc,
		withPagination,
		resultEntity,
		findFun,
		readFun,
		conFun,
		domains.ServiceUserModel{},
	)
	if id != 0 {
		return resultEntity, nil, err
	}
	return datasEntity, pagination, err
}
func (serviceUserService) Update(c context.Context, id uint, data domains.ServiceUserUpdateForm, files ...domains.EntityFileMap) (int, domains.ServiceUserEntity, error) {
	var dataModel domains.ServiceUserModel
	var serviceUserEntity domains.ServiceUserEntity
	fun := func(id uint) (int, domains.Model, error) {
		aff, updated, err := domains.RepoRegistry.ServiceUserRepo.Update(c, id, dataModel)
		return int(aff), updated, err
	}
	aff, err := basicUpdateService(true, c, id, data, &dataModel, &serviceUserEntity, fun)
	return aff, serviceUserEntity, err
}
func (serviceUserService) Delete(c context.Context, id uint) (int, int, error) {
	delFun := func() (int, int, error) {
		id, aff, err := domains.RepoRegistry.ServiceUserRepo.Delete(c, id)
		return int(id), int(aff), err
	}
	idResult, aff, err := basicDeleteService(true, c, id, delFun, domains.ServiceUserModel{})
	return int(idResult), int(aff), err
}

// ----- END OF SERVICE USER SERVICE -----
// ----- EMPLOYEE SERVICE -----
type employeeService struct{}

func NewEmployeeService() domains.EmployeeService {
	return &employeeService{}
}
func (employeeService) SetaEmployeeAvailability(c context.Context, id uint, isAvailable bool) (bool, error) {
	user, ok := c.Value("user").(domains.UserEntity)
	if !ok {
		return false, domains.ErrInvalidAccess
	}
	if user.Role != string(domains.AdminRole) {
		return false, domains.ErrInvalidAccess
	}
	status := domains.NotAvailableEmployeeStatus
	if isAvailable {
		status = domains.AvailableEmployeeStatus

	}
	data := domains.EmployeeModel{
		Status: status,
	}
	aff, _, err := domains.RepoRegistry.EmployeeRepo.Update(c, id, data)
	return aff >= 1, err
}
func (employeeService) Read(c context.Context, id uint, category string, employeeStatus domains.EmployeeStatusEnum, q string, page uint, orderBy string, isDesc bool, withPagination bool) (any, *domains.Pagination, error) {
	var datasEntity []domains.EmployeeEntity
	var resultEntity domains.EmployeeEntity

	findFun := func() (any, error) {
		return domains.RepoRegistry.EmployeeRepo.Find(c, id)
	}
	readFun := func() (any, uint, error) {
		return domains.RepoRegistry.EmployeeRepo.Read(c, category, employeeStatus, q, page, orderBy, isDesc, withPagination)
	}
	conFun := func(datas any) error {
		datasModel := datas.([]domains.EmployeeModel)
		for _, data := range datasModel {
			var dataEntity domains.EmployeeEntity
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
func (employeeService) Update(c context.Context, id uint, data domains.EmployeeUpdateForm, files ...domains.EntityFileMap) (int, domains.EmployeeEntity, error) {
	var dataModel domains.EmployeeModel
	var employeeEntity domains.EmployeeEntity
	fun := func(id uint) (int, domains.Model, error) {
		aff, updated, err := domains.RepoRegistry.EmployeeRepo.Update(c, id, dataModel)
		return int(aff), updated, err
	}
	aff, err := basicUpdateService(true, c, id, data, &dataModel, &employeeEntity, fun)
	return aff, employeeEntity, err
}
func (employeeService) Delete(c context.Context, id uint) (int, int, error) {
	delFun := func() (int, int, error) {
		id, aff, err := domains.RepoRegistry.EmployeeRepo.Delete(c, id)
		return int(id), int(aff), err
	}
	idResult, aff, err := basicDeleteService(true, c, id, delFun, domains.EmployeeModel{})
	return int(idResult), int(aff), err
}

// ----- END OF EMPLOYEE SERVICE -----
// ----- SUPERVISOR SERVICE -----
type supervisorService struct{}

func NewSupervisorService() domains.SupervisorService {
	return &supervisorService{}
}
func (supervisorService) SetaSupervisorAvailability(c context.Context, id uint, isAvailable bool) (bool, error) {
	user, ok := c.Value("user").(domains.UserEntity)
	if !ok {
		return false, domains.ErrInvalidAccess
	}
	if user.Role != string(domains.AdminRole) {
		return false, domains.ErrInvalidAccess
	}
	status := domains.NotAvailableEmployeeStatus
	if isAvailable {
		status = domains.AvailableEmployeeStatus

	}
	data := domains.SupervisorModel{
		Status: status,
	}
	aff, _, err := domains.RepoRegistry.SupervisorRepo.Update(c, id, data)
	return aff >= 1, err
}

func (supervisorService) Read(c context.Context, id uint, employeeStatus domains.EmployeeStatusEnum, q string, page uint, orderBy string, isDesc bool, withPagination bool) (any, *domains.Pagination, error) {
	var datasEntity []domains.SupervisorEntity
	var resultEntity domains.SupervisorEntity

	findFun := func() (any, error) {
		return domains.RepoRegistry.SupervisorRepo.Find(c, id)
	}
	readFun := func() (any, uint, error) {
		return domains.RepoRegistry.SupervisorRepo.Read(c, employeeStatus, q, page, orderBy, isDesc, withPagination)
	}
	conFun := func(datas any) error {
		datasModel := datas.([]domains.SupervisorModel)
		for _, data := range datasModel {
			var dataEntity domains.SupervisorEntity
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
		&resultEntity, findFun, readFun, conFun, domains.SupervisorModel{},
	)
	if id != 0 {
		return resultEntity, nil, err
	}
	return datasEntity, pagination, err
}
func (supervisorService) Update(c context.Context, id uint, data domains.SupervisorUpdateForm, files ...domains.EntityFileMap) (int, domains.SupervisorEntity, error) {
	var dataModel domains.SupervisorModel
	var supervisorEntity domains.SupervisorEntity
	fun := func(id uint) (int, domains.Model, error) {
		aff, updated, err := domains.RepoRegistry.SupervisorRepo.Update(c, id, dataModel)
		return int(aff), updated, err
	}
	aff, err := basicUpdateService(true, c, id, data, &dataModel, &supervisorEntity, fun)
	return aff, supervisorEntity, err
}
func (supervisorService) Delete(c context.Context, id uint) (int, int, error) {
	delFun := func() (int, int, error) {
		id, aff, err := domains.RepoRegistry.SupervisorRepo.Delete(c, id)
		return int(id), int(aff), err
	}
	idResult, aff, err := basicDeleteService(true, c, id, delFun, domains.SupervisorModel{})
	return int(idResult), int(aff), err
}

// ----- END OF SUPERVISOR SERVICE -----
// ----- ADMIN SERVICE -----
type adminService struct{}

func NewAdminService() domains.AdminService {
	return &adminService{}
}
func (adminService) Read(c context.Context, id uint, q string, page uint, orderBy string, isDesc bool, withPagination bool) (any, *domains.Pagination, error) {
	var datasEntity []domains.AdminEntity
	var resultEntity domains.AdminEntity

	findFun := func() (any, error) {
		return domains.RepoRegistry.AdminRepo.Find(c, id)
	}
	readFun := func() (any, uint, error) {
		return domains.RepoRegistry.AdminRepo.Read(c, q, page, orderBy, isDesc, withPagination)
	}
	conFun := func(datas any) error {
		datasModel := datas.([]domains.AdminModel)
		for _, data := range datasModel {
			var dataEntity domains.AdminEntity
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
		&resultEntity, findFun, readFun, conFun, domains.AdminModel{},
	)
	if id != 0 {
		return resultEntity, nil, err
	}
	return datasEntity, pagination, err
}
func (adminService) Update(c context.Context, id uint, data domains.AdminUpdateForm, files ...domains.EntityFileMap) (int, domains.AdminEntity, error) {
	var dataModel domains.AdminModel
	var adminEntity domains.AdminEntity
	fun := func(id uint) (int, domains.Model, error) {
		aff, updated, err := domains.RepoRegistry.AdminRepo.Update(c, id, dataModel)
		return int(aff), updated, err
	}
	aff, err := basicUpdateService(true, c, id, data, &dataModel, &adminEntity, fun)
	return aff, adminEntity, err
}
func (adminService) Delete(c context.Context, id uint) (int, int, error) {
	delFun := func() (int, int, error) {
		id, aff, err := domains.RepoRegistry.AdminRepo.Delete(c, id)
		return int(id), int(aff), err
	}
	idResult, aff, err := basicDeleteService(true, c, id, delFun, domains.AdminModel{})
	return int(idResult), int(aff), err
}

// ----- END OF ADMIN SERVICE -----
