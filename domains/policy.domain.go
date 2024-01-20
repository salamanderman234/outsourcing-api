package domains

// --> BASIC
type Policy interface {
	Create(user UserEntity, data any) bool
	Find(id uint, user UserEntity, data any) bool
	ReadAll(user UserEntity, datas any) bool
	Update(id uint, user UserEntity, data any) bool
	Delete(id uint, user UserEntity, data any) bool
}

type BasicAdminOnlyPolicy struct{}

func (BasicAdminOnlyPolicy) Create(user UserEntity, data any) bool {
	return user.Role == string(AdminRole)
}
func (BasicAdminOnlyPolicy) Find(id uint, user UserEntity, data any) bool {
	return true
}
func (BasicAdminOnlyPolicy) ReadAll(user UserEntity, datas any) bool {
	return true
}
func (BasicAdminOnlyPolicy) Update(id uint, user UserEntity, data any) bool {
	return user.Role == string(AdminRole)
}
func (BasicAdminOnlyPolicy) Delete(id uint, user UserEntity, data any) bool {
	return user.Role == string(AdminRole)
}

// ----- AUTH POLICY -----
type ServiceUserAuthPolicy struct{}

func (ServiceUserAuthPolicy) Register(user UserEntity) bool {
	return user.Role == string(AdminRole)
}

type UserAuthPolicy struct{}

func (UserAuthPolicy) Create(user UserEntity, data any) bool {
	return false
}
func (UserAuthPolicy) Find(id uint, user UserEntity, data any) bool {
	return user.Role == string(AdminRole) || user.ID == id
}
func (UserAuthPolicy) ReadAll(user UserEntity, datas any) bool {
	return user.Role == string(AdminRole)
}
func (UserAuthPolicy) Update(id uint, user UserEntity, data any) bool {
	return user.Role == string(AdminRole) || user.ID == id
}
func (UserAuthPolicy) Delete(id uint, user UserEntity, data any) bool {
	return user.Role == string(AdminRole)
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
type ServiceOrderPolicy struct{}

func (ServiceOrderPolicy) Create(user UserEntity, data any) bool {
	return user.Role == string(ServiceUserRole)
}
func (ServiceOrderPolicy) Find(id uint, user UserEntity, data any) bool {
	dataModel, ok := data.(ServiceOrderModel)
	if !ok {
		return false
	}

	return dataModel.ServiceUserID == &user.ID
}
func (ServiceOrderPolicy) ReadAll(user UserEntity, datas any) bool {
	datasModel, ok := datas.([]ServiceOrderModel)
	if !ok {
		return false
	}
	id := user.ID
	for _, data := range datasModel {
		if data.ServiceUserID != &id {
			return false
		}
	}
	return true
}
func (ServiceOrderPolicy) Update(id uint, user UserEntity, data any) bool {
	dataModel, ok := data.(ServiceOrderModel)
	if !ok {
		return false
	}
	role := user.Role
	status := *dataModel.Status
	if role == string(ServiceUserRole) {
		return ((status == CancelledOrderStatus) || (status == WaitingForConfirmationOrderStatus) &&
			(user.ID == *dataModel.ServiceUserID))
	}
	if role == string(AdminRole) {
		return (status != CancelledOrderStatus) && (status != WaitingForConfirmationOrderStatus)
	}
	return false
}
func (ServiceOrderPolicy) Delete(id uint, user UserEntity, data any) bool {
	return false
}

// ----- END OF SERVICE ORDER POLICY -----
// ----- PLACEMENT POLICY -----
type PlacementPolicy struct {
	BasicAdminOnlyPolicy
}

// ----- END OF PLACEMENT POLICY -----
