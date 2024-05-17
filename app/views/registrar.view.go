package views

import (
	"github.com/salamanderman234/outsourcing-api/app/providers"
)

func RegisterAllViews() {
	providers.ViewProvider.AuthView = NewAuthView()
	providers.ViewProvider.MasterProvinceView = NewMasterProvinceView()
	providers.ViewProvider.MasterRegencyView = NewRegencyMasterView()
	providers.ViewProvider.MasterCategoryView = NewCategoryMasterView()
	providers.ViewProvider.ApplicationServiceView = NewApplicationServiceView()
}
