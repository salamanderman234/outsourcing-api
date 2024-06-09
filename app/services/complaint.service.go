package services

import (
	"context"
	"strconv"

	service_domains "github.com/salamanderman234/outsourcing-api/app/domains/services"
	"github.com/salamanderman234/outsourcing-api/app/forms"
	"github.com/salamanderman234/outsourcing-api/app/models"
	"github.com/salamanderman234/outsourcing-api/app/policies"
	"github.com/salamanderman234/outsourcing-api/app/providers"
	"github.com/salamanderman234/outsourcing-api/app/types"
	"github.com/salamanderman234/outsourcing-api/configs"
)

type complaintService struct{}

func NewComplaintService() service_domains.ComplaintServiceInterface {
	return &complaintService{}
}

func (complaintService) Create(ctx context.Context, data forms.ComplaintCreateForm) (models.Complaint, error) {
	var complaint models.Complaint
	before := func() error {
		claims, _ := ctx.Value(configs.VarConfig.UserContextName).(types.JWTCLaims)
		id, _ := strconv.Atoi(claims.ID)
		existsUser, err := providers.ServiceProvider.UserService.Find(ctx, uint(id))
		if err != nil {
			return err
		}
		if existsUser.ServiceUserProfile == nil {
			return types.ErrForbiden
		}
		complaint.ServiceUserID = &existsUser.ServiceUserProfile.ID

		detail, err := providers.ServiceProvider.PlacementService.EmployeePlacementDetail(ctx,
			data.PlacementDetailEmployeeID,
		)
		if err != nil {
			return err
		}

		employeeID := detail.EmployeeID
		complaint.EmployeeID = employeeID

		return nil
	}
	err := baseCreateFunc(ctx, policies.ComplaintPolicy{}, &complaint, data, before)
	return complaint, err
}
func (complaintService) Find(ctx context.Context, id uint) (models.Complaint, error) {
	var complaint models.Complaint
	err := baseFindFunc(ctx, policies.ComplaintPolicy{}, id, &complaint,
		"ServiceUser",
		"PlacementDetailEmployee",
		"Employee",
		"Replies",
	)
	return complaint, err
}
func (complaintService) Read(ctx context.Context, q string, page uint) ([]models.Complaint, *types.Pagination, error) {
	var results []models.Complaint
	params := types.DBSearchParams{
		Page:           page,
		WithPagination: page > 0,
		Params:         []types.WhereQuery{},
		Query:          q,
		Model:          &models.Complaint{},
		Preloads: []string{
			"ServiceUser",
			"PlacementDetailEmployee",
			"Employee",
			"Replies",
		},
	}

	pagination, err := baseReadFunc(
		ctx,
		policies.ComplaintPolicy{},
		params,
		&results,
	)
	return results, pagination, err
}
func (complaintService) Update(ctx context.Context, id uint, data forms.ComplaintUpdateForm) (uint, models.Complaint, error) {
	var complaint models.Complaint
	err := baseUpdateFunc(ctx, policies.ComplaintPolicy{}, id, &complaint, data)
	return id, complaint, err
}
func (complaintService) Delete(ctx context.Context, id uint) (uint, error) {
	var complaint models.Complaint
	err := baseDeleteFunc(ctx, policies.ComplaintPolicy{}, id, &complaint)
	return id, err
}
func (complaintService) Reply(ctx context.Context, data forms.ReplyComplaintCreateForm) (models.ComplaintReply, error) {
	var reply models.ComplaintReply
	before := func() error {
		claims, _ := ctx.Value(configs.VarConfig.UserContextName).(types.JWTCLaims)
		id, _ := strconv.Atoi(claims.ID)
		existsUser, err := providers.ServiceProvider.UserService.Find(ctx, uint(id))
		if err != nil {
			return err
		}
		if existsUser.SupervisorProfile == nil {
			return types.ErrForbiden
		}
		reply.SupervisorID = &existsUser.SupervisorProfile.ID
		return nil
	}
	err := baseCreateFunc(ctx, policies.ComplaintReplyPolicy{}, &reply, data, before)
	return reply, err
}
func (complaintService) UpdateReply(ctx context.Context, id uint, data forms.ReplyComplaintUpdateForm) (uint, models.ComplaintReply, error) {
	var reply models.ComplaintReply
	err := baseUpdateFunc(ctx, policies.ComplaintReplyPolicy{}, id, &reply, data)
	return id, reply, err
}
func (complaintService) DeleteReply(ctx context.Context, id uint) (uint, error) {
	var reply models.ComplaintReply
	err := baseDeleteFunc(ctx, policies.ComplaintReplyPolicy{}, id, &reply)
	return id, err
}
