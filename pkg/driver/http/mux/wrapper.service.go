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

// HandleFunc wraps mux.Router.HandleFunc()
func (r *RouterImpl) HandleFunc(pathPattern string, handler Handler) Route {
	return r.MuxRouter.HandleFunc(pathPattern, handler)
}

// ServeHTTP wraps mux.Router.ServeHTTP()
func (r *RouterImpl) ServeHTTP(res goHttp.ResponseWriter, req *goHttp.Request) {
	r.MuxRouter.ServeHTTP(res, req)
	return
}

// Wrapper encapsulates mux methods
type Wrapper interface {
	NewRouter() Router
	Vars(*goHttp.Request) map[string]string
}

// WrapperImpl implements Wrapper
type WrapperImpl struct{}

// Check we implement the interface
var _ Wrapper = &WrapperImpl{}

// NewWrapperImpl is a constructor
func NewWrapperImpl() *WrapperImpl {
	return &WrapperImpl{}
}

// NewRouter wraps muc.NewRouter()
func (w *WrapperImpl) NewRouter() Router {
	router := mux.NewRouter()
	return &RouterImpl{MuxRouter: router}
}

// Vars wraps mux.Vars()
func (w *WrapperImpl) Vars(req *goHttp.Request) map[string]string {
	return mux.Vars(req)
}
