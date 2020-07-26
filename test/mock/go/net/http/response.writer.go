package http

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

// MockResponseWriter is for mocking
type MockResponseWriter struct {
	mock.Mock
}

var _ http.ResponseWriter = &MockResponseWriter{}

// Header is for mocking
func (r *MockResponseWriter) Header() http.Header {
	args := r.Called()
	return args.Get(0).(http.Header)
}

func (r *MockResponseWriter) Write(data []byte) (int, error) {
	args := r.Called(data)
	return args.Int(0), args.Error(1)
}

// WriteHeader is for mocking
func (r *MockResponseWriter) WriteHeader(statusCode int) {
	r.Called(statusCode)
	return
}
