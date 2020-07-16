package inventory_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"

	entityMocks "github.com/liampulles/matchstick-video/test/mock/pkg/domain/entity"
	inventoryMocks "github.com/liampulles/matchstick-video/test/mock/pkg/usecase/inventory"

	"github.com/liampulles/matchstick-video/pkg/domain/entity"
	"github.com/liampulles/matchstick-video/pkg/usecase/inventory"
)

type ServiceImplTestSuite struct {
	suite.Suite
	mockRepository    *inventoryMocks.MockRepository
	mockEntityFactory *inventoryMocks.MockEntityFactory
	mockEntityModifer *inventoryMocks.MockEntityModifier
	sut               *inventory.ServiceImpl
}

func TestServiceImplTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceImplTestSuite))
}

func (suite *ServiceImplTestSuite) SetupTest() {
	suite.mockRepository = &inventoryMocks.MockRepository{}
	suite.mockEntityFactory = &inventoryMocks.MockEntityFactory{}
	suite.mockEntityModifer = &inventoryMocks.MockEntityModifier{}
	suite.sut = inventory.NewServiceImpl(
		suite.mockRepository, suite.mockEntityFactory, suite.mockEntityModifer)
}

func (suite *ServiceImplTestSuite) TestCreate_WhenFactoryFails_ShouldFail() {
	// Setup fixture
	voFixture := &inventory.CreateItemVO{
		Name: "some.name",
	}

	// Setup mocks
	mockErr := fmt.Errorf("mock.error")
	suite.mockEntityFactory.On("CreateFromVO", voFixture).Return(nil, mockErr)

	// Setup expectations
	expectedErr := "could not create inventory item - factory error: mock.error"

	// Exercise SUT
	actual, err := suite.sut.Create(voFixture)

	// Verify results
	suite.Nil(actual)
	suite.EqualError(err, expectedErr)
}

func (suite *ServiceImplTestSuite) TestCreate_WhenRepositoryFails_ShouldFail() {
	// Setup fixture
	voFixture := &inventory.CreateItemVO{
		Name: "some.name",
	}

	// Setup mocks
	mockEntity := &entityMocks.InventoryItemMock{Data: "mock.data"}
	mockErr := fmt.Errorf("mock.error")
	suite.mockEntityFactory.On("CreateFromVO", voFixture).Return(mockEntity, nil)
	suite.mockRepository.On("Save", mockEntity).Return(nil, mockErr)

	// Setup expectations
	expectedErr := "could not create inventory item - repository error: mock.error"

	// Exercise SUT
	actual, err := suite.sut.Create(voFixture)

	// Verify results
	suite.Nil(actual)
	suite.EqualError(err, expectedErr)
}

func (suite *ServiceImplTestSuite) TestCreate_WhenDelegatesSucceed_ShouldReturnExpected() {
	// Setup fixture
	voFixture := &inventory.CreateItemVO{
		Name: "some.name",
	}

	// Setup mocks
	mockEntity1 := &entityMocks.InventoryItemMock{Data: "mock.data.1"}
	mockEntity2 := &entityMocks.InventoryItemMock{Data: "mock.data.2"}
	suite.mockEntityFactory.On("CreateFromVO", voFixture).Return(mockEntity1, nil)
	suite.mockRepository.On("Save", mockEntity1).Return(mockEntity2, nil)

	// Exercise SUT
	actual, err := suite.sut.Create(voFixture)

	// Verify results
	suite.NoError(err)
	suite.Equal(actual, mockEntity2)
}

func (suite *ServiceImplTestSuite) TestRead_WhenRepositoryFails_ShouldFail() {
	// Setup fixture
	idFixture := entity.ID(101)

	// Setup mocks
	mockErr := fmt.Errorf("mock.error")
	suite.mockRepository.On("FindByID", idFixture).Return(nil, mockErr)

	// Setup expectations
	expectedErr := "could not read inventory item - repository error: mock.error"

	// Exercise SUT
	actual, err := suite.sut.Read(idFixture)

	// Verify results
	suite.Nil(actual)
	suite.EqualError(err, expectedErr)
}

func (suite *ServiceImplTestSuite) TestRead_WhenDelegatesSucceed_ShouldReturnAsExpected() {
	// Setup fixture
	idFixture := entity.ID(101)

	// Setup mocks
	mockEntity := &entityMocks.InventoryItemMock{Data: "mock.data"}
	suite.mockRepository.On("FindByID", idFixture).Return(mockEntity, nil)

	// Exercise SUT
	actual, err := suite.sut.Read(idFixture)

	// Verify results
	suite.NoError(err)
	suite.Equal(actual, mockEntity)
}

func (suite *ServiceImplTestSuite) TestUpdate_WhenRepositoryFindFails_ShouldFail() {
	// Setup fixture
	voFixture := &inventory.UpdateItemVO{
		ID:   101,
		Name: "new.name",
	}

	// Setup mocks
	mockErr := fmt.Errorf("mock.error")
	suite.mockRepository.On("FindByID", entity.ID(101)).Return(nil, mockErr)

	// Setup expectations
	expectedErr := "could not update inventory item - repository error: mock.error"

	// Exercise SUT
	err := suite.sut.Update(voFixture)

	// Verify results
	suite.EqualError(err, expectedErr)
}

func (suite *ServiceImplTestSuite) TestUpdate_WhenModifierFails_ShouldFail() {
	// Setup fixture
	voFixture := &inventory.UpdateItemVO{
		ID:   101,
		Name: "new.name",
	}

	// Setup mocks
	mockErr := fmt.Errorf("mock.error")
	mockEntity := &entityMocks.InventoryItemMock{Data: "mock.data"}
	suite.mockRepository.On("FindByID", entity.ID(101)).Return(mockEntity, nil)
	suite.mockEntityModifer.On("ModifyWithUpdateItemVO", mockEntity, voFixture).Return(mockErr)

	// Setup expectations
	expectedErr := "could not update inventory item - modifier error: mock.error"

	// Exercise SUT
	err := suite.sut.Update(voFixture)

	// Verify results
	suite.EqualError(err, expectedErr)
}

func (suite *ServiceImplTestSuite) TestUpdate_WhenRepositorySaveFails_ShouldFail() {
	// Setup fixture
	voFixture := &inventory.UpdateItemVO{
		ID:   101,
		Name: "new.name",
	}

	// Setup expectations
	expectedErr := "could not update inventory item - repository error: mock.error"

	// Setup mocks
	mockEntity := &entityMocks.InventoryItemMock{Data: "mock.data"}
	mockErr := fmt.Errorf("mock.error")
	suite.mockRepository.On("FindByID", entity.ID(101)).Return(mockEntity, nil)
	suite.mockEntityModifer.On("ModifyWithUpdateItemVO", mockEntity, voFixture).Return(nil)
	suite.mockRepository.On("Save", mockEntity).Return(nil, mockErr)

	// Exercise SUT
	err := suite.sut.Update(voFixture)

	// Verify results
	suite.EqualError(err, expectedErr)
}

func (suite *ServiceImplTestSuite) TestUpdate_WhenDelegatesSucceed_ShouldReturnAsExpected() {
	// Setup fixture
	voFixture := &inventory.UpdateItemVO{
		ID:   101,
		Name: "new.name",
	}

	// Setup mocks
	mockEntity := &entityMocks.InventoryItemMock{Data: "mock.data"}
	suite.mockRepository.On("FindByID", entity.ID(101)).Return(mockEntity, nil)
	suite.mockEntityModifer.On("ModifyWithUpdateItemVO", mockEntity, voFixture).Return(nil)
	suite.mockRepository.On("Save", mockEntity).Return(nil, nil)

	// Exercise SUT
	err := suite.sut.Update(voFixture)

	// Verify results
	suite.NoError(err)
}

func (suite *ServiceImplTestSuite) TestDelete_WhenRepositoryFails_ShouldFail() {
	// Setup fixture
	idFixture := entity.ID(101)

	// Setup mocks
	mockErr := fmt.Errorf("mock.error")
	suite.mockRepository.On("DeleteByID", idFixture).Return(mockErr)

	// Setup expectations
	expectedErr := "could not delete inventory item - repository error: mock.error"

	// Exercise SUT
	err := suite.sut.Delete(idFixture)

	// Verify results
	suite.EqualError(err, expectedErr)
}

func (suite *ServiceImplTestSuite) TestDelete_WhenDelegatesSucceed_ShouldReturnAsExpected() {
	// Setup fixture
	idFixture := entity.ID(101)

	// Setup mocks
	suite.mockRepository.On("DeleteByID", idFixture).Return(nil)

	// Exercise SUT
	err := suite.sut.Delete(idFixture)

	// Verify results
	suite.NoError(err)
}

func (suite *ServiceImplTestSuite) TestIsAvailable_WhenRepositoryFails_ShouldFail() {
	// Setup fixture
	idFixture := entity.ID(101)

	// Setup mocks
	mockErr := fmt.Errorf("mock.error")
	suite.mockRepository.On("FindByID", idFixture).Return(nil, mockErr)

	// Setup expectations
	expectedErr := "could not determine if inventory item is available - repository error: mock.error"

	// Exercise SUT
	actual, err := suite.sut.IsAvailable(idFixture)

	// Verify results
	suite.False(actual)
	suite.EqualError(err, expectedErr)
}

func (suite *ServiceImplTestSuite) TestIsAvailable_WhenDelegatesSucceed_ShouldReturnAsExpected() {
	// Setup fixture
	idFixture := entity.ID(101)

	// Setup mocks
	mockEntity := &entityMocks.InventoryItemMock{Data: "mock.data.1"}
	mockEntity.On("IsAvailable").Return(true)
	suite.mockRepository.On("FindByID", idFixture).Return(mockEntity, nil)

	// Exercise SUT
	actual, err := suite.sut.IsAvailable(idFixture)

	// Verify results
	suite.NoError(err)
	suite.True(actual)
}

func (suite *ServiceImplTestSuite) TestCheckout_WhenRepositoryFindFails_ShouldFail() {
	// Setup fixture
	idFixture := entity.ID(101)

	// Setup mocks
	mockErr := fmt.Errorf("mock.error")
	suite.mockRepository.On("FindByID", idFixture).Return(nil, mockErr)

	// Setup expectations
	expectedErr := "could not checkout inventory item - repository error: mock.error"

	// Exercise SUT
	err := suite.sut.Checkout(idFixture)

	// Verify results
	suite.EqualError(err, expectedErr)
}

func (suite *ServiceImplTestSuite) TestCheckout_WhenEntityFails_ShouldFail() {
	// Setup fixture
	idFixture := entity.ID(101)

	// Setup mocks
	mockEntity := &entityMocks.InventoryItemMock{Data: "some.data"}
	mockErr := fmt.Errorf("mock.error")
	suite.mockRepository.On("FindByID", idFixture).Return(mockEntity, nil)
	mockEntity.On("Checkout").Return(mockErr)

	// Setup expectations
	expectedErr := "could not checkout inventory item - entity error: mock.error"

	// Exercise SUT
	err := suite.sut.Checkout(idFixture)

	// Verify results
	suite.EqualError(err, expectedErr)
}

func (suite *ServiceImplTestSuite) TestCheckout_WhenRepositorySaveFails_ShouldFail() {
	// Setup fixture
	idFixture := entity.ID(101)

	// Setup mocks
	mockEntity := &entityMocks.InventoryItemMock{Data: "some.data"}
	mockErr := fmt.Errorf("mock.error")
	suite.mockRepository.On("FindByID", idFixture).Return(mockEntity, nil)
	mockEntity.On("Checkout").Return(nil)
	suite.mockRepository.On("Save", mockEntity).Return(nil, mockErr)

	// Setup expectations
	expectedErr := "could not checkout inventory item - repository error: mock.error"

	// Exercise SUT
	err := suite.sut.Checkout(idFixture)

	// Verify results
	suite.EqualError(err, expectedErr)
}

func (suite *ServiceImplTestSuite) TestCheckout_WhenDelegatesSucceed_ShouldReturnAsExpected() {
	// Setup fixture
	idFixture := entity.ID(101)

	// Setup mocks
	mockEntity1 := &entityMocks.InventoryItemMock{Data: "some.data.1"}
	suite.mockRepository.On("FindByID", idFixture).Return(mockEntity1, nil)
	mockEntity1.On("Checkout").Return(nil)
	suite.mockRepository.On("Save", mockEntity1).Return(nil, nil)

	// Exercise SUT
	err := suite.sut.Checkout(idFixture)

	// Verify results
	suite.NoError(err)
}

func (suite *ServiceImplTestSuite) TestCheckIn_WhenRepositoryFindFails_ShouldFail() {
	// Setup fixture
	idFixture := entity.ID(101)

	// Setup mocks
	mockErr := fmt.Errorf("mock.error")
	suite.mockRepository.On("FindByID", idFixture).Return(nil, mockErr)

	// Setup expectations
	expectedErr := "could not check in inventory item - repository error: mock.error"

	// Exercise SUT
	err := suite.sut.CheckIn(idFixture)

	// Verify results
	suite.EqualError(err, expectedErr)
}

func (suite *ServiceImplTestSuite) TestCheckIn_WhenEntityFails_ShouldFail() {
	// Setup fixture
	idFixture := entity.ID(101)

	// Setup mocks
	mockEntity := &entityMocks.InventoryItemMock{Data: "some.data"}
	mockErr := fmt.Errorf("mock.error")
	suite.mockRepository.On("FindByID", idFixture).Return(mockEntity, nil)
	mockEntity.On("CheckIn").Return(mockErr)

	// Setup expectations
	expectedErr := "could not check in inventory item - entity error: mock.error"

	// Exercise SUT
	err := suite.sut.CheckIn(idFixture)

	// Verify results
	suite.EqualError(err, expectedErr)
}

func (suite *ServiceImplTestSuite) TestCheckIn_WhenRepositorySaveFails_ShouldFail() {
	// Setup fixture
	idFixture := entity.ID(101)

	// Setup mocks
	mockEntity := &entityMocks.InventoryItemMock{Data: "some.data"}
	mockErr := fmt.Errorf("mock.error")
	suite.mockRepository.On("FindByID", idFixture).Return(mockEntity, nil)
	mockEntity.On("CheckIn").Return(nil)
	suite.mockRepository.On("Save", mockEntity).Return(nil, mockErr)

	// Setup expectations
	expectedErr := "could not check in inventory item - repository error: mock.error"

	// Exercise SUT
	err := suite.sut.CheckIn(idFixture)

	// Verify results
	suite.EqualError(err, expectedErr)
}

func (suite *ServiceImplTestSuite) TestCheckIn_WhenDelegatesSucceed_ShouldReturnAsExpected() {
	// Setup fixture
	idFixture := entity.ID(101)

	// Setup mocks
	mockEntity := &entityMocks.InventoryItemMock{Data: "some.data"}
	suite.mockRepository.On("FindByID", idFixture).Return(mockEntity, nil)
	mockEntity.On("CheckIn").Return(nil)
	suite.mockRepository.On("Save", mockEntity).Return(nil, nil)

	// Exercise SUT
	err := suite.sut.CheckIn(idFixture)

	// Verify results
	suite.NoError(err)
}
