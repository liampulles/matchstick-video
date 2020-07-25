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
}

// DatabaseServiceImpl implements DatabaseService
type DatabaseServiceImpl struct {
	sqlDB *goSql.DB
}

var _ sql.DatabaseService = &DatabaseServiceImpl{}

// NewDatabaseServiceImpl is a constructor
func NewDatabaseServiceImpl(configStore config.Store) (*DatabaseServiceImpl, error) {
	driverSelection := configStore.GetDbDriver()
	provider, ok := dbProviders[driverSelection]
	if !ok {
		return nil, fmt.Errorf(
			"could not create database service - %s is not a valid db driver selection - valid options: %v",
			driverSelection, providerList(),
		)
	}

	db, err := provider.constructor(configStore)
	if err != nil {
		return nil, fmt.Errorf("could not create database service - could not init db: %w", err)
	}

	err = provider.migrator(configStore, db)
	if err != nil {
		return nil, fmt.Errorf("could not create database service - could not migrate db: %w", err)
	}

	return &DatabaseServiceImpl{
		sqlDB: db,
	}, nil
}

// Get implements DatabaseService
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
