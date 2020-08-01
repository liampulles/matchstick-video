package db

import (
	goSql "database/sql"
	"fmt"

	"github.com/liampulles/matchstick-video/pkg/adapter/config"
	"github.com/liampulles/matchstick-video/pkg/adapter/db/sql"
)

// DatabaseServiceImpl implements DatabaseService
type DatabaseServiceImpl struct {
	sqlDB *goSql.DB
}

var _ sql.DatabaseService = &DatabaseServiceImpl{}

// NewDatabaseServiceImpl is a constructor
func NewDatabaseServiceImpl(configStore config.Store) (*DatabaseServiceImpl, error) {
	// Bring up DB
	db, err := newPostgreSQLDB(configStore)
	if err != nil {
		return nil, fmt.Errorf("could not create database service - could not init db: %w", err)
	}

	// Perform migrations
	err = migratePostgreSQLDB(configStore, db)
	if err != nil {
		return nil, fmt.Errorf("could not create database service - could not migrate db: %w", err)
	}

	// Return ready-to-use DB
	return &DatabaseServiceImpl{
		sqlDB: db,
	}, nil
}

// Get returns a pre-configured database, which is ready to use.
func (d *DatabaseServiceImpl) Get() *goSql.DB {
	return d.sqlDB
}
