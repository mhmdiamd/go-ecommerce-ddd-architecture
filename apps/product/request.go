package product

type CreateProductRequestPayload struct {
  Id int `db:"id"`
  Name string `db:"name"`
  Stock int16 `db:"stock"`
  Price int `db:"price"`
}

type UpdateProductRequestPayload struct {
  Id int `db:"id"`
  Name string `db:"name"`
  Stock int16 `db:"stock"`
  Price int `db:"price"`
}

type ListProductRequestPayload struct {
  Cursor int `query:"cursor" json:"cursor"`
  Size int `query:"size" json:"size"`
}

func (l ListProductRequestPayload) GenerateDefaultValue() ListProductRequestPayload {
  if l.Cursor < 0 {
    l.Cursor = 0
  }

  if l.Size <= 0 {
    l.Size = 10
  }

  return l
}


