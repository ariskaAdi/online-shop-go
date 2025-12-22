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
)