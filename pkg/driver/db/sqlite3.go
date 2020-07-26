package db

import (
	"database/sql"
	"fmt"
	"math/rand"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"

	// Import file source in the background
	_ "github.com/golang-migrate/migrate/v4/source/file"
	// Import the SQLite3 driver in the background
	_ "github.com/mattn/go-sqlite3"

	"github.com/liampulles/matchstick-video/pkg/adapter/config"
)

func newTempSQLite3DB(cfg config.Store) (*sql.DB, error) {
	dbPath := tempDbPath()

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("could not create sqlite3 db - open error: %w", err)
	}

	return db, nil
}

func migrateSQLite3DB(cfg config.Store, sqlDB *sql.DB) error {
	// Get migration driver
	driver, err := sqlite3.WithInstance(sqlDB, &sqlite3.Config{})
	if err != nil {
		return fmt.Errorf("could not migrate sqlite3 db - driver error: %w", err)
	}

	// Get migration instance
	source := cfg.GetMigrationSource()
	m, err := migrate.NewWithDatabaseInstance(source, "sqlite3", driver)
	if err != nil {
		return fmt.Errorf("could not migrate sqlite3 db - migrate init error: %w", err)
	}

	// Run migrations
	if err = m.Up(); err != nil {
		return fmt.Errorf("could not migrate sqlite3 db - up error: %w", err)
	}

	// Display post-migration status
	v, dirty, err := m.Version()
	if err != nil {
		return fmt.Errorf("could not migrate sqlite3 db - version error: %w", err)
	}
	fmt.Printf("DB Migration Version: %d. Dirty: %v\n", v, dirty)

	return nil
}

func tempDbPath() string {
	rand.Seed(time.Now().UnixNano())
	r := rand.Int63()
	rStr := strconv.FormatInt(r, 36)
	filename := fmt.Sprintf("matchstick.video.%s.db", rStr)
	return path.Join(os.TempDir(), filename)
}
