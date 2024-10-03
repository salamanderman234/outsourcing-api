package jobs

import (
	"context"
	"time"

	"github.com/salamanderman234/outsourcing-api/app/domains"
	"github.com/salamanderman234/outsourcing-api/app/models"
	"github.com/salamanderman234/outsourcing-api/app/providers"
	"github.com/salamanderman234/outsourcing-api/app/types"
	"github.com/salamanderman234/outsourcing-api/app/types/enums"
)

type setTransactionCancelJob struct{}

func NewSetTransactionCancelJob() domains.JobInterface {
	return &setTransactionCancelJob{}
}

func (setTransactionCancelJob) Handle(err chan<- error) {
	var transactions []models.Transaction
	providers.RepoProvider.BaseRepo.ReadAll(context.Background(), &transactions, types.DBSearchParams{
		Params: []types.WhereQuery{
			{Field: "status", Operator: "=", Str: string(enums.WaitingForInitialPayment)},
		},
	})
	// if errs != nil {
	// 	helpers.Logger.Error(fmt.Sprintf("(Job) Failed to execute job <Set Transaction Cancel Job> : %s", errs.Error()))
	// 	err <- errs
	// }
	for _, transaction := range transactions {
		nextDeadline := transaction.NextPaymentDeadline

		if nextDeadline == nil {
			continue
		}
		if time.Now().Equal(nextDeadline.Add(time.Duration(30)*time.Minute)) || time.Now().After(nextDeadline.Add(time.Duration(30)*time.Minute)) {
			providers.RepoProvider.BaseRepo.Delete(context.Background(), []uint{
				transaction.ID,
			}, &models.Transaction{})
		}

	}
}

func (setTransactionCancelJob) GetData() map[string]any {
	return map[string]any{}
}

func (setTransactionCancelJob) SetData(data map[string]any) {}
