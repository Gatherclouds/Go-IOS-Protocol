package iostdb

import (
	"reflect"
	"github.com/golang/mock/gomock"
)

// NewMockTxPool creates a new mock instance
func NewMockTxPool(ctrl *gomock.Controller) *MockTxPool {
	mock := &MockTxPool{ctrl: ctrl}
	mock.recorder = &MockTxPoolMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTxPool) EXPECT() *MockTxPoolMockRecorder {
	return m.recorder
}

// Add mocks base method
func (m *MockTxPool) Add(arg0 iosbase.Tx) error {
	ret := m.ctrl.Call(m, "Add", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}
















