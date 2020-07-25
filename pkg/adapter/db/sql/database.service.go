package sql

import (
	goSql "database/sql"
)

// DatabaseService provides a ready-to-use SQL db.
type DatabaseService interface {
	Get() *goSql.DB
}
