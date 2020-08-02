package sql

import (
	goSql "database/sql"

	"github.com/stretchr/testify/mock"

	"github.com/liampulles/matchstick-video/pkg/adapter/db/sql"
	"github.com/liampulles/matchstick-video/pkg/domain/entity"
)

// MockHelperService is for mocking
type MockHelperService struct {
	mock.Mock
}

var _ sql.HelperService = &MockHelperService{}

// ExecForSingleItem is for mocking
func (s *MockHelperService) ExecForSingleItem(db *goSql.DB, query string, args ...interface{}) error {
	allArgs := make([]interface{}, 0)
	allArgs = append(allArgs, db, query)
	allArgs = append(allArgs, args...)
	a := s.Called(allArgs...)
	return a.Error(0)
}

// SingleRowQuery is for mocking
func (s *MockHelperService) SingleRowQuery(db *goSql.DB, query string, scanFunc sql.ScanFunc, _type string, args ...interface{}) error {
	allArgs := make([]interface{}, 0)
	allArgs = append(allArgs, db, query, scanFunc, _type)
	allArgs = append(allArgs, args...)
	a := s.Called(allArgs...)
	return a.Error(0)
}

// ManyRowsQuery is for mocking
func (s *MockHelperService) ManyRowsQuery(db *goSql.DB, query string, scanFunc sql.ScanFunc, _type string, args ...interface{}) error {
	allArgs := make([]interface{}, 0)
	allArgs = append(allArgs, db, query, scanFunc, _type)
	allArgs = append(allArgs, args...)
	a := s.Called(allArgs...)
	return a.Error(0)
}

// SingleQueryForID is for mocking
func (s *MockHelperService) SingleQueryForID(db *goSql.DB, query string, _type string, args ...interface{}) (entity.ID, error) {
	allArgs := make([]interface{}, 0)
	allArgs = append(allArgs, db, query, _type)
	allArgs = append(allArgs, args...)
	a := s.Called(allArgs...)
	return a.Get(0).(entity.ID), a.Error(1)
}

func safeArgsGetRow(args mock.Arguments, idx int) sql.Row {
	if val, ok := args.Get(idx).(sql.Row); ok {
		return val
	}
	return nil
}

func safeArgsGetRows(args mock.Arguments, idx int) sql.Rows {
	if val, ok := args.Get(idx).(sql.Rows); ok {
		return val
	}
	return nil
}

// RowMock is for mocking
type RowMock struct {
	mock.Mock
}

var _ sql.Row = &RowMock{}

// Scan is for mocking
func (s *RowMock) Scan(dest ...interface{}) error {
	args := s.Called(dest...)
	return args.Error(0)
}

// RowsMock is for mocking
type RowsMock struct {
	mock.Mock
}

var _ sql.Rows = &RowsMock{}

// Scan is for mocking
func (s *RowsMock) Scan(dest ...interface{}) error {
	args := s.Called(dest...)
	return args.Error(0)
}

// Close is for mocking
func (s *RowsMock) Close() error {
	args := s.Called()
	return args.Error(0)
}

// Err is for mocking
func (s *RowsMock) Err() error {
	args := s.Called()
	return args.Error(0)
}

// Next is for mocking
func (s *RowsMock) Next() bool {
	args := s.Called()
	return args.Bool(0)
}

// NextResultSet is for mocking
func (s *RowsMock) NextResultSet() bool {
	args := s.Called()
	return args.Bool(0)
}
