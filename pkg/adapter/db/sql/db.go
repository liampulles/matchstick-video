package sql

import (
	goSql "database/sql"

	"github.com/liampulles/matchstick-video/pkg/driver/db"
	"github.com/rs/zerolog/log"
)

// GetDB lazy loads the database
func GetDB() *goSql.DB {
	if DB == nil {
		db, err := Load()
		if err != nil {
			log.Err(err).Msg("could not load db")
			panic(err)
		}
		DB = db
	}
	return DB
}

// --- Backing ---

// DB to use. Should only override this in a test.
var DB *goSql.DB

// Load actually connects to a database, if not already connected.
// Default driver is Postgres.
var Load = db.NewPostgresDB
