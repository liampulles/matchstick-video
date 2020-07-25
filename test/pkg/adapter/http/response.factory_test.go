package http_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/liampulles/matchstick-video/pkg/adapter/http"
	"github.com/liampulles/matchstick-video/pkg/domain/commonerror"
	"github.com/liampulles/matchstick-video/pkg/domain/entity"
)

type ResponseFactoryImplTestSuite struct {
	suite.Suite
	sut *http.ResponseFactoryImpl
}

func TestResponseFactoryImplTestSuite(t *testing.T) {
	suite.Run(t, new(ResponseFactoryImplTestSuite))
}

func (suite *ResponseFactoryImplTestSuite) SetupTest() {
	suite.sut = http.NewResponseFactoryImpl()
}

func (suite *ResponseFactoryImplTestSuite) TestCreate_ShouldCreateResponse() {
	// Setup expectations
	expected := &http.Response{
		StatusCode: 501,
		Body:       []byte("some.data"),
	}

	// Exercise SUT
	actual := suite.sut.Create(501, []byte("some.data"))

	// Verify results
	suite.Equal(expected, actual)
}

func (suite *ResponseFactoryImplTestSuite) TestCreateFromError_WhenIsValidationError_ShouldReturnBadRequest() {
	// Setup fixture
	fixture := &commonerror.Validation{
		Field:   "id",
		Problem: "not numeric",
	}

	// Setup expectations
	expected := &http.Response{
		StatusCode: 400,
		Body:       []byte("validation error: field=[id], problem=[not numeric]"),
	}

	// Exercise SUT
	actual := suite.sut.CreateFromError(fixture)

	// Verify results
	suite.Equal(expected, actual)
}

func (suite *ResponseFactoryImplTestSuite) TestCreateFromError_WhenIsArbitraryError_ShouldReturnInternalServerError() {
	// Setup fixture
	fixture := fmt.Errorf("some.error")

	// Setup expectations
	expected := &http.Response{
		StatusCode: 500,
		Body:       []byte("some.error"),
	}

	// Exercise SUT
	actual := suite.sut.CreateFromError(fixture)

	// Verify results
	suite.Equal(expected, actual)
}

func (suite *ResponseFactoryImplTestSuite) TestCreateFromEntityID_ShouldReturnIdAsString() {
	// Setup expectations
	expected := &http.Response{
		StatusCode: 201,
		Body:       []byte("101"),
	}

	// Exercise SUT
	actual := suite.sut.CreateFromEntityID(201, entity.ID(101))

	// Verify results
	suite.Equal(expected, actual)
}
