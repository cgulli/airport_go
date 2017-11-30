// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cgulli/airport_go/airport (interfaces: Airport)

// Package mocks is a generated GoMock package.
package mocks

import (
	airport "github.com/cgulli/airport_go/airport"
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

// NewAirport mocks base method
func (m *MockAirport) NewAirport() airport.Airport {
	ret := m.ctrl.Call(m, "NewAirport")
	ret0, _ := ret[0].(airport.Airport)
	return ret0
}

// NewAirport indicates an expected call of NewAirport
func (mr *MockAirportMockRecorder) NewAirport() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewAirport", reflect.TypeOf((*MockAirport)(nil).NewAirport))
}
