package transaction

import (
	"context"
	"testing"

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
  helper.PanicIfError(err)

  db, err := database.ConnectPostgres(config.Cfg.DB)
  helper.PanicIfError(err)

  repo := newRepository(db)
  svc = newService(repo)
}

func TestCreateTransaction(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		req := CreateTransactionRequestPayload{
			ProductSKU:   "6d8c41f9-8697-4e6b-b0cc-5228b2dd9399",
			Amount:       2,
      UserPublicId : "e172a0f4-05bd-4dcc-99d6-5f39f2902754",
		}

		err := svc.CreateTransaction(context.Background(), req)
		require.Nil(t, err)
	})

	t.Run("product not found", func(t *testing.T) {
		req := CreateTransactionRequestPayload{
			ProductSKU:   "6d8c41f9-8697-4e6b-b0cc-5228b2dd9389",
			Amount:       2,
      UserPublicId : "e172a0f4-05bd-4dcc-99d6-5f39f2902754",
		}

		err := svc.CreateTransaction(context.Background(), req)
		require.NotNil(t, err)
		require.Equal(t, response.ErrNotFound, err)
	})

	t.Run("amount is greater than stock", func(t *testing.T) {
		req := CreateTransactionRequestPayload{
			ProductSKU:   "6d8c41f9-8697-4e6b-b0cc-5228b2dd9399",
			Amount:       100,
      UserPublicId : "e172a0f4-05bd-4dcc-99d6-5f39f2902754",
		}

		err := svc.CreateTransaction(context.Background(), req)
		require.NotNil(t, err)
		require.Equal(t, response.ErrAmountGreaterThanStock, err)
	})
}
