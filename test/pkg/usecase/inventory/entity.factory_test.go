package inventory_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	entityMocks "github.com/liampulles/matchstick-video/test/mock/pkg/domain/entity"

	"github.com/liampulles/matchstick-video/pkg/usecase/inventory"
)

type EntityFactoryTestSuite struct {
	suite.Suite
	mockConstructor *entityMocks.InventoryItemConstructorMock
	sut             *inventory.EntityFactoryImpl
}

func TestEntityFactoryTestSuite(t *testing.T) {
	suite.Run(t, new(EntityFactoryTestSuite))
}

func (suite *EntityFactoryTestSuite) SetupTest() {
	suite.mockConstructor = &entityMocks.InventoryItemConstructorMock{}
	suite.sut = inventory.NewEntityFactoryImpl(suite.mockConstructor)
}

func (suite *EntityFactoryTestSuite) TestCreateFromVO_ShouldCallConstructorAndReturnEntity() {
	// Setup fixture
	voFixture := &inventory.CreateItemVO{
		Name:     "some.name",
		Location: "some.location",
	}

	// Setup mocks
	mockEntity := &entityMocks.InventoryItemMock{}
	suite.mockConstructor.On("NewAvailable", "some.name", "some.location").Return(mockEntity)

	// Exercise SUT
	actual, err := suite.sut.CreateFromVO(voFixture)

	// Verify results
	suite.NoError(err)
	suite.Equal(actual, mockEntity)
}