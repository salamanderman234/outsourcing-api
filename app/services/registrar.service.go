package services

import (
	"github.com/salamanderman234/outsourcing-api/app/providers"
)

func RegisterAllServices() {
	providers.ServiceProvider.AuthService = NewAuthService()
	providers.ServiceProvider.MasterProvinceService = NewMasterProvinceService()
	providers.ServiceProvider.MasterRegencyService = NewRegencyMasterService()
	providers.ServiceProvider.MasterCategoryService = NewCategoryMasterService()
	providers.ServiceProvider.AppServiceService = NewApplicationServiceService()
	providers.ServiceProvider.AppPackageServiceService = NewApplicationPackageService()
	providers.ServiceProvider.TransactionService = NewTransactionService()
	providers.ServiceProvider.FileService = NewFileService()
}
