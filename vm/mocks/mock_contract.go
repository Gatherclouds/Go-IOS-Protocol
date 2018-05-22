package mocks

import (
	gomock "github.com/golang/mock/gomock"

)

// MockContract is a mock of Contract interface
type MockContract struct {
	ctrl     *gomock.Controller
	recorder *MockContractMockRecorder
}

// MockContractMockRecorder is the mock recorder for MockContract
type MockContractMockRecorder struct {
	mock *MockContract
}

// NewMockContract creates a new mock instance
func NewMockContract(ctrl *gomock.Controller) *MockContract {
	mock := &MockContract{ctrl: ctrl}
	mock.recorder = &MockContractMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockContract) EXPECT() *MockContractMockRecorder {
	return m.recorder
}

// API mocks base method
func (m *MockContract) API(arg0 string) (vm.Method, error) {
	ret := m.ctrl.Call(m, "API", arg0)
	ret0, _ := ret[0].(vm.Method)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddSigner mocks base method
func (m *MockContract) AddSigner(arg0 []byte) {
	m.ctrl.Call(m, "AddSigner", arg0)
}
