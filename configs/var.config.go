package configs

import "github.com/salamanderman234/outsourcing-api/app/types"

type varConfig struct {
	AuthCookieName    string
	UserContextName   types.AuthSessionKey
	AccessContextName types.AuthSessionKey
}

var VarConfig varConfig

func (v *varConfig) setVarConfig() {
	v.AuthCookieName = "session"
	v.UserContextName = types.UserContextKey
	v.AccessContextName = types.AccessContextKey
}
