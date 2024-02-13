package product

import (
	"testing"

	"github.com/mhmdiamd/go-ecommerce-ddd-architecture/infra/response"
	"github.com/stretchr/testify/require"
)

func TestValidateProduct(t *testing.T) {
  t.Run("success", func(t *testing.T) {
    product := Product{
      Name: "Baju Baru",
      Stock : 100,
      Price: 10_000,
    }

    err := product.Validate()
    require.Nil(t, err)
  })

  t.Run("product required", func(t *testing.T) {
    product := Product{
      Stock : 100,
      Price: 10_000,
    }

    err := product.Validate()
    require.NotNil(t, err)
    require.Equal(t, response.ErrProductRequired, err)
  })

  t.Run("product is invalid", func(t *testing.T) {
    product := Product{
      Name: "a",
      Stock : 100,
      Price: 10_000,
    }

    err := product.Validate()
    require.NotNil(t, err)
    require.Equal(t, response.ErrProductInvalid, err)
  })
}

func TestValidateStock(t *testing.T) {
  t.Run("stock is invalid", func(t *testing.T) {
    product := Product{
      Name: "Baju Baru",
      Stock : 0,
      Price: 10_000,
    }

    err := product.Validate()
    require.NotNil(t, err)
    require.Equal(t, response.ErrStockInvalid, err)
  })
}


func TestValidatePrice(t *testing.T) {
  t.Run("price is invalid", func(t *testing.T) {
    product := Product{
      Name: "Baju Baru",
      Stock : 4,
      Price: 0,
    }

    err := product.Validate()
    require.NotNil(t, err)
    require.Equal(t, response.ErrPriceInvalid, err)
  })

}
