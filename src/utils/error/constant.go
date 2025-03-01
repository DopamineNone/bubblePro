package error

const (
	UnableToQueryUserErrorCode = 40000 + iota
	UserAlreadyExistedErrorCode
	UnableToCreateUserErrorCode
	UnableToGenerateCipherErrorCode
	InvalidTokenErrorCode
	UnableToParseTokenErrorCode
)

const (
	UnableToGenerateCipherError = "fail to generate cipher"
	UnableToCreateUserError     = "fail to create user"
	UserAlreadyExistedError     = "user already existed"
	UnableToQueryUserError      = "fail to query user"
	InvalidTokenError           = "invalid token"
	UnableToParseTokenError     = "fail to parse token"
)
