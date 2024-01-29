package database

import (
	"testing"

	"github.com/mhmdiamd/go-ecommerce-ddd-architecture/internal/config"
	"github.com/stretchr/testify/require"
)

func init() {
  filename := "../../cmd/api/config.yaml"
  err := config.LoadConfig(filename)

  if err != nil {
    panic(err)
  }
}

func testConnectionPostgres(t *testing.T) {
  t.Run("Success", func(t *testing.T) {
    db, err := ConnectPostgres(config.Cfg.DB)
    require.Nil(t, err)
    require.NotNil(t, db)
  })
}
