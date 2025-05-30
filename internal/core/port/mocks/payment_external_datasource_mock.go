// Code generated by MockGen. DO NOT EDIT.
// Source: internal/core/port/payment_external_datasource_port.go
//
// Generated by this command:
//
//	mockgen -source=internal/core/port/payment_external_datasource_port.go -destination=internal/core/port/mocks/payment_external_datasource_mock.go
//

// Package mock_port is a generated GoMock package.
package mock_port

import (
	context "context"
	reflect "reflect"

	entity "github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/domain/entity"
	gomock "go.uber.org/mock/gomock"
)

// MockPaymentExternalDatasource is a mock of PaymentExternalDatasource interface.
type MockPaymentExternalDatasource struct {
	ctrl     *gomock.Controller
	recorder *MockPaymentExternalDatasourceMockRecorder
	isgomock struct{}
}

// MockPaymentExternalDatasourceMockRecorder is the mock recorder for MockPaymentExternalDatasource.
type MockPaymentExternalDatasourceMockRecorder struct {
	mock *MockPaymentExternalDatasource
}

// NewMockPaymentExternalDatasource creates a new mock instance.
func NewMockPaymentExternalDatasource(ctrl *gomock.Controller) *MockPaymentExternalDatasource {
	mock := &MockPaymentExternalDatasource{ctrl: ctrl}
	mock.recorder = &MockPaymentExternalDatasourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPaymentExternalDatasource) EXPECT() *MockPaymentExternalDatasourceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockPaymentExternalDatasource) Create(context context.Context, payment *entity.CreatePaymentExternalInput) (*entity.CreatePaymentExternalOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", context, payment)
	ret0, _ := ret[0].(*entity.CreatePaymentExternalOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockPaymentExternalDatasourceMockRecorder) Create(context, payment any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPaymentExternalDatasource)(nil).Create), context, payment)
}
