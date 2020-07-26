package mux

import (
	goHttp "net/http"

	"github.com/stretchr/testify/mock"

	"github.com/liampulles/matchstick-video/pkg/adapter/http"
	muxDriver "github.com/liampulles/matchstick-video/pkg/driver/http/mux"
)

// MockHandlerMapper is for mocking
type MockHandlerMapper struct {
	mock.Mock
}

var _ muxDriver.HandlerMapper = &MockHandlerMapper{}

// Map is for mocking
func (h *MockHandlerMapper) Map(handler http.Handler) muxDriver.Handler {
	args := h.Called(handler)
	return args.Get(0).(func(goHttp.ResponseWriter, *goHttp.Request))
}
