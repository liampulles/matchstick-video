package wire_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	goConfig "github.com/liampulles/go-config"

	"github.com/liampulles/matchstick-video/pkg/wire"
)

func TestCreateServer_GivenInvalidConfig_ShouldFail(t *testing.T) {
	// Setup fixture
	fixture := goConfig.MapSource(map[string]string{
		"PORT": "not.an.int",
	})

	// Setup expectations
	expectedErr := "could not fetch config: value of PORT property can not be converted to int (is not.an.int)"

	// Exercise SUT
	actual, err := wire.CreateServer(fixture)

	// Verify results
	assert.Nil(t, actual)
	assert.EqualError(t, err, expectedErr)
}

func TestCreateServer_GivenValidConfig_ShouldPass(t *testing.T) {
	// Setup fixture
	fixture := goConfig.MapSource(map[string]string{})

	// Exercise SUT
	actual, err := wire.CreateServer(fixture)

	// Verify results
	assert.NoError(t, err)
	assert.NotNil(t, actual)
}
