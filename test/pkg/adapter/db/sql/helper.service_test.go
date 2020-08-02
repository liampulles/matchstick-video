package sql_test

import (
	goSql "database/sql"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"github.com/liampulles/matchstick-video/test/mock/pkg/adapter/db"

	"github.com/liampulles/matchstick-video/pkg/adapter/db/sql"
	"github.com/liampulles/matchstick-video/pkg/domain/entity"
)

type HelperServiceTestSuite struct {
	suite.Suite
	db              *goSql.DB
	mockDb          sqlmock.Sqlmock
	mockErrorParser *db.MockErrorParser
	sut             *sql.HelperServiceImpl
}

func TestHelperServiceTestSuite(t *testing.T) {
	suite.Run(t, new(HelperServiceTestSuite))
}

func (suite *HelperServiceTestSuite) SetupTest() {
	d, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}
	suite.db = d
	suite.mockDb = mock
	suite.mockErrorParser = &db.MockErrorParser{}
	suite.sut = sql.NewHelperServiceImpl(
		suite.mockErrorParser,
	)
}

func (suite *HelperServiceTestSuite) TestExecForSingleItem_WhenPrepareContextFails_ShouldFail() {
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
	err := suite.sut.ExecForSingleItem(suite.db, queryFixture, arg1Fixture, arg2Fixture)

	// Verify results
	suite.EqualError(err, expectedErr)
}

func (suite *HelperServiceTestSuite) TestExecForSingleItem_WhenExecContextFails_ShouldFail() {
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
	err := suite.sut.ExecForSingleItem(suite.db, queryFixture, arg1Fixture, arg2Fixture)

	// Verify results
	suite.EqualError(err, expectedErr)
}

func (suite *HelperServiceTestSuite) TestExecForSingleItem_WhenRowsAffectedFails_ShouldFail() {
	// Setup fixture
	queryFixture := "some.query"
	arg1Fixture := "arg.1"
	arg2Fixture := 2

	// Setup expectations
	expectedErr := "cannot execute exec - db exec error: mock.error"

	// Setup mocks
	mockResult := &mockResult{}
	mockErr := fmt.Errorf("mock.error")
	suite.mockDb.ExpectPrepare(queryFixture).
		ExpectExec().
		WithArgs(arg1Fixture, arg2Fixture).
		WillReturnResult(mockResult)
	mockResult.On("RowsAffected").Return(int64(-1), mockErr)

	// Exercise SUT
	err := suite.sut.ExecForSingleItem(suite.db, queryFixture, arg1Fixture, arg2Fixture)

	// Verify results
	suite.EqualError(err, expectedErr)
}

func (suite *HelperServiceTestSuite) TestExecForSingleItem_WhenRowsAffectedIsZero_ShouldFail() {
	// Setup fixture
	queryFixture := "some.query"
	arg1Fixture := "arg.1"
	arg2Fixture := 2

	// Setup expectations
	expectedErr := "entity not found: type=[inventory item]"

	// Setup mocks
	mockResult := &mockResult{}
	suite.mockDb.ExpectPrepare(queryFixture).
		ExpectExec().
		WithArgs(arg1Fixture, arg2Fixture).
		WillReturnResult(mockResult)
	mockResult.On("RowsAffected").Return(int64(0), nil)

	// Exercise SUT
	err := suite.sut.ExecForSingleItem(suite.db, queryFixture, arg1Fixture, arg2Fixture)

	// Verify results
	suite.EqualError(err, expectedErr)
}

func (suite *HelperServiceTestSuite) TestExecForSingleItem_WhenRowsAffectedIsMoreThanOne_ShouldFail() {
	// Setup fixture
	queryFixture := "some.query"
	arg1Fixture := "arg.1"
	arg2Fixture := 2

	// Setup expectations
	expectedErr := "exec error: expected 1 entity to be affected, but was: 2"

	// Setup mocks
	mockResult := &mockResult{}
	suite.mockDb.ExpectPrepare(queryFixture).
		ExpectExec().
		WithArgs(arg1Fixture, arg2Fixture).
		WillReturnResult(mockResult)
	mockResult.On("RowsAffected").Return(int64(2), nil)

	// Exercise SUT
	err := suite.sut.ExecForSingleItem(suite.db, queryFixture, arg1Fixture, arg2Fixture)

	// Verify results
	suite.EqualError(err, expectedErr)
}

func (suite *HelperServiceTestSuite) TestExecForSingleItem_WhenRowsAffectedIsOne_ShouldPass() {
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
	mockResult.On("RowsAffected").Return(int64(1), nil)

	// Exercise SUT
	err := suite.sut.ExecForSingleItem(suite.db, queryFixture, arg1Fixture, arg2Fixture)

	// Verify results
	suite.NoError(err)
}

func (suite *HelperServiceTestSuite) TestSingleRowQuery_WhenPrepareContextFails_ShouldFail() {
	// Setup fixture
	queryFixture := "some.query"
	arg1Fixture := "arg.1"
	arg2Fixture := 2
	passingFunc := func(row sql.Row) error {
		return nil
	}

	// Setup expectations
	expectedErr := "cannot execute query - db prepare error: mock.error"

	// Setup mocks
	mockErr := fmt.Errorf("mock.error")
	suite.mockDb.ExpectPrepare(queryFixture).
		WillReturnError(mockErr)

	// Exercise SUT
	err := suite.sut.SingleRowQuery(suite.db, queryFixture, passingFunc, "some.type", arg1Fixture, arg2Fixture)

	// Verify results
	suite.EqualError(err, expectedErr)
}

func (suite *HelperServiceTestSuite) TestSingleRowQuery_WhenScanFuncFails_ShouldFail() {
	// Setup fixture
	queryFixture := "some.query"
	arg1Fixture := "arg.1"
	arg2Fixture := 2

	// Setup expectations
	expectedErr := "cannot execute query - db scan error: mock.parsed.error"

	// Setup mocks
	mockErr := fmt.Errorf("mock.error")
	mockParsedErr := fmt.Errorf("mock.parsed.error")
	mockRows := suite.mockDb.NewRows([]string{"some", "columns"}).
		FromCSVString("with,data")
	failingFunc := func(row sql.Row) error {
		return mockErr
	}
	suite.mockDb.ExpectPrepare(queryFixture).
		ExpectQuery().
		WithArgs(arg1Fixture, arg2Fixture).
		WillReturnRows(mockRows)
	suite.mockErrorParser.On("FromDBRowScan", mockErr, "some.type").
		Return(mockParsedErr)

	// Exercise SUT
	err := suite.sut.SingleRowQuery(suite.db, queryFixture, failingFunc, "some.type", arg1Fixture, arg2Fixture)

	// Verify results
	suite.EqualError(err, expectedErr)
}

func (suite *HelperServiceTestSuite) TestSingleRowQuery_WhenScanFuncPasses_ShouldPass() {
	// Setup fixture
	queryFixture := "some.query"
	arg1Fixture := "arg.1"
	arg2Fixture := 2

	// Setup mocks
	mockRows := suite.mockDb.NewRows([]string{"some", "columns"}).
		FromCSVString("with,data")
	passingFunc := func(row sql.Row) error {
		return nil
	}
	suite.mockDb.ExpectPrepare(queryFixture).
		ExpectQuery().
		WithArgs(arg1Fixture, arg2Fixture).
		WillReturnRows(mockRows)

	// Exercise SUT
	err := suite.sut.SingleRowQuery(suite.db, queryFixture, passingFunc, "some.type", arg1Fixture, arg2Fixture)

	// Verify results
	suite.NoError(err)
}

func (suite *HelperServiceTestSuite) TestSingleQueryForID_WhenPrepareContextFails_ShouldFail() {
	// Setup fixture
	queryFixture := "some.query"
	arg1Fixture := "arg.1"
	arg2Fixture := 2

	// Setup expectations
	expectedErr := "cannot execute query - db prepare error: mock.error"

	// Setup mocks
	mockErr := fmt.Errorf("mock.error")
	suite.mockDb.ExpectPrepare(queryFixture).
		WillReturnError(mockErr)

	// Exercise SUT
	actual, err := suite.sut.SingleQueryForID(suite.db, queryFixture, "some.type", arg1Fixture, arg2Fixture)

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
	expectedErr := "cannot execute query - db scan error: mock.parsed.error"

	// Setup mocks
	mockErr := fmt.Errorf("mock.error")
	mockParsedErr := fmt.Errorf("mock.parsed.error")
	suite.mockDb.ExpectPrepare(queryFixture).
		ExpectQuery().
		WithArgs(arg1Fixture, arg2Fixture).
		WillReturnError(mockErr)
	suite.mockErrorParser.On("FromDBRowScan", mockErr, "some.type").
		Return(mockParsedErr)

	// Exercise SUT
	actual, err := suite.sut.SingleQueryForID(suite.db, queryFixture, "some.type", arg1Fixture, arg2Fixture)

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
	actual, err := suite.sut.SingleQueryForID(suite.db, queryFixture, "some.type", arg1Fixture, arg2Fixture)

	// Verify results
	suite.NoError(err)
	suite.Equal(entity.ID(101), actual)
}

func (suite *HelperServiceTestSuite) TestManyRowsQuery_WhenPrepareContextFails_ShouldFail() {
	// Setup fixture
	queryFixture := "some.query"
	arg1Fixture := "arg.1"
	arg2Fixture := 2
	passingFunc := func(row sql.Row) error {
		return nil
	}

	// Setup expectations
	expectedErr := "cannot execute query - db prepare error: mock.error"

	// Setup mocks
	mockErr := fmt.Errorf("mock.error")
	suite.mockDb.ExpectPrepare(queryFixture).
		WillReturnError(mockErr)

	// Exercise SUT
	err := suite.sut.ManyRowsQuery(suite.db, queryFixture, passingFunc, "some.type", arg1Fixture, arg2Fixture)

	// Verify results
	suite.EqualError(err, expectedErr)
}

func (suite *HelperServiceTestSuite) TestManyRowsQuery_WhenQueryRowContextFails_ShouldFail() {
	// Setup fixture
	queryFixture := "some.query"
	arg1Fixture := "arg.1"
	arg2Fixture := 2
	passingFunc := func(row sql.Row) error {
		return nil
	}

	// Setup expectations
	expectedErr := "cannot execute query - db context error: mock.error"

	// Setup mocks
	mockErr := fmt.Errorf("mock.error")
	suite.mockDb.ExpectPrepare(queryFixture).
		ExpectQuery().
		WithArgs(arg1Fixture, arg2Fixture).
		WillReturnError(mockErr)

	// Exercise SUT
	err := suite.sut.ManyRowsQuery(suite.db, queryFixture, passingFunc, "some.type", arg1Fixture, arg2Fixture)

	// Verify results
	suite.EqualError(err, expectedErr)
}

func (suite *HelperServiceTestSuite) TestManyRowsQuery_WhenScanFuncFails_ShouldFail() {
	// Setup fixture
	queryFixture := "some.query"
	arg1Fixture := "arg.1"
	arg2Fixture := 2

	// Setup expectations
	expectedErr := "cannot execute query - db scan error: mock.parsed.error"

	// Setup mocks
	mockErr := fmt.Errorf("mock.error")
	mockParsedErr := fmt.Errorf("mock.parsed.error")
	failingFunc := func(row sql.Row) error {
		return fmt.Errorf("mock.error")
	}
	mockRows := suite.mockDb.NewRows([]string{"something"}).
		FromCSVString("some.data")
	suite.mockDb.ExpectPrepare(queryFixture).
		ExpectQuery().
		WithArgs(arg1Fixture, arg2Fixture).
		WillReturnRows(mockRows)
	suite.mockErrorParser.On("FromDBRowScan", mockErr, "some.type").
		Return(mockParsedErr)

	// Exercise SUT
	err := suite.sut.ManyRowsQuery(suite.db, queryFixture, failingFunc, "some.type", arg1Fixture, arg2Fixture)

	// Verify results
	suite.EqualError(err, expectedErr)
}

func (suite *HelperServiceTestSuite) TestManyRowsQuery_WhenFirstRowScanFails_ShouldFail() {
	// Setup fixture
	queryFixture := "some.query"
	arg1Fixture := "arg.1"
	arg2Fixture := 2
	var actual []mockData
	scanFunc := func(row sql.Row) error {
		item := mockData{}
		err := row.Scan(&item.Data)
		if err == nil {
			actual = append(actual, item)
		}
		return err
	}

	// Setup expectations
	expectedErr := "cannot execute query - db iteration error: mock.error"

	// Setup mocks
	mockErr := fmt.Errorf("mock.error")
	mockRows := suite.mockDb.NewRows([]string{"something"}).
		FromCSVString("some.data").
		RowError(0, mockErr)
	suite.mockDb.ExpectPrepare(queryFixture).
		ExpectQuery().
		WithArgs(arg1Fixture, arg2Fixture).
		WillReturnRows(mockRows)

	// Exercise SUT
	err := suite.sut.ManyRowsQuery(suite.db, queryFixture, scanFunc, "some.type", arg1Fixture, arg2Fixture)

	// Verify results
	suite.EqualError(err, expectedErr)
}

func (suite *HelperServiceTestSuite) TestManyRowsQuery_WhenSecondRowScanFails_ShouldFail() {
	// Setup fixture
	queryFixture := "some.query"
	arg1Fixture := "arg.1"
	arg2Fixture := 2
	var actual []mockData
	scanFunc := func(row sql.Row) error {
		item := mockData{}
		err := row.Scan(&item.Data)
		if err == nil {
			actual = append(actual, item)
		}
		return err
	}

	// Setup expectations
	expectedErr := "cannot execute query - db iteration error: mock.error"

	// Setup mocks
	mockErr := fmt.Errorf("mock.error")
	mockRows := suite.mockDb.NewRows([]string{"something"}).
		FromCSVString("some.data.1").
		FromCSVString("some.data.2").
		RowError(1, mockErr)
	suite.mockDb.ExpectPrepare(queryFixture).
		ExpectQuery().
		WithArgs(arg1Fixture, arg2Fixture).
		WillReturnRows(mockRows)

	// Exercise SUT
	err := suite.sut.ManyRowsQuery(suite.db, queryFixture, scanFunc, "some.type", arg1Fixture, arg2Fixture)

	// Verify results
	suite.EqualError(err, expectedErr)
}

func (suite *HelperServiceTestSuite) TestManyRowsQuery_WhenIterationPasses_ShouldPass() {
	// Setup fixture
	queryFixture := "some.query"
	arg1Fixture := "arg.1"
	arg2Fixture := 2
	var actual []mockData
	scanFunc := func(row sql.Row) error {
		item := mockData{}
		err := row.Scan(&item.Data)
		if err == nil {
			actual = append(actual, item)
		}
		return err
	}

	// Setup expectations
	expected := []mockData{
		mockData{Data: "some.data.1"},
		mockData{Data: "some.data.2"},
	}

	// Setup mocks
	mockRows := suite.mockDb.NewRows([]string{"something"}).
		FromCSVString("some.data.1").
		FromCSVString("some.data.2")
	suite.mockDb.ExpectPrepare(queryFixture).
		ExpectQuery().
		WithArgs(arg1Fixture, arg2Fixture).
		WillReturnRows(mockRows)

	// Exercise SUT
	err := suite.sut.ManyRowsQuery(suite.db, queryFixture, scanFunc, "some.type", arg1Fixture, arg2Fixture)

	// Verify results
	suite.NoError(err)
	suite.Equal(expected, actual)
}

type mockResult struct {
	mock.Mock
	Data string
}

type mockData struct {
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
