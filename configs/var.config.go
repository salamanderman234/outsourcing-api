package configs

import (
	auth_types "github.com/salamanderman234/outsourcing-api/app/domains/types/auth"
)

type varConfig struct {
	AuthCookieName  string
	UserContextName auth_types.AuthSessionKey
}

var VarConfig varConfig

func (v *varConfig) setVarConfig() {
	v.AuthCookieName = "token"
	v.UserContextName = auth_types.UserContextKey
}
