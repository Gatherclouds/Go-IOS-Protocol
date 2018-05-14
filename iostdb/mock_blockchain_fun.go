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

func Sign(msg []byte, passphrase string) (*bliss.Signature, error) {
	entropy, sk, err := newPrivateKey(passphrase)
	if err != nil {
		return nil, fmt.Errorf("Error: bad passphrase.")
	}

	signature, err := sk.SignAgainstSideChannel(msg, entropy)
	if err != nil {
		return nil, fmt.Errorf("Error: failed in signature.")
	} else {
		return signature, nil
	}
}









