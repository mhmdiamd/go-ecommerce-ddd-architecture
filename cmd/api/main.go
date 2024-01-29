package main

import (
	"log"

	"github.com/mhmdiamd/go-ecommerce-ddd-architecture/external/database"
	"github.com/mhmdiamd/go-ecommerce-ddd-architecture/internal/config"
)

func main() {
  filename := "config.yaml"
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
}
