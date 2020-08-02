package db_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/liampulles/matchstick-video/pkg/adapter/db"
)

type ErrorParserTestSuite struct {
	suite.Suite
	sut *db.ErrorParserImpl
}

func TestErrorParserTestSuite(t *testing.T) {
	suite.Run(t, new(ErrorParserTestSuite))
}

func (suite *ErrorParserTestSuite) SetupTest() {
	suite.sut = db.NewErrorParserImpl()
}

func (suite *ErrorParserTestSuite) TestFromDBRowScan_WhenIsUniquenessConstraint_ShouldReturnUniquenessConstraintError() {
	// Setup fixture
	fixture := fmt.Errorf("something violates unique constraint")

	// Setup expectations
	expectedErr := "uniqueness constraint error: something violates unique constraint"

	// Exercise SUT
	err := suite.sut.FromDBRowScan(fixture, "some.type")

	// Verify results
	suite.EqualError(err, expectedErr)
}

func (suite *ErrorParserTestSuite) TestFromDBRowScan_WhenIsNoRowsError_ShouldReturnNotFoundError() {
	// Setup fixture
	fixture := fmt.Errorf("there were no rows in result set")

	// Setup expectations
	expectedErr := "entity not found: type=[some.type]"

	// Exercise SUT
	err := suite.sut.FromDBRowScan(fixture, "some.type")

	// Verify results
	suite.EqualError(err, expectedErr)
}

func (suite *ErrorParserTestSuite) TestFromDBRowScan_WhenIsArbitraryError_ShouldReturnSameError() {
	// Setup fixture
	fixture := fmt.Errorf("some random error")

	// Setup expectations
	expectedErr := "some random error"

	// Exercise SUT
	err := suite.sut.FromDBRowScan(fixture, "some.type")

	// Verify results
	suite.EqualError(err, expectedErr)
}
