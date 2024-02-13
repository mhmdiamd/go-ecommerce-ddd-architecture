package transaction

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	infrafiber "github.com/mhmdiamd/go-ecommerce-ddd-architecture/infra/fiber"
)

func Init(router fiber.Router, db *sqlx.DB) {
  repo := newRepository(db)
  svc := newService(repo)
  handler := newHandler(svc)

  trxRoute := router.Group("transaction")
  {
    trxRoute.Use(infrafiber.CheckAuth())

    trxRoute.Post("/checkout", handler.CreateTransaction)
    trxRoute.Get("/user/histories", handler.GetTransactionByUser)
  }
}
