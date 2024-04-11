package domains

import (
	repository_domains "github.com/salamanderman234/outsourcing-api/app/domains/repositories"
	service_domains "github.com/salamanderman234/outsourcing-api/app/domains/services"
	view_domains "github.com/salamanderman234/outsourcing-api/app/domains/views"
	"gorm.io/gorm"
)

type repoRegistry struct {
	BaseRepo repository_domains.BaseRepositoryInterface
	UserRepo repository_domains.UserRepositoryInterface
}

type serviceRegistry struct {
	AuthService           service_domains.AuthServiceInterface
	UserService           service_domains.UserServiceInterface
	ProfileService        service_domains.ProfileServiceInterface
	MasterProvinceService service_domains.MasterProvinceServiceInterface
	MasterRegencyService  service_domains.MasterRegencyServiceInterface
}

type viewRegistry struct {
	AuthView           view_domains.AuthViewInterface
	MasterProvinceView view_domains.MasterProvinceViewInterface
	MasterRegencyView  view_domains.MasterRegencyViewInterface
}

// entity
var RepoRegistry repoRegistry
var ServiceRegistry serviceRegistry
var ViewRegistry viewRegistry

// database connection
var Connection *gorm.DB
