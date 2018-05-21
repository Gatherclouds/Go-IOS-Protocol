package mocks

import (
	gomock "github.com/golang/mock/gomock"

)

// MockContract is a mock of Contract interface
type MockContract struct {
	ctrl     *gomock.Controller
	recorder *MockContractMockRecorder
}
