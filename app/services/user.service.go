package services

import (
	"context"
	"fmt"
	"strconv"

	service_domains "github.com/salamanderman234/outsourcing-api/app/domains/services"
	"github.com/salamanderman234/outsourcing-api/app/forms"
	"github.com/salamanderman234/outsourcing-api/app/helpers"
	"github.com/salamanderman234/outsourcing-api/app/models"
	"github.com/salamanderman234/outsourcing-api/app/policies"
	"github.com/salamanderman234/outsourcing-api/app/providers"
	"github.com/salamanderman234/outsourcing-api/app/types"
	"github.com/salamanderman234/outsourcing-api/app/types/enums"
	"github.com/salamanderman234/outsourcing-api/configs"
	"golang.org/x/crypto/bcrypt"
)

type userService struct{}

func NewUserService() service_domains.UserServiceInterface {
	return &userService{}
}

func (userService) Read(ctx context.Context,
	q string,
	regencyID uint,
	role enums.UserRolesEnum,
	page uint,
) ([]models.User, *types.Pagination, error) {
	var results []models.User
	params := types.DBSearchParams{
		Page:           page,
		WithPagination: page > 0,
		Params: []types.WhereQuery{
			{Field: "users.role", Operator: "=", Str: string(role)},
			{Field: "users.email", Operator: "LIKE", Str: q},
		},
		Query:    q,
		Model:    &models.User{},
		Preloads: []string{},
		Joins:    []string{},
	}

	regencyIDStr := strconv.Itoa(int(regencyID))
	queries := []types.WhereQuery{}
	joins := []string{}
	preloads := []string{}
	if role == enums.AdminUserRole {
		joins = append(joins, "JOIN admins ON admins.user_id = users.id")
		if regencyID != 0 {
			queries = append(queries, types.WhereQuery{
				Field: "admins.regency_id", Operator: "=", Str: regencyIDStr,
			})
		}
		queries = append(queries, types.WhereQuery{Field: "admins.fullname", Operator: "LIKE", Str: q})
		preloads = append(preloads, "AdminProfile")

	} else if role == enums.EmployeeUserRole {
		if regencyID != 0 {
			queries = append(queries, types.WhereQuery{
				Field: "employees.regency_id", Operator: "=", Str: regencyIDStr,
			})
		}
		queries = append(queries, types.WhereQuery{Field: "employees.fullname", Operator: "LIKE", Str: q})
		joins = append(joins, "JOIN employees ON employees.user_id = users.id")
		preloads = append(preloads, "EmployeeProfile")
	} else if role == enums.SupervisorUserRole {
		if regencyID != 0 {
			queries = append(queries, types.WhereQuery{
				Field: "supervisors.regency_id", Operator: "=", Str: regencyIDStr,
			})
		}
		queries = append(queries, types.WhereQuery{Field: "supervisors.fullname", Operator: "LIKE", Str: q})
		joins = append(joins, "JOIN supervisors ON supervisors.user_id = users.id")
		preloads = append(preloads, "SupervisorProfile")
	} else if role == enums.ServiceUserRole {
		queries = append(queries, types.WhereQuery{Field: "service_users.fullname", Operator: "LIKE", Str: q})
		joins = append(joins, "JOIN service_users ON service_users.user_id = users.id")
		preloads = append(preloads, "ServiceUserProfile")
	} else if role == enums.SuperAdminUserRole {
		queries = append(queries, types.WhereQuery{Field: "super_admins.fullname", Operator: "LIKE", Str: q})
		joins = append(joins, "JOIN super_admins ON super_admins.user_id = users.id")
		preloads = append(preloads, "SuperAdminProfile")
	}
	params.Params = append(params.Params, queries...)
	params.Joins = append(params.Joins, joins...)
	params.Preloads = append(params.Preloads, preloads...)

	pagination, err := baseReadFunc(
		ctx,
		policies.UserPolicy{},
		params,
		&results,
	)
	return results, pagination, err
}
func (userService) Find(ctx context.Context, id uint) (models.User, error) {
	var user models.User
	err := baseFindFunc(
		ctx,
		&policies.UserPolicy{},
		id,
		&user,
		"AdminProfile",
		"SupervisorProfile",
		"EmployeeProfile",
		"ServiceUserProfile",
	)
	return user, err
}
func (userService) Update(ctx context.Context,
	id uint,
	data forms.UserUpdateForm,
) (uint, models.User, error) {
	var user models.User
	var result models.User
	claims, _ := ctx.Value(configs.VarConfig.UserContextName).(types.JWTCLaims)
	err := providers.RepoProvider.BaseRepo.Find(ctx, id, &user)
	if err != nil {
		return id, user, err
	}
	policy := policies.UserPolicy{}
	if !policy.Update(&user, claims) {
		helpers.Logger.Warning(
			fmt.Sprintf("(Forbidden) User: %s", claims.Email),
		)
		return id, user, types.ErrForbiden
	}
	user.SetUpdatedEmail(claims.Email)
	if err := helpers.Validator.Validate(data); err != nil {
		return id, user, err
	}
	if err := helpers.Translator.TranslateStruct(data, &result); err != nil {
		return id, user, err
	}
	profile := data.ProfilePic
	if result.Password != nil {
		hashed, err := bcrypt.GenerateFromPassword([]byte(*result.Password), 1)
		if err != nil {
			return id, result, err
		}
		strHashed := string(hashed)
		result.Password = &strHashed
	}
	if profile != nil {
		old := user.ProfilePic
		res, err := providers.ResourceProvider.CreateResource("users.profile")
		if err != nil {
			return id, user, err
		}
		res.SetData(&user)
		rest, err := providers.ServiceProvider.FileService.UploadFile(ctx, *profile, res)
		if err != nil {
			return id, user, err
		}
		if old != nil {
			go providers.ServiceProvider.FileService.DeleteFile(ctx, *old)
		}
		result.ProfilePic = &rest
	}
	result, err = providers.RepoProvider.UserRepo.UpdateUser(ctx, id, result)
	if err != nil {
		return id, result, err
	}
	result.SetID(id)
	return id, result, err
}
