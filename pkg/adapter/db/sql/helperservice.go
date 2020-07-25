package sql

import (
	"context"
	"database/sql"
	goSql "database/sql"
	"fmt"

	"github.com/liampulles/matchstick-video/pkg/domain/entity"
)

// Row encapsulates *goSql.Row for testing purposes
type Row interface {
	Scan(dest ...interface{}) error
}

// HelperService encapsulates some common methods on sql.DB.
type HelperService interface {
	ExecForError(db *goSql.DB, query string, args ...interface{}) error
	ExecForID(db *goSql.DB, query string, args ...interface{}) (entity.ID, error)
	SingleRowQuery(db *goSql.DB, query string, args ...interface{}) (Row, error)
}

// HelperServiceImpl implements the HelperService interface
type HelperServiceImpl struct{}

var _ HelperService = &HelperServiceImpl{}

// NewHelperServiceImpl is a constructor
func NewHelperServiceImpl() *HelperServiceImpl {
	return &HelperServiceImpl{}
}

// ExecForError will perform exec type SQL and format an error if needed.
func (s *HelperServiceImpl) ExecForError(db *goSql.DB, query string, args ...interface{}) error {
	_, err := s.exec(db, query, args...)
	if err != nil {
		return fmt.Errorf("cannot execute exec - db exec error: %w", err)
	}
	return nil
}

// ExecForID will perform exec type SQL and return the last insert id.
func (s *HelperServiceImpl) ExecForID(db *goSql.DB, query string, args ...interface{}) (entity.ID, error) {
	// Perform the exec
	res, err := s.exec(db, query, args...)
	if err != nil {
		return entity.InvalidID, fmt.Errorf("cannot execute exec - db exec error: %w", err)
	}

	// Return the last insert id
	id, err := res.LastInsertId()
	if err != nil {
		return entity.InvalidID, fmt.Errorf("cannot execute exec - result id error: %w", err)
	}
	return entity.ID(id), nil
}

// SingleRowQuery will run a query type SQL which gives a single Row
func (s *HelperServiceImpl) SingleRowQuery(db *goSql.DB, query string, args ...interface{}) (Row, error) {
	// Prepare the query
	ctx := context.TODO()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}

	// Run the query to get row
	return stmt.QueryRowContext(ctx, args...), nil
}

func (s *HelperServiceImpl) exec(db *goSql.DB, query string, args ...interface{}) (sql.Result, error) {
	// Prepare the exec
	ctx := context.TODO()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}

	// Run the exec
	res, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		return nil, err
	}
	return res, nil
}
