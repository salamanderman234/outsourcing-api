package domains

// ----- AUTH VIEW -----
type BasicAuthView interface {
	Login() error
	Register() error
	Verify() error
	Refresh() error
}

// ----- END OF AUTH VIEW -----
