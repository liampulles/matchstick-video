package mux

import (
	goHttp "net/http"

	"github.com/gorilla/mux"
)

// Route wraps mux.Route
type Route interface {
	Methods(...string) *mux.Route
}

// Router wraps mux.Router
type Router interface {
	goHttp.Handler
	HandleFunc(pathPattern string, handler Handler) Route
}

// RouterImpl implements Router
type RouterImpl struct {
	MuxRouter *mux.Router
}

var _ Router = &RouterImpl{}

// HandleFunc implements Router
func (r *RouterImpl) HandleFunc(pathPattern string, handler Handler) Route {
	return r.MuxRouter.HandleFunc(pathPattern, handler)
}

func (r *RouterImpl) ServeHTTP(res goHttp.ResponseWriter, req *goHttp.Request) {
	r.MuxRouter.ServeHTTP(res, req)
	return
}

// Wrapper encapsulates mux methods
type Wrapper interface {
	NewRouter() Router
}

// WrapperImpl implements Wrapper
type WrapperImpl struct{}

// Check we implement the interface
var _ Wrapper = &WrapperImpl{}

// NewWrapperImpl is a constructor
func NewWrapperImpl() *WrapperImpl {
	return &WrapperImpl{}
}

// NewRouter implements Wrapper
func (w *WrapperImpl) NewRouter() Router {
	router := mux.NewRouter()
	return &RouterImpl{MuxRouter: router}
}
