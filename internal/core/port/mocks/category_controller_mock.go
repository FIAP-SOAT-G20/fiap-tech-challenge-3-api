// Code generated by MockGen. DO NOT EDIT.
// Source: internal/core/port/category_controller_port.go
//
// Generated by this command:
//
//	mockgen -source=internal/core/port/category_controller_port.go -destination=internal/core/port/mocks/category_controller_mock.go
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

// MockCategoryController is a mock of CategoryController interface.
type MockCategoryController struct {
	ctrl     *gomock.Controller
	recorder *MockCategoryControllerMockRecorder
	isgomock struct{}
}

// MockCategoryControllerMockRecorder is the mock recorder for MockCategoryController.
type MockCategoryControllerMockRecorder struct {
	mock *MockCategoryController
}

// NewMockCategoryController creates a new mock instance.
func NewMockCategoryController(ctrl *gomock.Controller) *MockCategoryController {
	mock := &MockCategoryController{ctrl: ctrl}
	mock.recorder = &MockCategoryControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCategoryController) EXPECT() *MockCategoryControllerMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockCategoryController) Create(ctx context.Context, presenter port.Presenter, input dto.CreateCategoryInput) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, presenter, input)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockCategoryControllerMockRecorder) Create(ctx, presenter, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCategoryController)(nil).Create), ctx, presenter, input)
}

// Delete mocks base method.
func (m *MockCategoryController) Delete(ctx context.Context, presenter port.Presenter, input dto.DeleteCategoryInput) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, presenter, input)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockCategoryControllerMockRecorder) Delete(ctx, presenter, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCategoryController)(nil).Delete), ctx, presenter, input)
}

// Get mocks base method.
func (m *MockCategoryController) Get(ctx context.Context, presenter port.Presenter, input dto.GetCategoryInput) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, presenter, input)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockCategoryControllerMockRecorder) Get(ctx, presenter, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCategoryController)(nil).Get), ctx, presenter, input)
}

// List mocks base method.
func (m *MockCategoryController) List(ctx context.Context, presenter port.Presenter, input dto.ListCategoriesInput) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, presenter, input)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockCategoryControllerMockRecorder) List(ctx, presenter, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockCategoryController)(nil).List), ctx, presenter, input)
}

// Update mocks base method.
func (m *MockCategoryController) Update(ctx context.Context, presenter port.Presenter, input dto.UpdateCategoryInput) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, presenter, input)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockCategoryControllerMockRecorder) Update(ctx, presenter, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCategoryController)(nil).Update), ctx, presenter, input)
}
