package policies

type masterPolicy struct {
	baseAdminOnlyPolicy
}

var MasterPolicy = masterPolicy{}
