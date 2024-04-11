package configs

type routerConfig struct {
	MaxBodyLength string
	MaxTimeOut    int
}

var RouterConfig routerConfig

func (r *routerConfig) setRouterConfig() {
	r.MaxBodyLength = "15M"
	r.MaxTimeOut = 60
}
