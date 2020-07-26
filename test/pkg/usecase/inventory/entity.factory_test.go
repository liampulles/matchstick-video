package inventory_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"

	entityMocks "github.com/liampulles/matchstick-video/test/mock/pkg/domain/entity"

	"github.com/liampulles/matchstick-video/pkg/usecase/inventory"
)

type EntityFactoryTestSuite struct {
	suite.Suite
	mockConstructor *entityMocks.MockInventoryItemConstructor
	sut             *inventory.EntityFactoryImpl
}

func TestEntityFactoryTestSuite(t *testing.T) {
	suite.Run(t, new(EntityFactoryTestSuite))
}

func (suite *EntityFactoryTestSuite) SetupTest() {
	suite.mockConstructor = &entityMocks.MockInventoryItemConstructor{}
	suite.sut = inventory.NewEntityFactoryImpl(suite.mockConstructor)
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
	suite.mockConstructor.On("NewAvailable", "some.name", "some.location").Return(mockEntity, mockError)

	// Exercise SUT
	actual, err := suite.sut.CreateFromVO(voFixture)

	// Verify results
	suite.EqualError(err, "some.error")
	suite.Equal(actual, mockEntity)
}
