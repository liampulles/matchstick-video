package db

import (
	"database/sql"
	"fmt"
	"math/rand"
	"os"
	"path"
	"strconv"
	"time"

	// Import the SQLite3 driver in the background
	_ "github.com/mattn/go-sqlite3"
)

// NewTempSQLite3DB create a new SQLite3 database in the
// temp directory. It is effectively an embedded database.
func NewTempSQLite3DB() (*sql.DB, error) {
	dbPath := tempDbPath()

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("could not create sqlite3 db - open error: %w", err)
	}

	return db, nil
}

func tempDbPath() string {
	rand.Seed(time.Now().UnixNano())
	r := rand.Int63()
	rStr := strconv.FormatInt(r, 36)
	filename := fmt.Sprintf("matchstick.video.%s.db", rStr)
	return path.Join(os.TempDir(), filename)
}
