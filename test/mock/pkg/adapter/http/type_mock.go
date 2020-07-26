package http

import (
	"github.com/stretchr/testify/mock"

	"github.com/liampulles/matchstick-video/pkg/adapter/http"
)

// MockController is for mocking
type MockController struct {
	mock.Mock
}

var _ http.Controller = &MockController{}

// GetHandlers is for mocking
func (c *MockController) GetHandlers() map[http.HandlerPattern]http.Handler {
	args := c.Called()
	return args.Get(0).(map[http.HandlerPattern]http.Handler)
}
