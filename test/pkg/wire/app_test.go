package wire_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/liampulles/matchstick-video/pkg/wire"
)

func TestMain_ShouldPass(t *testing.T) {
	// Exercise SUT
	actual := wire.Run()

	// Verify results
	assert.Equal(t, 0, actual)
}
