package mux_test

import (
	"fmt"
	goHttp "net/http"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	httpMocks "github.com/liampulles/matchstick-video/test/mock/go/net/http"
	muxMocks "github.com/liampulles/matchstick-video/test/mock/pkg/driver/http/mux"

	"github.com/liampulles/matchstick-video/pkg/adapter/http"
	muxDriver "github.com/liampulles/matchstick-video/pkg/driver/http/mux"
)

type HandlerMapperImplTestSuite struct {
	suite.Suite
	mockIoMapper *muxMocks.MockIOMapper
	sut          *muxDriver.HandlerMapperImpl
}

func TestHandlerMapperImplTestSuite(t *testing.T) {
	suite.Run(t, new(HandlerMapperImplTestSuite))
}

func (suite *HandlerMapperImplTestSuite) SetupTest() {
	suite.mockIoMapper = &muxMocks.MockIOMapper{}
	suite.sut = muxDriver.NewHandlerMapperImpl(
		suite.mockIoMapper,
	)
}

func (suite *HandlerMapperImplTestSuite) TestMap_WhenIoMapperRequestFails_ShouldFail() {
	// Setup fixture
	requestFixture := &goHttp.Request{}

	// Setup mocks
	mockHandler := &mockHandlerStruct{}
	mockResponse := &httpMocks.MockResponseWriter{}
	mockErr := fmt.Errorf("mock.error")
	suite.mockIoMapper.On("MapRequest", requestFixture).
		Return(nil, mockErr)
	mockResponse.On("WriteHeader", 400).Return()
	mockResponse.On("Write", []byte("mock.error")).Return(0, nil)

	// Exercise SUT
	actual := suite.sut.Map(mockHandler.MockHandler)
	actual(mockResponse, requestFixture)

	// Verify results
	mockResponse.AssertExpectations(suite.T())
}

func (suite *HandlerMapperImplTestSuite) TestMap_WhenIoMapperResponseReturns_ShouldWriteResponseAsExpected() {
	// Setup fixture
	requestFixture := &goHttp.Request{}

	// Setup mocks
	mockHandler := &mockHandlerStruct{}
	mockResponse := &httpMocks.MockResponseWriter{}
	mockAdapterReq := &http.Request{Body: []byte("some.data")}
	mockAdapterResp := &http.Response{StatusCode: 200}
	suite.mockIoMapper.On("MapRequest", requestFixture).
		Return(mockAdapterReq, nil)
	mockHandler.On("MockHandler", mockAdapterReq).
		Return(mockAdapterResp)
	suite.mockIoMapper.On("MapResponse", mockAdapterResp, mockResponse).
		Return()

	// Exercise SUT
	actual := suite.sut.Map(mockHandler.MockHandler)
	actual(mockResponse, requestFixture)

	// Verify results
	suite.mockIoMapper.AssertExpectations(suite.T())
}

type mockHandlerStruct struct {
	mock.Mock
}

func (m *mockHandlerStruct) MockHandler(req *http.Request) *http.Response {
	args := m.Called(req)
	return safeArgsGetResponse(args, 0)
}

func safeArgsGetResponse(args mock.Arguments, idx int) *http.Response {
	if val, ok := args.Get(idx).(*http.Response); ok {
		return val
	}
	return nil
}
