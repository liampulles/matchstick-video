package sql_test

import (
	goSql "database/sql"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	sqlMocks "github.com/liampulles/matchstick-video/test/mock/pkg/adapter/db/sql"
	entityMocks "github.com/liampulles/matchstick-video/test/mock/pkg/domain/entity"

	"github.com/liampulles/matchstick-video/pkg/adapter/db/sql"
	"github.com/liampulles/matchstick-video/pkg/domain/entity"
)

type InventoryRepositoryTestSuite struct {
	suite.Suite
	db                *goSql.DB
	mockDb            sqlmock.Sqlmock
	mockDbService     *sqlMocks.MockDatabaseStore
	mockHelperService *sqlMocks.MockHelperService
	mockConstructor   *entityMocks.MockInventoryItemConstructor
	sut               *sql.InventoryRepositoryImpl
}

func TestInventoryRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(InventoryRepositoryTestSuite))
}

func (suite *InventoryRepositoryTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}
	suite.db = db
	suite.mockDb = mock
	suite.mockDbService = &sqlMocks.MockDatabaseStore{}
	suite.mockHelperService = &sqlMocks.MockHelperService{}
	suite.mockConstructor = &entityMocks.MockInventoryItemConstructor{}
	suite.sut = sql.NewInventoryRepositoryImpl(
		suite.mockDbService, suite.mockHelperService, suite.mockConstructor,
	)
}

func (suite *InventoryRepositoryTestSuite) TestFindByID_WhenHelperServiceFails_ShouldFail() {
	// Setup fixture
	idFixture := entity.ID(101)

	// Setup expectations
	expectedSql := `
	SELECT 
		id, 
		name, 
		location, 
		available 
	FROM inventory_item
	WHERE 
		id=$1;`
	expectedErr := "mock.error"

	// Setup mocks
	mockErr := fmt.Errorf("mock.error")
	suite.mockDbService.On("Get").Return(suite.db)
	suite.mockHelperService.
		On("SingleRowQuery", suite.db, expectedSql, mock.Anything, "inventory item", idFixture).
		Return(mockErr)

	// Exercise SUT
	_, err := suite.sut.FindByID(idFixture)

	// Verify results
	suite.EqualError(err, expectedErr)
}

func (suite *InventoryRepositoryTestSuite) TestFindAll_WhenHelperServiceFails_ShouldFail() {
	// Setup expectations
	expectedSql := `
	SELECT 
		id, 
		name, 
		location, 
		available 
	FROM inventory_item;`
	expectedErr := "mock.error"

	// Setup mocks
	mockErr := fmt.Errorf("mock.error")
	suite.mockDbService.On("Get").Return(suite.db)
	suite.mockHelperService.
		On("ManyRowsQuery", suite.db, expectedSql, mock.Anything, "inventory item").
		Return(mockErr)

	// Exercise SUT
	actual, err := suite.sut.FindAll()

	// Verify results
	suite.Nil(actual)
	suite.EqualError(err, expectedErr)
}

func (suite *InventoryRepositoryTestSuite) TestFindAll_WhenHelperServicePasses_ShouldPass() {
	// Setup expectations
	expectedSql := `
	SELECT 
		id, 
		name, 
		location, 
		available 
	FROM inventory_item;`

	// Setup mocks
	suite.mockDbService.On("Get").Return(suite.db)
	suite.mockHelperService.
		On("ManyRowsQuery", suite.db, expectedSql, mock.Anything, "inventory item").
		Return(nil)

	// Exercise SUT
	_, err := suite.sut.FindAll()

	// Verify results
	suite.NoError(err)
}

func (suite *InventoryRepositoryTestSuite) TestCreate_WhenHelperServiceFails_ShouldFail() {
	// Setup expectations
	expectedSql := `
	INSERT INTO inventory_item
		(
			name, 
			location, 
			available
		)
	VALUES ($1, $2, $3)
	RETURNING id;`
	expectedErr := "mock.error"

	// Setup mocks
	mockEntity := &entityMocks.MockInventoryItem{}
	mockErr := fmt.Errorf(expectedErr)
	suite.mockDbService.On("Get").Return(suite.db)
	mockEntity.On("Name").Return("some.name").
		On("Location").Return("some.location").
		On("IsAvailable").Return(true)
	suite.mockHelperService.On("SingleQueryForID", suite.db, expectedSql, "inventory item",
		"some.name",
		"some.location",
		true,
	).Return(entity.InvalidID, mockErr)

	// Exercise SUT
	actual, err := suite.sut.Create(mockEntity)

	// Verify results
	suite.Equal(actual, entity.InvalidID)
	suite.EqualError(err, expectedErr)
}

func (suite *InventoryRepositoryTestSuite) TestCreate_WhenHelperServiceSucceeds_ShouldReturnID() {
	// Setup expectations
	expectedSql := `
	INSERT INTO inventory_item
		(
			name, 
			location, 
			available
		)
	VALUES ($1, $2, $3)
	RETURNING id;`
	expectedID := entity.ID(101)

	// Setup mocks
	mockEntity := &entityMocks.MockInventoryItem{}
	suite.mockDbService.On("Get").Return(suite.db)
	mockEntity.On("Name").Return("some.name").
		On("Location").Return("some.location").
		On("IsAvailable").Return(true)
	suite.mockHelperService.On("SingleQueryForID", suite.db, expectedSql, "inventory item",
		"some.name",
		"some.location",
		true,
	).Return(expectedID, nil)

	// Exercise SUT
	actual, err := suite.sut.Create(mockEntity)

	// Verify results
	suite.NoError(err)
	suite.Equal(expectedID, actual)
}

func (suite *InventoryRepositoryTestSuite) TestDeleteByID_WhenHelperServiceFails_ShouldFail() {
	// Setup fixture
	idFixture := entity.ID(101)

	// Setup expectations
	expectedSql := `
	DELETE FROM inventory_item
	WHERE 
		id=$1;`
	expectedErr := "mock.error"

	// Setup mocks
	mockErr := fmt.Errorf("mock.error")
	suite.mockDbService.On("Get").Return(suite.db)
	suite.mockHelperService.On("ExecForSingleItem", suite.db, expectedSql, idFixture).
		Return(mockErr)

	// Exercise SUT
	err := suite.sut.DeleteByID(idFixture)

	// Verify results
	suite.EqualError(err, expectedErr)
}

func (suite *InventoryRepositoryTestSuite) TestDeleteByID_WhenHelperServicePasses_ShouldPass() {
	// Setup fixture
	idFixture := entity.ID(101)

	// Setup expectations
	expectedSql := `
	DELETE FROM inventory_item
	WHERE 
		id=$1;`

	// Setup mocks
	suite.mockDbService.On("Get").Return(suite.db)
	suite.mockHelperService.On("ExecForSingleItem", suite.db, expectedSql, idFixture).
		Return(nil)

	// Exercise SUT
	err := suite.sut.DeleteByID(idFixture)

	// Verify results
	suite.NoError(err)
}

func (suite *InventoryRepositoryTestSuite) TestUpdate_WhenHelperServiceFails_ShouldFail() {
	// Setup expectations
	expectedSql := `
	UPDATE inventory_item
	SET
		name=$1, location=$2, available=$3
	WHERE 
		id=$4;`
	expectedErr := "mock.error"

	// Setup mocks
	mockEntity := &entityMocks.MockInventoryItem{}
	mockErr := fmt.Errorf("mock.error")
	suite.mockDbService.On("Get").Return(suite.db)
	mockEntity.On("ID").Return(entity.ID(101)).
		On("Name").Return("some.name").
		On("Location").Return("some.location").
		On("IsAvailable").Return(true)
	suite.mockHelperService.On("ExecForSingleItem", suite.db, expectedSql,
		"some.name",
		"some.location",
		true,
		entity.ID(101),
	).Return(mockErr)

	// Exercise SUT
	err := suite.sut.Update(mockEntity)

	// Verify results
	suite.EqualError(err, expectedErr)
}

func (suite *InventoryRepositoryTestSuite) TestUpdate_WhenHelperServicePasses_ShouldPass() {
	// Setup expectations
	expectedSql := `
	UPDATE inventory_item
	SET
		name=$1, location=$2, available=$3
	WHERE 
		id=$4;`

	// Setup mocks
	mockEntity := &entityMocks.MockInventoryItem{}
	suite.mockDbService.On("Get").Return(suite.db)
	mockEntity.On("ID").Return(entity.ID(101)).
		On("Name").Return("some.name").
		On("Location").Return("some.location").
		On("IsAvailable").Return(true)
	suite.mockHelperService.On("ExecForSingleItem", suite.db, expectedSql,
		"some.name",
		"some.location",
		true,
		entity.ID(101),
	).Return(nil)

	// Exercise SUT
	err := suite.sut.Update(mockEntity)

	// Verify results
	suite.NoError(err)
}
