package product

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/mhmdiamd/go-ecommerce-ddd-architecture/apps/auth"
	infrafiber "github.com/mhmdiamd/go-ecommerce-ddd-architecture/infra/fiber"
)

func Init(router fiber.Router, db *sqlx.DB) {
  repo := newRepository(db)
  svc := newService(repo)
  handler := newHandler(svc)

  productRoute := router.Group("products") 
  {
    productRoute.Get("", handler.GetListProducts)
    productRoute.Get("/sku/:sku", handler.GetProductDetail)

    productRoute.Post("", 
      infrafiber.CheckAuth(),
      infrafiber.CheckRoles([]string{string(auth.ROLE_Admin)}),
      handler.CreateProduct,
    )
  }
}
