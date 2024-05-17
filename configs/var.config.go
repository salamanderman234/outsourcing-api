package configs

import "github.com/salamanderman234/outsourcing-api/app/types"

type varConfig struct {
	AuthCookieName  string
	UserContextName types.AuthSessionKey
}

var VarConfig varConfig

func (v *varConfig) setVarConfig() {
	v.AuthCookieName = "token"
	v.UserContextName = types.UserContextKey
}
