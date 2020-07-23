package dto_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	entityMocks "github.com/liampulles/matchstick-video/test/mock/pkg/domain/entity"

	"github.com/liampulles/matchstick-video/pkg/adapter/http/json/dto"
	"github.com/liampulles/matchstick-video/pkg/domain/entity"
)

type FactoryImplTestSuite struct {
	suite.Suite
	sut *dto.FactoryImpl
}

func TestFactoryImplTestSuite(t *testing.T) {
	suite.Run(t, new(FactoryImplTestSuite))
}

func (suite *FactoryImplTestSuite) SetupTest() {
	suite.sut = dto.NewFactoryImpl()
}

func (suite *FactoryImplTestSuite) TestCreateInventoryItemViewFromEntity_WhenUnmarshalFails_ShouldMap() {
	// Setup mocks
	entityMock := &entityMocks.InventoryItemMock{}
	entityMock.On("ID").Return(entity.ID(101))
	entityMock.On("Name").Return("some.name")
	entityMock.On("Location").Return("some.location")
	entityMock.On("IsAvailable").Return(true)

	// Setup expectations
	expected := &dto.InventoryItemView{
		ID:        101,
		Name:      "some.name",
		Location:  "some.location",
		Available: true,
	}

	// Exercise SUT
	actual := suite.sut.CreateInventoryItemViewFromEntity(entityMock)

	// Verify results
	suite.Equal(expected, actual)
}
