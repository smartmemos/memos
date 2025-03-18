// Code generated by MockGen. DO NOT EDIT.
// Source: dao.go

// Package user is a generated GoMock package.
package user

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/smartmemos/memos/internal/module/user/model"
)

// MockDAO is a mock of DAO interface.
type MockDAO struct {
	ctrl     *gomock.Controller
	recorder *MockDAOMockRecorder
}

// MockDAOMockRecorder is the mock recorder for MockDAO.
type MockDAOMockRecorder struct {
	mock *MockDAO
}

// NewMockDAO creates a new mock instance.
func NewMockDAO(ctrl *gomock.Controller) *MockDAO {
	mock := &MockDAO{ctrl: ctrl}
	mock.recorder = &MockDAOMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDAO) EXPECT() *MockDAOMockRecorder {
	return m.recorder
}

// CountUsers mocks base method.
func (m *MockDAO) CountUsers(ctx context.Context, filter *model.FindUserFilter) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountUsers", ctx, filter)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountUsers indicates an expected call of CountUsers.
func (mr *MockDAOMockRecorder) CountUsers(ctx, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountUsers", reflect.TypeOf((*MockDAO)(nil).CountUsers), ctx, filter)
}

// CreateUser mocks base method.
func (m_2 *MockDAO) CreateUser(ctx context.Context, m *model.User) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "CreateUser", ctx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockDAOMockRecorder) CreateUser(ctx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockDAO)(nil).CreateUser), ctx, m)
}

// FindSetting mocks base method.
func (m *MockDAO) FindSetting(ctx context.Context, filter *model.FindSettingFilter) (*model.Setting, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindSetting", ctx, filter)
	ret0, _ := ret[0].(*model.Setting)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindSetting indicates an expected call of FindSetting.
func (mr *MockDAOMockRecorder) FindSetting(ctx, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindSetting", reflect.TypeOf((*MockDAO)(nil).FindSetting), ctx, filter)
}

// FindSettings mocks base method.
func (m *MockDAO) FindSettings(ctx context.Context, filter *model.FindSettingFilter) ([]*model.Setting, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindSettings", ctx, filter)
	ret0, _ := ret[0].([]*model.Setting)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindSettings indicates an expected call of FindSettings.
func (mr *MockDAOMockRecorder) FindSettings(ctx, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindSettings", reflect.TypeOf((*MockDAO)(nil).FindSettings), ctx, filter)
}

// FindUser mocks base method.
func (m *MockDAO) FindUser(ctx context.Context, filter *model.FindUserFilter) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUser", ctx, filter)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUser indicates an expected call of FindUser.
func (mr *MockDAOMockRecorder) FindUser(ctx, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUser", reflect.TypeOf((*MockDAO)(nil).FindUser), ctx, filter)
}

// FindUserByID mocks base method.
func (m *MockDAO) FindUserByID(ctx context.Context, id int64) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByID", ctx, id)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByID indicates an expected call of FindUserByID.
func (mr *MockDAOMockRecorder) FindUserByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByID", reflect.TypeOf((*MockDAO)(nil).FindUserByID), ctx, id)
}

// FindUsers mocks base method.
func (m *MockDAO) FindUsers(ctx context.Context, filter *model.FindUserFilter) ([]*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUsers", ctx, filter)
	ret0, _ := ret[0].([]*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUsers indicates an expected call of FindUsers.
func (mr *MockDAOMockRecorder) FindUsers(ctx, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUsers", reflect.TypeOf((*MockDAO)(nil).FindUsers), ctx, filter)
}

// UpdateUser mocks base method.
func (m_2 *MockDAO) UpdateUser(ctx context.Context, m *model.User, update map[string]any) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "UpdateUser", ctx, m, update)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockDAOMockRecorder) UpdateUser(ctx, m, update interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockDAO)(nil).UpdateUser), ctx, m, update)
}

// UpdateUsers mocks base method.
func (m *MockDAO) UpdateUsers(ctx context.Context, filter *model.FindUserFilter, update map[string]any) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUsers", ctx, filter, update)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUsers indicates an expected call of UpdateUsers.
func (mr *MockDAOMockRecorder) UpdateUsers(ctx, filter, update interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUsers", reflect.TypeOf((*MockDAO)(nil).UpdateUsers), ctx, filter, update)
}
