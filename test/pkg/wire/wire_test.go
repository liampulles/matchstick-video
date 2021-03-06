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
	// Exercise SUT
	actual, err := wire.CreateServerFactory(fixture)

	// Verify results
	assert.Nil(t, actual)
	assert.Error(t, err)
}
