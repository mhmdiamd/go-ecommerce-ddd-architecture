package auth

import (
	"log"
	"testing"

	"github.com/mhmdiamd/go-ecommerce-ddd-architecture/infra/response"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestValidateAuthEntity(t *testing.T) {
  t.Run("success", func(t *testing.T) {
    authEntity := AuthEntity {
      Email : "am@gmail.com",
      Password: "password",
    }

    err := authEntity.Validate()

    require.Nil(t, err)
  })

  t.Run("Error required email", func (t *testing.T) {
    authEntity := AuthEntity {
      Email : "",
      Password: "password",
    }

    err := authEntity.Validate()

    require.NotNil(t, err)
    require.Equal(t, response.ErrEmailRequired, err)
  }) 

  t.Run("Error Invalid email", func (t *testing.T) {
    authEntity := AuthEntity {
      Email : "amgmail.com",
      Password: "password",
    }

    err := authEntity.Validate()

    require.NotNil(t, err)
    require.Equal(t, response.ErrEmailInvalid, err)
  }) 

  t.Run("Error required password", func (t *testing.T) {
    authEntity := AuthEntity {
      Email : "am@gmail.com",
      Password: "",
    }

    err := authEntity.Validate()

    require.NotNil(t, err)
    require.Equal(t, response.ErrPasswordRequired, err)
  }) 

  t.Run("Error invalid password", func (t *testing.T) {
    authEntity := AuthEntity {
      Email : "am@gmail.com",
      Password: "12345",
    }

    err := authEntity.Validate()

    require.NotNil(t, err)
    require.Equal(t, response.ErrPasswordInvalidLength, err)
  }) 
}

func TestEncryptPassword(t *testing.T) {
  t.Run("Success", func(t *testing.T) {
    authEntity := AuthEntity {
      Email : "am@gmail.com",
      Password: "123456",
    }

    err := authEntity.EncryptPassword(bcrypt.DefaultCost)
    require.Nil(t, err)

    log.Printf("%+v\n", authEntity)
  })
}











