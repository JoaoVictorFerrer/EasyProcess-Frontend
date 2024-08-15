package serverErrors

const (
	NoError                   = 0
	GenericError              = 1
	InvalidEmailOrPassword    = 100
	UserAlreadyExistsForEmail = 101
	ExpiredToken              = 102
)
