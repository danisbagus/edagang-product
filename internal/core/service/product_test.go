package service

import (
	"testing"

	"github.com/danisbagus/semimarket-product/internal/core/domain"
	"github.com/danisbagus/semimarket-product/internal/core/port"
	"github.com/danisbagus/semimarket-product/internal/dto"
	"github.com/danisbagus/semimarket-product/mocks/repo"
	"github.com/danisbagus/semimarket-product/pkg/errs"
	"github.com/golang/mock/gomock"
)

var mockRepo *repo.MockIProductRepo
var service port.IProducService

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockRepo = repo.NewMockIProductRepo(ctrl)
	service = NewProductService(mockRepo)

	return func() {
		service = nil
		defer ctrl.Finish()
	}
}

func Test_should_return_validation_error_response_when_request_is_not_validated(t *testing.T) {
	// Arrange
	request := dto.NewProductRequest{
		ProductName:     "Oppo A600",
		ProductCategory: "ELECTRONIC",
		Quantity:        0,
	}
	service := NewProductService(nil)

	// Act
	_, err := service.NewProduct(&request)

	// Assert
	if err == nil {
		t.Error("Testing failed while validate new product request")
	}
}

func Test_should_return_error_from_server_side_if_the_new_product_cannot_be_created(t *testing.T) {
	// Arrange
	returnSetup := setup(t)
	defer returnSetup()

	req := dto.NewProductRequest{
		ProductName:     "Oppo A600",
		ProductCategory: "ELECTRONIC",
		Quantity:        100,
	}

	product := domain.ProductModel{
		ProductName:     req.ProductName,
		ProductCategory: req.ProductCategory,
		Quantity:        req.Quantity,
	}

	mockRepo.EXPECT().Create(&product).Return(nil, errs.NewUnexpectedError("Unexpected database error"))

	// Act
	_, err := service.NewProduct(&req)

	// Assert
	if err == nil {
		t.Error("Test failed while validate error for new product")
	}
}

func Test_should_return_response_when_success_to_create_new_product(t *testing.T) {
	// Arrange
	returnSetup := setup(t)
	defer returnSetup()

	req := dto.NewProductRequest{
		ProductName:     "Oppo A600",
		ProductCategory: "ELECTRONIC",
		Quantity:        100,
	}

	product := domain.ProductModel{
		ProductName:     req.ProductName,
		ProductCategory: req.ProductCategory,
		Quantity:        req.Quantity,
	}

	response := product
	response.ProductID = 2

	mockRepo.EXPECT().Create(&product).Return(&response, nil)

	// Act
	NewProduct, err := service.NewProduct(&req)

	// Assert
	if err != nil {
		t.Error("Test failed while create new product")
	}

	if NewProduct.ProductID != response.ProductID {
		t.Error("Test failed while maching new product ID")
	}

}
