package iosbase

import "reflect"

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

// Init mocks base method
func (m *MockStatePool) Init() error {
	ret := m.ctrl.Call(m, "Init")
	ret0, _ := ret[0].(error)
	return ret0
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

// Add indicates an expected call of Add
func (mr *MockStatePoolMockRecorder) Add(state interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockStatePool)(nil).Add), state)
}

// Find mocks base method
func (m *MockStatePool) Find(stateHash []byte) (State, error) {
	ret := m.ctrl.Call(m, "Find", stateHash)
	ret0, _ := ret[0].(State)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockStatePoolMockRecorder) Find(stateHash interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockStatePool)(nil).Find), stateHash)
}

// Del mocks base method
func (m *MockStatePool) Del(StateHash []byte) error {
	ret := m.ctrl.Call(m, "Del", StateHash)
	ret0, _ := ret[0].(error)
	return ret0
}

// Del indicates an expected call of Del
func (mr *MockStatePoolMockRecorder) Del(StateHash interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Del", reflect.TypeOf((*MockStatePool)(nil).Del), StateHash)
}

// Transact mocks base method
func (m *MockStatePool) Transact(block *Block) error {
	ret := m.ctrl.Call(m, "Transact", block)
	ret0, _ := ret[0].(error)
	return ret0
}

// Transact indicates an expected call of Transact
func (mr *MockStatePoolMockRecorder) Transact(block interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transact", reflect.TypeOf((*MockStatePool)(nil).Transact), block)
}

