package http

import (
	"github.com/stretchr/testify/mock"

	"github.com/liampulles/matchstick-video/pkg/adapter/http"
)

// ControllerMock is for mocking
type ControllerMock struct {
	mock.Mock
}

var _ http.Controller = &ControllerMock{}

// GetHandlers is for mocking
func (c *ControllerMock) GetHandlers() map[http.HandlerPattern]http.Handler {
	args := c.Called()
	return args.Get(0).(map[http.HandlerPattern]http.Handler)
}
