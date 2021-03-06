// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cgulli/airport_go/airport (interfaces: Airport)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockAirport is a mock of Airport interface
type MockAirport struct {
	ctrl     *gomock.Controller
	recorder *MockAirportMockRecorder
}

// MockAirportMockRecorder is the mock recorder for MockAirport
type MockAirportMockRecorder struct {
	mock *MockAirport
}

// NewMockAirport creates a new mock instance
func NewMockAirport(ctrl *gomock.Controller) *MockAirport {
	mock := &MockAirport{ctrl: ctrl}
	mock.recorder = &MockAirportMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAirport) EXPECT() *MockAirportMockRecorder {
	return m.recorder
}

// CanLand mocks base method
func (m *MockAirport) CanLand() bool {
	ret := m.ctrl.Call(m, "CanLand")
	ret0, _ := ret[0].(bool)
	return ret0
}

// CanLand indicates an expected call of CanLand
func (mr *MockAirportMockRecorder) CanLand() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanLand", reflect.TypeOf((*MockAirport)(nil).CanLand))
}
