package jobs

import (
	"context"
	"fmt"
	"time"

	"github.com/salamanderman234/outsourcing-api/app/domains"
	"github.com/salamanderman234/outsourcing-api/app/helpers"
	"github.com/salamanderman234/outsourcing-api/app/models"
	"github.com/salamanderman234/outsourcing-api/app/providers"
	"github.com/salamanderman234/outsourcing-api/app/types"
	"github.com/salamanderman234/outsourcing-api/app/types/enums"
)

type generatePerformanceForm struct {
}

func NewGeneratePerformanceFormJob() domains.JobInterface {
	return &generatePerformanceForm{}
}

func (s generatePerformanceForm) Handle(err chan<- error) {
	date := time.Now()
	var placements []models.Placement
	_, errs := providers.RepoProvider.BaseRepo.ReadAll(context.Background(), &placements, types.DBSearchParams{
		Params: []types.WhereQuery{
			{Field: "status", Operator: "=", Str: string(enums.PlacementOngoingStatus)},
			{Field: "status", Operator: "=", Str: string(enums.PlacementSuspendStatus), IsOr: true},
		},
	})
	if errs != nil {
		helpers.Logger.Error(fmt.Sprintf("(Job) Failed to execute job <Performance Form Job> : %s", errs.Error()))
		err <- errs
	}
	for _, placement := range placements {
		schedule := placement.FormGenerateSchedule
		last := placement.LastFormDate

		if schedule == nil {
			continue
		}

		if last == nil {
			last = placement.StartDate
		}

		success := true
		placementID := placement.ID
		generateFunc := func() {
			errs := providers.RepoProvider.BaseRepo.Create(context.Background(), &models.PerformanceForm{
				PlacementID: &placementID,
				Date:        &date,
			})
			if errs != nil {
				err <- errs
				helpers.Logger.Error(fmt.Sprintf("(Job) Failed to execute job <Performance Form Job> : %s", errs.Error()))
				success = false
			}
		}
		if *schedule == string(enums.EndFormSchedule) && (time.Now().Day() == placement.EndDate.Day() &&
			time.Now().Month() == placement.EndDate.Month() &&
			time.Now().Year() == placement.EndDate.Year()) {

			generateFunc()
		} else if *schedule == string(enums.WeeklyFormSchedule) && (time.Now().Day() == last.Add((7*24)*time.Hour).Day() &&
			time.Now().Month() == last.Add((7*24)*time.Hour).Month() &&
			time.Now().Year() == last.Add((7*24)*time.Hour).Year()) {

			generateFunc()
		} else if *schedule == string(enums.MonthlyFormSchedule) && (time.Now().Day() == last.Add((30*24)*time.Hour).Day() &&
			time.Now().Month() == last.Add((30*24)*time.Hour).Month() &&
			time.Now().Year() == last.Add((30*24)*time.Hour).Year()) {

			generateFunc()
		}

		if success {
			placement.LastFormDate = &date
			errs := providers.RepoProvider.BaseRepo.Update(context.Background(), []uint{
				placementID,
			}, &placement)
			if errs != nil {
				err <- errs
				helpers.Logger.Error(fmt.Sprintf("(Job) Failed to execute job <Performance Form Job> : %s", errs.Error()))
			}
		}

	}
}

func (s generatePerformanceForm) GetData() map[string]any {
	return map[string]any{}
}
func (s *generatePerformanceForm) SetData(data map[string]any) {}

func init() {
	providers.JobProvider.RegisterJob(&generatePerformanceForm{})
}
