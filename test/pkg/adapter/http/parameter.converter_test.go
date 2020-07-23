package http_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/liampulles/matchstick-video/pkg/adapter/http"
	"github.com/liampulles/matchstick-video/pkg/domain/entity"
)

type ParameterConverterImplTestSuite struct {
	suite.Suite
	sut *http.ParameterConverterImpl
}

func TestParameterConverterImplTestSuite(t *testing.T) {
	suite.Run(t, new(ParameterConverterImplTestSuite))
}

func (suite *ParameterConverterImplTestSuite) SetupTest() {
	suite.sut = http.NewParameterConverterImpl()
}

func (suite *ParameterConverterImplTestSuite) TestToEntityID_WhenValueNotPresent_ShouldFail() {
	// Setup fixture
	mapFixture := map[string]string{
		"something": "else",
	}

	// Setup expectations
	expectedErr := "could not convert parameters to entity id - \"id\" is not in the parameter list"

	// Exercise SUT
	actual, err := suite.sut.ToEntityID(mapFixture, "id")

	// Verify results
	suite.Equal(entity.InvalidID, actual)
	suite.EqualError(err, expectedErr)
}

func (suite *ParameterConverterImplTestSuite) TestToEntityID_WhenValueIsWrongType_ShouldFail() {
	// Setup fixture
	mapFixture := map[string]string{
		"something": "else",
		"id":        "not.an.int",
		"and":       "another.thing",
	}

	// Setup expectations
	expectedErr := "could not convert parameters to entity id - cannot convert to int64"

	// Exercise SUT
	actual, err := suite.sut.ToEntityID(mapFixture, "id")

	// Verify results
	suite.Equal(entity.InvalidID, actual)
	suite.EqualError(err, expectedErr)
}

func (suite *ParameterConverterImplTestSuite) TestToEntityID_WhenValueIsRightType_ShouldReturnEntityID() {
	// Setup fixture
	mapFixture := map[string]string{
		"something": "else",
		"id":        "101",
		"and":       "another.thing",
	}

	// Setup expectations
	expected := entity.ID(101)

	// Exercise SUT
	actual, err := suite.sut.ToEntityID(mapFixture, "id")

	// Verify results
	suite.NoError(err)
	suite.Equal(expected, actual)
}
