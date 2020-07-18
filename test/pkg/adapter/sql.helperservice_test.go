package adapter_test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"github.com/liampulles/matchstick-video/pkg/adapter"
	"github.com/liampulles/matchstick-video/pkg/domain/entity"
)

type SQLDbHelperServiceTestSuite struct {
	suite.Suite
	db     *sql.DB
	mockDb sqlmock.Sqlmock
	sut    *adapter.SQLDbHelperServiceImpl
}

func TestSQLDbHelperServiceTestSuite(t *testing.T) {
	suite.Run(t, new(SQLDbHelperServiceTestSuite))
}

func (suite *SQLDbHelperServiceTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}
	suite.db = db
	suite.mockDb = mock
	suite.sut = adapter.NewSQLDbHelperServiceImpl()
}

func (suite *SQLDbHelperServiceTestSuite) TestExecForError_WhenPrepareContextFails_ShouldFail() {
	// Setup fixture
	queryFixture := "some.query"
	arg1Fixture := "arg.1"
	arg2Fixture := 2

	// Setup expectations
	expectedErr := "cannot execute exec - db exec error: mock.error"

	// Setup mocks
	mockErr := fmt.Errorf("mock.error")
	suite.mockDb.ExpectPrepare(queryFixture).
		WillReturnError(mockErr)

	// Exercise SUT
	err := suite.sut.ExecForError(suite.db, queryFixture, arg1Fixture, arg2Fixture)

	// Verify results
	suite.EqualError(err, expectedErr)
}

func (suite *SQLDbHelperServiceTestSuite) TestExecForError_WhenExecContextFails_ShouldFail() {
	// Setup fixture
	queryFixture := "some.query"
	arg1Fixture := "arg.1"
	arg2Fixture := 2

	// Setup expectations
	expectedErr := "cannot execute exec - db exec error: mock.error"

	// Setup mocks
	mockErr := fmt.Errorf("mock.error")
	suite.mockDb.ExpectPrepare(queryFixture).
		ExpectExec().
		WithArgs(arg1Fixture, arg2Fixture).
		WillReturnError(mockErr)

	// Exercise SUT
	err := suite.sut.ExecForError(suite.db, queryFixture, arg1Fixture, arg2Fixture)

	// Verify results
	suite.EqualError(err, expectedErr)
}

func (suite *SQLDbHelperServiceTestSuite) TestExecForError_WhenExecContextPasses_ShouldReturnAsExpected() {
	// Setup fixture
	queryFixture := "some.query"
	arg1Fixture := "arg.1"
	arg2Fixture := 2

	// Setup mocks
	suite.mockDb.ExpectPrepare(queryFixture).
		ExpectExec().
		WithArgs(arg1Fixture, arg2Fixture).
		WillReturnResult(&mockResult{})

	// Exercise SUT
	err := suite.sut.ExecForError(suite.db, queryFixture, arg1Fixture, arg2Fixture)

	// Verify results
	suite.NoError(err)
}

func (suite *SQLDbHelperServiceTestSuite) TestExecForID_WhenPrepareContextFails_ShouldFail() {
	// Setup fixture
	queryFixture := "some.query"
	arg1Fixture := "arg.1"
	arg2Fixture := 2

	// Setup expectations
	expectedErr := "cannot execute exec - db exec error: mock.error"

	// Setup mocks
	mockErr := fmt.Errorf("mock.error")
	suite.mockDb.ExpectPrepare(queryFixture).
		WillReturnError(mockErr)

	// Exercise SUT
	actual, err := suite.sut.ExecForID(suite.db, queryFixture, arg1Fixture, arg2Fixture)

	// Verify results
	suite.Equal(entity.InvalidID, actual)
	suite.EqualError(err, expectedErr)
}

func (suite *SQLDbHelperServiceTestSuite) TestExecForID_WhenExecContextFails_ShouldFail() {
	// Setup fixture
	queryFixture := "some.query"
	arg1Fixture := "arg.1"
	arg2Fixture := 2

	// Setup expectations
	expectedErr := "cannot execute exec - db exec error: mock.error"

	// Setup mocks
	mockErr := fmt.Errorf("mock.error")
	suite.mockDb.ExpectPrepare(queryFixture).
		ExpectExec().
		WithArgs(arg1Fixture, arg2Fixture).
		WillReturnError(mockErr)

	// Exercise SUT
	actual, err := suite.sut.ExecForID(suite.db, queryFixture, arg1Fixture, arg2Fixture)

	// Verify results
	suite.Equal(entity.InvalidID, actual)
	suite.EqualError(err, expectedErr)
}

func (suite *SQLDbHelperServiceTestSuite) TestExecForID_WhenLastInsertIdFails_ShouldFail() {
	// Setup fixture
	queryFixture := "some.query"
	arg1Fixture := "arg.1"
	arg2Fixture := 2

	// Setup expectations
	expectedErr := "cannot execute exec - result id error: mock.error"

	// Setup mocks
	mockResult := &mockResult{}
	mockErr := fmt.Errorf("mock.error")
	suite.mockDb.ExpectPrepare(queryFixture).
		ExpectExec().
		WithArgs(arg1Fixture, arg2Fixture).
		WillReturnResult(mockResult)
	mockResult.On("LastInsertId").Return(int64(-1), mockErr)

	// Exercise SUT
	actual, err := suite.sut.ExecForID(suite.db, queryFixture, arg1Fixture, arg2Fixture)

	// Verify results
	suite.Equal(entity.InvalidID, actual)
	suite.EqualError(err, expectedErr)
}

func (suite *SQLDbHelperServiceTestSuite) TestExecForID_WhenLastInsertIdPasses_ShouldReturnAsExpected() {
	// Setup fixture
	queryFixture := "some.query"
	arg1Fixture := "arg.1"
	arg2Fixture := 2

	// Setup expectations
	expectedId := entity.ID(101)

	// Setup mocks
	mockResult := &mockResult{}
	suite.mockDb.ExpectPrepare(queryFixture).
		ExpectExec().
		WithArgs(arg1Fixture, arg2Fixture).
		WillReturnResult(mockResult)
	mockResult.On("LastInsertId").Return(int64(101), nil)

	// Exercise SUT
	actual, err := suite.sut.ExecForID(suite.db, queryFixture, arg1Fixture, arg2Fixture)

	// Verify results
	suite.NoError(err)
	suite.Equal(expectedId, actual)
}

func (suite *SQLDbHelperServiceTestSuite) TestSingleRowQuery_WhenPrepareContextFails_ShouldFail() {
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

func (suite *SQLDbHelperServiceTestSuite) TestSingleRowQuery_WhenQueryRowContextPasses_ShouldReturnAsExpected() {
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
