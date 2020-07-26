package http

// Request defines everything a user can submit
// via HTTP for us to process
type Request struct {
	PathParam  map[string]string
	QueryParam map[string][]string
	Body       []byte
}

// Response defines what we return after
// a given HTTP request.
type Response struct {
	// TODO: include content-type
	StatusCode uint
	Body       []byte
}

// Handler handles an HTTP request
// and generates a response.
type Handler func(*Request) *Response

// HandlerPattern defines a unique set of HTTP
// request properties which one can then map
// to a handler.
type HandlerPattern struct {
	Method      string
	PathPattern string
}

// A Controller defines handlers
type Controller interface {
	GetHandlers() map[HandlerPattern]Handler
}

func addHandler(
	handlers map[HandlerPattern]Handler,
	method string,
	pathPattern string,
	handler Handler,
) {
	handlerPattern := HandlerPattern{
		Method:      method,
		PathPattern: pathPattern,
	}
	handlers[handlerPattern] = handler
}
