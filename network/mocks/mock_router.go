package mocks

import "github.com/golang/mock/gomock"

// MockRouter is a mock of Router interface
type MockRouter struct {
	ctrl     *gomock.Controller
	recorder *MockRouterMockRecorder
}

