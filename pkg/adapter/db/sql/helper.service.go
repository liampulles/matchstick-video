package sql

import (
	"context"
	"database/sql"
	goSql "database/sql"

	"github.com/liampulles/matchstick-video/pkg/domain/entity"
)

// Row encapsulates *goSql.Row for testing purposes
// TODO: Own types should go in their own file, or shared file if simple.
type Row interface {
	Scan(dest ...interface{}) error
}

// HelperService encapsulates some common methods on sql.DB.
type HelperService interface {
	ExecForRowsAffected(db *goSql.DB, query string, args ...interface{}) (int64, error)
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

// ExecForRowsAffected will perform exec type SQL and return the number of rows
// affected
func (s *HelperServiceImpl) ExecForRowsAffected(db *goSql.DB, query string, args ...interface{}) (int64, error) {
	// Perform the exec
	res, err := s.exec(db, query, args...)
	if err != nil {
		return -1, err
	}

	// Return the number of rows affected
	return res.RowsAffected()
}

// ExecForID will perform exec type SQL and return the last insert id.
func (s *HelperServiceImpl) ExecForID(db *goSql.DB, query string, args ...interface{}) (entity.ID, error) {
	// Perform the exec
	res, err := s.exec(db, query, args...)
	if err != nil {
		return entity.InvalidID, err
	}

	// Return the last insert id
	id, err := res.LastInsertId()
	if err != nil {
		return entity.InvalidID, err
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
