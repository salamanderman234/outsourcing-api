package domains

// --> BASIC
type Policy interface {
	Create(user UserEntity) bool
	Find(id uint, user UserEntity) bool
	ReadAll(user UserEntity, id ...uint) bool
	Update(id uint, user UserEntity) bool
	Delete(id uint, user UserEntity) bool
}

type BasicAdminOnlyPolicy struct{}

func (BasicAdminOnlyPolicy) Create(user UserEntity) bool {
	return user.Role == string(AdminRole)
}
func (BasicAdminOnlyPolicy) Find(id uint, user UserEntity) bool {
	return true
}
func (BasicAdminOnlyPolicy) ReadAll(user UserEntity, id ...uint) bool {
	return true
}
func (BasicAdminOnlyPolicy) Update(id uint, user UserEntity) bool {
	return user.Role == string(AdminRole)
}
func (BasicAdminOnlyPolicy) Delete(id uint, user UserEntity) bool {
	return user.Role == string(AdminRole)
}

type BasicOnlyServiceUser struct{}

func (BasicOnlyServiceUser) Create(user UserEntity) bool {
	return user.Role == string(ServiceUserRole)
}
func (BasicOnlyServiceUser) Find(id uint, user UserEntity) bool {
	return id == user.ID
}
func (BasicOnlyServiceUser) ReadAll(user UserEntity, id ...uint) bool {
	if len(id) == 1 {
		return user.Role == string(AdminRole) || id[0] == user.ID
	}
	return user.Role == string(AdminRole)
}
func (BasicOnlyServiceUser) Update(id uint, user UserEntity) bool {
	return user.Role == string(AdminRole)
}
func (BasicOnlyServiceUser) Delete(id uint, user UserEntity) bool {
	return id == user.ID
}

// ----- AUTH POLICY -----
type ServiceUserAuthPolicy struct{}

func (ServiceUserAuthPolicy) Register(user UserEntity) bool {
	return true
}

// ----- END OF AUTH POLICY -----
// ----- MASTER POLICY -----
type CategoryPolicy struct {
	BasicAdminOnlyPolicy
}

type DistrictPolicy struct {
	BasicAdminOnlyPolicy
}

type SubDistrictPolicy struct {
	BasicAdminOnlyPolicy
}

type VillagePolicy struct {
	BasicAdminOnlyPolicy
}

// ----- END OF MASTER POLICY -----
// ----- APP SERVICE POLICY -----
type ServiceItemPolicy struct {
	BasicAdminOnlyPolicy
}
type ServicePolicy struct {
	BasicAdminOnlyPolicy
}

// ----- END OF APP SERVICE POLICY -----
// ----- SERVICE ORDER POLICY -----
type ServiceOrderPolicy struct {
	BasicOnlyServiceUser
}

// ----- END OF SERVICE ORDER POLICY -----
