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

// ExecForError implements the HelperService interface
func (s *HelperServiceImpl) ExecForError(db *goSql.DB, query string, args ...interface{}) error {
	_, err := s.exec(db, query, args...)
	if err != nil {
		return fmt.Errorf("cannot execute exec - db exec error: %w", err)
	}
	return nil
}

// ExecForID implements the HelperService interface
func (s *HelperServiceImpl) ExecForID(db *goSql.DB, query string, args ...interface{}) (entity.ID, error) {
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

// SingleRowQuery implements the HelperService interface
func (s *HelperServiceImpl) SingleRowQuery(db *goSql.DB, query string, args ...interface{}) (Row, error) {
	ctx := context.TODO()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}

	return stmt.QueryRowContext(ctx, args...), nil
}

func (s *HelperServiceImpl) exec(db *goSql.DB, query string, args ...interface{}) (sql.Result, error) {
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
