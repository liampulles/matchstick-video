package http

import (
	"github.com/stretchr/testify/mock"

	"github.com/liampulles/matchstick-video/pkg/adapter/http"
	"github.com/liampulles/matchstick-video/pkg/domain/entity"
)

// ResponseFactoryMock is for mocking
type ResponseFactoryMock struct {
	mock.Mock
}

var _ http.ResponseFactory = &ResponseFactoryMock{}

// CreateJSON is for mocking
func (r *ResponseFactoryMock) CreateJSON(statusCode uint, body []byte) *http.Response {
	args := r.Called(statusCode, body)
	return args.Get(0).(*http.Response)
}

// CreateFromError is for mocking
func (r *ResponseFactoryMock) CreateFromError(err error) *http.Response {
	args := r.Called(err)
	return args.Get(0).(*http.Response)
}

// CreateFromEntityID is for mocking
func (r *ResponseFactoryMock) CreateFromEntityID(statusCode uint, id entity.ID) *http.Response {
	args := r.Called(statusCode, id)
	return args.Get(0).(*http.Response)
}
