package mux

import (
	"fmt"
	goHttp "net/http"

	"github.com/liampulles/matchstick-video/pkg/adapter/config"
	"github.com/liampulles/matchstick-video/pkg/adapter/http"
	"github.com/liampulles/matchstick-video/pkg/domain"
)

// ServerConfigurationImpl implements ServerConfiguration with gorilla/mux
type ServerConfigurationImpl struct {
	configStore   config.Store
	handlerMapper HandlerMapper
	muxWrapper    Wrapper
}

// Check we implement the interface
var _ http.ServerConfiguration = &ServerConfigurationImpl{}

// NewServerConfigurationImpl is a constructor
func NewServerConfigurationImpl(
	configStore config.Store,
	handlerMapper HandlerMapper,
	muxWrapper Wrapper,
) *ServerConfigurationImpl {

	return &ServerConfigurationImpl{
		configStore:   configStore,
		handlerMapper: handlerMapper,
		muxWrapper:    muxWrapper,
	}
}

// CreateRunnable registers handlers with mux and configures a server
// to create a runnable.
func (m *ServerConfigurationImpl) CreateRunnable(handlers map[http.HandlerPattern]http.Handler) domain.Runnable {

	r := m.muxWrapper.NewRouter()
	// Register each handler with mux
	for pattern, handler := range handlers {
		method := pattern.Method
		pathPattern := pattern.PathPattern

		muxHandler := m.handlerMapper.Map(handler)

		r.HandleFunc(pathPattern, muxHandler).
			Methods(method)
	}

	// Create a server configuration
	port := m.getPort()
	server := goHttp.Server{Addr: port, Handler: r}

	// Run the server!
	return server.ListenAndServe
}

func (m *ServerConfigurationImpl) getPort() string {
	return fmt.Sprintf(":%d", m.configStore.GetPort())
}
