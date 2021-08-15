package dto

import (
	"github.com/danisbagus/semimarket-product/internal/core/domain"
	"github.com/danisbagus/semimarket-product/pkg/errs"
	validation "github.com/go-ozzo/ozzo-validation"
)

var ValidProductCategories = map[string]bool{
	"ELECTRONIC": true,
	"CLOTHES":    true,
	"VEHICLE":    true,
}

type ProductResponse struct {
	ProductID       int64  `json:"product_id"`
	ProductName     string `json:"product_name"`
	ProductCategory string `json:"product_category"`
	Quantity        int64  `json:"quantity"`
}

type ProductListResponse struct {
	Products []ProductResponse `json:"data"`
}

type NewProductRequest struct {
	ProductName     string `json:"product_name"`
	ProductCategory string `json:"product_category"`
	Quantity        int64  `json:"quantity"`
}

type NewProductResponse struct {
	ProductID int64 `json:"product_id"`
}

func NewGetListProductResponse(data []domain.ProductModel) *ProductListResponse {
	dataList := make([]ProductResponse, len(data))

	for k, v := range data {
		dataList[k] = ProductResponse{
			ProductID:       v.ProductID,
			ProductName:     v.ProductName,
			ProductCategory: v.ProductCategory,
			Quantity:        v.Quantity,
		}
	}
	return &ProductListResponse{Products: dataList}
}

func NewGetDetailProductResponse(data *domain.ProductModel) *ProductResponse {
	result := &ProductResponse{
		ProductID:       data.ProductID,
		ProductName:     data.ProductName,
		ProductCategory: data.ProductCategory,
		Quantity:        data.Quantity,
	}
	return result
}

func NewNewProductResponse(data *domain.ProductModel) *NewProductResponse {
	result := &NewProductResponse{
		ProductID: data.ProductID,
	}

	return result
}

func (r NewProductRequest) Validate() *errs.AppError {

	if err := validation.Validate(r.ProductName, validation.Required); err != nil {
		return errs.NewBadRequestError("Product name is required")

	}

	if err := validation.Validate(r.ProductCategory, validation.Required); err != nil {
		return errs.NewBadRequestError("Product category name is required")

	}

	if err := validation.Validate(r.Quantity, validation.Required); err != nil {
		return errs.NewBadRequestError("Product quantity is required")

	}

	if !ValidProductCategories[r.ProductCategory] {
		return errs.NewValidationError("Product category is not valid")
	}

	if r.Quantity <= 0 {
		return errs.NewValidationError("Product quantity must more than 0")
	}
	return nil
}
