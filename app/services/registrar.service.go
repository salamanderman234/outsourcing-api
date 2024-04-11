package services

import "github.com/salamanderman234/outsourcing-api/app/domains"

func RegisterAllServices() {
	domains.ServiceRegistry.AuthService = NewAuthService()
	domains.ServiceRegistry.MasterProvinceService = NewMasterProvinceService()
}
