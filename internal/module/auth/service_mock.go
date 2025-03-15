// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package auth is a generated GoMock package.
package auth

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/smartmemos/memos/internal/module/auth/model"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// Authenticate mocks base method.
func (m *MockService) Authenticate(ctx context.Context, token string) (*model.AccessToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authenticate", ctx, token)
	ret0, _ := ret[0].(*model.AccessToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Authenticate indicates an expected call of Authenticate.
func (mr *MockServiceMockRecorder) Authenticate(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authenticate", reflect.TypeOf((*MockService)(nil).Authenticate), ctx, token)
}

// DeleteAccessToken mocks base method.
func (m *MockService) DeleteAccessToken(ctx context.Context, userId int64, token string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAccessToken", ctx, userId, token)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAccessToken indicates an expected call of DeleteAccessToken.
func (mr *MockServiceMockRecorder) DeleteAccessToken(ctx, userId, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccessToken", reflect.TypeOf((*MockService)(nil).DeleteAccessToken), ctx, userId, token)
}

// SignIn mocks base method.
func (m *MockService) SignIn(ctx context.Context, req *model.SignInRequest) (*model.AccessToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignIn", ctx, req)
	ret0, _ := ret[0].(*model.AccessToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignIn indicates an expected call of SignIn.
func (mr *MockServiceMockRecorder) SignIn(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignIn", reflect.TypeOf((*MockService)(nil).SignIn), ctx, req)
}
