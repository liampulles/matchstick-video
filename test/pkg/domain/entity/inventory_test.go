package entity_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/liampulles/matchstick-video/pkg/domain/entity"
)

func TestInventoryItem_IsAvailable_FalseCase(t *testing.T) {
	// Setup fixture
	fixture := &entity.InventoryItem{
		Available: false,
	}

	// Exercise SUT
	actual := fixture.IsAvailable()

	// Verify results
	assert.False(t, actual)
}

func TestInventoryItem_IsAvailable_TrueCase(t *testing.T) {
	// Setup fixture
	fixture := &entity.InventoryItem{
		Available: true,
	}

	// Exercise SUT
	actual := fixture.IsAvailable()

	// Verify results
	assert.True(t, actual)
}

func TestInventoryItem_Checkout_WhenUnavailable_ShouldFail(t *testing.T) {
	// Setup fixture
	fixture := &entity.InventoryItem{
		Available: false,
	}

	// Exercise SUT
	err := fixture.Checkout()

	// Verify results
	assert.Error(t, err)
}

func TestInventoryItem_Checkout_WhenAvailable_ShouldPass(t *testing.T) {
	// Setup fixture
	fixture := &entity.InventoryItem{
		Available: true,
	}

	// Exercise SUT
	err := fixture.Checkout()

	// Verify results
	assert.NoError(t, err)
}

func TestInventoryItem_CheckIn_WhenAvailable_ShouldFail(t *testing.T) {
	// Setup fixture
	fixture := &entity.InventoryItem{
		Available: true,
	}

	// Exercise SUT
	err := fixture.CheckIn()

	// Verify results
	assert.Error(t, err)
}

func TestInventoryItem_CheckIn_WhenUnavailable_ShouldPass(t *testing.T) {
	// Setup fixture
	fixture := &entity.InventoryItem{
		Available: false,
	}

	// Exercise SUT
	err := fixture.CheckIn()

	// Verify results
	assert.NoError(t, err)
}
