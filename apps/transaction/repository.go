package transaction

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/mhmdiamd/go-ecommerce-ddd-architecture/infra/response"
)

type repository struct {
  db *sqlx.DB
}

func newRepository(db *sqlx.DB) repository {
  return repository{
    db: db,
  }
}

func (r repository) Begin(ctx context.Context) (tx *sqlx.Tx, err error) {
  tx, err = r.db.BeginTxx(ctx, &sql.TxOptions{})

  return
}

func (r repository) Commit(ctx context.Context, tx *sqlx.Tx) (err error) {
  return tx.Commit()
}

func (r repository) Rollback(ctx context.Context, tx *sqlx.Tx) (err error) {
  return tx.Rollback()
}

func(r repository) GetTransactionsByUserPublicId(ctx context.Context, userPublcId string) (trxs []Transaction, err error){

  query := `
    SELECT 
      id, user_public_id, product_id, product_price
      , amount, sub_total, platform_fee
      , grand_total, status, product_snapshot
      , created_at, updated_at
    FROM transactions
    WHERE user_public_id=$1
  `

  err = r.db.SelectContext(ctx, &trxs, query, userPublcId)
  if err != nil {
    if err == sql.ErrNoRows {
      err = response.ErrorNotFound
      return
    }
    return 
  }

  return
}

func (r repository) CreateTransactionWithTx(ctx context.Context, tx *sqlx.Tx, trx Transaction) (err error) {
  query := `
    INSERT INTO transactions (
      user_public_id, product_id, product_price
			, amount, sub_total, platform_fee
			, grand_total, status, product_snapshot
			, created_at, updated_at
    ) VALUES (
      :user_public_id, :product_id, :product_price
			, :amount, :sub_total, :platform_fee
			, :grand_total, :status, :product_snapshot
			, :created_at, :updated_at
    )
  `

  stmt, err := tx.PrepareNamedContext(ctx, query)
  if err != nil {
    return
  }

  defer stmt.Close()

  _, err = stmt.ExecContext(ctx, trx)

  return
}

func(r repository) GetProductBySKU(ctx context.Context, productSKU string) (product Product, err error){

  query := `
    SELECT 
      id, sku, name, stock, price
    FROM products
    WHERE sku=$1
  `

  err = r.db.GetContext(ctx, &product, query, productSKU)
  if err != nil {
    if err == sql.ErrNoRows {
      return Product{}, response.ErrNotFound
    }
    return 
  }

  return
}

// UpdateProductStockWithTx implements repository
func (r repository) UpdateProductStockWithTx(ctx context.Context, tx *sqlx.Tx, product Product) (err error){
  
  query := `
    UPDATE products
    SET stock=:stock
    WHERE id=:id
  `

  stmt, err := r.db.PrepareNamedContext(ctx, query)
  if err != nil {
    return
  }

  _, err = stmt.ExecContext(ctx, product)
  defer stmt.Close()

  return
}










