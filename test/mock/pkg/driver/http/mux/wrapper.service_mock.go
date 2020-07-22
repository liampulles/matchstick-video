package mux

import (
	"github.com/gorilla/mux"

	goHttp "net/http"

	"github.com/stretchr/testify/mock"

	muxDriver "github.com/liampulles/matchstick-video/pkg/driver/http/mux"
)

// RouteMock is for mocking
type RouteMock struct {
	mock.Mock
}

var _ muxDriver.Route = &RouteMock{}

// Methods is for mocking
func (r *RouteMock) Methods(methods ...string) *mux.Route {
	args := r.Called(methods)
	return safeArgsGetMuxRoute(args, 0)
}

// RouterMock is for mocking
type RouterMock struct {
	mock.Mock
}

var _ muxDriver.Router = &RouterMock{}

// HandleFunc is for mocking
func (r *RouterMock) HandleFunc(pathPattern string, handler muxDriver.Handler) muxDriver.Route {
	args := r.Called(pathPattern, handler)
	return safeArgsGetRouteMock(args, 0)
}

// ServeHTTP is for mocking
func (r *RouterMock) ServeHTTP(res goHttp.ResponseWriter, req *goHttp.Request) {
	r.Called(res, req)
	return
}

// WrapperMock is for mocking
type WrapperMock struct {
	mock.Mock
}

var _ muxDriver.Wrapper = &WrapperMock{}

// NewRouter is for mocking
func (w *WrapperMock) NewRouter() muxDriver.Router {
	args := w.Called()
	return safeArgsGetRouterMock(args, 0)
}

// Vars is for mocking
func (w *WrapperMock) Vars(req *goHttp.Request) map[string]string {
	args := w.Called(req)
	return args.Get(0).(map[string]string)
}

func safeArgsGetMuxRoute(args mock.Arguments, idx int) *mux.Route {
	if val, ok := args.Get(idx).(*mux.Route); ok {
		return val
	}
	return nil
}

func safeArgsGetRouteMock(args mock.Arguments, idx int) *RouteMock {
	if val, ok := args.Get(idx).(*RouteMock); ok {
		return val
	}
	return nil
}

func safeArgsGetRouterMock(args mock.Arguments, idx int) *RouterMock {
	if val, ok := args.Get(idx).(*RouterMock); ok {
		return val
	}
	return nil
}
