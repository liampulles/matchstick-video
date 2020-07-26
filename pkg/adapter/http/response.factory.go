package http

import (
	"strconv"

	"github.com/liampulles/matchstick-video/pkg/domain/commonerror"
	"github.com/liampulles/matchstick-video/pkg/domain/entity"
)

const (
	jsonContentType = "application/json"
	textContentType = "text/plain; charset=utf-8"
)

// ResponseFactory constructs responses from various
// return types
type ResponseFactory interface {
	CreateJSON(statusCode uint, body []byte) *Response
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

// CreateJSON implements ResponseFactory
func (r *ResponseFactoryImpl) CreateJSON(statusCode uint, body []byte) *Response {
	return &Response{
		ContentType: jsonContentType,
		StatusCode:  statusCode,
		Body:        body,
	}
}

// CreateFromError implements ResponseFactory
func (r *ResponseFactoryImpl) CreateFromError(err error) *Response {
	switch v := err.(type) {
	// TODO: Add your specific error handlers for controllers here.
	case *commonerror.Validation:
		return r.createText(400, v.Error())
	case *commonerror.NotImplemented:
		return r.createText(501, v.Error())
	default:
		return r.createText(500, v.Error())
	}
}

// CreateFromEntityID implements ResponseFactory
func (r *ResponseFactoryImpl) CreateFromEntityID(statusCode uint, id entity.ID) *Response {
	str := strconv.FormatInt(int64(id), 10)
	return r.createText(statusCode, str)
}

func (r *ResponseFactoryImpl) createText(statusCode uint, body string) *Response {
	return &Response{
		ContentType: textContentType,
		StatusCode:  statusCode,
		Body:        []byte(body),
	}
}
