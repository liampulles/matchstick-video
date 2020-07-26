package http

import (
	"github.com/stretchr/testify/mock"

	"github.com/liampulles/matchstick-video/pkg/adapter/http"
	"github.com/liampulles/matchstick-video/pkg/domain/entity"
)

// MockResponseFactory is for mocking
type MockResponseFactory struct {
	mock.Mock
}

var _ http.ResponseFactory = &MockResponseFactory{}

// CreateJSON is for mocking
func (r *MockResponseFactory) CreateJSON(statusCode uint, body []byte) *http.Response {
	args := r.Called(statusCode, body)
	return args.Get(0).(*http.Response)
}

// CreateFromError is for mocking
func (r *MockResponseFactory) CreateFromError(err error) *http.Response {
	args := r.Called(err)
	return args.Get(0).(*http.Response)
}

// CreateFromEntityID is for mocking
func (r *MockResponseFactory) CreateFromEntityID(statusCode uint, id entity.ID) *http.Response {
	args := r.Called(statusCode, id)
	return args.Get(0).(*http.Response)
}
