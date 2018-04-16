// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/libopenstorage/openstorage/secrets (interfaces: Secrets)

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockSecrets is a mock of Secrets interface
type MockSecrets struct {
	ctrl     *gomock.Controller
	recorder *MockSecretsMockRecorder
}

// MockSecretsMockRecorder is the mock recorder for MockSecrets
type MockSecretsMockRecorder struct {
	mock *MockSecrets
}

// NewMockSecrets creates a new mock instance
func NewMockSecrets(ctrl *gomock.Controller) *MockSecrets {
	mock := &MockSecrets{ctrl: ctrl}
	mock.recorder = &MockSecretsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSecrets) EXPECT() *MockSecretsMockRecorder {
	return m.recorder
}

// CheckLogin mocks base method
func (m *MockSecrets) CheckLogin() error {
	ret := m.ctrl.Call(m, "CheckLogin")
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckLogin indicates an expected call of CheckLogin
func (mr *MockSecretsMockRecorder) CheckLogin() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckLogin", reflect.TypeOf((*MockSecrets)(nil).CheckLogin))
}

// Get mocks base method
func (m *MockSecrets) Get(arg0 string) (interface{}, error) {
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockSecretsMockRecorder) Get(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSecrets)(nil).Get), arg0)
}

// GetDefaultSecretKey mocks base method
func (m *MockSecrets) GetDefaultSecretKey() (interface{}, error) {
	ret := m.ctrl.Call(m, "GetDefaultSecretKey")
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDefaultSecretKey indicates an expected call of GetDefaultSecretKey
func (mr *MockSecretsMockRecorder) GetDefaultSecretKey() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDefaultSecretKey", reflect.TypeOf((*MockSecrets)(nil).GetDefaultSecretKey))
}

// Login mocks base method
func (m *MockSecrets) Login(arg0 string, arg1 map[string]string) error {
	ret := m.ctrl.Call(m, "Login", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Login indicates an expected call of Login
func (mr *MockSecretsMockRecorder) Login(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockSecrets)(nil).Login), arg0, arg1)
}

// Set mocks base method
func (m *MockSecrets) Set(arg0 string, arg1 interface{}) error {
	ret := m.ctrl.Call(m, "Set", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set
func (mr *MockSecretsMockRecorder) Set(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockSecrets)(nil).Set), arg0, arg1)
}

// SetDefaultSecretKey mocks base method
func (m *MockSecrets) SetDefaultSecretKey(arg0 string, arg1 bool) error {
	ret := m.ctrl.Call(m, "SetDefaultSecretKey", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetDefaultSecretKey indicates an expected call of SetDefaultSecretKey
func (mr *MockSecretsMockRecorder) SetDefaultSecretKey(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDefaultSecretKey", reflect.TypeOf((*MockSecrets)(nil).SetDefaultSecretKey), arg0, arg1)
}