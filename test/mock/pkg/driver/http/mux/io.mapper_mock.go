package mux

import (
	goHttp "net/http"

	"github.com/stretchr/testify/mock"

	adapterHttp "github.com/liampulles/matchstick-video/pkg/adapter/http"
	muxDriver "github.com/liampulles/matchstick-video/pkg/driver/http/mux"
)

// IOMapperMock is for mocking
type IOMapperMock struct {
	mock.Mock
}

var _ muxDriver.IOMapper = &IOMapperMock{}

// MapRequest is for mocking
func (i *IOMapperMock) MapRequest(req *goHttp.Request) (*adapterHttp.Request, error) {
	args := i.Called(req)
	return safeArgsGetRequest(args, 0), args.Error(1)
}

// MapResponse is for mocking
func (i *IOMapperMock) MapResponse(adapterRes *adapterHttp.Response, goRes goHttp.ResponseWriter) {
	i.Called(adapterRes, goRes)
	return
}

func safeArgsGetRequest(args mock.Arguments, idx int) *adapterHttp.Request {
	if val, ok := args.Get(idx).(*adapterHttp.Request); ok {
		return val
	}
	return nil
}
