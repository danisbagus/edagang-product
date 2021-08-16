package port

import (
	"github.com/danisbagus/semimarket-product/internal/core/domain"
	"github.com/danisbagus/semimarket-product/internal/dto"
	"github.com/danisbagus/semimarket-product/pkg/errs"
)

//go:generate mockgen -destination=../../../mocks/service/mockProductService.go -package=service github.com/danisbagus/semimarket-product/internal/core/port IProducService

type IProductRepo interface {
	FindAll() ([]domain.ProductModel, *errs.AppError)
	FindOneByID(productID int64) (*domain.ProductModel, *errs.AppError)
	Create(data *domain.ProductModel) (*domain.ProductModel, *errs.AppError)
}

type IProducService interface {
	GetAll() (*dto.ProductListResponse, *errs.AppError)
	GetDetail(productID int64) (*dto.ProductResponse, *errs.AppError)
	NewProduct(data *dto.NewProductRequest) (*dto.NewProductResponse, *errs.AppError)
}
