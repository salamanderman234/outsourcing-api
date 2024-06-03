package resources

import (
	"github.com/salamanderman234/outsourcing-api/app/domains"
	"github.com/salamanderman234/outsourcing-api/app/models"
	"github.com/salamanderman234/outsourcing-api/app/policies"
	"github.com/salamanderman234/outsourcing-api/app/types"
	"github.com/salamanderman234/outsourcing-api/configs"
)

type mainImagePackageResource struct {
	Resource
}

func NewMainPackageResource() domains.ResourceInterface {
	return &mainImagePackageResource{
		Resource: Resource{
			path:       "/package/main_image",
			model:      &models.Package{},
			fileConfig: configs.ResourceConfig.GetFileConfig(types.ImageConfig),
			field:      "MainImage",
			policy:     &policies.PackagePolicy{},
		},
	}
}
