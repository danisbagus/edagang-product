package service

import (
	"time"

	"github.com/danisbagus/semimarket-lib/errs"
	"github.com/danisbagus/semimarket-product/internal/core/domain"
	"github.com/danisbagus/semimarket-product/internal/core/port"
	"github.com/danisbagus/semimarket-product/internal/dto"
)

const dbTSLayout = "2006-01-02 15:04:05"

type TransactionService struct {
	repo        port.ITransactionRepo
	productRepo port.IProductRepo
}

func NewTransactionService(repo port.ITransactionRepo, productRepo port.IProductRepo) port.ITransactionService {
	return &TransactionService{
		repo:        repo,
		productRepo: productRepo,
	}
}

func (r TransactionService) NewTransaction(req *dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError) {

	// incoming request validation
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	// server side validation
	product, err := r.productRepo.FindOneByID(req.ProductID)
	if err != nil {
		return nil, err
	}
	if product.Quantity < req.ProductID {
		return nil, errs.NewValidationError("Insufficient product quantity")
	}

	form := domain.TransactionModel{
		ProductID:       req.ProductID,
		Quantity:        req.Quantity,
		TransactionDate: time.Now().Format(dbTSLayout),
	}

	newData, err := r.repo.Create(&form)
	if err != nil {
		return nil, err
	}
	response := dto.NewNewTransactionResponse(newData)

	return response, nil
}
