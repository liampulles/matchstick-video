package wire_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/liampulles/matchstick-video/pkg/wire"
)

func TestMain_WhenGivenBadConfig_ShouldFail(t *testing.T) {
	// Setup
	// -> This is just an example of a config that will fail
	// - we're not going to test every possibility here.
	prev := os.Getenv("PORT")
	os.Setenv("PORT", "not a port")

	// Exercise SUT
	actual := wire.Run()

	// Verify results
	assert.Equal(t, actual, 1)

	// Teardown
	if prev != "" {
		os.Setenv("PORT", prev)
	}
}
