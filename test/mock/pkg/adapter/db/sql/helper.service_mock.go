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

// ExecForRowsAffected is for mocking
func (s *MockHelperService) ExecForRowsAffected(db *goSql.DB, query string, args ...interface{}) (int64, error) {
	allArgs := make([]interface{}, 0)
	allArgs = append(allArgs, db, query)
	allArgs = append(allArgs, args...)
	a := s.Called(allArgs...)
	return a.Get(0).(int64), a.Error(1)
}

// SingleRowQuery is for mocking
func (s *MockHelperService) SingleRowQuery(db *goSql.DB, query string, args ...interface{}) (sql.Row, error) {
	allArgs := make([]interface{}, 0)
	allArgs = append(allArgs, db, query)
	allArgs = append(allArgs, args...)
	a := s.Called(allArgs...)
	return safeArgsGetRow(a, 0), a.Error(1)
}

// SingleQueryForID is for mocking
func (s *MockHelperService) SingleQueryForID(db *goSql.DB, query string, args ...interface{}) (entity.ID, error) {
	allArgs := make([]interface{}, 0)
	allArgs = append(allArgs, db, query)
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
