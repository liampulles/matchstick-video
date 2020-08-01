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
	mockRepository     *inventoryMocks.MockRepository
	mockEntityFactory  *inventoryMocks.MockEntityFactory
	mockEntityModifier *inventoryMocks.MockEntityModifier
	mockVoFactory      *inventoryMocks.MockVOFactory
	sut                *inventory.ServiceImpl
}

func TestServiceImplTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceImplTestSuite))
}

func (suite *ServiceImplTestSuite) SetupTest() {
	suite.mockRepository = &inventoryMocks.MockRepository{}
	suite.mockEntityFactory = &inventoryMocks.MockEntityFactory{}
	suite.mockEntityModifier = &inventoryMocks.MockEntityModifier{}
	suite.mockVoFactory = &inventoryMocks.MockVOFactory{}
	suite.sut = inventory.NewServiceImpl(
		suite.mockRepository,
		suite.mockEntityFactory,
		suite.mockEntityModifier,
		suite.mockVoFactory,
	)
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
	suite.Equal(actual, entity.InvalidID)
	suite.EqualError(err, expectedErr)
}

func (suite *ServiceImplTestSuite) TestCreate_WhenRepositoryFails_ShouldFail() {
	// Setup fixture
	voFixture := &inventory.CreateItemVO{
		Name: "some.name",
	}

	// Setup mocks
	mockEntity := &entityMocks.MockInventoryItem{Data: "mock.data"}
	mockErr := fmt.Errorf("mock.error")
	suite.mockEntityFactory.On("CreateFromVO", voFixture).Return(mockEntity, nil)
	suite.mockRepository.On("Create", mockEntity).Return(entity.InvalidID, mockErr)

	// Setup expectations
	expectedErr := "could not create inventory item - repository error: mock.error"

	// Exercise SUT
	actual, err := suite.sut.Create(voFixture)

	// Verify results
	suite.Equal(actual, entity.InvalidID)
	suite.EqualError(err, expectedErr)
}

func (suite *ServiceImplTestSuite) TestCreate_WhenDelegatesSucceed_ShouldReturnExpected() {
	// Setup fixture
	voFixture := &inventory.CreateItemVO{
		Name: "some.name",
	}

	// Setup expectations
	expected := entity.ID(101)

	// Setup mocks
	mockEntity := &entityMocks.MockInventoryItem{Data: "mock.data"}
	suite.mockEntityFactory.On("CreateFromVO", voFixture).Return(mockEntity, nil)
	suite.mockRepository.On("Create", mockEntity).Return(expected, nil)

	// Exercise SUT
	actual, err := suite.sut.Create(voFixture)

	// Verify results
	suite.NoError(err)
	suite.Equal(actual, expected)
}

func (suite *ServiceImplTestSuite) TestReadDetails_WhenRepositoryFails_ShouldFail() {
	// Setup fixture
	idFixture := entity.ID(101)

	// Setup mocks
	mockErr := fmt.Errorf("mock.error")
	suite.mockRepository.On("FindByID", idFixture).Return(nil, mockErr)

	// Setup expectations
	expectedErr := "could not read inventory item - repository error: mock.error"

	// Exercise SUT
	actual, err := suite.sut.ReadDetails(idFixture)

	// Verify results
	suite.Nil(actual)
	suite.EqualError(err, expectedErr)
}

func (suite *ServiceImplTestSuite) TestReadDetails_WhenDelegatesSucceed_ShouldReturnAsExpected() {
	// Setup fixture
	idFixture := entity.ID(101)

	// Setup expectations
	expected := &inventory.ViewVO{
		Name: "some.name",
	}

	// Setup mocks
	mockEntity := &entityMocks.MockInventoryItem{Data: "mock.data"}
	suite.mockRepository.On("FindByID", idFixture).Return(mockEntity, nil)
	suite.mockVoFactory.On("CreateViewVOFromEntity", mockEntity).Return(expected)

	// Exercise SUT
	actual, err := suite.sut.ReadDetails(idFixture)

	// Verify results
	suite.NoError(err)
	suite.Equal(actual, expected)
}

func (suite *ServiceImplTestSuite) TestReadAll_WhenRepositoryFails_ShouldFail() {
	// Setup mocks
	mockErr := fmt.Errorf("mock.error")
	suite.mockRepository.On("FindAll").Return(nil, mockErr)

	// Setup expectations
	expectedErr := "could not read inventory items - repository error: mock.error"

	// Exercise SUT
	actual, err := suite.sut.ReadAll()

	// Verify results
	suite.Nil(actual)
	suite.EqualError(err, expectedErr)
}

func (suite *ServiceImplTestSuite) TestReadAll_WhenDelegatesSucceed_ShouldReturnAsExpected() {
	// Setup expectations
	expected := []inventory.ViewVO{
		inventory.ViewVO{
			Name: "some.name",
		},
	}

	// Setup mocks
	mockEntity := &entityMocks.MockInventoryItem{Data: "mock.data"}
	mockEntities := []entity.InventoryItem{mockEntity}
	suite.mockRepository.On("FindAll").Return(mockEntities, nil)
	suite.mockVoFactory.On("CreateViewVOsFromEntities", mockEntities).Return(expected)

	// Exercise SUT
	actual, err := suite.sut.ReadAll()

	// Verify results
	suite.NoError(err)
	suite.Equal(actual, expected)
}

func (suite *ServiceImplTestSuite) TestUpdate_WhenRepositoryFindFails_ShouldFail() {
	// Setup fixture
	idFixture := entity.ID(101)
	voFixture := &inventory.UpdateItemVO{
		Name: "new.name",
	}

	// Setup mocks
	mockErr := fmt.Errorf("mock.error")
	suite.mockRepository.On("FindByID", entity.ID(101)).Return(nil, mockErr)

	// Setup expectations
	expectedErr := "could not update inventory item - repository error: mock.error"

	// Exercise SUT
	err := suite.sut.Update(idFixture, voFixture)

	// Verify results
	suite.EqualError(err, expectedErr)
}

func (suite *ServiceImplTestSuite) TestUpdate_WhenModifierFails_ShouldFail() {
	// Setup fixture
	idFixture := entity.ID(101)
	voFixture := &inventory.UpdateItemVO{
		Name: "new.name",
	}

	// Setup mocks
	mockErr := fmt.Errorf("mock.error")
	mockEntity := &entityMocks.MockInventoryItem{Data: "mock.data"}
	suite.mockRepository.On("FindByID", entity.ID(101)).Return(mockEntity, nil)
	suite.mockEntityModifier.On("ModifyWithUpdateItemVO", mockEntity, voFixture).Return(mockErr)

	// Setup expectations
	expectedErr := "could not update inventory item - modifier error: mock.error"

	// Exercise SUT
	err := suite.sut.Update(idFixture, voFixture)

	// Verify results
	suite.EqualError(err, expectedErr)
}

func (suite *ServiceImplTestSuite) TestUpdate_WhenRepositoryUpdateFails_ShouldFail() {
	// Setup fixture
	idFixture := entity.ID(101)
	voFixture := &inventory.UpdateItemVO{
		Name: "new.name",
	}

	// Setup expectations
	expectedErr := "could not update inventory item - repository error: mock.error"

	// Setup mocks
	mockEntity := &entityMocks.MockInventoryItem{Data: "mock.data"}
	mockErr := fmt.Errorf("mock.error")
	suite.mockRepository.On("FindByID", entity.ID(101)).Return(mockEntity, nil)
	suite.mockEntityModifier.On("ModifyWithUpdateItemVO", mockEntity, voFixture).Return(nil)
	suite.mockRepository.On("Update", mockEntity).Return(mockErr)

	// Exercise SUT
	err := suite.sut.Update(idFixture, voFixture)

	// Verify results
	suite.EqualError(err, expectedErr)
}

func (suite *ServiceImplTestSuite) TestUpdate_WhenDelegatesSucceed_ShouldReturnAsExpected() {
	// Setup fixture
	idFixture := entity.ID(101)
	voFixture := &inventory.UpdateItemVO{
		Name: "new.name",
	}

	// Setup mocks
	mockEntity := &entityMocks.MockInventoryItem{Data: "mock.data"}
	suite.mockRepository.On("FindByID", entity.ID(101)).Return(mockEntity, nil)
	suite.mockEntityModifier.On("ModifyWithUpdateItemVO", mockEntity, voFixture).Return(nil)
	suite.mockRepository.On("Update", mockEntity).Return(nil)

	// Exercise SUT
	err := suite.sut.Update(idFixture, voFixture)

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
	mockEntity := &entityMocks.MockInventoryItem{Data: "some.data"}
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

func (suite *ServiceImplTestSuite) TestCheckout_WhenRepositoryUpdateFails_ShouldFail() {
	// Setup fixture
	idFixture := entity.ID(101)

	// Setup mocks
	mockEntity := &entityMocks.MockInventoryItem{Data: "some.data"}
	mockErr := fmt.Errorf("mock.error")
	suite.mockRepository.On("FindByID", idFixture).Return(mockEntity, nil)
	mockEntity.On("Checkout").Return(nil)
	suite.mockRepository.On("Update", mockEntity).Return(mockErr)

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
	mockEntity1 := &entityMocks.MockInventoryItem{Data: "some.data.1"}
	suite.mockRepository.On("FindByID", idFixture).Return(mockEntity1, nil)
	mockEntity1.On("Checkout").Return(nil)
	suite.mockRepository.On("Update", mockEntity1).Return(nil)

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
	mockEntity := &entityMocks.MockInventoryItem{Data: "some.data"}
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

func (suite *ServiceImplTestSuite) TestCheckIn_WhenRepositoryUpdateFails_ShouldFail() {
	// Setup fixture
	idFixture := entity.ID(101)

	// Setup mocks
	mockEntity := &entityMocks.MockInventoryItem{Data: "some.data"}
	mockErr := fmt.Errorf("mock.error")
	suite.mockRepository.On("FindByID", idFixture).Return(mockEntity, nil)
	mockEntity.On("CheckIn").Return(nil)
	suite.mockRepository.On("Update", mockEntity).Return(mockErr)

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
	mockEntity := &entityMocks.MockInventoryItem{Data: "some.data"}
	suite.mockRepository.On("FindByID", idFixture).Return(mockEntity, nil)
	mockEntity.On("CheckIn").Return(nil)
	suite.mockRepository.On("Update", mockEntity).Return(nil)

	// Exercise SUT
	err := suite.sut.CheckIn(idFixture)

	// Verify results
	suite.NoError(err)
}
