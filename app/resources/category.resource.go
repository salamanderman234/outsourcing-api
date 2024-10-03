package resources

import (
	"github.com/salamanderman234/outsourcing-api/app/domains"
	"github.com/salamanderman234/outsourcing-api/app/models"
	"github.com/salamanderman234/outsourcing-api/app/policies"
	"github.com/salamanderman234/outsourcing-api/app/types"
	"github.com/salamanderman234/outsourcing-api/configs"
)

type iconCategoryResource struct {
	Resource
}

func NewIconCategoryResource() domains.ResourceInterface {
	return &iconCategoryResource{
		Resource: Resource{
			path:       "/category/icon",
			model:      &models.Category{},
			fileConfig: configs.ResourceConfig.GetFileConfig(types.ImageConfig),
			field:      "Icon",
			policy:     &policies.MasterPolicy{},
		},
	}
}
