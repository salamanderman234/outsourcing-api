package domains

// --> BASIC
type Policy interface {
	Create(payload JWTClaims) bool
	Find(id uint, payload JWTClaims) bool
	ReadAll(payload JWTClaims) bool
	Update(id uint, payload JWTClaims) bool
	Delete(id uint, payload JWTClaims) bool
}

// ----- AUTH POLICY -----
type ServiceUserAuthPolicy struct{}

func (ServiceUserAuthPolicy) Register(payload JWTPayload) bool {
	return true
}

// ----- END OF AUTH POLICY -----
// ----- MASTER POLICY -----
type CategoryPolicy struct{}

func (CategoryPolicy) Create(payload JWTClaims) bool {
	return *payload.Role == string(AdminRole)
}
func (CategoryPolicy) Find(id uint, payload JWTClaims) bool {
	return true
}
func (CategoryPolicy) ReadAll(payload JWTClaims) bool {
	return true
}
func (CategoryPolicy) Update(id uint, payload JWTClaims) bool {
	return *payload.Role == string(AdminRole)
}
func (CategoryPolicy) Delete(id uint, payload JWTClaims) bool {
	return *payload.Role == string(AdminRole)
}

type DistrictPolicy struct{}

func (DistrictPolicy) Create(payload JWTClaims) bool {
	return *payload.Role == string(AdminRole)
}
func (DistrictPolicy) Find(id uint, payload JWTClaims) bool {
	return true
}
func (DistrictPolicy) ReadAll(payload JWTClaims) bool {
	return true
}
func (DistrictPolicy) Update(id uint, payload JWTClaims) bool {
	return *payload.Role == string(AdminRole)
}
func (DistrictPolicy) Delete(id uint, payload JWTClaims) bool {
	return *payload.Role == string(AdminRole)
}

type SubDistrictPolicy struct{}

func (SubDistrictPolicy) Create(payload JWTClaims) bool {
	return *payload.Role == string(AdminRole)
}
func (SubDistrictPolicy) Find(id uint, payload JWTClaims) bool {
	return true
}
func (SubDistrictPolicy) ReadAll(payload JWTClaims) bool {
	return true
}
func (SubDistrictPolicy) Update(id uint, payload JWTClaims) bool {
	return *payload.Role == string(AdminRole)
}
func (SubDistrictPolicy) Delete(id uint, payload JWTClaims) bool {
	return *payload.Role == string(AdminRole)
}

type VillagePolicy struct{}

func (VillagePolicy) Create(payload JWTClaims) bool {
	return *payload.Role == string(AdminRole)
}
func (VillagePolicy) Find(id uint, payload JWTClaims) bool {
	return true
}
func (VillagePolicy) ReadAll(payload JWTClaims) bool {
	return true
}
func (VillagePolicy) Update(id uint, payload JWTClaims) bool {
	return *payload.Role == string(AdminRole)
}
func (VillagePolicy) Delete(id uint, payload JWTClaims) bool {
	return *payload.Role == string(AdminRole)
}

// ----- END OF MASTER POLICY -----
