package port

import (
	"github.com/danisbagus/semimarket-lib/errs"
	"github.com/danisbagus/semimarket-product/internal/core/domain"
	"github.com/danisbagus/semimarket-product/internal/dto"
)

type ITransactionRepo interface {
	Create(data *domain.TransactionModel) (*domain.TransactionModel, *errs.AppError)
}

type ITransactionService interface {
	NewTransaction(data *dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError)
}
