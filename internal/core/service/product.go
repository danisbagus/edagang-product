package service

import (
	"github.com/danisbagus/semimarket-product/internal/core/domain"
	"github.com/danisbagus/semimarket-product/internal/core/port"
	"github.com/danisbagus/semimarket-product/internal/dto"
	"github.com/danisbagus/semimarket-product/pkg/errs"
)

type ProductService struct {
	repo port.IProductRepo
}

func NewProductService(repo port.IProductRepo) port.IProducService {
	return &ProductService{
		repo: repo,
	}
}

func (r ProductService) GetAll() (*dto.ProductListResponse, *errs.AppError) {

	dataList, err := r.repo.FindAll()
	if err != nil {
		return nil, err
	}

	response := dto.NewGetListProductResponse(dataList)

	return response, nil
}

func (r ProductService) GetDetail(productID int64) (*dto.ProductResponse, *errs.AppError) {
	data, err := r.repo.FindOneByID(productID)
	if err != nil {
		return nil, err
	}

	response := dto.NewGetDetailProductResponse(data)

	return response, nil
}

func (r ProductService) NewProduct(req *dto.NewProductRequest) (*dto.NewProductResponse, *errs.AppError) {

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	form := domain.ProductModel{
		ProductName:     req.ProductName,
		ProductCategory: req.ProductCategory,
		Quantity:        req.Quantity,
	}

	newData, err := r.repo.Create(&form)
	if err != nil {
		return nil, err
	}
	response := dto.NewNewProductResponse(newData)

	return response, nil
}
