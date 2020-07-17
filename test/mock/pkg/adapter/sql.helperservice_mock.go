package adapter

import (
	"database/sql"

	"github.com/stretchr/testify/mock"

	"github.com/liampulles/matchstick-video/pkg/adapter"
	"github.com/liampulles/matchstick-video/pkg/domain/entity"
)

// SQLDbHelperServiceMock is for mocking
type SQLDbHelperServiceMock struct {
	mock.Mock
}

var _ adapter.SQLDbHelperService = &SQLDbHelperServiceMock{}

// ExecForError is for mocking
func (s *SQLDbHelperServiceMock) ExecForError(db *sql.DB, query string, args ...interface{}) error {
	allArgs := make([]interface{}, 0)
	allArgs = append(allArgs, db, query)
	allArgs = append(allArgs, args...)
	a := s.Called(allArgs...)
	return a.Error(0)
}

// ExecForID is for mocking
func (s *SQLDbHelperServiceMock) ExecForID(db *sql.DB, query string, args ...interface{}) (entity.ID, error) {
	allArgs := make([]interface{}, 0)
	allArgs = append(allArgs, db, query)
	allArgs = append(allArgs, args...)
	a := s.Called(allArgs...)
	return a.Get(0).(entity.ID), a.Error(1)
}

// SingleRowQuery is for mocking
func (s *SQLDbHelperServiceMock) SingleRowQuery(db *sql.DB, query string, args ...interface{}) (adapter.SQLRow, error) {
	allArgs := make([]interface{}, 0)
	allArgs = append(allArgs, db, query)
	allArgs = append(allArgs, args...)
	a := s.Called(allArgs...)
	return safeArgsGetRow(a, 0), a.Error(1)
}

func safeArgsGetRow(args mock.Arguments, idx int) adapter.SQLRow {
	if val, ok := args.Get(idx).(adapter.SQLRow); ok {
		return val
	}
	return nil
}

// SQLRowMock is for mocking
type SQLRowMock struct {
	mock.Mock
}

var _ adapter.SQLRow = &SQLRowMock{}

// Scan is for mocking
func (s *SQLRowMock) Scan(dest ...interface{}) error {
	args := s.Called(dest...)
	return args.Error(0)
}
