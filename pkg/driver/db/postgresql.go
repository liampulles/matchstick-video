package db

import (
	goSql "database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"

	// Import file source in the background
	_ "github.com/golang-migrate/migrate/v4/source/file"
	// Import the PostgreSQL driver in the background
	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/liampulles/matchstick-video/pkg/adapter/config"
)

// NewPostgresDB is a constructor
func NewPostgresDB() (*goSql.DB, error) {
	// Bring up DB
	db, err := newPostgreSQLDB()
	if err != nil {
		return nil, fmt.Errorf("could not create database service - could not init db: %w", err)
	}

	// Perform migrations
	err = migratePostgreSQLDB(db)
	if err != nil {
		return nil, fmt.Errorf("could not create database service - could not migrate db: %w", err)
	}

	// Return ready-to-use DB
	return db, nil
}

func newPostgreSQLDB() (*goSql.DB, error) {
	connStr := getConnectionString()
	db, err := goSql.Open("pgx", connStr)
	if err != nil {
		return nil, fmt.Errorf("could not create postgres db - open error: %w", err)
	}

	return db, nil
}

func migratePostgreSQLDB(sqlDB *goSql.DB) error {
	// Get migration driver
	driver, err := postgres.WithInstance(sqlDB, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("could not migrate postgres db - driver error: %w", err)
	}

	// Get migration instance
	source := config.GetMigrationSource()
	m, err := migrate.NewWithDatabaseInstance(source, "postgres", driver)
	if err != nil {
		return fmt.Errorf("could not migrate postgres db - migrate init error: %w", err)
	}

	// Run migrations
	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("could not migrate postgres db - up error: %w", err)
	}

	// Display post-migration status
	v, dirty, err := m.Version()
	if err != nil {
		return fmt.Errorf("could not migrate postgres db - version error: %w", err)
	}
	fmt.Printf("DB Migration Version: %d. Dirty: %v\n", v, dirty)

	return nil
}

func getConnectionString() string {
	host := config.GetDbHost()
	port := config.GetDbPort()
	user := config.GetDbUser()
	pass := config.GetDbPassword()
	name := config.GetDbName()

	return fmt.Sprintf("user=%s password=%s host=%s port=%d database=%s sslmode=disable",
		user, pass, host, port, name,
	)
}
