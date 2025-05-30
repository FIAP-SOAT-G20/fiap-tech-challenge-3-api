// Code generated by MockGen. DO NOT EDIT.
// Source: internal/core/port/category_gateway_port.go
//
// Generated by this command:
//
//	mockgen -source=internal/core/port/category_gateway_port.go -destination=internal/core/port/mocks/category_gateway_mock.go
//

// Package mock_port is a generated GoMock package.
package mock_port

import (
	context "context"
	reflect "reflect"

	entity "github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/domain/entity"
	gomock "go.uber.org/mock/gomock"
)

// MockCategoryGateway is a mock of CategoryGateway interface.
type MockCategoryGateway struct {
	ctrl     *gomock.Controller
	recorder *MockCategoryGatewayMockRecorder
	isgomock struct{}
}

// MockCategoryGatewayMockRecorder is the mock recorder for MockCategoryGateway.
type MockCategoryGatewayMockRecorder struct {
	mock *MockCategoryGateway
}

// NewMockCategoryGateway creates a new mock instance.
func NewMockCategoryGateway(ctrl *gomock.Controller) *MockCategoryGateway {
	mock := &MockCategoryGateway{ctrl: ctrl}
	mock.recorder = &MockCategoryGatewayMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCategoryGateway) EXPECT() *MockCategoryGatewayMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockCategoryGateway) Create(ctx context.Context, category *entity.Category) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, category)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockCategoryGatewayMockRecorder) Create(ctx, category any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCategoryGateway)(nil).Create), ctx, category)
}

// Delete mocks base method.
func (m *MockCategoryGateway) Delete(ctx context.Context, id uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockCategoryGatewayMockRecorder) Delete(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCategoryGateway)(nil).Delete), ctx, id)
}

// FindAll mocks base method.
func (m *MockCategoryGateway) FindAll(ctx context.Context, name string, page, limit int) ([]*entity.Category, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", ctx, name, page, limit)
	ret0, _ := ret[0].([]*entity.Category)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// FindAll indicates an expected call of FindAll.
func (mr *MockCategoryGatewayMockRecorder) FindAll(ctx, name, page, limit any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockCategoryGateway)(nil).FindAll), ctx, name, page, limit)
}

// FindByID mocks base method.
func (m *MockCategoryGateway) FindByID(ctx context.Context, id uint64) (*entity.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", ctx, id)
	ret0, _ := ret[0].(*entity.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockCategoryGatewayMockRecorder) FindByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockCategoryGateway)(nil).FindByID), ctx, id)
}

// Update mocks base method.
func (m *MockCategoryGateway) Update(ctx context.Context, category *entity.Category) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, category)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockCategoryGatewayMockRecorder) Update(ctx, category any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCategoryGateway)(nil).Update), ctx, category)
}
