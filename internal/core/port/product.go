package port

import (
	"github.com/danisbagus/edagang-pkg/errs"
	"github.com/danisbagus/edagang-product/internal/core/domain"
	"github.com/danisbagus/edagang-product/internal/dto"
)

//go:generate mockgen -destination=../../../mocks/repo/mockProductRepo.go -package=repo github.com/danisbagus/edagang-product/internal/core/port IProductRepo
type IProductRepo interface {
	FindAll() ([]domain.ProductModel, *errs.AppError)
	FindOneByID(productID int64) (*domain.ProductModel, *errs.AppError)
	Create(data *domain.ProductModel) (*domain.ProductModel, *errs.AppError)
	Delete(productID int64) *errs.AppError
	Update(productID int64, data *domain.ProductModel) *errs.AppError
}

//go:generate mockgen -destination=../../../mocks/service/mockProductService.go -package=service github.com/danisbagus/edagang-product/internal/core/port IProducService
type IProducService interface {
	GetAll() (*dto.ProductListResponse, *errs.AppError)
	GetDetail(productID int64) (*dto.ProductResponse, *errs.AppError)
	NewProduct(data *dto.NewProductRequest) (*dto.NewProductResponse, *errs.AppError)
	RemoveProduct(ProductID int64) *errs.AppError
	UpdateProduct(productID int64, data *dto.NewProductRequest) (*dto.ProductResponse, *errs.AppError)
}
