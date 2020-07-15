package entity_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/liampulles/matchstick-video/pkg/domain/entity"
)

func TestNotFoundError_Error_ShouldReturnCorrectMessage(t *testing.T) {
	// Setup fixture
	err := entity.NewNotFoundError("some.type", entity.ID(101))

	// Setup expectations
	expected := "entity not found for type=[some.type] and id=[101]"

	// Verify results
	assert.EqualError(t, err, expected)
}
