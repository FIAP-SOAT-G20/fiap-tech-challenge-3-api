// Code generated by MockGen. DO NOT EDIT.
// Source: internal/core/port/jwt_service_port.go
//
// Generated by this command:
//
//	mockgen -source=internal/core/port/jwt_service_port.go -destination=internal/core/port/mocks/jwt_service_mock.go
//

// Package mock_port is a generated GoMock package.
package mock_port

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockJWTService is a mock of JWTService interface.
type MockJWTService struct {
	ctrl     *gomock.Controller
	recorder *MockJWTServiceMockRecorder
	isgomock struct{}
}

// MockJWTServiceMockRecorder is the mock recorder for MockJWTService.
type MockJWTServiceMockRecorder struct {
	mock *MockJWTService
}

// NewMockJWTService creates a new mock instance.
func NewMockJWTService(ctrl *gomock.Controller) *MockJWTService {
	mock := &MockJWTService{ctrl: ctrl}
	mock.recorder = &MockJWTServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJWTService) EXPECT() *MockJWTServiceMockRecorder {
	return m.recorder
}

// GenerateToken mocks base method.
func (m *MockJWTService) GenerateToken(customerID uint64) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken", customerID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateToken indicates an expected call of GenerateToken.
func (mr *MockJWTServiceMockRecorder) GenerateToken(customerID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockJWTService)(nil).GenerateToken), customerID)
}

// ValidateToken mocks base method.
func (m *MockJWTService) ValidateToken(token string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateToken", token)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateToken indicates an expected call of ValidateToken.
func (mr *MockJWTServiceMockRecorder) ValidateToken(token any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateToken", reflect.TypeOf((*MockJWTService)(nil).ValidateToken), token)
}
