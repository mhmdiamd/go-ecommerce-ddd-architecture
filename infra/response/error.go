package response

import "errors"

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
