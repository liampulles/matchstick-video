package http

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

// ResponseWriterMock is for mocking
type ResponseWriterMock struct {
	mock.Mock
}

var _ http.ResponseWriter = &ResponseWriterMock{}

// Header is for mocking
func (r *ResponseWriterMock) Header() http.Header {
	args := r.Called()
	return args.Get(0).(http.Header)
}

func (r *ResponseWriterMock) Write(data []byte) (int, error) {
	args := r.Called(data)
	return args.Int(0), args.Error(1)
}

// WriteHeader is for mocking
func (r *ResponseWriterMock) WriteHeader(statusCode int) {
	r.Called(statusCode)
	return
}
