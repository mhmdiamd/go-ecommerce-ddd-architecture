package infrafiber

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/mhmdiamd/go-ecommerce-ddd-architecture/infra/response"
	"github.com/mhmdiamd/go-ecommerce-ddd-architecture/internal/config"
	"github.com/mhmdiamd/go-ecommerce-ddd-architecture/utility"
)

func CheckAuth() fiber.Handler {
  return func (c *fiber.Ctx) error {
    authorization := c.Get("Authorization")

    if authorization == "" {
      return NewResponse(
        WithError(response.ErrorUnauthorized),
      ).Send(c)
    }

    bearer := strings.Split(authorization, "Bearer ")
    if len(bearer) != 2 {
      log.Println("Token invalid")
      return NewResponse(
        WithError(response.ErrorUnauthorized),
      ).Send(c)
    }

    token := bearer[1]

    publicId, role, err := utility.ValidateToken(token, config.Cfg.App.Encryption.JWTSecret)
    if err != nil {
      log.Println(err.Error())
      return NewResponse(
        WithError(response.ErrorUnauthorized),
      ).Send(c)
    }

    c.Locals("ROLE", role)
    c.Locals("PUBLIC_ID", publicId)

    return c.Next()
  }
}
