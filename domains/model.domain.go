package domains

// Registry
func ModelRegistry() []any {
	return nil
}

// ----- AUTH MODEL -----
type EmployeeModel struct {
}

func (EmployeeModel) TableName() string {
	return "employees"
}

type ServiceUserModel struct {
}

func (ServiceUserModel) TableName() string {
	return "service_users"
}

type SupervisorModel struct {
}

func (SupervisorModel) TableName() string {
	return "supervisors"
}

type AdminModel struct {
}

func (AdminModel) TableName() string {
	return "admins"
}

// ----- END OF AUTH MODEL -----
