package services

import (
	"github.com/salamanderman234/outsourcing-api/app/providers"
)

func RegisterAllServices() {
	providers.ServiceProvider.AuthService = NewAuthService()
	providers.ServiceProvider.MasterProvinceService = NewMasterProvinceService()
	providers.ServiceProvider.MasterRegencyService = NewRegencyMasterService()
	providers.ServiceProvider.MasterCategoryService = NewCategoryMasterService()
	providers.ServiceProvider.MasterQuestionService = NewQuestionMasterService()
	providers.ServiceProvider.MasterPaymentConfigService = NewPaymentConfigMasterService()
	providers.ServiceProvider.AppServiceService = NewApplicationServiceService()
	providers.ServiceProvider.AppPackageServiceService = NewApplicationPackageService()
	providers.ServiceProvider.TransactionService = NewTransactionService()
	providers.ServiceProvider.FileService = NewFileService()
	providers.ServiceProvider.MidtransService = NewMidtransService()
	providers.ServiceProvider.PlacementService = NewPlacementService()
	providers.ServiceProvider.UserService = NewUserService()
	providers.ServiceProvider.FeedbackService = NewFeedbackService()
	providers.ServiceProvider.ComplaintService = NewComplaintService()
	providers.ServiceProvider.PerformanceService = NewPerformanceService()
}
