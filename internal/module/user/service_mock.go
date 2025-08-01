// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package user is a generated GoMock package.
package user

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/smartmemos/memos/internal/module/user/model"
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

// CreateAccessToken mocks base method.
func (m *MockService) CreateAccessToken(ctx context.Context, req *model.CreateAccessTokenRequest) (*model.AccessToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccessToken", ctx, req)
	ret0, _ := ret[0].(*model.AccessToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccessToken indicates an expected call of CreateAccessToken.
func (mr *MockServiceMockRecorder) CreateAccessToken(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccessToken", reflect.TypeOf((*MockService)(nil).CreateAccessToken), ctx, req)
}

// CreateUser mocks base method.
func (m *MockService) CreateUser(arg0 context.Context, arg1 *model.CreateUserRequest) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockServiceMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockService)(nil).CreateUser), arg0, arg1)
}

// DeleteAccessToken mocks base method.
func (m *MockService) DeleteAccessToken(ctx context.Context, req *model.DeleteAccessTokenRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAccessToken", ctx, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAccessToken indicates an expected call of DeleteAccessToken.
func (mr *MockServiceMockRecorder) DeleteAccessToken(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccessToken", reflect.TypeOf((*MockService)(nil).DeleteAccessToken), ctx, req)
}

// GetUserByID mocks base method.
func (m *MockService) GetUserByID(arg0 context.Context, arg1 int64) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", arg0, arg1)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockServiceMockRecorder) GetUserByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockService)(nil).GetUserByID), arg0, arg1)
}

// ListAccessTokens mocks base method.
func (m *MockService) ListAccessTokens(ctx context.Context, req *model.ListAccessTokensRequest) ([]*model.AccessToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAccessTokens", ctx, req)
	ret0, _ := ret[0].([]*model.AccessToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAccessTokens indicates an expected call of ListAccessTokens.
func (mr *MockServiceMockRecorder) ListAccessTokens(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccessTokens", reflect.TypeOf((*MockService)(nil).ListAccessTokens), ctx, req)
}

// ListAllUserStats mocks base method.
func (m *MockService) ListAllUserStats(ctx context.Context, req *model.ListAllUserStatsRequest) (*model.Stats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllUserStats", ctx, req)
	ret0, _ := ret[0].(*model.Stats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllUserStats indicates an expected call of ListAllUserStats.
func (mr *MockServiceMockRecorder) ListAllUserStats(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllUserStats", reflect.TypeOf((*MockService)(nil).ListAllUserStats), ctx, req)
}

// UpdateUser mocks base method.
func (m *MockService) UpdateUser(arg0 context.Context, arg1 *model.UpdateUserRequest) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockServiceMockRecorder) UpdateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockService)(nil).UpdateUser), arg0, arg1)
}
