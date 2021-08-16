package repo

import (
	"github.com/danisbagus/semimarket-lib/errs"
	"github.com/danisbagus/semimarket-lib/logger"
	"github.com/danisbagus/semimarket-product/internal/core/domain"
	"github.com/danisbagus/semimarket-product/internal/core/port"
	"github.com/jmoiron/sqlx"
)

type TransactionRepo struct {
	db *sqlx.DB
}

func NewTransactionRepo(db *sqlx.DB) port.ITransactionRepo {
	return &TransactionRepo{
		db: db,
	}
}

func (r TransactionRepo) Create(data *domain.TransactionModel) (*domain.TransactionModel, *errs.AppError) {
	// starting the database transaction block
	tx, err := r.db.Begin()
	if err != nil {
		logger.Error("Error when starting new transaction for product transaction " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	// inserting product transaction
	result, err := tx.Exec(`insert into transactions (product_id, quantity, transaction_date) 
											values (?, ?, ?)`, data.ProductID, data.Quantity, data.TransactionDate)

	if err != nil {
		tx.Rollback()
		logger.Error("Error while create new transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	// updating product quantity
	_, err = tx.Exec(`update products set quantity = quantity - ? where product_id = ?`, data.Quantity, data.ProductID)

	// if error, rollback the transaction
	if err != nil {
		tx.Rollback()
		logger.Error("Error while update product quantity: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	// commit the transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("Error while commiting transaction for product: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	transactionId, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting the last transaction id: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	data.TransactionID = transactionId

	return data, nil
}
