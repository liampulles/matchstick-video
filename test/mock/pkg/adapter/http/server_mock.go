package http

import (
	"github.com/stretchr/testify/mock"

	"github.com/liampulles/matchstick-video/pkg/adapter/http"
	"github.com/liampulles/matchstick-video/pkg/domain"
)

// ServerConfigurationMock is for mocking
type ServerConfigurationMock struct {
	mock.Mock
}

var _ http.ServerConfiguration = &ServerConfigurationMock{}

// CreateRunnable is for mocking
func (s *ServerConfigurationMock) CreateRunnable(handlers map[http.HandlerPattern]http.Handler) domain.Runnable {
	args := s.Called(handlers)
	return args.Get(0).(domain.Runnable)
}

// ServerFactoryMock is for mocking
type ServerFactoryMock struct {
	mock.Mock
}

var _ http.ServerFactory = &ServerFactoryMock{}

// Create is for mocking
func (s *ServerFactoryMock) Create() domain.Runnable {
	args := s.Called(0)
	return args.Get(0).(domain.Runnable)
}
