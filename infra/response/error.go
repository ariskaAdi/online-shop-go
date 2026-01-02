package response

import (
	"errors"
	"net/http"
)

// error general
var (
	ErrNotFound = errors.New("data not found")
	ErrUnauthorized = errors.New("unauthorized")
	ErrForbiddenAccess = errors.New("forbiddden access")
)

var (
	// AUTH RESPONSE ERROR
	ErrEmailRequired    = errors.New("email is required")
	ErrEmailInvalid     = errors.New("email is invalid")
	ErrPasswordRequired = errors.New("password is required")
	ErrPasswordInvalid  = errors.New("password is must be at least 6 characters")
	ErrAuthIsNotExist = errors.New("auth is not exist")
	ErrEmailAlreadyExist = errors.New("email already exist")
	ErrPasswordNotMatch = errors.New("password not match")

	// PRODUCT RESPONSE ERROR
	ErrProductRequired    = errors.New("Product is required")
	ErrProductInvalid     = errors.New("Product is must be at least 4 characters")
	ErrStockInvalid     = errors.New("stock is must greater than 0")
	ErrPriceInvalid     = errors.New("Price is must greater than 0")

	// TRANSACTION RESPONSE ERROR
	ErrAmountInvalid     = errors.New("invalid amount")
	ErrAmountGreaterThanStock = errors.New("amount greater tahn stock")
)

type Error struct {
	Message string
	Code    string
	HttpCode int
}

func NewError(message string, code string, httpCode int) Error {
	return Error{Message: message, Code: code, HttpCode: httpCode}
}

func (e Error) Error() string {
	return e.Message
}

var (
	ErrorGeneral = NewError("general error", "99999", http.StatusInternalServerError)
	ErrorBadRequest = NewError("bad request", "40000", http.StatusBadRequest)
	ErrorNotFound = NewError(ErrNotFound.Error(), "40400", http.StatusNotFound)
	ErrorUnauthorized    = NewError(ErrUnauthorized.Error(), "40100", http.StatusUnauthorized)
	ErrorForbiddenAccess = NewError(ErrForbiddenAccess.Error(), "40100", http.StatusForbidden)
)

var (
	// error bad request
	ErrorEmailRequired = NewError(ErrEmailRequired.Error(), "40001",  http.StatusBadRequest)
	ErrorEmailInvalid = NewError(ErrEmailInvalid.Error(), "40002",  http.StatusBadRequest)
	ErrorPasswordRequired = NewError(ErrPasswordRequired.Error(), "40003",  http.StatusBadRequest)
	ErrorPasswordInvalid = NewError(ErrPasswordInvalid.Error(), "40004",  http.StatusBadRequest)
	ErrorProductRequired = NewError(ErrProductRequired.Error(), "40005",  http.StatusBadRequest)
	ErrorProductInvalid = NewError(ErrProductInvalid.Error(), "40006",  http.StatusBadRequest)
	ErrorStockInvalid = NewError(ErrStockInvalid.Error(), "40007",  http.StatusBadRequest)
	ErrorPriceInvalid = NewError(ErrPriceInvalid.Error(), "40008",  http.StatusBadRequest)


	ErrorAuthIsNotExist = NewError(ErrAuthIsNotExist.Error(), "40401", http.StatusNotFound)
	ErrorEmailAlreadyExist = NewError(ErrEmailAlreadyExist.Error(), "40901", http.StatusConflict)
	ErrorPasswordNotMatch = NewError(ErrPasswordNotMatch.Error(), "40101", http.StatusUnauthorized)

)

var (
	ErrorMapping = map[string]Error {
		ErrNotFound.Error(): ErrorNotFound,
		ErrEmailRequired.Error(): ErrorEmailRequired,
		ErrEmailInvalid.Error(): ErrorEmailInvalid,
		ErrPasswordRequired.Error(): ErrorPasswordRequired,
		ErrPasswordInvalid.Error(): ErrorPasswordInvalid,
		ErrAuthIsNotExist.Error(): ErrorAuthIsNotExist,
		ErrEmailAlreadyExist.Error(): ErrorEmailAlreadyExist,
		ErrPasswordNotMatch.Error(): ErrorPasswordNotMatch,
		ErrUnauthorized.Error():          ErrorUnauthorized,
		ErrForbiddenAccess.Error():       ErrorForbiddenAccess,
	}
)