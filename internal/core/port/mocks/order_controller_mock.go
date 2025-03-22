// Code generated by MockGen. DO NOT EDIT.
// Source: internal/core/port/order_controller_port.go
//
// Generated by this command:
//
//	mockgen -source=internal/core/port/order_controller_port.go -destination=internal/core/port/mocks/order_controller_mock.go
//

// Package mock_port is a generated GoMock package.
package mock_port

import (
	context "context"
	reflect "reflect"

	dto "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/dto"
	port "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
	gomock "go.uber.org/mock/gomock"
)

// MockOrderController is a mock of OrderController interface.
type MockOrderController struct {
	ctrl     *gomock.Controller
	recorder *MockOrderControllerMockRecorder
	isgomock struct{}
}

// MockOrderControllerMockRecorder is the mock recorder for MockOrderController.
type MockOrderControllerMockRecorder struct {
	mock *MockOrderController
}

// NewMockOrderController creates a new mock instance.
func NewMockOrderController(ctrl *gomock.Controller) *MockOrderController {
	mock := &MockOrderController{ctrl: ctrl}
	mock.recorder = &MockOrderControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderController) EXPECT() *MockOrderControllerMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockOrderController) Create(ctx context.Context, presenter port.Presenter, input dto.CreateOrderInput) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, presenter, input)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockOrderControllerMockRecorder) Create(ctx, presenter, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockOrderController)(nil).Create), ctx, presenter, input)
}

// Delete mocks base method.
func (m *MockOrderController) Delete(ctx context.Context, presenter port.Presenter, input dto.DeleteOrderInput) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, presenter, input)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockOrderControllerMockRecorder) Delete(ctx, presenter, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockOrderController)(nil).Delete), ctx, presenter, input)
}

// Get mocks base method.
func (m *MockOrderController) Get(ctx context.Context, presenter port.Presenter, input dto.GetOrderInput) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, presenter, input)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockOrderControllerMockRecorder) Get(ctx, presenter, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockOrderController)(nil).Get), ctx, presenter, input)
}

// List mocks base method.
func (m *MockOrderController) List(ctx context.Context, presenter port.Presenter, input dto.ListOrdersInput) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, presenter, input)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockOrderControllerMockRecorder) List(ctx, presenter, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockOrderController)(nil).List), ctx, presenter, input)
}

// Update mocks base method.
func (m *MockOrderController) Update(ctx context.Context, presenter port.Presenter, input dto.UpdateOrderInput) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, presenter, input)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockOrderControllerMockRecorder) Update(ctx, presenter, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockOrderController)(nil).Update), ctx, presenter, input)
}
