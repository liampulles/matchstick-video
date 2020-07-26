package db_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	configMocks "github.com/liampulles/matchstick-video/test/mock/pkg/adapter/config"

	"github.com/liampulles/matchstick-video/pkg/driver/db"
)

func TestNewDatabaseServiceImpl_GivenUnknownDbDriver_ShouldFail(t *testing.T) {
	// Setup mocks
	cfgMock := &configMocks.MockStore{}
	cfgMock.On("GetDbDriver").Return("not.a.driver")

	// Exercise SUT
	actual, err := db.NewDatabaseServiceImpl(cfgMock)

	// Verify results
	assert.Nil(t, actual)
	assert.Error(t, err)
}

func TestNewDatabaseServiceImpl_GivenSqlite3Driver_ShouldPass(t *testing.T) {
	// Setup mocks
	cfgMock := &configMocks.MockStore{}
	cfgMock.On("GetDbDriver").Return("sqlite3")
	cfgMock.On("GetMigrationSource").Return("file://../../../../migrations")

	// Exercise SUT
	actual, err := db.NewDatabaseServiceImpl(cfgMock)

	// Verify results
	assert.NotNil(t, actual)
	assert.NoError(t, err)
	assert.NotNil(t, actual.Get())
}
