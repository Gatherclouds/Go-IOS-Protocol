package iostdb

import (
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














