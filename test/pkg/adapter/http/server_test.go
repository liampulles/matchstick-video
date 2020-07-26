package http_test

import (
	goHttp "net/http"
	"testing"

	"github.com/stretchr/testify/suite"

	httpMocks "github.com/liampulles/matchstick-video/test/mock/pkg/adapter/http"

	"github.com/liampulles/matchstick-video/pkg/adapter/http"
	"github.com/liampulles/matchstick-video/pkg/domain"
)

type ServerFactoryTestSuite struct {
	suite.Suite
	mockInventoryController *httpMocks.MockController
	mockServerConfiguration *httpMocks.MockServerConfiguration
	sut                     *http.ServerFactoryImpl
}

func TestServerFactoryTestSuite(t *testing.T) {
	suite.Run(t, new(ServerFactoryTestSuite))
}

func (suite *ServerFactoryTestSuite) SetupTest() {
	suite.mockInventoryController = &httpMocks.MockController{}
	suite.mockServerConfiguration = &httpMocks.MockServerConfiguration{}
	suite.sut = http.NewServerFactoryImpl(
		suite.mockInventoryController,
		suite.mockServerConfiguration,
	)
}

func (suite *ServerFactoryTestSuite) TestCreate_ShouldCreateRunnableFromHandlers() {
	// Setup expectations
	data := "previous"
	expectedRunnable := domain.Runnable(func() error {
		data = "after"
		return nil
	})

	// Setup mocks
	mockHandlers := map[http.HandlerPattern]http.Handler{
		http.HandlerPattern{
			Method:      goHttp.MethodGet,
			PathPattern: "some.path.pattern",
		}: mockHandler,
	}
	suite.mockInventoryController.On("GetHandlers").
		Return(mockHandlers)
	suite.mockServerConfiguration.On("CreateRunnable", mockHandlers).
		Return(expectedRunnable)

	// Exercise SUT
	actual := suite.sut.Create()

	// Verify results
	actual()
	suite.Equal(data, "after")
}

func mockHandler(*http.Request) *http.Response {
	return nil
}

func mockRunnable() error {
	return nil
}
