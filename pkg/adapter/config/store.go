package config

import (
	"fmt"

	goConfig "github.com/liampulles/go-config"
)

// Store encapsulates configuration properties
// to be injected
type Store interface {
	GetPort() int
	GetMigrationSource() string
	GetDbUser() string
	GetDbPassword() string
	GetDbHost() string
	GetDbPort() int
	GetDbName() string
}

// StoreImpl implements store
type StoreImpl struct {
	port            int
	dbDriver        string
	migrationSource string
	dbUser          string
	dbPassword      string
	dbHost          string
	dbPort          int
	dbName          string
}

// Check we implement the interface
var _ Store = &StoreImpl{}

// NewStoreImpl is a constructor
func NewStoreImpl(source goConfig.Source) (*StoreImpl, error) {
	typedSource := goConfig.NewTypedSource(source)
	// Set defaults
	store := &StoreImpl{
		port:            8080,
		migrationSource: "file://migrations",
		dbUser:          "matchvid",
		dbPassword:      "password",
		dbHost:          "localhost",
		dbPort:          5432,
		dbName:          "matchvid",
	}

	// Read in from source
	if err := goConfig.LoadProperties(typedSource,
		goConfig.IntProp("PORT", &store.port, false),
		goConfig.StrProp("DB_DRIVER", &store.dbDriver, false),
		goConfig.StrProp("MIGRATION_SOURCE", &store.migrationSource, false),
		goConfig.StrProp("DB_USER", &store.dbUser, false),
		goConfig.StrProp("DB_PASSWORD", &store.dbPassword, false),
		goConfig.StrProp("DB_HOST", &store.dbHost, false),
		goConfig.IntProp("DB_PORT", &store.dbPort, false),
		goConfig.StrProp("DB_NAME", &store.dbName, false),
	); err != nil {
		return nil, fmt.Errorf("could not fetch config: %w", err)
	}

	return store, nil
}

// GetPort returns the configured port for the server
func (s *StoreImpl) GetPort() int {
	return s.port
}

// GetMigrationSource returns the source for database migrations to run
func (s *StoreImpl) GetMigrationSource() string {
	return s.migrationSource
}

// GetDbUser returns the database user
func (s *StoreImpl) GetDbUser() string {
	return s.dbUser
}

// GetDbPassword returns the database password
func (s *StoreImpl) GetDbPassword() string {
	return s.dbPassword
}

// GetDbHost returns the database host
func (s *StoreImpl) GetDbHost() string {
	return s.dbHost
}

// GetDbPort returns the database port
func (s *StoreImpl) GetDbPort() int {
	return s.dbPort
}

// GetDbName returns the database name
func (s *StoreImpl) GetDbName() string {
	return s.dbName
}
