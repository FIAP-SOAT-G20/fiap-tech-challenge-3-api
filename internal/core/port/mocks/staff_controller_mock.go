// Code generated by MockGen. DO NOT EDIT.
// Source: internal/core/port/staff_controller_port.go
//
// Generated by this command:
//
//	mockgen -source=internal/core/port/staff_controller_port.go -destination=internal/core/port/mocks/staff_controller_mock.go
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

// MockStaffController is a mock of StaffController interface.
type MockStaffController struct {
	ctrl     *gomock.Controller
	recorder *MockStaffControllerMockRecorder
	isgomock struct{}
}

// MockStaffControllerMockRecorder is the mock recorder for MockStaffController.
type MockStaffControllerMockRecorder struct {
	mock *MockStaffController
}

// NewMockStaffController creates a new mock instance.
func NewMockStaffController(ctrl *gomock.Controller) *MockStaffController {
	mock := &MockStaffController{ctrl: ctrl}
	mock.recorder = &MockStaffControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStaffController) EXPECT() *MockStaffControllerMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockStaffController) Create(ctx context.Context, presenter port.Presenter, input dto.CreateStaffInput) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, presenter, input)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockStaffControllerMockRecorder) Create(ctx, presenter, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockStaffController)(nil).Create), ctx, presenter, input)
}

// Delete mocks base method.
func (m *MockStaffController) Delete(ctx context.Context, presenter port.Presenter, input dto.DeleteStaffInput) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, presenter, input)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockStaffControllerMockRecorder) Delete(ctx, presenter, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockStaffController)(nil).Delete), ctx, presenter, input)
}

// Get mocks base method.
func (m *MockStaffController) Get(ctx context.Context, presenter port.Presenter, input dto.GetStaffInput) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, presenter, input)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockStaffControllerMockRecorder) Get(ctx, presenter, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockStaffController)(nil).Get), ctx, presenter, input)
}

// List mocks base method.
func (m *MockStaffController) List(ctx context.Context, presenter port.Presenter, input dto.ListStaffsInput) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, presenter, input)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockStaffControllerMockRecorder) List(ctx, presenter, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockStaffController)(nil).List), ctx, presenter, input)
}

// Update mocks base method.
func (m *MockStaffController) Update(ctx context.Context, presenter port.Presenter, input dto.UpdateStaffInput) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, presenter, input)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockStaffControllerMockRecorder) Update(ctx, presenter, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockStaffController)(nil).Update), ctx, presenter, input)
}
