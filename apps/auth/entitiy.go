package auth

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/mhmdiamd/go-ecommerce-ddd-architecture/infra/response"
	"github.com/mhmdiamd/go-ecommerce-ddd-architecture/utility"
	"golang.org/x/crypto/bcrypt"
)

type Role string

const (
  ROLE_Admin Role = "admin"
  ROLE_User Role = "user"
)

type AuthEntity struct {
  Id int `db:"id"`
  Email string `db:"email"`
  PublicId uuid.UUID `db:"public_id"`
  Password string `db:"password"`
  Role Role `db:"role"` 

  CreatedAt time.Time `db:"created_at"`
  UpdatedAt time.Time `db:"updated_at"`
}

func NewFromRegisterRequest(req RegisterRequestPayload) AuthEntity {
  return AuthEntity{
    Email : req.Email,
    Password: req.Password,
    PublicId: uuid.New(),
    Role: ROLE_User,
    CreatedAt: time.Now(),
    UpdatedAt: time.Now(),
  }
}

func NewFromLoginRequest(req LoginRequestPayload) AuthEntity {
  return AuthEntity{
    Email : req.Email,
    Password: req.Password,
  }
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

func (a AuthEntity) IsExists() bool {
  return a.Id != 0
}

func (a *AuthEntity) EncryptPassword(salt int) (err error) {
  encryptedPass, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)
  if err != nil {
    return
  }

  a.Password = string(encryptedPass)
  return 
}

func (a AuthEntity) VerifyPasswordFromEncrypted(plain string) (err error) {
  return bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(plain))
}

func (a AuthEntity) VerifyPasswordFromPlain(encrypted string) (err error) {
  return bcrypt.CompareHashAndPassword([]byte(encrypted), []byte(a.Password))
}

func (a AuthEntity) GenerateToken(secret string) (tokenString string, err error) {
 return utility.GenerateToken(a.PublicId.String(), string(a.Role), secret)
}














