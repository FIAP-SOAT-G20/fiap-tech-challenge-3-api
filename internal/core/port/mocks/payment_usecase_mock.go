// Code generated by MockGen. DO NOT EDIT.
// Source: internal/core/port/payment_usecase_port.go
//
// Generated by this command:
//
//	mockgen -source=internal/core/port/payment_usecase_port.go -destination=internal/core/port/mocks/payment_usecase_mock.go
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

// MockPaymentUseCase is a mock of PaymentUseCase interface.
type MockPaymentUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockPaymentUseCaseMockRecorder
	isgomock struct{}
}

// MockPaymentUseCaseMockRecorder is the mock recorder for MockPaymentUseCase.
type MockPaymentUseCaseMockRecorder struct {
	mock *MockPaymentUseCase
}

// NewMockPaymentUseCase creates a new mock instance.
func NewMockPaymentUseCase(ctrl *gomock.Controller) *MockPaymentUseCase {
	mock := &MockPaymentUseCase{ctrl: ctrl}
	mock.recorder = &MockPaymentUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPaymentUseCase) EXPECT() *MockPaymentUseCaseMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockPaymentUseCase) Create(ctx context.Context, input dto.CreatePaymentInput) (*entity.Payment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, input)
	ret0, _ := ret[0].(*entity.Payment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockPaymentUseCaseMockRecorder) Create(ctx, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPaymentUseCase)(nil).Create), ctx, input)
}

// Get mocks base method.
func (m *MockPaymentUseCase) Get(ctx context.Context, payment dto.GetPaymentInput) (*entity.Payment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, payment)
	ret0, _ := ret[0].(*entity.Payment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockPaymentUseCaseMockRecorder) Get(ctx, payment any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPaymentUseCase)(nil).Get), ctx, payment)
}

// Update mocks base method.
func (m *MockPaymentUseCase) Update(ctx context.Context, payment dto.UpdatePaymentInput) (*entity.Payment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, payment)
	ret0, _ := ret[0].(*entity.Payment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockPaymentUseCaseMockRecorder) Update(ctx, payment any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPaymentUseCase)(nil).Update), ctx, payment)
}
