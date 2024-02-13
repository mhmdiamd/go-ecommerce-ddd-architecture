package product

import (
	"context"
	"log"
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

func TestCreateProduct_Success(t *testing.T) {
  req := CreateProductRequestPayload{
    Name: "Baju Baru",
    Stock: 3,
    Price: 50_000,
  }

  err := svc.CreateProduct(context.Background(), req)
  require.Nil(t, err)
}

func TestCreateProduct_Fail(t *testing.T) {
  t.Run("name is required", func(t *testing.T) {
    req := CreateProductRequestPayload{
      Stock: 3,
      Price: 50_000,
    }

    err := svc.CreateProduct(context.Background(), req)
    require.NotNil(t, err)
    require.Equal(t, response.ErrProductRequired, err)
  })
  
  t.Run("name is invalid", func(t *testing.T) {
    req := CreateProductRequestPayload{
      Name: "ad",
      Stock: 3,
      Price: 50_000,
    }

    err := svc.CreateProduct(context.Background(), req)
    require.NotNil(t, err)
    require.Equal(t, response.ErrProductInvalid, err)
  })

  t.Run("stock is invalid", func(t *testing.T) {
    req := CreateProductRequestPayload{
      Name: "Baju New",
      Price: 50_000,
    }

    err := svc.CreateProduct(context.Background(), req)
    require.NotNil(t, err)
    require.Equal(t, response.ErrStockInvalid, err)
  })

  t.Run("price is invalid", func(t *testing.T) {
    req := CreateProductRequestPayload{
      Name: "Baju New",
      Stock: 3,
    }

    err := svc.CreateProduct(context.Background(), req)
    require.NotNil(t, err)
    require.Equal(t, response.ErrPriceInvalid, err)
  })
}

func TestListProduct_Success(t *testing.T) {
  pagination := ListProductRequestPayload{
    Cursor: 0,
    Size: 10,
  }

  products, err := svc.ListProducts(context.Background(), pagination)
  require.NotNil(t, products)
  require.Nil(t, err)
  log.Printf("%+v", products)
}

func TestListProductDetail_Success(t *testing.T) {
  req := CreateProductRequestPayload{
    Name : "Baju Baru",
    Stock : 10,
    Price : 10_000,
  }

  ctx := context.Background()

  err := svc.CreateProduct(ctx, req)

  require.Nil(t, err)

  products, err := svc.ListProducts(ctx, ListProductRequestPayload{
    Cursor: 0,
    Size: 10,
  })

  require.Nil(t, err)
  require.NotNil(t, products)
  require.Greater(t, len(products), 0)

  product, err := svc.ProductDetail(ctx, products[0].SKU)
  require.Nil(t, err)
  require.NotEmpty(t, product)
}





