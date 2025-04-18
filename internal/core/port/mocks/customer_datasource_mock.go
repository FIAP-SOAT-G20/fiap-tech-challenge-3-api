// Code generated by MockGen. DO NOT EDIT.
// Source: internal/core/port/customer_datasource_port.go
//
// Generated by this command:
//
//	mockgen -source=internal/core/port/customer_datasource_port.go -destination=internal/core/port/mocks/customer_datasource_mock.go
//

// Package mock_port is a generated GoMock package.
package mock_port

import (
	context "context"
	reflect "reflect"

	entity "github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/domain/entity"
	gomock "go.uber.org/mock/gomock"
)

// MockCustomerDataSource is a mock of CustomerDataSource interface.
type MockCustomerDataSource struct {
	ctrl     *gomock.Controller
	recorder *MockCustomerDataSourceMockRecorder
	isgomock struct{}
}

// MockCustomerDataSourceMockRecorder is the mock recorder for MockCustomerDataSource.
type MockCustomerDataSourceMockRecorder struct {
	mock *MockCustomerDataSource
}

// NewMockCustomerDataSource creates a new mock instance.
func NewMockCustomerDataSource(ctrl *gomock.Controller) *MockCustomerDataSource {
	mock := &MockCustomerDataSource{ctrl: ctrl}
	mock.recorder = &MockCustomerDataSourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCustomerDataSource) EXPECT() *MockCustomerDataSourceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockCustomerDataSource) Create(ctx context.Context, product *entity.Customer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, product)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockCustomerDataSourceMockRecorder) Create(ctx, product any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCustomerDataSource)(nil).Create), ctx, product)
}

// Delete mocks base method.
func (m *MockCustomerDataSource) Delete(ctx context.Context, id uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockCustomerDataSourceMockRecorder) Delete(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCustomerDataSource)(nil).Delete), ctx, id)
}

// FindAll mocks base method.
func (m *MockCustomerDataSource) FindAll(ctx context.Context, filters map[string]any, page, limit int) ([]*entity.Customer, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", ctx, filters, page, limit)
	ret0, _ := ret[0].([]*entity.Customer)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// FindAll indicates an expected call of FindAll.
func (mr *MockCustomerDataSourceMockRecorder) FindAll(ctx, filters, page, limit any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockCustomerDataSource)(nil).FindAll), ctx, filters, page, limit)
}

// FindByCPF mocks base method.
func (m *MockCustomerDataSource) FindByCPF(ctx context.Context, cpf string) (*entity.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByCPF", ctx, cpf)
	ret0, _ := ret[0].(*entity.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByCPF indicates an expected call of FindByCPF.
func (mr *MockCustomerDataSourceMockRecorder) FindByCPF(ctx, cpf any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByCPF", reflect.TypeOf((*MockCustomerDataSource)(nil).FindByCPF), ctx, cpf)
}

// FindByID mocks base method.
func (m *MockCustomerDataSource) FindByID(ctx context.Context, id uint64) (*entity.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", ctx, id)
	ret0, _ := ret[0].(*entity.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockCustomerDataSourceMockRecorder) FindByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockCustomerDataSource)(nil).FindByID), ctx, id)
}

// Transaction mocks base method.
func (m *MockCustomerDataSource) Transaction(ctx context.Context, fn func(context.Context) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Transaction", ctx, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// Transaction indicates an expected call of Transaction.
func (mr *MockCustomerDataSourceMockRecorder) Transaction(ctx, fn any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transaction", reflect.TypeOf((*MockCustomerDataSource)(nil).Transaction), ctx, fn)
}

// Update mocks base method.
func (m *MockCustomerDataSource) Update(ctx context.Context, product *entity.Customer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, product)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockCustomerDataSourceMockRecorder) Update(ctx, product any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCustomerDataSource)(nil).Update), ctx, product)
}
