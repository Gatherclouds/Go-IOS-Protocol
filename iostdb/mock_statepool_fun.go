package iostdb

import (
	"reflect"
	"github.com/golang/mock/gomock"
)

// NewMockStatePool creates a new mock instance
func NewMockStatePool(ctrl *gomock.Controller) *MockStatePool {
	mock := &MockStatePool{ctrl: ctrl}
	mock.recorder = &MockStatePoolMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStatePool) EXPECT() *MockStatePoolMockRecorder {
	return m.recorder
}

// Init indicates an expected call of Init
func (mr *MockStatePoolMockRecorder) Init() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockStatePool)(nil).Init))
}

// Close mocks base method
func (m *MockStatePool) Close() error {
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockStatePoolMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockStatePool)(nil).Close))
}

// Add mocks base method
func (m *MockStatePool) Add(state State) error {
	ret := m.ctrl.Call(m, "Add", state)
	ret0, _ := ret[0].(error)
	return ret0
}








