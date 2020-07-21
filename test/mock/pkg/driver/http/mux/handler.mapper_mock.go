package mux

import (
	goHttp "net/http"

	"github.com/stretchr/testify/mock"

	"github.com/liampulles/matchstick-video/pkg/adapter/http"
	muxDriver "github.com/liampulles/matchstick-video/pkg/driver/http/mux"
)

// HandlerMapperMock is for mocking
type HandlerMapperMock struct {
	mock.Mock
}

var _ muxDriver.HandlerMapper = &HandlerMapperMock{}

// Map is for mocking
func (h *HandlerMapperMock) Map(handler http.Handler) muxDriver.Handler {
	args := h.Called(handler)
	return args.Get(0).(func(goHttp.ResponseWriter, *goHttp.Request))
}
