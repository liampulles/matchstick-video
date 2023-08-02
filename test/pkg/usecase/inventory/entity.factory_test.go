package inventory_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"

	entityMocks "github.com/liampulles/matchstick-video/test/mock/pkg/domain/entity"

	"github.com/liampulles/matchstick-video/pkg/domain/entity"
	"github.com/liampulles/matchstick-video/pkg/usecase/inventory"
)

type EntityFactoryTestSuite struct {
	suite.Suite
	sut *inventory.EntityFactoryImpl
}

func TestEntityFactoryTestSuite(t *testing.T) {
	suite.Run(t, new(EntityFactoryTestSuite))
}

func (suite *EntityFactoryTestSuite) SetupTest() {
	suite.sut = inventory.NewEntityFactoryImpl()
}

func (suite *EntityFactoryTestSuite) TestCreateFromVO_ShouldCallConstructorAndReturnEntityAndError() {
	// Setup fixture
	voFixture := &inventory.CreateItemVO{
		Name:     "some.name",
		Location: "some.location",
	}

	// Setup mocks
	mockEntity := &entityMocks.MockInventoryItem{}
	mockError := fmt.Errorf("some.error")
	entity.NewAvailableInventory = func(name, location string) (entity.InventoryItem, error) {
		suite.Equal("some.name", name)
		suite.Equal("some.location", location)
		return mockEntity, mockError
	}

	// Exercise SUT
	actual, err := suite.sut.CreateFromVO(voFixture)

	// Verify results
	suite.EqualError(err, "some.error")
	suite.Equal(actual, mockEntity)
}
