package entity_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/liampulles/matchstick-video/pkg/domain/entity"
)

type InventoryItemConstructorTestSuite struct {
	suite.Suite
	sut *entity.InventoryItemConstructorImpl
}

func TestInventoryItemConstructorTestSuite(t *testing.T) {
	suite.Run(t, new(InventoryItemConstructorTestSuite))
}

func (suite *InventoryItemConstructorTestSuite) SetupTest() {
	suite.sut = entity.NewInventoryItemConstructorImpl()
}

func (suite *InventoryItemConstructorTestSuite) TestNewAvailable_WhenNameValidationFails_ShouldFail() {
	// Setup fixture
	nameFixture := "some.name "
	locationFixture := "some.location"

	// Setup expectations
	expectedErr := "validation error: field=[name], problem=[must not have whitespace at the beginning or the end]"

	// Exercise SUT
	actual, err := suite.sut.NewAvailable(nameFixture, locationFixture)

	// Verify results
	suite.EqualError(err, expectedErr)
	suite.Nil(actual)
}

func (suite *InventoryItemConstructorTestSuite) TestNewAvailable_WhenLocationValidationFails_ShouldFail() {
	// Setup fixture
	nameFixture := "some.name"
	locationFixture := "some.location "

	// Setup expectations
	expectedErr := "validation error: field=[location], problem=[must not have whitespace at the beginning or the end]"

	// Exercise SUT
	actual, err := suite.sut.NewAvailable(nameFixture, locationFixture)

	// Verify results
	suite.EqualError(err, expectedErr)
	suite.Nil(actual)
}

func (suite *InventoryItemConstructorTestSuite) TestNewAvailable_WhenValidationPasses_ShouldCreateAvailableEntity() {
	// Setup fixture
	nameFixture := "some.name"
	locationFixture := "some.location"

	// Exercise SUT
	actual, err := suite.sut.NewAvailable(nameFixture, locationFixture)

	// Verify results
	suite.NoError(err)
	suite.Equal(actual.ID(), entity.InvalidID)
	suite.Equal(actual.Name(), nameFixture)
	suite.Equal(actual.Location(), locationFixture)
	suite.True(actual.IsAvailable())
}

func (suite *InventoryItemConstructorTestSuite) TestReincarnate_ShouldCreateGivenEntity() {
	// Setup fixture
	idFixture := entity.ID(101)
	nameFixture := "some.name"
	locationFixture := "some.location"
	availableFixture := true

	// Exercise SUT
	actual := suite.sut.Reincarnate(idFixture, nameFixture, locationFixture, availableFixture)

	// Verify results
	suite.Equal(actual.ID(), idFixture)
	suite.Equal(actual.Name(), nameFixture)
	suite.Equal(actual.Location(), locationFixture)
	suite.True(actual.IsAvailable())
}
