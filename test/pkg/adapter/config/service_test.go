package config_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	goConfig "github.com/liampulles/go-config"

	"github.com/liampulles/matchstick-video/pkg/adapter/config"
)

func TestStore_NewStoreImpl_WhenConfigIsWrongType_ShouldFail(t *testing.T) {
	// Setup fixture
	fixture := goConfig.MapSource(map[string]string{
		"PORT": "not.an.int",
	})

	// Setup expectations
	expectedErr := "could not fetch config: value of PORT property can not be converted to int (is not.an.int)"

	// Exercise SUT
	actual, err := config.NewStoreImpl(fixture)

	// Verify results
	assert.Nil(t, actual)
	assert.EqualError(t, err, expectedErr)
}

func TestStore_NewStoreImpl_WhenConfigIsValid_ShouldPass(t *testing.T) {
	// Setup fixture
	fixture := goConfig.MapSource(map[string]string{
		"PORT": "9001",
	})

	// Exercise SUT
	actual, err := config.NewStoreImpl(fixture)

	// Verify results
	assert.NoError(t, err)
	assert.NotNil(t, actual)
}

func TestStore_GetPort_ShouldReturnPort(t *testing.T) {
	// Setup fixture
	fixture := goConfig.MapSource(map[string]string{
		"PORT": "9001",
	})
	sut, _ := config.NewStoreImpl(fixture)

	// Exercise SUT
	actual := sut.GetPort()

	// Verify results
	assert.Equal(t, 9001, actual)
}

func TestStore_GetDbDriver_ShouldReturnDbDriver(t *testing.T) {
	// Setup fixture
	fixture := goConfig.MapSource(map[string]string{
		"DB_DRIVER": "some.driver",
	})
	sut, _ := config.NewStoreImpl(fixture)

	// Exercise SUT
	actual := sut.GetDbDriver()

	// Verify results
	assert.Equal(t, "some.driver", actual)
}

func TestStore_GetMigrationSource_ShouldReturnMigrationSource(t *testing.T) {
	// Setup fixture
	fixture := goConfig.MapSource(map[string]string{
		"MIGRATION_SOURCE": "some.migration.source",
	})
	sut, _ := config.NewStoreImpl(fixture)

	// Exercise SUT
	actual := sut.GetMigrationSource()

	// Verify results
	assert.Equal(t, "some.migration.source", actual)
}
