// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/danisbagus/edagang-product/internal/core/port (interfaces: IProducService)

// Package service is a generated GoMock package.
package service

import (
	reflect "reflect"

	errs "github.com/danisbagus/edagang-pkg/errs"
	dto "github.com/danisbagus/edagang-product/internal/dto"
	gomock "github.com/golang/mock/gomock"
)

// MockIProducService is a mock of IProducService interface.
type MockIProducService struct {
	ctrl     *gomock.Controller
	recorder *MockIProducServiceMockRecorder
}

// MockIProducServiceMockRecorder is the mock recorder for MockIProducService.
type MockIProducServiceMockRecorder struct {
	mock *MockIProducService
}

// NewMockIProducService creates a new mock instance.
func NewMockIProducService(ctrl *gomock.Controller) *MockIProducService {
	mock := &MockIProducService{ctrl: ctrl}
	mock.recorder = &MockIProducServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIProducService) EXPECT() *MockIProducServiceMockRecorder {
	return m.recorder
}

// GetAll mocks base method.
func (m *MockIProducService) GetAll() (*dto.ProductListResponse, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].(*dto.ProductListResponse)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockIProducServiceMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockIProducService)(nil).GetAll))
}

// GetDetail mocks base method.
func (m *MockIProducService) GetDetail(arg0 int64) (*dto.ProductResponse, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDetail", arg0)
	ret0, _ := ret[0].(*dto.ProductResponse)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// GetDetail indicates an expected call of GetDetail.
func (mr *MockIProducServiceMockRecorder) GetDetail(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDetail", reflect.TypeOf((*MockIProducService)(nil).GetDetail), arg0)
}

// NewProduct mocks base method.
func (m *MockIProducService) NewProduct(arg0 *dto.NewProductRequest) (*dto.NewProductResponse, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewProduct", arg0)
	ret0, _ := ret[0].(*dto.NewProductResponse)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// NewProduct indicates an expected call of NewProduct.
func (mr *MockIProducServiceMockRecorder) NewProduct(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewProduct", reflect.TypeOf((*MockIProducService)(nil).NewProduct), arg0)
}
