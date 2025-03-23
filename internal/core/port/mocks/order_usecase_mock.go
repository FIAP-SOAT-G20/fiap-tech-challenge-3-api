// Code generated by MockGen. DO NOT EDIT.
// Source: internal/core/port/order_usecase_port.go
//
// Generated by this command:
//
//	mockgen -source=internal/core/port/order_usecase_port.go -destination=internal/core/port/mocks/order_usecase_mock.go
//

// Package mock_port is a generated GoMock package.
package mock_port

import (
	context "context"
	reflect "reflect"

	entity "github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/domain/entity"
	dto "github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/dto"
	gomock "go.uber.org/mock/gomock"
)

// MockOrderUseCase is a mock of OrderUseCase interface.
type MockOrderUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockOrderUseCaseMockRecorder
	isgomock struct{}
}

// MockOrderUseCaseMockRecorder is the mock recorder for MockOrderUseCase.
type MockOrderUseCaseMockRecorder struct {
	mock *MockOrderUseCase
}

// NewMockOrderUseCase creates a new mock instance.
func NewMockOrderUseCase(ctrl *gomock.Controller) *MockOrderUseCase {
	mock := &MockOrderUseCase{ctrl: ctrl}
	mock.recorder = &MockOrderUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderUseCase) EXPECT() *MockOrderUseCaseMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockOrderUseCase) Create(ctx context.Context, input dto.CreateOrderInput) (*entity.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, input)
	ret0, _ := ret[0].(*entity.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockOrderUseCaseMockRecorder) Create(ctx, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockOrderUseCase)(nil).Create), ctx, input)
}

// Delete mocks base method.
func (m *MockOrderUseCase) Delete(ctx context.Context, input dto.DeleteOrderInput) (*entity.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, input)
	ret0, _ := ret[0].(*entity.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockOrderUseCaseMockRecorder) Delete(ctx, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockOrderUseCase)(nil).Delete), ctx, input)
}

// Get mocks base method.
func (m *MockOrderUseCase) Get(ctx context.Context, input dto.GetOrderInput) (*entity.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, input)
	ret0, _ := ret[0].(*entity.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockOrderUseCaseMockRecorder) Get(ctx, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockOrderUseCase)(nil).Get), ctx, input)
}

// List mocks base method.
func (m *MockOrderUseCase) List(ctx context.Context, input dto.ListOrdersInput) ([]*entity.Order, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, input)
	ret0, _ := ret[0].([]*entity.Order)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockOrderUseCaseMockRecorder) List(ctx, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockOrderUseCase)(nil).List), ctx, input)
}

// Update mocks base method.
func (m *MockOrderUseCase) Update(ctx context.Context, input dto.UpdateOrderInput) (*entity.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, input)
	ret0, _ := ret[0].(*entity.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockOrderUseCaseMockRecorder) Update(ctx, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockOrderUseCase)(nil).Update), ctx, input)
}
