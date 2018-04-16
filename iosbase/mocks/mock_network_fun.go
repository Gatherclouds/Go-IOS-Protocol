package iosbase_mock

import "reflect"

// NewMockNetwork creates a new mock instance
func NewMockNetwork(ctrl *gomock.Controller) *MockNetwork {
	mock := &MockNetwork{ctrl: ctrl}
	mock.recorder = &MockNetworkMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNetwork) EXPECT() *MockNetworkMockRecorder {
	return m.recorder
}

// Send mocks base method
func (m *MockNetwork) Send(req Request) chan Response {
	ret := m.ctrl.Call(m, "Send", req)
	ret0, _ := ret[0].(chan Response)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockNetworkMockRecorder) Send(req interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockNetwork)(nil).Send), req)
}

// Listen mocks base method
func (m *MockNetwork) Listen(port uint16) (chan Request, chan Response, error) {
	ret := m.ctrl.Call(m, "Listen", port)
	ret0, _ := ret[0].(chan Request)
	ret1, _ := ret[1].(chan Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Listen indicates an expected call of Listen
func (mr *MockNetworkMockRecorder) Listen(port interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Listen", reflect.TypeOf((*MockNetwork)(nil).Listen), port)
}

// Close mocks base method
func (m *MockNetwork) Close(port uint16) error {
	ret := m.ctrl.Call(m, "Close", port)
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockNetworkMockRecorder) Close(port interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockNetwork)(nil).Close), port)
}

