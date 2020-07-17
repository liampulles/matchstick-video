package adapter

import (
	"database/sql"

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
