package resources

import (
	"github.com/salamanderman234/outsourcing-api/app/domains"
	"github.com/salamanderman234/outsourcing-api/app/models"
	"github.com/salamanderman234/outsourcing-api/app/policies"
	"github.com/salamanderman234/outsourcing-api/app/types"
	"github.com/salamanderman234/outsourcing-api/configs"
)

type mouTransactionResource struct {
	Resource
}

func NewTransactionMouResource() domains.ResourceInterface {
	return &mouTransactionResource{
		Resource{
			path:       "/transaction/mou",
			model:      &models.Transaction{},
			field:      "MOU",
			fileConfig: configs.ResourceConfig.GetFileConfig(types.PDFConfig),
			policy:     &policies.TransactionPolicy{},
		},
	}
}
