package mux_test

import (
	goHttp "net/http"
	"net/url"

	"testing"

	"github.com/stretchr/testify/assert"

	gorillaMux "github.com/gorilla/mux"

	"github.com/liampulles/matchstick-video/pkg/driver/http/mux"
)

func TestRouterImpl_HandleFunc_ShouldReturnMuxRoute(t *testing.T) {
	// Setup fixture
	sut := mux.RouterImpl{
		MuxRouter: gorillaMux.NewRouter(),
	}

	// Exercise SUT
	actual := sut.HandleFunc("/pattern", MockMuxHandler)

	// Verify results
	assert.IsType(t, &gorillaMux.Route{}, actual)
}

func TestRouterImpl_ServeHTTP_ShouldPass(t *testing.T) {
	// Setup fixture
	sut := mux.RouterImpl{
		MuxRouter: gorillaMux.NewRouter(),
	}
	requestFixture := &goHttp.Request{
		URL: &url.URL{},
	}

	// Exercise SUT
	sut.ServeHTTP(&testResponseWriter{}, requestFixture)
}

func TestWrapperImpl_NewRouter_ShouldReturnRouterImplWrappingMuxRouter(t *testing.T) {
	// Setup fixture
	sut := mux.NewWrapperImpl()

	// Setup expectations
	expectedMux := gorillaMux.NewRouter()
	expected := &mux.RouterImpl{
		MuxRouter: expectedMux,
	}

	// Exercise SUT
	actual := sut.NewRouter()

	// Verify results
	assert.Equal(t, expected, actual)
}

type testResponseWriter struct {
	data string
}

var _ goHttp.ResponseWriter = &testResponseWriter{}

func (t *testResponseWriter) Header() goHttp.Header {
	m := make(map[string][]string)
	return goHttp.Header(m)
}

func (t *testResponseWriter) Write(data []byte) (int, error) {
	return len(data), nil
}

func (t *testResponseWriter) WriteHeader(statusCode int) {
	return
}
