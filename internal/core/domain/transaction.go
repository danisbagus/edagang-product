package domain

type TransactionModel struct {
	TransactionID   int64  `db:"transaction_id"`
	ProductID       int64  `db:"transaction_id"`
	Quantity        int64  `db:"quantity"`
	TransactionDate string `db:"transaction_date"`
}
