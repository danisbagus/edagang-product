package domain

type ProductModel struct {
	ProductID       int64  `db:"product_id"`
	ProductName     string `db:"product_name"`
	ProductCategory string `db:"product_category"`
	Quantity        int64  `db:"quantity"`
}
