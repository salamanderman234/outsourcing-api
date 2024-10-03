package enums

type TokenType string

const (
	ResetPasswordTokenType  TokenType = "reset_password"
	AuthenticationTokenType TokenType = "auth"
	AccessTokenType         TokenType = "access"
)
