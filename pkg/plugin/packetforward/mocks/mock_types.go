// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/microsoft/retina/pkg/plugin/packetforward (interfaces: IMap)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIMap is a mock of IMap interface.
type MockIMap struct {
	ctrl     *gomock.Controller
	recorder *MockIMapMockRecorder
}

// MockIMapMockRecorder is the mock recorder for MockIMap.
type MockIMapMockRecorder struct {
	mock *MockIMap
}

// NewMockIMap creates a new mock instance.
func NewMockIMap(ctrl *gomock.Controller) *MockIMap {
	mock := &MockIMap{ctrl: ctrl}
	mock.recorder = &MockIMapMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIMap) EXPECT() *MockIMapMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockIMap) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockIMapMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockIMap)(nil).Close))
}

// Lookup mocks base method.
func (m *MockIMap) Lookup(arg0, arg1 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Lookup", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Lookup indicates an expected call of Lookup.
func (mr *MockIMapMockRecorder) Lookup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Lookup", reflect.TypeOf((*MockIMap)(nil).Lookup), arg0, arg1)
}
