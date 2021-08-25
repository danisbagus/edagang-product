package dto

import (
	"github.com/danisbagus/edagang-package/errs"
	"github.com/danisbagus/edagang-product/internal/core/domain"
)

type NewTransactionRequest struct {
	ProductID int64 `json:"product_id"`
	Quantity  int64 `json:"quantity"`
}

type NewTransactionResponse struct {
	TransactionID   int64  `json:"transaction_id"`
	TransactionDate string `json:"transaction_date"`
}

func NewNewTransactionResponse(data *domain.TransactionModel) *NewTransactionResponse {
	result := &NewTransactionResponse{
		TransactionID:   data.TransactionID,
		TransactionDate: data.TransactionDate,
	}

	return result
}

func (r NewTransactionRequest) Validate() *errs.AppError {
	if r.Quantity <= 0 {
		return errs.NewValidationError("Product quantity must more than 0")
	}
	return nil
}
