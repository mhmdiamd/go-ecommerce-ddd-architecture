package auth

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	infrafiber "github.com/mhmdiamd/go-ecommerce-ddd-architecture/infra/fiber"
	"github.com/mhmdiamd/go-ecommerce-ddd-architecture/infra/response"
)

type handler struct {
 svc service
}

func newHandler(s service) handler {
  return handler{
    svc: s,
  }
}

func (h handler) register(ctx *fiber.Ctx) error {
  var req = RegisterRequestPayload{}
  
  if err := ctx.BodyParser(&req); err != nil {
    myErr := response.ErrorBadRequest 
    return infrafiber.NewResponse(
      infrafiber.WithMessage(err.Error()),
      infrafiber.WithError(myErr),
      infrafiber.WithHttpCode(http.StatusBadRequest),
    ).Send(ctx)
  }

  if err := h.svc.register(ctx.UserContext(), req); err != nil {
    myErr, ok := response.ErrorMapping[err.Error()]

    if !ok {
      return response.ErrorGeneral
    }

    return infrafiber.NewResponse(
      infrafiber.WithMessage("register fail"),
      infrafiber.WithError(myErr),
      infrafiber.WithHttpCode(http.StatusBadRequest),
    ).Send(ctx)
  }

  return infrafiber.NewResponse(
    infrafiber.WithHttpCode(http.StatusOK),
    infrafiber.WithMessage("register success"),
  ).Send(ctx)

}

func (h handler) login(ctx *fiber.Ctx) error {
  var req = LoginRequestPayload{}
  
  if err := ctx.BodyParser(&req); err != nil {
    myErr := response.ErrorBadRequest 
    return infrafiber.NewResponse(
      infrafiber.WithMessage(err.Error()),
      infrafiber.WithError(myErr),
    ).Send(ctx)
  }

  token, err := h.svc.login(ctx.UserContext(), req);
  if err != nil {
    myErr, ok := response.ErrorMapping[err.Error()]

    if !ok {
      return response.ErrorGeneral
    }

    return infrafiber.NewResponse(
      infrafiber.WithMessage("login fail"),
      infrafiber.WithError(myErr),
    ).Send(ctx)
  }

  return infrafiber.NewResponse(
    infrafiber.WithMessage("login success"),
    infrafiber.WithHttpCode(http.StatusOK),
    infrafiber.WithPayload(map[string]interface{}{
      "access_token" : token, 
    }),
  ).Send(ctx)

}


