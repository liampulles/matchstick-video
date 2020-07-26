package db_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/liampulles/matchstick-video/pkg/adapter/db"
	"github.com/liampulles/matchstick-video/pkg/domain/entity"
)

func TestNotFoundError_Error_ShouldReturnCorrectMessage(t *testing.T) {
	// Setup fixture
	err := db.NewNotFoundError("some.type", entity.ID(101))

	// Setup expectations
	expected := "entity not found: type=[some.type], id=[101]"

	// Verify results
	assert.EqualError(t, err, expected)
}
