package product

import (
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
    svc: svc,
  }
}

func (h handler) CreateProduct(ctx *fiber.Ctx) error {
  var req = CreateProductRequestPayload{}

  if err := ctx.BodyParser(&req); err != nil  {
    return infrafiber.NewResponse(
      infrafiber.WithMessage("invalid payload"),
      infrafiber.WithError(response.ErrorBadRequest),
    ).Send(ctx)
  }

  if err := h.svc.CreateProduct(ctx.UserContext(), req); err != nil {

    myErr, ok := response.ErrorMapping[err.Error()]
    if !ok {
      myErr = response.ErrorGeneral
    }

    return infrafiber.NewResponse(
      infrafiber.WithError(myErr),
      infrafiber.WithMessage(err.Error()),
    ).Send(ctx)
  }

  return infrafiber.NewResponse(
    infrafiber.WithHttpCode(http.StatusCreated),
    infrafiber.WithMessage("create product success"),
  ).Send(ctx)
}

func (h handler) GetListProducts(ctx *fiber.Ctx) error {
  var req = ListProductRequestPayload{}

  if err := ctx.QueryParser(&req); err != nil  {
    return infrafiber.NewResponse(
      infrafiber.WithMessage("invalid payload"),
      infrafiber.WithError(response.ErrorBadRequest),
    ).Send(ctx)
  }

  products, err := h.svc.ListProducts(ctx.UserContext(), req)

  if err != nil {
    myErr, ok := response.ErrorMapping[err.Error()]
    if !ok {
      myErr = response.ErrorGeneral
    }

    return infrafiber.NewResponse(
      infrafiber.WithError(myErr),
      infrafiber.WithMessage(err.Error()),
    ).Send(ctx)
  }

  productListResponse := NewProductListResponseFromEntity(products)

  return infrafiber.NewResponse(
    infrafiber.WithHttpCode(http.StatusOK),
    infrafiber.WithMessage("get list product success"),
    infrafiber.WithPayload(productListResponse),
    infrafiber.WithQuery(req),
  ).Send(ctx)
}

func (h handler) GetProductDetail(ctx *fiber.Ctx) error {

  sku := ctx.Params("sku", "")
  if sku == "" {
    return infrafiber.NewResponse(
      infrafiber.WithMessage("invalid payload"),
      infrafiber.WithError(response.ErrorBadRequest),
    ).Send(ctx)
  }

  product, err := h.svc.ProductDetail(ctx.UserContext(), sku)

  if err != nil {
    myErr, ok := response.ErrorMapping[err.Error()]
    if !ok {
      myErr = response.ErrorGeneral
    }

    return infrafiber.NewResponse(
      infrafiber.WithError(myErr),
      infrafiber.WithMessage(err.Error()),
    ).Send(ctx)
  }

  toProductDetail := product.ToProductDetailResponse()

  return infrafiber.NewResponse(
    infrafiber.WithHttpCode(http.StatusOK),
    infrafiber.WithMessage("get detail product success"),
    infrafiber.WithPayload(toProductDetail),
  ).Send(ctx)
}

