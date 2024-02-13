package response

import (
	"errors"
	"net/http"
)

var (
  ErrNotFound = errors.New("not found")
  ErrUnauthorized = errors.New("unauthorized")
)

var (
  ErrEmailRequired = errors.New("email is required")
  ErrEmailInvalid = errors.New("email is invalid")
  ErrPasswordRequired = errors.New("password is required")
  ErrPasswordInvalidLength = errors.New("password must have minimum 6 characters")

  // Product
  ErrProductRequired = errors.New("product is required")
  ErrProductInvalid = errors.New("product must have 4 characters")
  ErrStockInvalid = errors.New("stock must be greater than 0")
  ErrPriceInvalid = errors.New("price must be greater than 0")

  // Login
  ErrPasswordNotMatch = errors.New("password is not a match")

  ErrAuthIsNotExists = errors.New("auth is not exists")
  ErrEmailAlreadyUsed = errors.New("email already used")

  // Transaction
  ErrAmountInvalid = errors.New("invalid amount")
  ErrAmountGreaterThanStock = errors.New("amount greater than stock")
)

type Error struct {
  Message string
  Code string
  HttpCode int
}

func NewError(msg string, code string, httpCode int) Error {
  return Error {
    Message: msg,
    Code: code,
    HttpCode: httpCode,
  }
}

func(e Error) Error() string {
  return e.Message
}

// General Error 
var (
  ErrorGeneral = NewError("internal server error", "99999", http.StatusInternalServerError)
  ErrorBadRequest = NewError("bad request", "40000", http.StatusBadRequest)
  ErrorNotFound = NewError(ErrNotFound.Error(), "40401", http.StatusNotFound)
  ErrorUnauthorized = NewError(ErrUnauthorized.Error(), "40100", http.StatusUnauthorized)
)

var (
  // Authentication
  ErrorEmailRequired = NewError(ErrEmailRequired.Error(), "40001", http.StatusBadRequest)
  ErrorEmailInvalid = NewError(ErrEmailInvalid.Error(), "40002", http.StatusBadRequest)
  ErrorPasswordRequired = NewError(ErrPasswordRequired.Error(), "40003", http.StatusBadRequest)
  ErrorPasswordInvalidLength = NewError(ErrPasswordInvalidLength.Error(), "40004", http.StatusBadRequest)

  ErrorPasswordNotMatch = NewError(ErrPasswordNotMatch.Error(), "40101", http.StatusUnauthorized)

  ErrorAuthIsNotExists = NewError(ErrAuthIsNotExists.Error(), "40401", http.StatusNotFound)
  ErrorEmailAlreadyUsed = NewError(ErrEmailAlreadyUsed.Error(), "40901", http.StatusConflict)

  // Product 
  ErrorProductRequired = NewError(ErrProductRequired.Error(), "40005", http.StatusBadRequest)
  ErrorProductInvalid = NewError(ErrProductInvalid.Error(), "40006", http.StatusBadRequest)
  ErrorStockInvalid = NewError(ErrStockInvalid.Error(), "40007", http.StatusBadRequest)
  ErrorPriceInvalid = NewError(ErrPriceInvalid.Error(), "40008", http.StatusBadRequest)
  
  // Transaction
  ErrorAmountInvalid = NewError(ErrAmountInvalid.Error(), "40009", http.StatusBadRequest)
  ErrorAmountGreaterThanStock = NewError(ErrAmountGreaterThanStock.Error(), "40009", http.StatusBadRequest)
)

var (
  ErrorMapping = map[string]Error {
    ErrorNotFound.Error() : ErrorNotFound,
    ErrorUnauthorized.Error() : ErrorUnauthorized,

    ErrorEmailRequired.Error() : ErrorEmailRequired,
    ErrorEmailInvalid.Error() : ErrorEmailInvalid,    
    ErrorPasswordRequired.Error() : ErrorPasswordRequired,
    ErrorPasswordInvalidLength.Error() : ErrorPasswordInvalidLength,

    ErrorPasswordNotMatch.Error() : ErrorPasswordNotMatch,

    ErrorAuthIsNotExists.Error() : ErrorAuthIsNotExists,
    ErrorEmailAlreadyUsed.Error() : ErrorEmailAlreadyUsed,

    // Product error
    ErrorProductRequired.Error() : ErrorProductRequired,
    ErrorProductInvalid.Error() : ErrorProductInvalid,
    ErrorStockInvalid.Error() : ErrorStockInvalid,
    ErrorPriceInvalid.Error() : ErrorPriceInvalid,

    // Transaction
    ErrorAmountInvalid.Error() : ErrorAmountInvalid,
    ErrorAmountGreaterThanStock.Error() : ErrorAmountGreaterThanStock,

  }
)








