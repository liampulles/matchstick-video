package sql

import (
	"context"
	"database/sql"
	goSql "database/sql"
	"fmt"

	"github.com/liampulles/matchstick-video/pkg/adapter/db"
	"github.com/liampulles/matchstick-video/pkg/domain/entity"
)

// Row encapsulates *goSql.Row for testing purposes
type Row interface {
	Scan(dest ...interface{}) error
}

// Rows encapsulates *goSql.Rows for testing purposes
type Rows interface {
	Row
	Close() error
	Err() error
	Next() bool
	NextResultSet() bool
}

// ScanFunc scans a row and returns any errors
type ScanFunc func(row Row) error

// ExecForSingleItem will perform exec type SQL and verify a single row
// is affected.
var ExecForSingleItem = func(d *goSql.DB, query string, args ...interface{}) error {
	// Run exec to get rows affected
	rows, err := execForRowsAffected(d, query, args...)
	if err != nil {
		return fmt.Errorf("cannot execute exec - db exec error: %w", err)
	}

	// Verify rows affected is 1
	if rows == 0 {
		return db.NewNotFoundError("inventory item")
	}
	if rows != 1 {
		return fmt.Errorf("exec error: expected 1 entity to be affected, but was: %d", rows)
	}
	return nil
}

// SingleRowQuery will run a query type SQL which gives a single Row
var SingleRowQuery = func(goDB *goSql.DB, query string, scanFunc ScanFunc, _type string, args ...interface{}) error {
	// Prepare the query
	ctx := context.TODO()
	stmt, err := goDB.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("cannot execute query - db prepare error: %w", err)
	}

	// Run the query to get row
	row := stmt.QueryRowContext(ctx, args...)

	// Scan the row
	if err = scanFunc(row); err != nil {
		err = db.FromDBRowScan(err, _type)
		return fmt.Errorf("cannot execute query - db scan error: %w", err)
	}
	return nil
}

// ManyRowsQuery will run a query type SQL which gives many rows
var ManyRowsQuery = func(goDB *goSql.DB, query string, scanFunc ScanFunc, _type string, args ...interface{}) error {
	// Prepare the query
	ctx := context.TODO()
	stmt, err := goDB.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("cannot execute query - db prepare error: %w", err)
	}

	// Run the query to get rows
	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		return fmt.Errorf("cannot execute query - db context error: %w", err)
	}

	// Extract data from the row
	for rows.Next() {
		err := scanFunc(rows)
		if err != nil {
			err = db.FromDBRowScan(err, _type)
			return fmt.Errorf("cannot execute query - db scan error: %w", err)
		}
	}
	if err = rows.Err(); err != nil {
		return fmt.Errorf("cannot execute query - db iteration error: %w", err)
	}
	return nil
}

// SingleQueryForID will run SQL which returns an id, and return the entity form of
// the id
var SingleQueryForID = func(db *goSql.DB, query string, _type string, args ...interface{}) (entity.ID, error) {
	var id entity.ID

	// Map the ID, if we can
	err := SingleRowQuery(db, query, func(row Row) error {
		return row.Scan(&id)
	}, _type, args...)

	if err != nil {
		return entity.InvalidID, err
	}

	return id, nil
}

func execForRowsAffected(db *goSql.DB, query string, args ...interface{}) (int64, error) {
	// Perform the exec
	res, err := exec(db, query, args...)
	if err != nil {
		return -1, err
	}

	// Return the number of rows affected
	return res.RowsAffected()
}

func exec(db *goSql.DB, query string, args ...interface{}) (sql.Result, error) {
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
