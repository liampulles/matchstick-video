package http

import (
	"strconv"

	"github.com/liampulles/matchstick-video/pkg/domain/commonerror"
	"github.com/liampulles/matchstick-video/pkg/domain/entity"
)

// ResponseFactory constructs responses from various
// return types
type ResponseFactory interface {
	Create(statusCode uint, body []byte) *Response
	CreateFromError(error) *Response
	CreateFromEntityID(statusCode uint, id entity.ID) *Response
}

// ResponseFactoryImpl implements ResponseFactory
type ResponseFactoryImpl struct{}

// Check that we implement the interface
var _ ResponseFactory = &ResponseFactoryImpl{}

// NewResponseFactoryImpl is a constructor
func NewResponseFactoryImpl() *ResponseFactoryImpl {
	return &ResponseFactoryImpl{}
}

// Create implements ResponseFactory
func (r *ResponseFactoryImpl) Create(statusCode uint, body []byte) *Response {
	return &Response{
		StatusCode: statusCode,
		Body:       body,
	}
}

// CreateFromError implements ResponseFactory
func (r *ResponseFactoryImpl) CreateFromError(err error) *Response {
	switch v := err.(type) {
	// TODO: Add your specific error handlers for controllers here.
	case *commonerror.Validation:
		return r.create(400, v.Error())
	case *commonerror.NotImplemented:
		return r.create(501, v.Error())
	default:
		return r.create(500, v.Error())
	}
}

// CreateFromEntityID implements ResponseFactory
func (r *ResponseFactoryImpl) CreateFromEntityID(statusCode uint, id entity.ID) *Response {
	str := strconv.FormatInt(int64(id), 10)
	return r.create(statusCode, str)
}

func (r *ResponseFactoryImpl) create(statusCode uint, body string) *Response {
	return r.Create(statusCode, []byte(body))
}
