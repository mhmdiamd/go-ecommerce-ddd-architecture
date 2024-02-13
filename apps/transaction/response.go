package transaction

import (
	"encoding/json"
	"time"
)

type TransactionHistoryResponse struct {
  Id int `json:"id"`
  UserPublicId string `json:"user_public_id"`
  ProductId uint `json:"product_id"`
  ProductPrice uint `json:"product_price"`
  Amount uint8 `json:"amount"`
  SubTotal uint `json:"sub_total"`
  PlatformFee uint `json:"platform_fee"`
  GrandTotal uint `json:"grand_total"`
  Status TransactionStatus `json:"status"`
  ProductJSON json.RawMessage `json:"product_snapshot"`
  CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`

  Product Product `json:"product"`
}
