package policies

type servicePolicy struct {
	baseAdminOnlyPolicy
}

type packagePolicy struct {
	baseAdminOnlyPolicy
}

var ServicePolicy = servicePolicy{}
var PackagePolicy = packagePolicy{}
