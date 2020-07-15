package entity_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/liampulles/matchstick-video/pkg/domain/entity"
)

func TestInventoryItem_ID_ShouldReturnID(t *testing.T) {
	// Setup fixture
	fixture := entity.NewAvailableInventoryItem(101, "", "")

	// Exercise SUT
	actual := fixture.ID()

	// Verify results
	assert.Equal(t, actual, entity.ID(101))
}

func TestInventoryItem_IsAvailable_FalseCase(t *testing.T) {
	// Setup fixture
	fixture := entity.NewUnavailableInventoryItem(-1, "", "")

	// Exercise SUT
	actual := fixture.IsAvailable()

	// Verify results
	assert.False(t, actual)
}

func TestInventoryItem_IsAvailable_TrueCase(t *testing.T) {
	// Setup fixture
	fixture := entity.NewAvailableInventoryItem(-1, "", "")

	// Exercise SUT
	actual := fixture.IsAvailable()

	// Verify results
	assert.True(t, actual)
}

func TestInventoryItem_Checkout_WhenUnavailable_ShouldFail(t *testing.T) {
	// Setup fixture
	fixture := entity.NewUnavailableInventoryItem(-1, "", "")

	// Exercise SUT
	err := fixture.Checkout()

	// Verify results
	assert.Error(t, err)
}

func TestInventoryItem_Checkout_WhenAvailable_ShouldPass(t *testing.T) {
	// Setup fixture
	fixture := entity.NewAvailableInventoryItem(-1, "", "")

	// Exercise SUT
	err := fixture.Checkout()

	// Verify results
	assert.NoError(t, err)
}

func TestInventoryItem_CheckIn_WhenAvailable_ShouldFail(t *testing.T) {
	// Setup fixture
	fixture := entity.NewAvailableInventoryItem(-1, "", "")

	// Exercise SUT
	err := fixture.CheckIn()

	// Verify results
	assert.Error(t, err)
}

func TestInventoryItem_CheckIn_WhenUnavailable_ShouldPass(t *testing.T) {
	// Setup fixture
	fixture := entity.NewUnavailableInventoryItem(-1, "", "")

	// Exercise SUT
	err := fixture.CheckIn()

	// Verify results
	assert.NoError(t, err)
}
