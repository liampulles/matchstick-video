package config

import (
	goConfig "github.com/liampulles/go-config"
	"github.com/rs/zerolog/log"
)

// GetPort returns the configured port for the server
func GetPort() int {
	return lazyGet().port
}

// GetMigrationSource returns the source for database migrations to run
func GetMigrationSource() string {
	return lazyGet().migrationSource
}

// GetDbUser returns the database user
func GetDbUser() string {
	return lazyGet().dbUser
}

// GetDbPassword returns the database password
func GetDbPassword() string {
	return lazyGet().dbPassword
}

// GetDbHost returns the database host
func GetDbHost() string {
	return lazyGet().dbHost
}

// GetDbPort returns the database port
func GetDbPort() int {
	return lazyGet().dbPort
}

// GetDbName returns the database name
func GetDbName() string {
	return lazyGet().dbName
}

// --- Backing ---

// Config state
type state struct {
	port            int
	migrationSource string
	dbUser          string
	dbPassword      string
	dbHost          string
	dbPort          int
	dbName          string
}

var (
	defaultSource = goConfig.NewEnvSource()
	loaded        = false
	loadedState   state
)

func lazyGet() state {
	if !loaded {
		if err := Load(defaultSource); err != nil {
			log.Err(err).Msg("could not load config")
			panic(err)
		}
	}

	return loadedState
}

// Load the global config.
func Load(source goConfig.Source) error {
	typedSource := goConfig.NewTypedSource(source)
	// Set defaults
	s := state{
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
		goConfig.IntProp("PORT", &s.port, false),
		goConfig.StrProp("MIGRATION_SOURCE", &s.migrationSource, false),
		goConfig.StrProp("DB_USER", &s.dbUser, false),
		goConfig.StrProp("DB_PASSWORD", &s.dbPassword, false),
		goConfig.StrProp("DB_HOST", &s.dbHost, false),
		goConfig.IntProp("DB_PORT", &s.dbPort, false),
		goConfig.StrProp("DB_NAME", &s.dbName, false),
	); err != nil {
		return err
	}

	loadedState, loaded = s, true
	return nil
}
