package config

import (
	"fmt"

	goConfig "github.com/liampulles/go-config"
)

// Store encapsulates configuration properties
// to be injected
type Store interface {
	GetPort() int
	GetDbDriver() string
	GetMigrationSource() string
}

// StoreImpl implements store
type StoreImpl struct {
	port            int
	dbDriver        string
	migrationSource string
}

// Check we implement the interface
var _ Store = &StoreImpl{}

// NewStoreImpl is a constructor
func NewStoreImpl(source goConfig.Source) (*StoreImpl, error) {
	typedSource := goConfig.NewTypedSource(source)
	// Set defaults
	store := &StoreImpl{
		port:            8080,
		dbDriver:        "sqlite3",
		migrationSource: "file://migrations",
	}

	// Read in from source
	if err := goConfig.LoadProperties(typedSource,
		goConfig.IntProp("PORT", &store.port, false),
		goConfig.StrProp("DB_DRIVER", &store.dbDriver, false),
		goConfig.StrProp("MIGRATION_SOURCE", &store.migrationSource, false),
	); err != nil {
		return nil, fmt.Errorf("could not fetch config: %w", err)
	}

	return store, nil
}

// GetPort returns the configured port for the server
func (s *StoreImpl) GetPort() int {
	return s.port
}

// GetDbDriver returns the set database driver
func (s *StoreImpl) GetDbDriver() string {
	return s.dbDriver
}

// GetMigrationSource returns the source for database migrations to run
func (s *StoreImpl) GetMigrationSource() string {
	return s.migrationSource
}
