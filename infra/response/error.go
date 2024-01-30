package response

import (
	"errors"
	"net/http"
)

var (
  ErrNotFound = errors.New("not found")
)

var (
  ErrEmailRequired = errors.New("email is required")
  ErrEmailInvalid = errors.New("email is invalid")
  ErrPasswordRequired = errors.New("password is required")
  ErrPasswordInvalidLength = errors.New("password must have minimum 6 characters")

  // Login
  ErrPasswordNotMatch = errors.New("password is not a match")

  ErrAuthIsNotExists = errors.New("auth is not exists")
  ErrEmailAlreadyUsed = errors.New("email already used")
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

var (
  ErrorGeneral = NewError("internal server error", "99999", http.StatusInternalServerError)
  ErrorBadRequest = NewError("bad request", "40000", http.StatusBadRequest)

  ErrorEmailRequired = NewError(ErrEmailRequired.Error(), "40001", http.StatusBadRequest)
  ErrorEmailInvalid = NewError(ErrEmailInvalid.Error(), "40002", http.StatusBadRequest)
  ErrorPasswordRequired = NewError(ErrPasswordRequired.Error(), "40003", http.StatusBadRequest)
  ErrorPasswordInvalidLength = NewError(ErrPasswordInvalidLength.Error(), "40004", http.StatusBadRequest)

  ErrorPasswordNotMatch = NewError(ErrPasswordNotMatch.Error(), "40101", http.StatusUnauthorized)

  ErrorAuthIsNotExists = NewError(ErrAuthIsNotExists.Error(), "40401", http.StatusNotFound)
  ErrorEmailAlreadyUsed = NewError(ErrEmailAlreadyUsed.Error(), "40901", http.StatusConflict)

  ErrorNotFound = NewError(ErrNotFound.Error(), "40401", http.StatusNotFound)
)

var (
  ErrorMapping = map[string]Error {
    ErrorNotFound.Error() : ErrorNotFound,
    ErrorEmailRequired.Error() : ErrorEmailRequired,
    ErrorEmailInvalid.Error() : ErrorEmailInvalid,    
    ErrorPasswordRequired.Error() : ErrorPasswordRequired,
    ErrorPasswordInvalidLength.Error() : ErrorPasswordInvalidLength,

    ErrorPasswordNotMatch.Error() : ErrorPasswordNotMatch,

    ErrorAuthIsNotExists.Error() : ErrorAuthIsNotExists,
    ErrorEmailAlreadyUsed.Error() : ErrorEmailAlreadyUsed,
  }
)

