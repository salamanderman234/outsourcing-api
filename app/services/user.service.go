package services

import (
	"context"
	"strconv"

	service_domains "github.com/salamanderman234/outsourcing-api/app/domains/services"
	"github.com/salamanderman234/outsourcing-api/app/forms"
	"github.com/salamanderman234/outsourcing-api/app/models"
	"github.com/salamanderman234/outsourcing-api/app/policies"
	"github.com/salamanderman234/outsourcing-api/app/types"
	"github.com/salamanderman234/outsourcing-api/app/types/enums"
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
			{Field: "users.email", Operator: "LIKE", Str: q},
			{Field: "admins.full_name", Operator: "LIKE", Str: q, IsOr: true},
			{Field: "service_users.full_name", Operator: "LIKE", Str: q, IsOr: true},
			{Field: "employees.full_name", Operator: "LIKE", Str: q, IsOr: true},
			{Field: "super_admins.full_name", Operator: "LIKE", Str: q, IsOr: true},
			{Field: "supervisor.full_name", Operator: "LIKE", Str: q, IsOr: true},
		},
		Query: q,
		Model: &models.User{},
		Preloads: []string{
			"AdminProfile",
			"SupervisorProfile",
			"EmployeeProfile",
			"ServiceUserProfile",
			"SuperAdminProfile",
		},
		Joins: []string{
			"JOIN admins ON admins.user_id = users.id",
			"JOIN service_users ON service_users.user_id = users.id",
			"JOIN supervisors ON supervisors.user_id = users.id",
			"JOIN super_admins ON super_admins.user_id = users.id",
			"JOIN employees ON employees.user_id = users.id",
		},
	}

	if regencyID != 0 {
		regencyIDStr := strconv.Itoa(int(regencyID))
		queries := []types.WhereQuery{}
		if role == enums.AdminUserRole {
			queries = append(queries, types.WhereQuery{
				Field: "admins.regency_id", Operator: "=", Str: regencyIDStr,
			})
		} else if role == enums.EmployeeUserRole {
			queries = append(queries, types.WhereQuery{
				Field: "employees.regency_id", Operator: "=", Str: regencyIDStr,
			})
		} else if role == enums.SupervisorUserRole {
			queries = append(queries, types.WhereQuery{
				Field: "supervisors.regency_id", Operator: "=", Str: regencyIDStr,
			})
		}
		params.Params = append(params.Params, queries...)
	}
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
	err := baseUpdateFunc(ctx, &policies.UserPolicy{}, id, &user, data)
	return id, user, err
}
