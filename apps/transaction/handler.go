package transaction

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	infrafiber "github.com/mhmdiamd/go-ecommerce-ddd-architecture/infra/fiber"
	"github.com/mhmdiamd/go-ecommerce-ddd-architecture/infra/response"
)

type handler struct {
  svc service
}

func newHandler(svc service) handler {
  return handler{
    svc : svc,
  }
}

func (h handler) CreateTransaction(ctx *fiber.Ctx) error {

  var req = CreateTransactionRequestPayload{}

  if err := ctx.BodyParser(&req); err != nil {
    myErr := response.ErrorBadRequest

    return infrafiber.NewResponse(
      infrafiber.WithError(myErr),
      infrafiber.WithHttpCode(myErr.HttpCode),
      infrafiber.WithMessage("invalid request"),
    ).Send(ctx)
  }

  userPublicId := ctx.Locals("PUBLIC_ID")
  req.UserPublicId = fmt.Sprintf("%v", userPublicId)

  if err := h.svc.CreateTransaction(ctx.UserContext(), req); err != nil {

    myErr, ok := response.ErrorMapping[err.Error()];

    if !ok {
      myErr = response.ErrorGeneral
    }

    return infrafiber.NewResponse(
      infrafiber.WithError(myErr),
      infrafiber.WithMessage(err.Error()),
    ).Send(ctx)
  }

  return infrafiber.NewResponse(
    infrafiber.WithHttpCode(http.StatusOK),
    infrafiber.WithMessage("create transaction success"),
  ).Send(ctx)
}

func (h handler) GetTransactionByUser(ctx *fiber.Ctx) error {
  userPublicId := fmt.Sprintf("%v", ctx.Locals("PUBLIC_ID"))

  trxs, err := h.svc.TransactionHistories(ctx.UserContext(), userPublicId)
  if err != nil {
    return err
  }

  var response = []TransactionHistoryResponse{}
  
  for _, trx := range trxs {
    response = append(response, trx.ToTransactionHistoryResponse())
  }

  return infrafiber.NewResponse(
    infrafiber.WithHttpCode(http.StatusOK),
    infrafiber.WithMessage("get transaction history success"),
    infrafiber.WithPayload(response),
  ).Send(ctx)

}









