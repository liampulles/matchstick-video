//go:build integration
// +build integration

package integration

import (
	"testing"

	"github.com/stretchr/testify/assert"

	goConfig "github.com/liampulles/go-config"

	"github.com/liampulles/matchstick-video/pkg/adapter/config"
	"github.com/liampulles/matchstick-video/pkg/wire"
)

func TestCreateServerFactory_GivenValidIntegrationConfig_ShouldPass(t *testing.T) {
	// Setup fixture
	config.Load(goConfig.MapSource{
		"PORT":             "9010",
		"MIGRATION_SOURCE": "file://../../migrations",
		"DB_USER":          "integration",
		"DB_PASSWORD":      "integration",
		"DB_NAME":          "integration",
		"DB_PORT":          "5050",
	})

	// Exercise SUT
	actual, err := wire.CreateServerFactory()

	// Verify results
	assert.NotNil(t, actual)
	assert.NoError(t, err)
}

func TestCreateApp_SmokeTest(t *testing.T) {
	// Setup fixture
	config.Load(goConfig.MapSource{
		"PORT":             "9010",
		"MIGRATION_SOURCE": "file://../../migrations",
		"DB_USER":          "integration",
		"DB_PASSWORD":      "integration",
		"DB_NAME":          "integration",
		"DB_PORT":          "5050",
	})

	// Exercise SUT
	wire.CreateApp()
}
