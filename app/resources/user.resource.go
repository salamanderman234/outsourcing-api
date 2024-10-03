package resources

import (
	"github.com/salamanderman234/outsourcing-api/app/domains"
	"github.com/salamanderman234/outsourcing-api/app/models"
	"github.com/salamanderman234/outsourcing-api/app/policies"
	"github.com/salamanderman234/outsourcing-api/app/types"
	"github.com/salamanderman234/outsourcing-api/configs"
)

type userResource struct {
	Resource
}

func NewUserProfileResource() domains.ResourceInterface {
	return &userResource{
		Resource{
			path:       "/user/profile",
			model:      &models.User{},
			field:      "ProfilePic",
			fileConfig: configs.ResourceConfig.GetFileConfig(types.ImageConfig),
			policy:     &policies.UserPolicy{},
		},
	}
}
