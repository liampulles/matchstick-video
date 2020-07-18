package entity_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/liampulles/matchstick-video/pkg/domain/entity"
)

func TestInventoryItem_ID_ShouldReturnID(t *testing.T) {
	// Setup fixture
	fixture := entity.TestInventoryItemImplConstructor(101, "", "", true)

	// Exercise SUT
	actual := fixture.ID()

	// Verify results
	assert.Equal(t, actual, entity.ID(101))
}

func TestInventoryItem_Name_ShouldReturnName(t *testing.T) {
	// Setup fixture
	fixture := entity.TestInventoryItemImplConstructor(101, "some.name", "", true)

	// Exercise SUT
	actual := fixture.Name()

	// Verify results
	assert.Equal(t, actual, "some.name")
}

func TestInventoryItem_Location_ShouldReturnLocation(t *testing.T) {
	// Setup fixture
	fixture := entity.TestInventoryItemImplConstructor(101, "", "some.location", true)

	// Exercise SUT
	actual := fixture.Location()

	// Verify results
	assert.Equal(t, actual, "some.location")
}

func TestInventoryItem_IsAvailable_FalseCase(t *testing.T) {
	// Setup fixture
	fixture := entity.TestInventoryItemImplConstructor(101, "", "", false)

	// Exercise SUT
	actual := fixture.IsAvailable()

	// Verify results
	assert.False(t, actual)
}

func TestInventoryItem_IsAvailable_TrueCase(t *testing.T) {
	// Setup fixture
	fixture := entity.TestInventoryItemImplConstructor(101, "", "", true)

	// Exercise SUT
	actual := fixture.IsAvailable()

	// Verify results
	assert.True(t, actual)
}

func TestInventoryItem_Checkout_WhenUnavailable_ShouldFail(t *testing.T) {
	// Setup fixture
	fixture := entity.TestInventoryItemImplConstructor(101, "", "", false)

	// Exercise SUT
	err := fixture.Checkout()

	// Verify results
	assert.Error(t, err)
}

func TestInventoryItem_Checkout_WhenAvailable_ShouldPass(t *testing.T) {
	// Setup fixture
	fixture := entity.TestInventoryItemImplConstructor(101, "", "", true)

	// Exercise SUT
	err := fixture.Checkout()

	// Verify results
	assert.NoError(t, err)
	assert.False(t, fixture.IsAvailable())
}

func TestInventoryItem_CheckIn_WhenAvailable_ShouldFail(t *testing.T) {
	// Setup fixture
	fixture := entity.TestInventoryItemImplConstructor(101, "", "", true)

	// Exercise SUT
	err := fixture.CheckIn()

	// Verify results
	assert.Error(t, err)
}

func TestInventoryItem_CheckIn_WhenUnavailable_ShouldPass(t *testing.T) {
	// Setup fixture
	fixture := entity.TestInventoryItemImplConstructor(101, "", "", false)

	// Exercise SUT
	err := fixture.CheckIn()

	// Verify results
	assert.NoError(t, err)
	assert.True(t, fixture.IsAvailable())
}

func TestInventoryItem_ChangeName_WhenGivenNameIsBlank_ShouldFail(t *testing.T) {
	// Setup fixture
	sut := entity.TestInventoryItemImplConstructor(101, "", "", true)
	nameFixture := ""

	// Setup expectations
	expectedErr := "validation error: field=[name], problem=[must not be blank]"

	// Exercise SUT
	err := sut.ChangeName(nameFixture)

	// Verify results
	assert.EqualError(t, err, expectedErr)
}

func TestInventoryItem_ChangeName_WhenGivenNameIsNotTrimmed_ShouldFail(t *testing.T) {
	// Setup fixture
	sut := entity.TestInventoryItemImplConstructor(101, "", "", true)
	nameFixture := " duck"

	// Setup expectations
	expectedErr := "validation error: field=[name], problem=[must not have whitespace at the beginning or the end]"

	// Exercise SUT
	err := sut.ChangeName(nameFixture)

	// Verify results
	assert.EqualError(t, err, expectedErr)
}

func TestInventoryItem_ChangeName_WhenGivenNamePassesValidation_ShouldChange(t *testing.T) {
	// Setup fixture
	sut := entity.TestInventoryItemImplConstructor(101, "", "", true)
	nameFixture := "duck"

	// Exercise SUT
	err := sut.ChangeName(nameFixture)

	// Verify results
	assert.NoError(t, err)
	assert.Equal(t, sut.Name(), nameFixture)
}

func TestInventoryItem_ChangeLocation_WhenGivenLocationIsBlank_ShouldFail(t *testing.T) {
	// Setup fixture
	sut := entity.TestInventoryItemImplConstructor(101, "", "", true)
	locationFixture := ""

	// Setup expectations
	expectedErr := "validation error: field=[location], problem=[must not be blank]"

	// Exercise SUT
	err := sut.ChangeLocation(locationFixture)

	// Verify results
	assert.EqualError(t, err, expectedErr)
}

func TestInventoryItem_ChangeLocation_WhenGivenLocationIsNotTrimmed_ShouldFail(t *testing.T) {
	// Setup fixture
	sut := entity.TestInventoryItemImplConstructor(101, "", "", true)
	locationFixture := " duck"

	// Setup expectations
	expectedErr := "validation error: field=[location], problem=[must not have whitespace at the beginning or the end]"

	// Exercise SUT
	err := sut.ChangeLocation(locationFixture)

	// Verify results
	assert.EqualError(t, err, expectedErr)
}

func TestInventoryItem_ChangeLocation_WhenGivenLocationPassesValidation_ShouldChange(t *testing.T) {
	// Setup fixture
	sut := entity.TestInventoryItemImplConstructor(101, "", "", true)
	locationFixture := "duck"

	// Exercise SUT
	err := sut.ChangeLocation(locationFixture)

	// Verify results
	assert.NoError(t, err)
	assert.Equal(t, sut.Location(), locationFixture)
}
