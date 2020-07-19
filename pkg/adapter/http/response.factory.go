package http

import (
	"github.com/liampulles/matchstick-video/pkg/domain/entity"
)

// ResponseFactory constructs responses from various
// return types
type ResponseFactory interface {
	CreateFromError(error) *Response
	CreateFromEntityID(statusCode uint, id entity.ID) *Response
}
