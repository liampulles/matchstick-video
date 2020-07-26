package inventory_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"

	entityMocks "github.com/liampulles/matchstick-video/test/mock/pkg/domain/entity"

	"github.com/liampulles/matchstick-video/pkg/usecase/inventory"
)

type EntityModifierTestSuite struct {
	suite.Suite
	sut *inventory.EntityModifierImpl
}

func TestEntityModifierTestSuite(t *testing.T) {
	suite.Run(t, new(EntityModifierTestSuite))
}

func (suite *EntityModifierTestSuite) SetupTest() {
	suite.sut = inventory.NewEntityModifierImpl()
}

func (suite *EntityModifierTestSuite) TestModifyWithUpdateItemVO_WhenEntityChangeNameFails_ShouldFail() {
	// Setup fixture
	voFixture := &inventory.UpdateItemVO{
		Name: "some.name",
	}

	// Setup mocks
	mockEntity := &entityMocks.MockInventoryItem{}
	mockErr := fmt.Errorf("mock.error")
	mockEntity.On("ChangeName", "some.name").Return(mockErr)

	// Setup expectations
	expectedErr := "could not modify entity with update vo - entity name change error: mock.error"

	// Exercise SUT
	err := suite.sut.ModifyWithUpdateItemVO(mockEntity, voFixture)

	// Verify results
	suite.EqualError(err, expectedErr)
}

func (suite *EntityModifierTestSuite) TestModifyWithUpdateItemVO_WhenEntityChangeLocationFails_ShouldFail() {
	// Setup fixture
	voFixture := &inventory.UpdateItemVO{
		Name:     "some.name",
		Location: "some.location",
	}

	// Setup mocks
	mockEntity := &entityMocks.MockInventoryItem{}
	mockErr := fmt.Errorf("mock.error")
	mockEntity.On("ChangeName", "some.name").Return(nil)
	mockEntity.On("ChangeLocation", "some.location").Return(mockErr)

	// Setup expectations
	expectedErr := "could not modify entity with update vo - entity location change error: mock.error"

	// Exercise SUT
	err := suite.sut.ModifyWithUpdateItemVO(mockEntity, voFixture)

	// Verify results
	suite.EqualError(err, expectedErr)
}

func (suite *EntityModifierTestSuite) TestModifyWithUpdateItemVO_WhenChangesSucceed_ShouldReturnAsExpected() {
	// Setup fixture
	voFixture := &inventory.UpdateItemVO{
		Name:     "some.name",
		Location: "some.location",
	}

	// Setup mocks
	mockEntity := &entityMocks.MockInventoryItem{}
	mockEntity.On("ChangeName", "some.name").Return(nil)
	mockEntity.On("ChangeLocation", "some.location").Return(nil)

	// Exercise SUT
	err := suite.sut.ModifyWithUpdateItemVO(mockEntity, voFixture)

	// Verify results
	suite.NoError(err)
}
