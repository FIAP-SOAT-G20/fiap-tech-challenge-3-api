// Code generated by MockGen. DO NOT EDIT.
// Source: internal/core/port/order_product_gateway_port.go
//
// Generated by this command:
//
//	mockgen -source=internal/core/port/order_product_gateway_port.go -destination=internal/core/port/mocks/order_product_gateway_mock.go
//

// Package mock_port is a generated GoMock package.
package mock_port

import (
	context "context"
	reflect "reflect"

	entity "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	gomock "go.uber.org/mock/gomock"
)

// MockOrderProductGateway is a mock of OrderProductGateway interface.
type MockOrderProductGateway struct {
	ctrl     *gomock.Controller
	recorder *MockOrderProductGatewayMockRecorder
	isgomock struct{}
}

// MockOrderProductGatewayMockRecorder is the mock recorder for MockOrderProductGateway.
type MockOrderProductGatewayMockRecorder struct {
	mock *MockOrderProductGateway
}

// NewMockOrderProductGateway creates a new mock instance.
func NewMockOrderProductGateway(ctrl *gomock.Controller) *MockOrderProductGateway {
	mock := &MockOrderProductGateway{ctrl: ctrl}
	mock.recorder = &MockOrderProductGatewayMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderProductGateway) EXPECT() *MockOrderProductGatewayMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockOrderProductGateway) Create(ctx context.Context, orderProduct *entity.OrderProduct) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, orderProduct)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockOrderProductGatewayMockRecorder) Create(ctx, orderProduct any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockOrderProductGateway)(nil).Create), ctx, orderProduct)
}

// Delete mocks base method.
func (m *MockOrderProductGateway) Delete(ctx context.Context, orderId, productId uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, orderId, productId)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockOrderProductGatewayMockRecorder) Delete(ctx, orderId, productId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockOrderProductGateway)(nil).Delete), ctx, orderId, productId)
}

// FindAll mocks base method.
func (m *MockOrderProductGateway) FindAll(ctx context.Context, orderId, productId uint64, page, limit int) ([]*entity.OrderProduct, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", ctx, orderId, productId, page, limit)
	ret0, _ := ret[0].([]*entity.OrderProduct)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// FindAll indicates an expected call of FindAll.
func (mr *MockOrderProductGatewayMockRecorder) FindAll(ctx, orderId, productId, page, limit any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockOrderProductGateway)(nil).FindAll), ctx, orderId, productId, page, limit)
}

// FindByID mocks base method.
func (m *MockOrderProductGateway) FindByID(ctx context.Context, orderId, productId uint64) (*entity.OrderProduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", ctx, orderId, productId)
	ret0, _ := ret[0].(*entity.OrderProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockOrderProductGatewayMockRecorder) FindByID(ctx, orderId, productId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockOrderProductGateway)(nil).FindByID), ctx, orderId, productId)
}

// Update mocks base method.
func (m *MockOrderProductGateway) Update(ctx context.Context, orderProduct *entity.OrderProduct) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, orderProduct)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockOrderProductGatewayMockRecorder) Update(ctx, orderProduct any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockOrderProductGateway)(nil).Update), ctx, orderProduct)
}
