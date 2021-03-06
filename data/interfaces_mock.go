// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package encoder is a generated GoMock package.
package data

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockParser is a mock of Parser interface
type MockParser struct {
	ctrl     *gomock.Controller
	recorder *MockParserMockRecorder
}

// MockParserMockRecorder is the mock recorder for MockParser
type MockParserMockRecorder struct {
	mock *MockParser
}

// NewMockParser creates a new mock instance
func NewMockParser(ctrl *gomock.Controller) *MockParser {
	mock := &MockParser{ctrl: ctrl}
	mock.recorder = &MockParserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockParser) EXPECT() *MockParserMockRecorder {
	return m.recorder
}

// ParseMapToData mocks base method
func (m *MockParser) ParseMapToData(input map[string]interface{}) (*Data, error) {
	ret := m.ctrl.Call(m, "ParseMapToData", input)
	ret0, _ := ret[0].(*Data)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseMapToData indicates an expected call of ParseMapToData
func (mr *MockParserMockRecorder) ParseMapToData(input interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseMapToData", reflect.TypeOf((*MockParser)(nil).ParseMapToData), input)
}
