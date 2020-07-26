package db_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/liampulles/matchstick-video/pkg/adapter/db"
)

func TestNotFoundError_Error_ShouldReturnCorrectMessage(t *testing.T) {
	// Setup fixture
	err := db.NewNotFoundError("some.type")

	// Setup expectations
	expected := "entity not found: type=[some.type]"

	// Verify results
	assert.EqualError(t, err, expected)
}
