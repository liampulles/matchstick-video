package db

import (
	goSql "database/sql"
	"fmt"

	"github.com/liampulles/matchstick-video/pkg/adapter/config"
	"github.com/liampulles/matchstick-video/pkg/adapter/db/sql"
)

type dbProvider struct {
	constructor func(config.Store) (*goSql.DB, error)
	migrator    func(config.Store, *goSql.DB) error
}

var dbProviders map[string]dbProvider = map[string]dbProvider{
	"sqlite3": {
		constructor: newTempSQLite3DB,
		migrator:    migrateSQLite3DB,
	},
	"postgres": {
		constructor: newPostgreSQLDB,
		migrator:    migratePostgreSQLDB,
	},
}

// DatabaseServiceImpl implements DatabaseService
type DatabaseServiceImpl struct {
	sqlDB *goSql.DB
}

var _ sql.DatabaseService = &DatabaseServiceImpl{}

// NewDatabaseServiceImpl is a constructor
func NewDatabaseServiceImpl(configStore config.Store) (*DatabaseServiceImpl, error) {
	// Select configured provider
	driverSelection := configStore.GetDbDriver()
	provider, ok := dbProviders[driverSelection]
	if !ok {
		return nil, fmt.Errorf(
			"could not create database service - %s is not a valid db driver selection - valid options: %v",
			driverSelection, providerList(),
		)
	}

	// Bring up DB
	db, err := provider.constructor(configStore)
	if err != nil {
		return nil, fmt.Errorf("could not create database service - could not init db: %w", err)
	}

	// Perform migrations
	err = provider.migrator(configStore, db)
	if err != nil {
		return nil, fmt.Errorf("could not create database service - could not migrate db: %w", err)
	}

	// Return ready-to-use DB
	return &DatabaseServiceImpl{
		sqlDB: db,
	}, nil
}

// Get retusn a pre-configured database, which is ready to use.
func (d *DatabaseServiceImpl) Get() *goSql.DB {
	return d.sqlDB
}

func providerList() []string {
	var providers []string
	for k := range dbProviders {
		providers = append(providers, k)
	}
	return providers
}
