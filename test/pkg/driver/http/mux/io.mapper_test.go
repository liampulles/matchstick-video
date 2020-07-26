package mux_test

import (
	"bytes"
	"errors"
	"io/ioutil"
	goHttp "net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/suite"

	httpMocks "github.com/liampulles/matchstick-video/test/mock/go/net/http"
	muxMocks "github.com/liampulles/matchstick-video/test/mock/pkg/driver/http/mux"

	adapterHttp "github.com/liampulles/matchstick-video/pkg/adapter/http"
	muxDriver "github.com/liampulles/matchstick-video/pkg/driver/http/mux"
)

type IOMapperImplTestSuite struct {
	suite.Suite
	mockMuxWrapper *muxMocks.MockWrapper
	sut            *muxDriver.IOMapperImpl
}

func TestIOMapperImplTestSuite(t *testing.T) {
	suite.Run(t, new(IOMapperImplTestSuite))
}

func (suite *IOMapperImplTestSuite) SetupTest() {
	suite.mockMuxWrapper = &muxMocks.MockWrapper{}
	suite.sut = muxDriver.NewIOMapperImpl(
		suite.mockMuxWrapper,
	)
}

func (suite *IOMapperImplTestSuite) TestMapRequest_WhenExtractBodyFails_ShouldFail() {
	// Setup fixture
	requestFixture := &goHttp.Request{
		URL: &url.URL{
			RawQuery: "something",
		},
		Body: errReader(0),
	}

	// Setup expectations
	expectedPathParams := map[string]string{
		"path": "param",
	}
	expectedErr := "could not extract body: test error"

	// Setup mocks
	suite.mockMuxWrapper.On("Vars", requestFixture).
		Return(expectedPathParams)

	// Exercise SUT
	actual, err := suite.sut.MapRequest(requestFixture)

	// Verify mocks
	suite.Nil(actual)
	suite.EqualError(err, expectedErr)
}

func (suite *IOMapperImplTestSuite) TestMapRequest_WhenExtractBodyPasses_ShouldReturnAdapterRequest() {
	// Setup fixture
	body := ioutil.NopCloser(bytes.NewReader([]byte("some.data")))
	requestFixture := &goHttp.Request{
		URL: &url.URL{
			RawQuery: "something",
		},
		Body: body,
	}

	// Setup expectations
	expectedPathParams := map[string]string{
		"path": "param",
	}
	expected := &adapterHttp.Request{
		PathParam:  expectedPathParams,
		QueryParam: map[string][]string{"something": []string{""}},
		Body:       []byte("some.data"),
	}

	// Setup mocks
	suite.mockMuxWrapper.On("Vars", requestFixture).
		Return(expectedPathParams)

	// Exercise SUT
	actual, err := suite.sut.MapRequest(requestFixture)

	// Verify mocks
	suite.Nil(err)
	suite.Equal(expected, actual)
}

func (suite *IOMapperImplTestSuite) TestMapResponse_ShouldWriteToResponse() {
	// Setup fixture
	respFixture := &adapterHttp.Response{
		ContentType: "some.content.type",
		StatusCode:  101,
		Body:        []byte("some.data"),
	}

	// Setup mocks
	mockResponse := &httpMocks.MockResponseWriter{}
	mockHeaders := goHttp.Header(make(map[string][]string))
	mockResponse.On("Header").Return(mockHeaders)
	mockResponse.On("WriteHeader", 101).Return()
	mockResponse.On("Write", []byte("some.data")).Return(0, nil)

	// Exercise SUT
	suite.sut.MapResponse(respFixture, mockResponse)

	// Verify mocks
	mockResponse.AssertExpectations(suite.T())
	suite.Equal([]string{"some.content.type"}, mockHeaders["Content-Type"])
}

type errReader int

func (errReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("test error")
}

func (errReader) Close() error {
	return nil
}
