package transaction

import (
	"log"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)


func TestSetSubTotal(t *testing.T) {
  var trx = Transaction {
    ProductPrice : 10_000,
    Amount: 2,
  }

  expected := uint(20_000) 
  trx.SetSubTotal()
  trx.SetSubTotal()
  trx.SetSubTotal()
  trx.SetSubTotal()
  trx.SetSubTotal()
    
  require.Equal(t, expected , trx.SubTotal)
}

func TestGrandTotal(t *testing.T) {

  t.Run("without set sub total first", func(t *testing.T) {
    var trx = Transaction {
      ProductPrice : 10_000,
      Amount: 2,
    }

    expected := uint(20_000) 
    trx.SetGrandTotal()
    
    require.Equal(t, expected , trx.GrandTotal)
  })

  t.Run("without platform fee", func(t *testing.T) {
    var trx = Transaction {
      ProductPrice : 10_000,
      Amount: 2,
    }

    expected := uint(20_000) 
    trx.SetSubTotal()
    trx.SetGrandTotal()
    
    require.Equal(t, expected , trx.GrandTotal)
  })

  t.Run("with platform fee", func(t *testing.T) {
    var trx = Transaction {
      ProductPrice : 10_000,
      Amount: 2,
      PlatformFee: 1_000, 
    }

    expected := uint(21_000) 
    trx.SetSubTotal()
    trx.SetGrandTotal()
    
    require.Equal(t, expected , trx.GrandTotal)
  })
}

func TestProductJSON(t *testing.T) {
  product := Product{
    Id: 1,
    SKU: uuid.NewString(),
    Name: "Product 1",
    Price: 10_000,
  }

  var trx = Transaction{}
  err := trx.SetProductJSON(product)
  require.Nil(t, err)
  require.NotNil(t, trx.ProductJSON)

  productFromTsx, err := trx.GetProduct()
  require.Nil(t, err)
  require.NotEmpty(t, productFromTsx)

  require.Equal(t, product, productFromTsx)

  log.Println(string(trx.ProductJSON))
}

func TestTransactionStatus(t *testing.T) {

  type tabletest struct {
    title string
    expected string
    trx Transaction
  }

  var tableTests = []tabletest{
    {
      title: "status created",
      expected: TRX_CREATED,
      trx: Transaction{Status : TransactionStatus_Created},
    },
    {
      title: "status progress",
      expected: TRX_PROGRESS,
      trx: Transaction{Status : TransactionStatus_Progress},
    },
    {
      title: "status in delivery",
      expected: TRX_IN_DELIVERY,
      trx: Transaction{Status : TransactionStatus_InDelivery},
    },
    {
      title: "status completed",
      expected: TRX_COMPLETED,
      trx: Transaction{Status : TransactionStatus_Completed},
    },
    {
      title: "status unknown",
      expected: TRX_UNKNOWN,
      trx: Transaction{},
    },
  }

  for _, test:= range tableTests{
    t.Run(test.title, func(t *testing.T) {
      trx := test.trx

      require.Equal(t, test.expected, trx.GetStatus())
    })
  }

}






