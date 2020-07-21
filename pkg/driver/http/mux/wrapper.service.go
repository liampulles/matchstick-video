package mux

import (
	goHttp "net/http"
)

// Route wraps mux.Route
type Route interface {
	Methods(...string) Route
}

// Router wraps mux.Router
type Router interface {
	goHttp.Handler
	HandleFunc(pathPattern string, handler Handler) Route
}

// Wrapper encapsulates mux methods
type Wrapper interface {
	NewRouter() Router
}
