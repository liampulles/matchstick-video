package entity_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/liampulles/matchstick-video/pkg/domain/entity"
)

type InventoryConstructorsTestSuite struct {
	suite.Suite
}

func TestInventoryConstructorsTestSuite(t *testing.T) {
	suite.Run(t, new(InventoryConstructorsTestSuite))
}

func (suite *InventoryConstructorsTestSuite) TestNewAvailable_WhenNameValidationFails_ShouldFail() {
	// Setup fixture
	nameFixture := "some.name "
	locationFixture := "some.location"

	// Setup expectations
	expectedErr := "validation error: field=[name], problem=[must not have whitespace at the beginning or the end]"

	// Exercise SUT
	actual, err := entity.NewAvailableInventory(nameFixture, locationFixture)

	// Verify results
	suite.EqualError(err, expectedErr)
	suite.Nil(actual)
}

func (suite *InventoryConstructorsTestSuite) TestNewAvailable_WhenLocationValidationFails_ShouldFail() {
	// Setup fixture
	nameFixture := "some.name"
	locationFixture := "some.location "

	// Setup expectations
	expectedErr := "validation error: field=[location], problem=[must not have whitespace at the beginning or the end]"

	// Exercise SUT
	actual, err := entity.NewAvailableInventory(nameFixture, locationFixture)

	// Verify results
	suite.EqualError(err, expectedErr)
	suite.Nil(actual)
}

func (suite *InventoryConstructorsTestSuite) TestNewAvailable_WhenValidationPasses_ShouldCreateAvailableEntity() {
	// Setup fixture
	nameFixture := "some.name"
	locationFixture := "some.location"

	// Exercise SUT
	actual, err := entity.NewAvailableInventory(nameFixture, locationFixture)

	// Verify results
	suite.NoError(err)
	suite.Equal(actual.ID(), entity.InvalidID)
	suite.Equal(actual.Name(), nameFixture)
	suite.Equal(actual.Location(), locationFixture)
	suite.True(actual.IsAvailable())
}

func (suite *InventoryConstructorsTestSuite) TestReincarnate_ShouldCreateGivenEntity() {
	// Setup fixture
	idFixture := entity.ID(101)
	nameFixture := "some.name"
	locationFixture := "some.location"
	availableFixture := true

	// Exercise SUT
	actual := entity.ReincarnateInventory(idFixture, nameFixture, locationFixture, availableFixture)

	// Verify results
	suite.Equal(actual.ID(), idFixture)
	suite.Equal(actual.Name(), nameFixture)
	suite.Equal(actual.Location(), locationFixture)
	suite.True(actual.IsAvailable())
}
