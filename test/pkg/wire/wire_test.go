package wire_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	goConfig "github.com/liampulles/go-config"

	"github.com/liampulles/matchstick-video/pkg/wire"
)

func TestCreateServerFactory_GivenInvalidConfig_ShouldFail(t *testing.T) {
	// Setup fixture
	fixture := goConfig.MapSource(map[string]string{
		"PORT": "not.an.int",
	})

	// Setup expectations
	expectedErr := "could not fetch config: value of PORT property can not be converted to int (is not.an.int)"

	// Exercise SUT
	actual, err := wire.CreateServerFactory(fixture)

	// Verify results
	assert.Nil(t, actual)
	assert.EqualError(t, err, expectedErr)
}

func TestCreateServerFactory_GivenBadDBConfig_ShouldFail(t *testing.T) {
	// Setup fixture
	fixture := goConfig.MapSource(map[string]string{
		"DB_HOST": "not.a.url",
	})

	// Setup expectations
	expectedErr := "could not create database service - could not migrate db: could not migrate postgres db - driver error: failed to connect to `host=not.a.url user=matchvid database=matchvid`: hostname resolving error (lookup not.a.url: no such host)"

	// Exercise SUT
	actual, err := wire.CreateServerFactory(fixture)

	// Verify results
	assert.Nil(t, actual)
	assert.EqualError(t, err, expectedErr)
}
