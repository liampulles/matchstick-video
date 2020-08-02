package db_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/liampulles/matchstick-video/pkg/adapter/db"
)

func TestUniqueConstraintError_Error_ShouldReturnCorrectMessage(t *testing.T) {
	// Setup fixture
	err := db.NewUniqueConstraintError(fmt.Errorf("some.error"))

	// Setup expectations
	expected := "uniqueness constraint error: some.error"

	// Verify results
	assert.EqualError(t, err, expected)
}
