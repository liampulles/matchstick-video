package inventory_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	entityMocks "github.com/liampulles/matchstick-video/test/mock/pkg/domain/entity"

	"github.com/liampulles/matchstick-video/pkg/domain/entity"
	"github.com/liampulles/matchstick-video/pkg/usecase/inventory"
)

type VOFactoryImplTestSuite struct {
	suite.Suite
	sut *inventory.VOFactoryImpl
}

func TestVOFactoryImplTestSuite(t *testing.T) {
	suite.Run(t, new(VOFactoryImplTestSuite))
}

func (suite *VOFactoryImplTestSuite) SetupTest() {
	suite.sut = inventory.NewVOFactoryImpl()
}

func (suite *VOFactoryImplTestSuite) TestCreateViewVOFromEntity_ShouldMapFields() {
	// Setup mocks
	mockEntity := &entityMocks.InventoryItemMock{}
	mockEntity.On("ID").Return(entity.ID(101))
	mockEntity.On("Name").Return("some.name")
	mockEntity.On("Location").Return("some.location")
	mockEntity.On("IsAvailable").Return(true)

	// Setup expectations
	expected := &inventory.ViewVO{
		ID:        entity.ID(101),
		Name:      "some.name",
		Location:  "some.location",
		Available: true,
	}

	// Exercise SUT
	actual := suite.sut.CreateViewVOFromEntity(mockEntity)

	// Verify results
	suite.Equal(actual, expected)
}
