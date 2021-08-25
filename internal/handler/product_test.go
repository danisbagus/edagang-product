package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/danisbagus/edagang-package/errs"
	"github.com/danisbagus/edagang-product/internal/dto"
	"github.com/danisbagus/edagang-product/mocks/service"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
)

var router *mux.Router
var productHandler ProductHandler
var mockService *service.MockIProducService

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockService = service.NewMockIProducService(ctrl)

	productHandler = ProductHandler{mockService}
	router = mux.NewRouter()

	router.HandleFunc("/products", productHandler.GetProductList)

	return func() {
		router = nil
		defer ctrl.Finish()
	}
}

func Test_should_return_products_with_status_200(t *testing.T) {
	// Arrange
	setupReturn := setup(t)
	defer setupReturn()

	dummyProducts := dto.ProductListResponse{
		Products: []dto.ProductResponse{
			{ProductID: 1, ProductName: "Oppo V1", ProductCategory: "ELECTRONICS", Quantity: 100},
			{ProductID: 2, ProductName: "Xiaomi Redmi 4x", ProductCategory: "ELECTRONICS", Quantity: 50},
		},
	}

	mockService.EXPECT().GetAll().Return(&dummyProducts, nil)
	request, _ := http.NewRequest(http.MethodGet, "/products", nil)

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusOK {
		t.Error("Failed when the status code")
	}
}

func Test_should_return_products_with_status_500_with_error_message(t *testing.T) {
	setupReturn := setup(t)
	defer setupReturn()

	mockService.EXPECT().GetAll().Return(nil, errs.NewUnexpectedError("database error"))
	request, _ := http.NewRequest(http.MethodGet, "/products", nil)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusInternalServerError {
		t.Error("Failed when the status code")
	}
}
