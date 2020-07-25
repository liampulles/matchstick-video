package db_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/liampulles/matchstick-video/pkg/driver/db"
)

func TestNewTempSQLite3DB_ShouldPass(t *testing.T) {
	// Setup fixture

	// Exercise SUT
	actual, err := db.NewTempSQLite3DB()

	// Verify results
	assert.NoError(t, err)
	assert.NotNil(t, actual)

	stmt, err := actual.Prepare(`CREATE TABLE contacts (
		contact_id INTEGER PRIMARY KEY,
		first_name TEXT NOT NULL,
		last_name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		phone TEXT NOT NULL UNIQUE
	);`)
	assert.NoError(t, err)

	res, err := stmt.Exec()
	assert.NoError(t, err)

	affect, err := res.RowsAffected()
	assert.NoError(t, err)

	fmt.Println(affect)

	err = actual.Close()
	assert.NoError(t, err)
}
