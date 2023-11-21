package domains

// ----- AUTH POLICY -----
type ServiceUserAuthPolicy struct{}

func (ServiceUserAuthPolicy) Register(payload JWTPayload) bool {
	return true
}

// ----- END OF AUTH POLICY -----
