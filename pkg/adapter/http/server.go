package http

import (
	"github.com/liampulles/matchstick-video/pkg/domain"
)

// ServerConfiguration encapsulates the specifics of
// a given HTTP server implementation. Implementations
// cna be found in the driver layer.
type ServerConfiguration interface {
	CreateRunnable(handlers map[HandlerPattern]Handler) domain.Runnable
}

// ServerFactory creates servers
type ServerFactory interface {
	Create() domain.Runnable
}

// ServerFactoryImpl implements ServerFactory
type ServerFactoryImpl struct {
	inventoryController Controller
	serverConfiguration ServerConfiguration
}

// Check we implement the interface
var _ ServerFactory = &ServerFactoryImpl{}

// NewServerFactoryImpl is a constructor
func NewServerFactoryImpl(inventoryController Controller, serverConfiguration ServerConfiguration) *ServerFactoryImpl {
	return &ServerFactoryImpl{
		inventoryController: inventoryController,
		serverConfiguration: serverConfiguration,
	}
}

// Create implements the ServerFactory interface
func (s *ServerFactoryImpl) Create() domain.Runnable {
	handlers := s.inventoryController.GetHandlers()
	return s.serverConfiguration.CreateRunnable(handlers)
}
