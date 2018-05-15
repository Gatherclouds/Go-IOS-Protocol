package iostdb

import (
	"reflect"
	"github.com/golang/mock/gomock"
	"github.com/LoCCS/bliss/sampler"
	"github.com/LoCCS/bliss"
	"fmt"
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













