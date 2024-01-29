package auth

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/google/uuid"
	"github.com/mhmdiamd/go-ecommerce-ddd-architecture/external/database"
	"github.com/mhmdiamd/go-ecommerce-ddd-architecture/helper"
	"github.com/mhmdiamd/go-ecommerce-ddd-architecture/infra/response"
	"github.com/mhmdiamd/go-ecommerce-ddd-architecture/internal/config"
	"github.com/stretchr/testify/require"
)

var svc service 

func init() {
  filename := "../../cmd/api/config.yaml"
  err := config.LoadConfig(filename)

  db, err := database.ConnectPostgres(config.Cfg.DB)
  helper.PanicIfError(err)

  repo := newRepository(db)
  svc = NewService(repo)
}

func TestRegister_Success(t *testing.T) {
  t.Run("Success", func(t *testing.T) {
    req := RegisterRequestPayload{
      Email: fmt.Sprintf("%vam@gmail.com", uuid.NewString()),
      Password: "password",
    }

    err := svc.register(context.Background(), req)
    require.Nil(t, err)
  })
}

func TestRegister_Fail(t *testing.T) {
  t.Run("Email already registred", func(t *testing.T) {

    // Preparation for duplicate Email
    email := fmt.Sprintf("%+vam@gmail.com", uuid.NewString())
    req := RegisterRequestPayload{
      Email: email,
      Password: "password",
    }

    err := svc.register(context.Background(), req)
    require.Nil(t, err)
    // End Preparation

    err = svc.register(context.Background(), req)
    require.Equal(t, response.ErrEmailAlreadyUsed, err)
  })
}

func TestLogin_Success(t *testing.T) {
  email := fmt.Sprintf("%v@noobee.id", uuid.NewString())
  pass := "password"
  req := RegisterRequestPayload {
    Email: email,
    Password: pass,
  }

  err := svc.register(context.Background(), req)
  require.Nil(t, err)

  reqLogin := LoginRequestPayload{
    Email : email,
    Password: pass,
  }

  token, err := svc.login(context.Background(), reqLogin)
  require.Nil(t, err)
  require.NotEmpty(t, token)
  log.Println(token)
}

