package views

import "github.com/salamanderman234/outsourcing-api/app/domains"

func RegisterAllViews() {
	domains.ViewRegistry.AuthView = NewAuthView()
	domains.ViewRegistry.MasterProvinceView = NewMasterProvinceView()
	domains.ViewRegistry.MasterRegencyView = NewRegencyMasterView()
}
