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
	providers.ViewProvider.ApplicationPackageServiceView = NewApplicationPackageServiceView()
	providers.ViewProvider.TransactionView = NewTransactionView()
	providers.ViewProvider.PaymentView = NewPaymentView()
	providers.ViewProvider.FileView = NewFileView()
	providers.ViewProvider.UserView = NewUserView()
}
