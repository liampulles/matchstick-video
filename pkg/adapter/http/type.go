package http

// Response defines what we return after
// a given HTTP request.
type Response struct {
	StatusCode uint
	Body       []byte
}

// Handler handles an HTTP request
// and generates a response.
type Handler func(
	pathParam []string,
	queryParam map[string]string,
	body []byte,
) *Response

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
