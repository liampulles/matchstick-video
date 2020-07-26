package http_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/liampulles/matchstick-video/pkg/adapter/db"
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

func (suite *ResponseFactoryImplTestSuite) TestCreateEmpty_ShouldCreateResponse() {
	// Setup expectations
	expected := &http.Response{
		StatusCode: 501,
	}

	// Exercise SUT
	actual := suite.sut.CreateEmpty(501)

	// Verify results
	suite.Equal(expected, actual)
}

func (suite *ResponseFactoryImplTestSuite) TestCreateJSON_ShouldCreateResponse() {
	// Setup expectations
	expected := &http.Response{
		ContentType: "application/json",
		StatusCode:  501,
		Body:        []byte("some.data"),
	}

	// Exercise SUT
	actual := suite.sut.CreateJSON(501, []byte("some.data"))

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
		ContentType: "text/plain; charset=utf-8",
		StatusCode:  400,
		Body:        []byte("validation error: field=[id], problem=[not numeric]"),
	}

	// Exercise SUT
	actual := suite.sut.CreateFromError(fixture)

	// Verify results
	suite.Equal(expected, actual)
}

func (suite *ResponseFactoryImplTestSuite) TestCreateFromError_WhenIsNotImplementedError_ShouldReturnNotImplemented() {
	// Setup fixture
	fixture := commonerror.NewNotImplemented("some.package", "some.struct", "some.method")

	// Setup expectations
	expected := &http.Response{
		ContentType: "text/plain; charset=utf-8",
		StatusCode:  501,
		Body:        []byte("method not implemented for package=[some.package], struct=[some.struct], method=[some.method]"),
	}

	// Exercise SUT
	actual := suite.sut.CreateFromError(fixture)

	// Verify results
	suite.Equal(expected, actual)
}

func (suite *ResponseFactoryImplTestSuite) TestCreateFromError_WhenIsNotFoundError_ShouldReturnNotFound() {
	// Setup fixture
	fixture := db.NewNotFoundError("some.type")

	// Setup expectations
	expected := &http.Response{
		ContentType: "text/plain; charset=utf-8",
		StatusCode:  404,
		Body:        []byte("entity not found: type=[some.type]"),
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
		ContentType: "text/plain; charset=utf-8",
		StatusCode:  500,
		Body:        []byte("some.error"),
	}

	// Exercise SUT
	actual := suite.sut.CreateFromError(fixture)

	// Verify results
	suite.Equal(expected, actual)
}

func (suite *ResponseFactoryImplTestSuite) TestCreateFromEntityID_ShouldReturnIdAsString() {
	// Setup expectations
	expected := &http.Response{
		ContentType: "text/plain; charset=utf-8",
		StatusCode:  201,
		Body:        []byte("101"),
	}

	// Exercise SUT
	actual := suite.sut.CreateFromEntityID(201, entity.ID(101))

	// Verify results
	suite.Equal(expected, actual)
}
