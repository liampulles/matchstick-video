package sql_test

import (
	goSql "database/sql"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"github.com/liampulles/matchstick-video/pkg/adapter/db/sql"
	"github.com/liampulles/matchstick-video/pkg/domain/entity"
)

type HelperServiceTestSuite struct {
	suite.Suite
	db     *goSql.DB
	mockDb sqlmock.Sqlmock
	sut    *sql.HelperServiceImpl
}

func TestHelperServiceTestSuite(t *testing.T) {
	suite.Run(t, new(HelperServiceTestSuite))
}

func (suite *HelperServiceTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}
	suite.db = db
	suite.mockDb = mock
	suite.sut = sql.NewHelperServiceImpl()
}

func (suite *HelperServiceTestSuite) TestExecForRowsAffected_WhenPrepareContextFails_ShouldFail() {
	// Setup fixture
	queryFixture := "some.query"
	arg1Fixture := "arg.1"
	arg2Fixture := 2

	// Setup expectations
	expectedErr := "mock.error"

	// Setup mocks
	mockErr := fmt.Errorf("mock.error")
	suite.mockDb.ExpectPrepare(queryFixture).
		WillReturnError(mockErr)

	// Exercise SUT
	actual, err := suite.sut.ExecForRowsAffected(suite.db, queryFixture, arg1Fixture, arg2Fixture)

	// Verify results
	suite.Equal(int64(-1), actual)
	suite.EqualError(err, expectedErr)
}

func (suite *HelperServiceTestSuite) TestExecForRowsAffected_WhenExecContextFails_ShouldFail() {
	// Setup fixture
	queryFixture := "some.query"
	arg1Fixture := "arg.1"
	arg2Fixture := 2

	// Setup expectations
	expectedErr := "mock.error"

	// Setup mocks
	mockErr := fmt.Errorf("mock.error")
	suite.mockDb.ExpectPrepare(queryFixture).
		ExpectExec().
		WithArgs(arg1Fixture, arg2Fixture).
		WillReturnError(mockErr)

	// Exercise SUT
	actual, err := suite.sut.ExecForRowsAffected(suite.db, queryFixture, arg1Fixture, arg2Fixture)

	// Verify results
	suite.Equal(int64(-1), actual)
	suite.EqualError(err, expectedErr)
}

func (suite *HelperServiceTestSuite) TestExecForRowsAffected_WhenRowsAffectedFails_ShouldFail() {
	// Setup fixture
	queryFixture := "some.query"
	arg1Fixture := "arg.1"
	arg2Fixture := 2

	// Setup expectations
	expectedErr := "mock.error"

	// Setup mocks
	mockResult := &mockResult{}
	mockErr := fmt.Errorf("mock.error")
	suite.mockDb.ExpectPrepare(queryFixture).
		ExpectExec().
		WithArgs(arg1Fixture, arg2Fixture).
		WillReturnResult(mockResult)
	mockResult.On("RowsAffected").Return(int64(-1), mockErr)

	// Exercise SUT
	actual, err := suite.sut.ExecForRowsAffected(suite.db, queryFixture, arg1Fixture, arg2Fixture)

	// Verify results
	suite.Equal(int64(-1), actual)
	suite.EqualError(err, expectedErr)
}

func (suite *HelperServiceTestSuite) TestExecForRowsAffected_WhenRowsAffectedPasses_ShouldReturnAsExpected() {
	// Setup fixture
	queryFixture := "some.query"
	arg1Fixture := "arg.1"
	arg2Fixture := 2

	// Setup mocks
	mockResult := &mockResult{}
	suite.mockDb.ExpectPrepare(queryFixture).
		ExpectExec().
		WithArgs(arg1Fixture, arg2Fixture).
		WillReturnResult(mockResult)
	mockResult.On("RowsAffected").Return(int64(25), nil)

	// Exercise SUT
	actual, err := suite.sut.ExecForRowsAffected(suite.db, queryFixture, arg1Fixture, arg2Fixture)

	// Verify results
	suite.NoError(err)
	suite.Equal(int64(25), actual)
}

func (suite *HelperServiceTestSuite) TestSingleRowQuery_WhenPrepareContextFails_ShouldFail() {
	// Setup fixture
	queryFixture := "some.query"
	arg1Fixture := "arg.1"
	arg2Fixture := 2

	// Setup expectations
	expectedErr := "mock.error"

	// Setup mocks
	mockErr := fmt.Errorf(expectedErr)
	suite.mockDb.ExpectPrepare(queryFixture).
		WillReturnError(mockErr)

	// Exercise SUT
	actual, err := suite.sut.SingleRowQuery(suite.db, queryFixture, arg1Fixture, arg2Fixture)

	// Verify results
	suite.Nil(actual)
	suite.EqualError(err, expectedErr)
}

func (suite *HelperServiceTestSuite) TestSingleRowQuery_WhenQueryRowContextPasses_ShouldReturnAsExpected() {
	// Setup fixture
	queryFixture := "some.query"
	arg1Fixture := "arg.1"
	arg2Fixture := 2

	// Setup mocks
	mockRows := suite.mockDb.NewRows([]string{"some", "columns"}).
		FromCSVString("with,data")
	suite.mockDb.ExpectPrepare(queryFixture).
		ExpectQuery().
		WithArgs(arg1Fixture, arg2Fixture).
		WillReturnRows(mockRows)

	// Exercise SUT
	actual, err := suite.sut.SingleRowQuery(suite.db, queryFixture, arg1Fixture, arg2Fixture)

	// Verify results
	suite.NoError(err)
	suite.NotNil(actual)
}

func (suite *HelperServiceTestSuite) TestSingleQueryForID_WhenPrepareContextFails_ShouldFail() {
	// Setup fixture
	queryFixture := "some.query"
	arg1Fixture := "arg.1"
	arg2Fixture := 2

	// Setup expectations
	expectedErr := "mock.error"

	// Setup mocks
	mockErr := fmt.Errorf(expectedErr)
	suite.mockDb.ExpectPrepare(queryFixture).
		WillReturnError(mockErr)

	// Exercise SUT
	actual, err := suite.sut.SingleQueryForID(suite.db, queryFixture, arg1Fixture, arg2Fixture)

	// Verify results
	suite.Equal(entity.InvalidID, actual)
	suite.EqualError(err, expectedErr)
}

func (suite *HelperServiceTestSuite) TestSingleQueryForID_WhenScanFails_ShouldFail() {
	// Setup fixture
	queryFixture := "some.query"
	arg1Fixture := "arg.1"
	arg2Fixture := 2

	// Setup expectations
	expectedErr := "mock.error"

	// Setup mocks
	mockErr := fmt.Errorf(expectedErr)
	suite.mockDb.ExpectPrepare(queryFixture).
		ExpectQuery().
		WithArgs(arg1Fixture, arg2Fixture).
		WillReturnError(mockErr)

	// Exercise SUT
	actual, err := suite.sut.SingleQueryForID(suite.db, queryFixture, arg1Fixture, arg2Fixture)

	// Verify results
	suite.Equal(entity.InvalidID, actual)
	suite.EqualError(err, expectedErr)
}

func (suite *HelperServiceTestSuite) TestSingleQueryForID_WhenScanPasses_ShouldReturnID() {
	// Setup fixture
	queryFixture := "some.query"
	arg1Fixture := "arg.1"
	arg2Fixture := 2

	// Setup mocks
	mockRows := suite.mockDb.NewRows([]string{"id"}).
		FromCSVString("101")
	suite.mockDb.ExpectPrepare(queryFixture).
		ExpectQuery().
		WithArgs(arg1Fixture, arg2Fixture).
		WillReturnRows(mockRows)

	// Exercise SUT
	actual, err := suite.sut.SingleQueryForID(suite.db, queryFixture, arg1Fixture, arg2Fixture)

	// Verify results
	suite.NoError(err)
	suite.Equal(entity.ID(101), actual)
}

type mockResult struct {
	mock.Mock
	Data string
}

// LastInsertId is for mocking
func (m *mockResult) LastInsertId() (int64, error) {
	args := m.Called()
	return args.Get(0).(int64), args.Error(1)
}

// RowsAffected is for mocking
func (m *mockResult) RowsAffected() (int64, error) {
	args := m.Called()
	return args.Get(0).(int64), args.Error(1)
}
