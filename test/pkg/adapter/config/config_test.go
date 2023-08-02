package config_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	goConfig "github.com/liampulles/go-config"

	"github.com/liampulles/matchstick-video/pkg/adapter/config"
)

func TestLoad_WhenConfigIsWrongType_ShouldFail(t *testing.T) {
	// Setup fixture
	fixture := goConfig.MapSource(map[string]string{
		"PORT": "not.an.int",
	})

	// Setup expectations
	expectedErr := "value of PORT property can not be converted to int (is not.an.int)"

	// Exercise SUT
	err := config.Load(fixture)

	// Verify results
	assert.EqualError(t, err, expectedErr)
}

func TestLoad_WhenConfigIsValid_ShouldPass(t *testing.T) {
	// Setup fixture
	fixture := goConfig.MapSource(map[string]string{
		"PORT": "9001",
	})

	// Exercise SUT
	err := config.Load(fixture)

	// Verify results
	assert.NoError(t, err)
	assert.Equal(t, config.GetPort(), 9001)
}

func TestGetPort_OnceLoaded_ShouldReturnPort(t *testing.T) {
	// Setup fixture
	fixture := goConfig.MapSource(map[string]string{
		"PORT": "9001",
	})
	err := config.Load(fixture)
	assert.NoError(t, err)

	// Exercise SUT
	actual := config.GetPort()

	// Verify results
	assert.Equal(t, 9001, actual)
}
