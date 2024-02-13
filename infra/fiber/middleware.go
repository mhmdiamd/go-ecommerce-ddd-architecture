package infrafiber

import (
	"fmt"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/mhmdiamd/go-ecommerce-ddd-architecture/infra/response"
	"github.com/mhmdiamd/go-ecommerce-ddd-architecture/internal/config"
	"github.com/mhmdiamd/go-ecommerce-ddd-architecture/utility"
)

func Trace() fiber.Handler {
  return func(c *fiber.Ctx) error {
    // ctx := c.UserContext()   

    // Get Request
    // logger.Log.Infof(ctx, "incoming request")
    
    err := c.Next()
  
    // Finish request
    // logger.Log.Infof(ctx, "incoming request")

    return err
  }
}

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
      return NewResponse(
        WithError(response.ErrorUnauthorized),
      ).Send(c)
    }

    c.Locals("ROLE", role)
    c.Locals("PUBLIC_ID", publicId)

    return c.Next()
  }
}

func CheckRoles(authorizedRoles []string) fiber.Handler {
  return func(c *fiber.Ctx) error {

    role := fmt.Sprintf("%v", c.Locals("ROLE"))

    isExists := false
    for _, authorizedRole := range authorizedRoles {
      if role == authorizedRole {
        isExists = true
        break
      }
    } 

    if !isExists {
      return NewResponse(
        WithError(response.ErrorForbiddenAccess),
      ).Send(c)
    }

    return c.Next()
  }
}










