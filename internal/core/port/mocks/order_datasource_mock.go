// Code generated by MockGen. DO NOT EDIT.
// Source: internal/core/port/order_datasource_port.go
//
// Generated by this command:
//
//	mockgen -source=internal/core/port/order_datasource_port.go -destination=internal/core/port/mocks/order_datasource_mock.go
//

// Package mock_port is a generated GoMock package.
package mock_port

import (
	context "context"
	reflect "reflect"

	entity "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	gomock "go.uber.org/mock/gomock"
)

// MockOrderDataSource is a mock of OrderDataSource interface.
type MockOrderDataSource struct {
	ctrl     *gomock.Controller
	recorder *MockOrderDataSourceMockRecorder
	isgomock struct{}
}

// MockOrderDataSourceMockRecorder is the mock recorder for MockOrderDataSource.
type MockOrderDataSourceMockRecorder struct {
	mock *MockOrderDataSource
}

// NewMockOrderDataSource creates a new mock instance.
func NewMockOrderDataSource(ctrl *gomock.Controller) *MockOrderDataSource {
	mock := &MockOrderDataSource{ctrl: ctrl}
	mock.recorder = &MockOrderDataSourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderDataSource) EXPECT() *MockOrderDataSourceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockOrderDataSource) Create(ctx context.Context, order *entity.Order) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, order)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockOrderDataSourceMockRecorder) Create(ctx, order any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockOrderDataSource)(nil).Create), ctx, order)
}

// Delete mocks base method.
func (m *MockOrderDataSource) Delete(ctx context.Context, id uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockOrderDataSourceMockRecorder) Delete(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockOrderDataSource)(nil).Delete), ctx, id)
}

// FindAll mocks base method.
func (m *MockOrderDataSource) FindAll(ctx context.Context, filters map[string]any, sort string, page, limit int) ([]*entity.Order, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", ctx, filters, sort, page, limit)
	ret0, _ := ret[0].([]*entity.Order)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// FindAll indicates an expected call of FindAll.
func (mr *MockOrderDataSourceMockRecorder) FindAll(ctx, filters, sort, page, limit any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockOrderDataSource)(nil).FindAll), ctx, filters, sort, page, limit)
}

// FindByID mocks base method.
func (m *MockOrderDataSource) FindByID(ctx context.Context, id uint64) (*entity.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", ctx, id)
	ret0, _ := ret[0].(*entity.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockOrderDataSourceMockRecorder) FindByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockOrderDataSource)(nil).FindByID), ctx, id)
}

// Transaction mocks base method.
func (m *MockOrderDataSource) Transaction(ctx context.Context, fn func(context.Context) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Transaction", ctx, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// Transaction indicates an expected call of Transaction.
func (mr *MockOrderDataSourceMockRecorder) Transaction(ctx, fn any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transaction", reflect.TypeOf((*MockOrderDataSource)(nil).Transaction), ctx, fn)
}

// Update mocks base method.
func (m *MockOrderDataSource) Update(ctx context.Context, order *entity.Order) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, order)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockOrderDataSourceMockRecorder) Update(ctx, order any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockOrderDataSource)(nil).Update), ctx, order)
}
