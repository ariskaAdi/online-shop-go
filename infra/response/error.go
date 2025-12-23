package response

import "errors"

// error general
var (
	ErrNotFound = errors.New("data not found")
)

var (
	ErrEmailRequired    = errors.New("email is required")
	ErrEmailInvalid     = errors.New("email is invalid")
	ErrPasswordRequired = errors.New("password is required")
	ErrPasswordInvalid  = errors.New("password is must be at least 6 characters")
	ErrAuthIsNotExist = errors.New("auth is not exist")
	ErrEmailAlreadyExist = errors.New("email already exist")
	ErrPasswordNotMatch = errors.New("password not match")
)

type Error struct {
	Message string
	Code    string
}

func NewError(message string, code string) Error {
	return Error{Message: message, Code: code}
}

func (e Error) Error() string {
	return e.Message
}

var (
	ErrorGeneral = NewError("general error", "99999")
	ErrorBadRequest = NewError("bad request", "40000")
)

var (
	ErrorEmailRequired = NewError(ErrEmailRequired.Error(), "40001")
	ErrorEmailInvalid = NewError(ErrEmailInvalid.Error(), "40002")
	ErrorPasswordRequired = NewError(ErrPasswordRequired.Error(), "40003")
	ErrorPasswordInvalid = NewError(ErrPasswordInvalid.Error(), "40004")
	ErrorAuthIsNotExist = NewError(ErrAuthIsNotExist.Error(), "40401")
	ErrorEmailAlreadyExist = NewError(ErrEmailAlreadyExist.Error(), "40901")
	ErrorPasswordNotMatch = NewError(ErrPasswordNotMatch.Error(), "40101")
)

var (
	ErrorMapping = map[string]Error {
		ErrEmailRequired.Error(): ErrorEmailRequired,
		ErrEmailInvalid.Error(): ErrorEmailInvalid,
		ErrPasswordRequired.Error(): ErrorPasswordRequired,
		ErrPasswordInvalid.Error(): ErrorPasswordInvalid,
		ErrAuthIsNotExist.Error(): ErrorAuthIsNotExist,
		ErrEmailAlreadyExist.Error(): ErrorEmailAlreadyExist,
		ErrPasswordNotMatch.Error(): ErrorPasswordNotMatch,
	}
)