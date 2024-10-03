package providers

import (
	"github.com/salamanderman234/outsourcing-api/app/domains"
	"github.com/salamanderman234/outsourcing-api/app/resources"
	"github.com/salamanderman234/outsourcing-api/app/types"
)

type fun func() domains.ResourceInterface

type resourceProvider struct {
	modelMap map[string]fun
}

func (m *resourceProvider) CreateResource(alias string) (domains.ResourceInterface, error) {
	fun, ok := m.modelMap[alias]
	if !ok {
		return nil, types.ErrRecordNotFound
	}
	resource := fun()
	return resource, nil
}

var ResourceProvider = resourceProvider{
	modelMap: map[string]fun{
		"packages.main_image": resources.NewMainPackageResource,
		"services.main_image": resources.NewServiceMainImageResource,
		"services.icon":       resources.NewServiceIconResource,
		"transactions.mou":    resources.NewTransactionMouResource,
		"users.profile":       resources.NewUserProfileResource,
		"categories.icon":     resources.NewIconCategoryResource,
	},
}
