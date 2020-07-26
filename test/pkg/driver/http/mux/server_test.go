package mux_test

import (
	goHttp "net/http"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	configMocks "github.com/liampulles/matchstick-video/test/mock/pkg/adapter/config"
	muxMocks "github.com/liampulles/matchstick-video/test/mock/pkg/driver/http/mux"

	"github.com/liampulles/matchstick-video/pkg/adapter/http"
	muxDriver "github.com/liampulles/matchstick-video/pkg/driver/http/mux"
)

type ServerConfigurationImplTestSuite struct {
	suite.Suite
	mockConfigStore   *configMocks.MockStore
	mockHandlerMapper *muxMocks.MockHandlerMapper
	mockMuxWrapper    *muxMocks.MockWrapper
	sut               *muxDriver.ServerConfigurationImpl
}

func TestServerConfigurationImplTestSuite(t *testing.T) {
	suite.Run(t, new(ServerConfigurationImplTestSuite))
}

func (suite *ServerConfigurationImplTestSuite) SetupTest() {
	suite.mockConfigStore = &configMocks.MockStore{}
	suite.mockHandlerMapper = &muxMocks.MockHandlerMapper{}
	suite.mockMuxWrapper = &muxMocks.MockWrapper{}
	suite.sut = muxDriver.NewServerConfigurationImpl(
		suite.mockConfigStore,
		suite.mockHandlerMapper,
		suite.mockMuxWrapper,
	)
}

func (suite *ServerConfigurationImplTestSuite) TestCreateRunnable_ShouldMapAllHandlers() {
	// Setup fixture
	fixture := map[http.HandlerPattern]http.Handler{
		http.HandlerPattern{
			Method:      "method.1",
			PathPattern: "path.pattern.1",
		}: mockHander1,
		http.HandlerPattern{
			Method:      "method.2",
			PathPattern: "path.pattern.2",
		}: mockHander2,
	}

	// Setup mocks
	mockRouter := &muxMocks.RouterMock{}
	mockRoute1 := &muxMocks.RouteMock{}
	mockRoute2 := &muxMocks.RouteMock{}
	suite.mockMuxWrapper.On("NewRouter").
		Return(mockRouter)
	suite.mockHandlerMapper.On("Map", mock.Anything).
		Return(MockMuxHandler)
	mockRouter.On("HandleFunc", "path.pattern.1", mock.Anything).
		Return(mockRoute1)
	mockRoute1.On("Methods", []string{"method.1"}).
		Return(nil)
	mockRouter.On("HandleFunc", "path.pattern.2", mock.Anything).
		Return(mockRoute2)
	mockRoute2.On("Methods", []string{"method.2"}).
		Return(nil)
	suite.mockConfigStore.On("GetPort").
		Return(101)

	// Exercise SUT
	suite.sut.CreateRunnable(fixture)

	// Verify mocks

}

func mockHander1(req *http.Request) *http.Response {
	return nil
}

func mockHander2(req *http.Request) *http.Response {
	return nil
}

func MockMuxHandler(res goHttp.ResponseWriter, req *goHttp.Request) {
	return
}
