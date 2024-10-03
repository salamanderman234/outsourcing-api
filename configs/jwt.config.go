package configs

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

type jwtConfig struct {
	Secret        string
	Exp           int
	SigningMethod jwt.SigningMethod
}

var JWTConfig jwtConfig

func (j *jwtConfig) setJWTConfig() {
	j.Exp = viper.GetInt("JWT_EXP")
	j.Secret = viper.GetString("JWT_SECRET")
	j.SigningMethod = jwt.SigningMethodHS256
}

func (j *jwtConfig) GetDefaultEXP() int {
	return j.Exp
}

func (j *jwtConfig) GetSecret() string {
	return j.Secret
}

func (j *jwtConfig) GetSigningMethod() jwt.SigningMethod {
	return j.SigningMethod
}
