package adapter

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/liampulles/matchstick-video/pkg/domain/entity"
)

// SQLRow encapsulates *sql.Row for testing purposes
type SQLRow interface {
	Scan(dest ...interface{}) error
}

// SQLDbHelperService encapsulates some common methods on sql.DB.
type SQLDbHelperService interface {
	ExecForError(db *sql.DB, query string, args ...interface{}) error
	ExecForID(db *sql.DB, query string, args ...interface{}) (entity.ID, error)
	SingleRowQuery(db *sql.DB, query string, args ...interface{}) (SQLRow, error)
}

// SQLDbHelperServiceImpl implements the SQLDbHelperService interface
type SQLDbHelperServiceImpl struct{}

var _ SQLDbHelperService = &SQLDbHelperServiceImpl{}

// NewSQLDbHelperServiceImpl is a constructor
func NewSQLDbHelperServiceImpl() *SQLDbHelperServiceImpl {
	return &SQLDbHelperServiceImpl{}
}

// ExecForError implements the SQLDbHelperService interface
func (s *SQLDbHelperServiceImpl) ExecForError(db *sql.DB, query string, args ...interface{}) error {
	_, err := s.exec(db, query, args...)
	if err != nil {
		return fmt.Errorf("cannot execute exec - db exec error: %w", err)
	}
	return nil
}

// ExecForID implements the SQLDbHelperService interface
func (s *SQLDbHelperServiceImpl) ExecForID(db *sql.DB, query string, args ...interface{}) (entity.ID, error) {
	res, err := s.exec(db, query, args...)
	if err != nil {
		return entity.InvalidID, fmt.Errorf("cannot execute exec - db exec error: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return entity.InvalidID, fmt.Errorf("cannot execute exec - result id error: %w", err)
	}
	return entity.ID(id), nil
}

// SingleRowQuery implements the SQLDbHelperService interface
func (s *SQLDbHelperServiceImpl) SingleRowQuery(db *sql.DB, query string, args ...interface{}) (SQLRow, error) {
	ctx := context.TODO()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}

	return stmt.QueryRowContext(ctx, args...), nil
}

func (s *SQLDbHelperServiceImpl) exec(db *sql.DB, query string, args ...interface{}) (sql.Result, error) {
	ctx := context.TODO()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}

	res, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		return nil, err
	}
	return res, nil
}
