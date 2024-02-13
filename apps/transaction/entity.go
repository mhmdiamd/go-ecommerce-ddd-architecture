package transaction

import (
	"encoding/json"
	"time"

	"github.com/mhmdiamd/go-ecommerce-ddd-architecture/infra/response"
)

type TransactionStatus uint8

const (
  TransactionStatus_Created TransactionStatus = 1
  TransactionStatus_Progress TransactionStatus = 10
  TransactionStatus_InDelivery TransactionStatus = 15
  TransactionStatus_Completed TransactionStatus = 20

  TRX_CREATED = "CREATED"
  TRX_PROGRESS = "PROGRESS"
  TRX_IN_DELIVERY = "IN DELIVERY"
  TRX_COMPLETED = "COMPLETED"
  TRX_UNKNOWN = "UNKNOWN"
)

var (
  MappingTransactionStatus = map[TransactionStatus]string {
    TransactionStatus_Created : TRX_CREATED,
    TransactionStatus_Progress : TRX_PROGRESS,
    TransactionStatus_InDelivery : TRX_IN_DELIVERY,
    TransactionStatus_Completed : TRX_COMPLETED,
  }
)

type Transaction struct {
  Id int `db:"id"`
  UserPublicId string `db:"user_public_id"`
  ProductId uint `db:"product_id"`
  ProductPrice uint `db:"product_price"`
  Amount uint8 `db:"amount"`
  SubTotal uint `db:"sub_total"`
  PlatformFee uint `db:"platform_fee"`
  GrandTotal uint `db:"grand_total"`
  Status TransactionStatus `db:"status"`
  ProductJSON json.RawMessage `db:"product_snapshot"`
  CreatedAt time.Time `db:"created_at"`
  UpdatedAt time.Time `db:"updated_at"`
}

func NewTransaction(userPublicId string) Transaction {
  return Transaction{
    UserPublicId : userPublicId,
    Status : TransactionStatus_Created,
    CreatedAt : time.Now(),
    UpdatedAt : time.Now(),
  }
}

func NewTransactionFromCreateRequest(req CreateTransactionRequestPayload) Transaction {
  return Transaction{
    UserPublicId : req.UserPublicId,
    Status: TransactionStatus_Created,
    Amount: req.Amount,
    CreatedAt: time.Now(),
    UpdatedAt: time.Now(),
  }
}

func (t Transaction) Validate() (err error) {
  if t.Amount == 0 {
    return response.ErrAmountInvalid
  }

  return
}

func (t Transaction) ValidateStock(productStock uint8) (err error) {
  if t.Amount > productStock {
    return response.ErrAmountGreaterThanStock  
  }

  return
}

func (t *Transaction) SetSubTotal() {
  if t.SubTotal == 0 {
    t.SubTotal = t.ProductPrice * uint(t.Amount)
  }
}

func (t *Transaction) SetGrandTotal() *Transaction {
  if t.GrandTotal == 0 {
    t.SetSubTotal()

    t.GrandTotal = t.SubTotal + t.PlatformFee
  }

  return t
}

func (t *Transaction) SetPlatformFee(platformfee uint) *Transaction {
  t.PlatformFee = platformfee

  return t
}

// set product id, price and json
func (t *Transaction) FromProduct(product Product) *Transaction {
  t.ProductId = uint(product.Id)
  t.ProductPrice = uint(product.Price)

  t.SetProductJSON(product)

  return t
}

func (t *Transaction) SetProductJSON(product Product) (err error) {
  productJSON, err := json.Marshal(product)
  if err != nil {
    return
  }

  t.ProductJSON = productJSON
  return
}

func (t Transaction) GetProduct() (product Product, err error) {
  err = json.Unmarshal(t.ProductJSON, &product)
  if err != nil {
    return
  }

  return
}

func (t Transaction) GetStatus() string {
  status, ok := MappingTransactionStatus[t.Status]
  if !ok {
    return TRX_UNKNOWN
  }

  return status
}




