// Code generated by MockGen. DO NOT EDIT.
// Source: internal/core/port/product_controller_port.go
//
// Generated by this command:
//
//	mockgen -source=internal/core/port/product_controller_port.go -destination=internal/core/port/mocks/product_controller_mock.go
//

// Package mock_port is a generated GoMock package.
package mock_port

import (
	context "context"
	reflect "reflect"

	dto "github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/dto"
	port "github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/port"
	gomock "go.uber.org/mock/gomock"
)

// MockProductController is a mock of ProductController interface.
type MockProductController struct {
	ctrl     *gomock.Controller
	recorder *MockProductControllerMockRecorder
	isgomock struct{}
}

// MockProductControllerMockRecorder is the mock recorder for MockProductController.
type MockProductControllerMockRecorder struct {
	mock *MockProductController
}

// NewMockProductController creates a new mock instance.
func NewMockProductController(ctrl *gomock.Controller) *MockProductController {
	mock := &MockProductController{ctrl: ctrl}
	mock.recorder = &MockProductControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductController) EXPECT() *MockProductControllerMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockProductController) Create(ctx context.Context, presenter port.Presenter, input dto.CreateProductInput) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, presenter, input)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockProductControllerMockRecorder) Create(ctx, presenter, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockProductController)(nil).Create), ctx, presenter, input)
}

// Delete mocks base method.
func (m *MockProductController) Delete(ctx context.Context, presenter port.Presenter, input dto.DeleteProductInput) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, presenter, input)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockProductControllerMockRecorder) Delete(ctx, presenter, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockProductController)(nil).Delete), ctx, presenter, input)
}

// Get mocks base method.
func (m *MockProductController) Get(ctx context.Context, presenter port.Presenter, input dto.GetProductInput) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, presenter, input)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockProductControllerMockRecorder) Get(ctx, presenter, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockProductController)(nil).Get), ctx, presenter, input)
}

// List mocks base method.
func (m *MockProductController) List(ctx context.Context, presenter port.Presenter, input dto.ListProductsInput) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, presenter, input)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockProductControllerMockRecorder) List(ctx, presenter, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockProductController)(nil).List), ctx, presenter, input)
}

// Update mocks base method.
func (m *MockProductController) Update(ctx context.Context, presenter port.Presenter, input dto.UpdateProductInput) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, presenter, input)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockProductControllerMockRecorder) Update(ctx, presenter, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockProductController)(nil).Update), ctx, presenter, input)
}
