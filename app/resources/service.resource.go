package resources

import (
	"github.com/salamanderman234/outsourcing-api/app/domains"
	"github.com/salamanderman234/outsourcing-api/app/models"
	"github.com/salamanderman234/outsourcing-api/app/policies"
	"github.com/salamanderman234/outsourcing-api/app/types"
	"github.com/salamanderman234/outsourcing-api/configs"
)

type iconServiceResource struct {
	Resource
}

type mainImageServiceResource struct {
	Resource
}

func NewServiceIconResource() domains.ResourceInterface {
	return &iconServiceResource{
		Resource: Resource{
			path:       "/service/icon",
			model:      &models.Service{},
			fileConfig: configs.ResourceConfig.GetFileConfig(types.ImageConfig),
			field:      "Icon",
			policy:     &policies.ServicePolicy{},
		},
	}
}

func NewServiceMainImageResource() domains.ResourceInterface {
	return &mainImageServiceResource{
		Resource: Resource{
			path:       "/service/main_image",
			model:      &models.Service{},
			fileConfig: configs.ResourceConfig.GetFileConfig(types.ImageConfig),
			field:      "MainImage",
			policy:     &policies.ServicePolicy{},
		},
	}
}
