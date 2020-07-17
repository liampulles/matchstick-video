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

func (suite *InventoryItemConstructorTestSuite) TestNewAvailable_ShouldCreateAvailableEntity() {
	// Setup fixture
	nameFixture := "some.name"
	locationFixture := "some.location"

	// Exercise SUT
	actual := suite.sut.NewAvailable(nameFixture, locationFixture)

	// Verify results
	suite.Equal(actual.ID(), entity.InvalidID)
	suite.Equal(actual.Name(), nameFixture)
	suite.Equal(actual.Location(), locationFixture)
	suite.True(actual.IsAvailable())
}
