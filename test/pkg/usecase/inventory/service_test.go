package inventory_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"

	inventoryMocks "github.com/liampulles/matchstick-video/test/mock/pkg/usecase/inventory"

	"github.com/liampulles/matchstick-video/pkg/domain/entity"
	"github.com/liampulles/matchstick-video/pkg/usecase/inventory"
)

var nilInventoryItem *entity.InventoryItem = nil

type ServiceImplTestSuite struct {
	suite.Suite
	mockValidator     *inventoryMocks.MockValidator
	mockRepository    *inventoryMocks.MockRepository
	mockEntityFactory *inventoryMocks.MockEntityFactory
	sut               *inventory.ServiceImpl
}

func TestServiceImplTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceImplTestSuite))
}

func (suite *ServiceImplTestSuite) SetupTest() {
	suite.mockValidator = &inventoryMocks.MockValidator{}
	suite.mockRepository = &inventoryMocks.MockRepository{}
	suite.mockEntityFactory = &inventoryMocks.MockEntityFactory{}
	suite.sut = inventory.NewServiceImpl(
		suite.mockValidator, suite.mockRepository, suite.mockEntityFactory)
}

func (suite *ServiceImplTestSuite) TestCreate_WhenValidatorFails_ShouldFail() {
	// Setup fixture
	voFixture := &inventory.CreateItemVO{
		Name: "some.name",
	}

	// Setup mocks
	mockErr := fmt.Errorf("mock.error")
	suite.mockValidator.On("ValidateCreateItemVO", voFixture).Return(mockErr)

	// Setup expectations
	expectedErr := "could not create inventory item - validation error: mock.error"

	// Exercise SUT
	actual, err := suite.sut.Create(voFixture)

	// Verify results
	suite.Nil(actual)
	suite.EqualError(err, expectedErr)
}

func (suite *ServiceImplTestSuite) TestCreate_WhenFactoryFails_ShouldFail() {
	// Setup fixture
	voFixture := &inventory.CreateItemVO{
		Name: "some.name",
	}

	// Setup mocks
	mockErr := fmt.Errorf("mock.error")
	suite.mockValidator.On("ValidateCreateItemVO", voFixture).Return(nil)
	suite.mockEntityFactory.On("CreateFromVO", voFixture).Return(nilInventoryItem, mockErr)

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
	mockEntity := &entity.InventoryItem{Name: "mock.name"}
	mockErr := fmt.Errorf("mock.error")
	suite.mockValidator.On("ValidateCreateItemVO", voFixture).Return(nil)
	suite.mockEntityFactory.On("CreateFromVO", voFixture).Return(mockEntity, nil)
	suite.mockRepository.On("Save", mockEntity).Return(nilInventoryItem, mockErr)

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
	mockEntity1 := &entity.InventoryItem{Name: "mock.name.1"}
	mockEntity2 := &entity.InventoryItem{Name: "mock.name.2"}
	suite.mockValidator.On("ValidateCreateItemVO", voFixture).Return(nil)
	suite.mockEntityFactory.On("CreateFromVO", voFixture).Return(mockEntity1, nil)
	suite.mockRepository.On("Save", mockEntity1).Return(mockEntity2, nil)

	// Exercise SUT
	actual, err := suite.sut.Create(voFixture)

	// Verify results
	suite.Nil(err)
	suite.Equal(actual, mockEntity2)
}

func (suite *ServiceImplTestSuite) TestRead_WhenRepositoryFails_ShouldFail() {
	// Setup fixture
	idFixture := entity.ID(101)

	// Setup mocks
	mockErr := fmt.Errorf("mock.error")
	suite.mockRepository.On("FindByID", idFixture).Return(nilInventoryItem, mockErr)

	// Setup expectations
	expectedErr := "could not read inventory item - repository error: mock.error"

	// Exercise SUT
	actual, err := suite.sut.Read(idFixture)

	// Verify results
	suite.Nil(actual)
	suite.EqualError(err, expectedErr)
}

func (suite *ServiceImplTestSuite) TestRead_WhenDelegatesSucceed_ShouldReturnAsExcpected() {
	// Setup fixture
	idFixture := entity.ID(101)

	// Setup mocks
	mockEntity := &entity.InventoryItem{Name: "mock.name"}
	suite.mockRepository.On("FindByID", idFixture).Return(mockEntity, nil)

	// Exercise SUT
	actual, err := suite.sut.Read(idFixture)

	// Verify results
	suite.Nil(err)
	suite.Equal(actual, mockEntity)
}
