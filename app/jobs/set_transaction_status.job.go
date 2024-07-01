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

type setTransactionStatusJob struct{}

func NewSetTransactionStatusJob() domains.JobInterface {
	return &setTransactionStatusJob{}
}

func (setTransactionStatusJob) Handle(err chan<- error) {
	var transactions []models.Transaction
	_, errs := providers.RepoProvider.BaseRepo.ReadAll(context.Background(), &transactions, types.DBSearchParams{
		Params: []types.WhereQuery{
			{Field: "status", Operator: "=", Str: string(enums.Ongoing)},
		},
	})
	if errs != nil {
		helpers.Logger.Error(fmt.Sprintf("(Job) Failed to execute job <Set Transaction Status Job> : %s", errs.Error()))
		err <- errs
	}

	for _, transaction := range transactions {
		method := transaction.PaymentMethod

		if method == nil {
			continue
		}

		if *method == string(enums.DpPayment) {
			nextDeadline := transaction.NextPaymentDeadline
			if time.Now().Day() == nextDeadline.Day() &&
				time.Now().Month() == nextDeadline.Month() &&
				time.Now().Year() == nextDeadline.Year() {

			}
		} else if *method == string(enums.ThreeTermin) {

		}
	}
}

func (setTransactionStatusJob) GetData() map[string]any {
	return map[string]any{}
}

func (setTransactionStatusJob) SetData(data map[string]any) {}
