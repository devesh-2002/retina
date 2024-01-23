// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package metrics is a generated GoMock package.
package metrics

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	prometheus "github.com/prometheus/client_golang/prometheus"
	io_prometheus_client "github.com/prometheus/client_model/go"
)

// MockICounterVec is a mock of ICounterVec interface.
type MockICounterVec struct {
	ctrl     *gomock.Controller
	recorder *MockICounterVecMockRecorder
}

// MockICounterVecMockRecorder is the mock recorder for MockICounterVec.
type MockICounterVecMockRecorder struct {
	mock *MockICounterVec
}

// NewMockICounterVec creates a new mock instance.
func NewMockICounterVec(ctrl *gomock.Controller) *MockICounterVec {
	mock := &MockICounterVec{ctrl: ctrl}
	mock.recorder = &MockICounterVecMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICounterVec) EXPECT() *MockICounterVecMockRecorder {
	return m.recorder
}

// GetMetricWithLabelValues mocks base method.
func (m *MockICounterVec) GetMetricWithLabelValues(lvs ...string) (prometheus.Counter, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range lvs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMetricWithLabelValues", varargs...)
	ret0, _ := ret[0].(prometheus.Counter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMetricWithLabelValues indicates an expected call of GetMetricWithLabelValues.
func (mr *MockICounterVecMockRecorder) GetMetricWithLabelValues(lvs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMetricWithLabelValues", reflect.TypeOf((*MockICounterVec)(nil).GetMetricWithLabelValues), lvs...)
}

// WithLabelValues mocks base method.
func (m *MockICounterVec) WithLabelValues(lvs ...string) prometheus.Counter {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range lvs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "WithLabelValues", varargs...)
	ret0, _ := ret[0].(prometheus.Counter)
	return ret0
}

// WithLabelValues indicates an expected call of WithLabelValues.
func (mr *MockICounterVecMockRecorder) WithLabelValues(lvs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithLabelValues", reflect.TypeOf((*MockICounterVec)(nil).WithLabelValues), lvs...)
}

// MockIGaugeVec is a mock of IGaugeVec interface.
type MockIGaugeVec struct {
	ctrl     *gomock.Controller
	recorder *MockIGaugeVecMockRecorder
}

// MockIGaugeVecMockRecorder is the mock recorder for MockIGaugeVec.
type MockIGaugeVecMockRecorder struct {
	mock *MockIGaugeVec
}

// NewMockIGaugeVec creates a new mock instance.
func NewMockIGaugeVec(ctrl *gomock.Controller) *MockIGaugeVec {
	mock := &MockIGaugeVec{ctrl: ctrl}
	mock.recorder = &MockIGaugeVecMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIGaugeVec) EXPECT() *MockIGaugeVecMockRecorder {
	return m.recorder
}

// GetMetricWithLabelValues mocks base method.
func (m *MockIGaugeVec) GetMetricWithLabelValues(lvs ...string) (prometheus.Gauge, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range lvs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMetricWithLabelValues", varargs...)
	ret0, _ := ret[0].(prometheus.Gauge)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMetricWithLabelValues indicates an expected call of GetMetricWithLabelValues.
func (mr *MockIGaugeVecMockRecorder) GetMetricWithLabelValues(lvs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMetricWithLabelValues", reflect.TypeOf((*MockIGaugeVec)(nil).GetMetricWithLabelValues), lvs...)
}

// WithLabelValues mocks base method.
func (m *MockIGaugeVec) WithLabelValues(lvs ...string) prometheus.Gauge {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range lvs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "WithLabelValues", varargs...)
	ret0, _ := ret[0].(prometheus.Gauge)
	return ret0
}

// WithLabelValues indicates an expected call of WithLabelValues.
func (mr *MockIGaugeVecMockRecorder) WithLabelValues(lvs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithLabelValues", reflect.TypeOf((*MockIGaugeVec)(nil).WithLabelValues), lvs...)
}

// MockIHistogramVec is a mock of IHistogramVec interface.
type MockIHistogramVec struct {
	ctrl     *gomock.Controller
	recorder *MockIHistogramVecMockRecorder
}

// MockIHistogramVecMockRecorder is the mock recorder for MockIHistogramVec.
type MockIHistogramVecMockRecorder struct {
	mock *MockIHistogramVec
}

// NewMockIHistogramVec creates a new mock instance.
func NewMockIHistogramVec(ctrl *gomock.Controller) *MockIHistogramVec {
	mock := &MockIHistogramVec{ctrl: ctrl}
	mock.recorder = &MockIHistogramVecMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIHistogramVec) EXPECT() *MockIHistogramVecMockRecorder {
	return m.recorder
}

// Observe mocks base method.
func (m *MockIHistogramVec) Observe(arg0 float64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Observe", arg0)
}

// Observe indicates an expected call of Observe.
func (mr *MockIHistogramVecMockRecorder) Observe(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Observe", reflect.TypeOf((*MockIHistogramVec)(nil).Observe), arg0)
}

// Write mocks base method.
func (m *MockIHistogramVec) Write(arg0 *io_prometheus_client.Metric) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Write indicates an expected call of Write.
func (mr *MockIHistogramVecMockRecorder) Write(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockIHistogramVec)(nil).Write), arg0)
}
