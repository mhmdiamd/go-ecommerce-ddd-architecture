package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/mhmdiamd/go-ecommerce-ddd-architecture/apps/auth"
	"github.com/mhmdiamd/go-ecommerce-ddd-architecture/apps/product"
	"github.com/mhmdiamd/go-ecommerce-ddd-architecture/apps/transaction"
	"github.com/mhmdiamd/go-ecommerce-ddd-architecture/external/database"
	"github.com/mhmdiamd/go-ecommerce-ddd-architecture/internal/config"
)

func main() {
  filename := "./cmd/api/config.yaml"
  if err := config.LoadConfig(filename); err != nil {
    panic(err)
  }

  db, err := database.ConnectPostgres(config.Cfg.DB)

  if err != nil {
    panic(err)
  }

  if db != nil {
    log.Println("db connected")
  }

  router := fiber.New(fiber.Config{
    Prefork: true,
    AppName: config.Cfg.App.Name,
  })

  auth.Init(router, db)
  product.Init(router, db)
  transaction.Init(router, db)

  router.Listen(config.Cfg.App.Port)

}
