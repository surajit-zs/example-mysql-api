// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package services is a generated GoMock package.
package services

import (
	models "mysql-api/models"
	reflect "reflect"

	gofr "developer.zopsmart.com/go/gofr/pkg/gofr"
	gomock "github.com/golang/mock/gomock"
)

// MockServices is a mock of Services interface.
type MockServices struct {
	ctrl     *gomock.Controller
	recorder *MockServicesMockRecorder
}

// MockServicesMockRecorder is the mock recorder for MockServices.
type MockServicesMockRecorder struct {
	mock *MockServices
}

// NewMockServices creates a new mock instance.
func NewMockServices(ctrl *gomock.Controller) *MockServices {
	mock := &MockServices{ctrl: ctrl}
	mock.recorder = &MockServicesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServices) EXPECT() *MockServicesMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockServices) Create(ctx *gofr.Context, c models.Cat) (models.Cat, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, c)
	ret0, _ := ret[0].(models.Cat)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockServicesMockRecorder) Create(ctx, c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockServices)(nil).Create), ctx, c)
}

// Delete mocks base method.
func (m *MockServices) Delete(ctx *gofr.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockServicesMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockServices)(nil).Delete), ctx, id)
}

// Get mocks base method.
func (m *MockServices) Get(ctx *gofr.Context) ([]models.Cat, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx)
	ret0, _ := ret[0].([]models.Cat)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockServicesMockRecorder) Get(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockServices)(nil).Get), ctx)
}

// GetByID mocks base method.
func (m *MockServices) GetByID(ctx *gofr.Context, id string) (models.Cat, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id)
	ret0, _ := ret[0].(models.Cat)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockServicesMockRecorder) GetByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockServices)(nil).GetByID), ctx, id)
}

// Update mocks base method.
func (m *MockServices) Update(ctx *gofr.Context, c models.Cat) (models.Cat, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, c)
	ret0, _ := ret[0].(models.Cat)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockServicesMockRecorder) Update(ctx, c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockServices)(nil).Update), ctx, c)
}
