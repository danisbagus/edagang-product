package port

import (
	"github.com/danisbagus/edagang-package/errs"
	"github.com/danisbagus/edagang-product/internal/core/domain"
	"github.com/danisbagus/edagang-product/internal/dto"
)

type ITransactionRepo interface {
	Create(data *domain.TransactionModel) (*domain.TransactionModel, *errs.AppError)
}

type ITransactionService interface {
	NewTransaction(data *dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError)
}
