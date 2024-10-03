package providers

import (
	repository_domains "github.com/salamanderman234/outsourcing-api/app/domains/repositories"
	service_domains "github.com/salamanderman234/outsourcing-api/app/domains/services"
	view_domains "github.com/salamanderman234/outsourcing-api/app/domains/views"
)

type repoProvider struct {
	BaseRepo        repository_domains.BaseRepositoryInterface
	UserRepo        repository_domains.UserRepositoryInterface
	PerformanceRepo repository_domains.PerformanceRepositoryInterface
}

type serviceProvider struct {
	AuthService                service_domains.AuthServiceInterface
	UserService                service_domains.UserServiceInterface
	MasterProvinceService      service_domains.MasterProvinceServiceInterface
	MasterRegencyService       service_domains.MasterRegencyServiceInterface
	MasterCategoryService      service_domains.MasterCategoryServiceInterface
	MasterQuestionService      service_domains.MasterQuestionServiceInterface
	MasterPaymentConfigService service_domains.MasterPaymentConfigServiceInterface
	AppServiceService          service_domains.ApplicationServiceServiceInterface
	AppPackageServiceService   service_domains.ApplicationPackageInterface
	TransactionService         service_domains.TransactionServiceInterface
	FileService                service_domains.FileServiceInterface
	MidtransService            service_domains.MidtransServiceInterface
	PlacementService           service_domains.PlacementServiceInterface
	FeedbackService            service_domains.FeedbackServiceInterface
	ComplaintService           service_domains.ComplaintServiceInterface
	PerformanceService         service_domains.PerformanceServiceInterface
}

type viewProvider struct {
	AuthView                      view_domains.AuthViewInterface
	UserView                      view_domains.UserViewInterface
	MasterProvinceView            view_domains.MasterProvinceViewInterface
	MasterRegencyView             view_domains.MasterRegencyViewInterface
	MasterCategoryView            view_domains.MasterCategoryViewInterface
	MasterQuestionView            view_domains.MasterQuestionViewInterface
	MasterPaymentConfigView       view_domains.MasterPaymentConfigViewInterface
	ApplicationServiceView        view_domains.ApplicationServiceViewInterface
	ApplicationPackageServiceView view_domains.ApplicationPackageServiceViewInterface
	TransactionView               view_domains.TransactionViewInterface
	PaymentView                   view_domains.PaymentViewInterface
	FileView                      view_domains.FileViewInterface
	PlacementView                 view_domains.PlacementViewInterface
	FeedbackView                  view_domains.FeedbackViewInterface
	ComplaintView                 view_domains.ComplaintViewInterface
	PerformanceView               view_domains.PerformanceViewInterface
	UIView                        view_domains.UIViewInterface
}

// entity
var RepoProvider repoProvider
var ServiceProvider serviceProvider
var ViewProvider viewProvider
