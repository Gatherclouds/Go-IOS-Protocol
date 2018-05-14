package iostdb

import (
	"reflect"
	"github.com/golang/mock/gomock"
	"github.com/LoCCS/bliss/sampler"
	"github.com/LoCCS/bliss"
)

// NewMockBlockChain creates a new mock instance
func NewMockBlockChain(ctrl *gomock.Controller) *MockBlockChain {
	mock := &MockBlockChain{ctrl: ctrl}
	mock.recorder = &MockBlockChainMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBlockChain) EXPECT() *MockBlockChainMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockBlockChain) Get(layer int) (*Block, error) {
	ret := m.ctrl.Call(m, "Get", layer)
	ret0, _ := ret[0].(*Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockBlockChainMockRecorder) Get(layer interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockBlockChain)(nil).Get), layer)
}

// Push mocks base method
func (m *MockBlockChain) Push(block *Block) error {
	ret := m.ctrl.Call(m, "Push", block)
	ret0, _ := ret[0].(error)
	return ret0
}

// Push indicates an expected call of Push
func (mr *MockBlockChainMockRecorder) Push(block interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Push", reflect.TypeOf((*MockBlockChain)(nil).Push), block)
}

// Length mocks base method
func (m *MockBlockChain) Length() int {
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(int)
	return ret0
}

// Length indicates an expected call of Length
func (mr *MockBlockChainMockRecorder) Length() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockBlockChain)(nil).Length))
}








