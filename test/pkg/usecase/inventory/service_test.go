package inventory_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"

	entityMocks "github.com/liampulles/matchstick-video/test/mock/pkg/domain/entity"
	inventoryMocks "github.com/liampulles/matchstick-video/test/mock/pkg/usecase/inventory"

	"github.com/liampulles/matchstick-video/pkg/adapter/db/sql"
	"github.com/liampulles/matchstick-video/pkg/domain/entity"
	"github.com/liampulles/matchstick-video/pkg/usecase/inventory"
)

type ServiceImplTestSuite struct {
	suite.Suite
	mockEntityFactory  *inventoryMocks.MockEntityFactory
	mockEntityModifier *inventoryMocks.MockEntityModifier
	mockVoFactory      *inventoryMocks.MockVOFactory
	sut                *inventory.ServiceImpl
}

func TestServiceImplTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceImplTestSuite))
}

func (suite *ServiceImplTestSuite) SetupTest() {
	suite.mockEntityFactory = &inventoryMocks.MockEntityFactory{}
	suite.mockEntityModifier = &inventoryMocks.MockEntityModifier{}
	suite.mockVoFactory = &inventoryMocks.MockVOFactory{}
	suite.sut = inventory.NewServiceImpl(
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
	sql.Create = func(e entity.InventoryItem) (entity.ID, error) {
		suite.Equal(mockEntity, e)
		return entity.InvalidID, mockErr
	}

	// Setup expectations
	expectedErr := "could not create inventory item - repository create error: mock.error"

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
	sql.Create = func(e entity.InventoryItem) (entity.ID, error) {
		suite.Equal(mockEntity, e)
		return expected, nil
	}

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
	sql.FindByID = func(id entity.ID) (entity.InventoryItem, error) {
		suite.Equal(idFixture, id)
		return nil, mockErr
	}

	// Setup expectations
	expectedErr := "could not read inventory item - repository find error: mock.error"

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
	sql.FindByID = func(id entity.ID) (entity.InventoryItem, error) {
		suite.Equal(idFixture, id)
		return mockEntity, nil
	}
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
	sql.FindAll = func() ([]entity.InventoryItem, error) {
		return nil, mockErr
	}

	// Setup expectations
	expectedErr := "could not read inventory items - repository find error: mock.error"

	// Exercise SUT
	actual, err := suite.sut.ReadAll()

	// Verify results
	suite.Nil(actual)
	suite.EqualError(err, expectedErr)
}

func (suite *ServiceImplTestSuite) TestReadAll_WhenDelegatesSucceed_ShouldReturnAsExpected() {
	// Setup expectations
	expected := []inventory.ThinViewVO{
		inventory.ThinViewVO{
			Name: "some.name",
		},
	}

	// Setup mocks
	mockEntity := &entityMocks.MockInventoryItem{Data: "mock.data"}
	mockEntities := []entity.InventoryItem{mockEntity}
	sql.FindAll = func() ([]entity.InventoryItem, error) {
		return mockEntities, nil
	}
	suite.mockVoFactory.On("CreateThinViewVOsFromEntities", mockEntities).Return(expected)

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
	sql.FindByID = func(id entity.ID) (entity.InventoryItem, error) {
		suite.Equal(idFixture, id)
		return nil, mockErr
	}

	// Setup expectations
	expectedErr := "could not update inventory item - repository find error: mock.error"

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
	sql.FindByID = func(id entity.ID) (entity.InventoryItem, error) {
		suite.Equal(idFixture, id)
		return mockEntity, nil
	}
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
	expectedErr := "could not update inventory item - repository update error: mock.error"

	// Setup mocks
	mockEntity := &entityMocks.MockInventoryItem{Data: "mock.data"}
	mockErr := fmt.Errorf("mock.error")
	sql.FindByID = func(id entity.ID) (entity.InventoryItem, error) {
		suite.Equal(idFixture, id)
		return mockEntity, nil
	}
	suite.mockEntityModifier.On("ModifyWithUpdateItemVO", mockEntity, voFixture).Return(nil)
	sql.Update = func(e entity.InventoryItem) error {
		suite.Equal(mockEntity, e)
		return mockErr
	}

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
	sql.FindByID = func(id entity.ID) (entity.InventoryItem, error) {
		suite.Equal(idFixture, id)
		return mockEntity, nil
	}
	suite.mockEntityModifier.On("ModifyWithUpdateItemVO", mockEntity, voFixture).Return(nil)
	sql.Update = func(e entity.InventoryItem) error {
		suite.Equal(mockEntity, e)
		return nil
	}

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
	sql.DeleteByID = func(id entity.ID) error {
		suite.Equal(idFixture, id)
		return mockErr
	}

	// Setup expectations
	expectedErr := "could not delete inventory item - repository delete error: mock.error"

	// Exercise SUT
	err := suite.sut.Delete(idFixture)

	// Verify results
	suite.EqualError(err, expectedErr)
}

func (suite *ServiceImplTestSuite) TestDelete_WhenDelegatesSucceed_ShouldReturnAsExpected() {
	// Setup fixture
	idFixture := entity.ID(101)

	// Setup mocks
	sql.DeleteByID = func(id entity.ID) error {
		suite.Equal(idFixture, id)
		return nil
	}

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
	sql.FindByID = func(id entity.ID) (entity.InventoryItem, error) {
		suite.Equal(idFixture, id)
		return nil, mockErr
	}

	// Setup expectations
	expectedErr := "could not checkout inventory item - repository find error: mock.error"

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
	sql.FindByID = func(id entity.ID) (entity.InventoryItem, error) {
		suite.Equal(idFixture, id)
		return mockEntity, nil
	}
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
	sql.FindByID = func(id entity.ID) (entity.InventoryItem, error) {
		suite.Equal(idFixture, id)
		return mockEntity, nil
	}
	mockEntity.On("Checkout").Return(nil)
	sql.Update = func(e entity.InventoryItem) error {
		suite.Equal(mockEntity, e)
		return mockErr
	}

	// Setup expectations
	expectedErr := "could not checkout inventory item - repository update error: mock.error"

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
	sql.FindByID = func(id entity.ID) (entity.InventoryItem, error) {
		suite.Equal(idFixture, id)
		return mockEntity1, nil
	}
	mockEntity1.On("Checkout").Return(nil)
	sql.Update = func(e entity.InventoryItem) error {
		suite.Equal(mockEntity1, e)
		return nil
	}

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
	sql.FindByID = func(id entity.ID) (entity.InventoryItem, error) {
		suite.Equal(idFixture, id)
		return nil, mockErr
	}

	// Setup expectations
	expectedErr := "could not check in inventory item - repository find error: mock.error"

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
	sql.FindByID = func(id entity.ID) (entity.InventoryItem, error) {
		suite.Equal(idFixture, id)
		return mockEntity, nil
	}
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
	sql.FindByID = func(id entity.ID) (entity.InventoryItem, error) {
		suite.Equal(idFixture, id)
		return mockEntity, nil
	}
	mockEntity.On("CheckIn").Return(nil)
	sql.Update = func(e entity.InventoryItem) error {
		suite.Equal(mockEntity, e)
		return mockErr
	}

	// Setup expectations
	expectedErr := "could not check in inventory item - repository update error: mock.error"

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
	sql.FindByID = func(id entity.ID) (entity.InventoryItem, error) {
		suite.Equal(idFixture, id)
		return mockEntity, nil
	}
	mockEntity.On("CheckIn").Return(nil)
	sql.Update = func(e entity.InventoryItem) error {
		suite.Equal(mockEntity, e)
		return nil
	}

	// Exercise SUT
	err := suite.sut.CheckIn(idFixture)

	// Verify results
	suite.NoError(err)
}
