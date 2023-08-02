package db_test

import (
	"fmt"
	"testing"

	"github.com/liampulles/matchstick-video/pkg/adapter/db"
	"github.com/stretchr/testify/suite"
)

type ErrorParserTestSuite struct {
	suite.Suite
}

func TestErrorParserTestSuite(t *testing.T) {
	suite.Run(t, new(ErrorParserTestSuite))
}

func (suite *ErrorParserTestSuite) SetupTest() {}

func (suite *ErrorParserTestSuite) TestFromDBRowScan_WhenIsUniquenessConstraint_ShouldReturnUniquenessConstraintError() {
	// Setup fixture
	fixture := fmt.Errorf("something violates unique constraint")

	// Setup expectations
	expectedErr := "uniqueness constraint error: something violates unique constraint"

	// Exercise SUT
	err := db.FromDBRowScan(fixture, "some.type")

	// Verify results
	suite.EqualError(err, expectedErr)
}

func (suite *ErrorParserTestSuite) TestFromDBRowScan_WhenIsNoRowsError_ShouldReturnNotFoundError() {
	// Setup fixture
	fixture := fmt.Errorf("there were no rows in result set")

	// Setup expectations
	expectedErr := "entity not found: type=[some.type]"

	// Exercise SUT
	err := db.FromDBRowScan(fixture, "some.type")

	// Verify results
	suite.EqualError(err, expectedErr)
}

func (suite *ErrorParserTestSuite) TestFromDBRowScan_WhenIsArbitraryError_ShouldReturnSameError() {
	// Setup fixture
	fixture := fmt.Errorf("some random error")

	// Setup expectations
	expectedErr := "some random error"

	// Exercise SUT
	err := db.FromDBRowScan(fixture, "some.type")

	// Verify results
	suite.EqualError(err, expectedErr)
}
