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

type setTransactionStatusJob struct{}

func NewSetTransactionStatusJob() domains.JobInterface {
	return &setTransactionStatusJob{}
}

func (setTransactionStatusJob) Handle(err chan<- error) {
	var transactions []models.Transaction
	providers.RepoProvider.BaseRepo.ReadAll(context.Background(), &transactions, types.DBSearchParams{
		Params: []types.WhereQuery{
			{Field: "status", Operator: "=", Str: string(enums.Ongoing)},
		},
	})
	// if errs != nil {
	// 	helpers.Logger.Error(fmt.Sprintf("(Job) Failed to execute job <Set Transaction Status Job> : %s", errs.Error()))
	// 	err <- errs
	// }
	for _, transaction := range transactions {
		method := transaction.PaymentMethod
		nextDeadline := transaction.NextPaymentDeadline

		if method == nil {
			continue
		}

		if nextDeadline == nil {
			continue
		}
		if time.Now().Day() == nextDeadline.Day() &&
			time.Now().Month() == nextDeadline.Month() &&
			time.Now().Year() == nextDeadline.Year() {

			status := string(enums.WaitingForFurtherPayments)
			if *method == string(enums.DpPayment) {
				dpStatus := string(enums.WaitingForRemainingDP)
				transaction.Status = &status
				transaction.DPStatus = &dpStatus

				providers.RepoProvider.BaseRepo.Update(context.Background(), []uint{
					transaction.ID,
				}, &transaction)
			} else if *method == string(enums.ThreeTermin) {
				terminStatus := transaction.TerminStatus
				if *terminStatus == string(enums.FirstTerminCompleted) {
					changeTerminStatus := string(enums.WaitingForSecondTermin)
					transaction.TerminStatus = &changeTerminStatus
					transaction.Status = &status

					providers.RepoProvider.BaseRepo.Update(context.Background(), []uint{
						transaction.ID,
					}, &transaction)
				} else if *terminStatus == string(enums.SecondTerminCompleted) {
					changeTerminStatus := string(enums.WaitingForThirdTermin)
					transaction.TerminStatus = &changeTerminStatus
					transaction.Status = &status

					providers.RepoProvider.BaseRepo.Update(context.Background(), []uint{
						transaction.ID,
					}, &transaction)
				}
			}

		}

	}
}

func (setTransactionStatusJob) GetData() map[string]any {
	return map[string]any{}
}

func (setTransactionStatusJob) SetData(data map[string]any) {}
