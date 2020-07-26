package mux

import (
	goHttp "net/http"

	"github.com/liampulles/matchstick-video/pkg/adapter/http"
)

// Handler is a handler that mux accepts
type Handler func(goHttp.ResponseWriter, *goHttp.Request)

// HandlerMapper wraps adapter Handlers into mux handlers.
type HandlerMapper interface {
	Map(http.Handler) Handler
}

// HandlerMapperImpl implements HandlerMapper
type HandlerMapperImpl struct {
	ioMapper IOMapper
}

var _ HandlerMapper = &HandlerMapperImpl{}

// NewHandlerMapperImpl is a constructor
func NewHandlerMapperImpl(ioMapper IOMapper) *HandlerMapperImpl {
	return &HandlerMapperImpl{
		ioMapper: ioMapper,
	}
}

// Map implements HandlerMapper
func (h *HandlerMapperImpl) Map(handler http.Handler) Handler {
	return func(res goHttp.ResponseWriter, req *goHttp.Request) {
		// Convert go request to adapter request
		adapterReq, err := h.ioMapper.MapRequest(req)
		if err != nil {
			badRequest(res, err)
			return
		}

		// Run handler
		adapterRes := handler(adapterReq)

		// Convert adapter response to go response
		h.ioMapper.MapResponse(adapterRes, res)
	}
}

func badRequest(res goHttp.ResponseWriter, err error) {
	res.WriteHeader(400)
	res.Write([]byte(err.Error()))
}
