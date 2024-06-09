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

type feedbackService struct{}

func NewFeedbackService() service_domains.FeedbackServiceInterface {
	return &feedbackService{}
}

func (feedbackService) Create(ctx context.Context, data forms.FeedbackCreateForm) (models.Feedback, error) {
	var feedback models.Feedback
	before := func() error {
		var existsUser models.User
		claims, _ := ctx.Value(configs.VarConfig.UserContextName).(types.JWTCLaims)
		id, _ := strconv.Atoi(claims.ID)
		err := providers.RepoProvider.BaseRepo.Find(ctx, uint(id), &existsUser, "ServiceUserProfile")
		if err != nil {
			return err
		}
		if existsUser.ServiceUserProfile == nil {
			return types.ErrForbiden
		}
		var exists []models.Feedback
		serviceUserIDStr := strconv.Itoa(int(existsUser.ServiceUserProfile.ID))
		transIDStr := strconv.Itoa(int(data.TransactionID))
		providers.RepoProvider.BaseRepo.ReadAll(ctx, &exists, types.DBSearchParams{
			Params: []types.WhereQuery{
				{Field: "transaction_id", Operator: "=", Str: transIDStr},
				{Field: "service_user_id", Operator: "=", Str: serviceUserIDStr},
			},
		})
		if len(exists) > 0 {
			return types.ErrDuplicateEntries.SetCustomMsg(
				"you have provided feedback on this transaction",
			)
		}
		feedback.ServiceUserID = &existsUser.ServiceUserProfile.ID
		return nil
	}
	err := baseCreateFunc(ctx, policies.FeedbackPolicy{}, &feedback, data, before)
	return feedback, err
}
func (feedbackService) Find(ctx context.Context, id uint) (models.Feedback, error) {
	var feedback models.Feedback
	err := baseFindFunc(ctx, policies.FeedbackPolicy{}, id, &feedback, "ServiceUser", "Transaction")
	return feedback, err
}
func (feedbackService) Read(ctx context.Context, q string, page uint) ([]models.Feedback, *types.Pagination, error) {
	var results []models.Feedback
	params := types.DBSearchParams{
		Page:           page,
		WithPagination: page > 0,
		Params: []types.WhereQuery{
			{Field: "comment", Operator: "LIKE", Str: q},
		},
		Query: q,
		Model: &models.Feedback{},
		Preloads: []string{
			"Transaction",
			"ServiceUser",
		},
	}

	pagination, err := baseReadFunc(
		ctx,
		policies.FeedbackPolicy{},
		params,
		&results,
	)
	return results, pagination, err
}
func (feedbackService) Update(ctx context.Context, id uint, data forms.FeedbackUpdateForm) (uint, models.Feedback, error) {
	var feedback models.Feedback
	err := baseUpdateFunc(ctx, policies.FeedbackPolicy{}, id, &feedback, data)
	return id, feedback, err
}
func (feedbackService) Delete(ctx context.Context, id uint) (uint, error) {
	var feedback models.Feedback
	err := baseDeleteFunc(ctx, policies.FeedbackPolicy{}, id, &feedback)
	return id, err
}
