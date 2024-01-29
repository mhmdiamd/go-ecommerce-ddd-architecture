package auth

import (
	"strings"

	"github.com/mhmdiamd/go-ecommerce-ddd-architecture/infra/response"
)

type Role string

const (
  ROLE_Admin Role = "admin"
  ROLE_User Role = "user"
)

type AuthEntity struct {
  Id string
  Email string
  Password string
  Role Role
}

func (a AuthEntity) Validate() (err error) {
  
  if err = a.EmailValidate(); err != nil {
    return 
  }

  if err = a.PasswordValidate(); err != nil {
    return 
  }

  return 
}

func (a AuthEntity) EmailValidate() (err error) {

  if a.Email == "" {
    return response.ErrEmailRequired
  }

  if len(strings.Split(a.Email, "@")) != 2 {
    return response.ErrEmailInvalid
  }

  return
}

func (a AuthEntity) PasswordValidate() (err error) {
 
  if a.Password == "" {
    return response.ErrPasswordRequired
  }

  if len(a.Password) < 6 {
    return response.ErrPasswordInvalidLength
  }

  return
}














