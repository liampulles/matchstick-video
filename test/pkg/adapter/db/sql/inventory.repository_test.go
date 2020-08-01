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
	expectedErr := "cannot execute query - db get row error: mock.error"

	// Setup mocks
	mockErr := fmt.Errorf("mock.error")
	suite.mockDbService.On("Get").Return(suite.db)
	suite.mockHelperService.
		On("SingleRowQuery", suite.db, expectedSql, idFixture).
		Return(nil, mockErr)

	// Exercise SUT
	actual, err := suite.sut.FindByID(idFixture)

	// Verify results
	suite.Nil(actual)
	suite.EqualError(err, expectedErr)
}

func (suite *InventoryRepositoryTestSuite) TestFindByID_WhenScanFails_ShouldFailWithNotFound() {
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
	expectedErr := "entity not found: type=[inventory item]"

	// Setup mocks
	// -> missing element
	mockRow := &sqlMocks.RowMock{}
	mockErr := fmt.Errorf("mock.error")
	suite.mockDbService.On("Get").Return(suite.db)
	suite.mockHelperService.
		On("SingleRowQuery", suite.db, expectedSql, idFixture).
		Return(mockRow, nil)
	mockRow.
		On("Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(mockErr)
	// Exercise SUT
	actual, err := suite.sut.FindByID(idFixture)

	// Verify results
	suite.Nil(actual)
	suite.EqualError(err, expectedErr)
}

func (suite *InventoryRepositoryTestSuite) TestFindByID_WhenDBSucceeds_ShouldReturnEntity() {
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

	// Setup mocks
	// -> missing element
	mockRow := &sqlMocks.RowMock{}
	mockEntity := &entityMocks.MockInventoryItem{Data: "some.data"}
	suite.mockDbService.On("Get").Return(suite.db)
	suite.mockHelperService.
		On("SingleRowQuery", suite.db, expectedSql, idFixture).
		Return(mockRow, nil)
	mockRow.
		On("Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(nil)
	suite.mockConstructor.
		On("Reincarnate", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(mockEntity)
	// Exercise SUT
	actual, err := suite.sut.FindByID(idFixture)

	// Verify results
	suite.NoError(err)
	suite.Equal(mockEntity, actual)
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
	suite.mockHelperService.On("SingleQueryForID", suite.db, expectedSql,
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
	suite.mockHelperService.On("SingleQueryForID", suite.db, expectedSql,
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
	expectedErr := "cannot execute exec - db exec error: mock.error"

	// Setup mocks
	mockErr := fmt.Errorf("mock.error")
	suite.mockDbService.On("Get").Return(suite.db)
	suite.mockHelperService.On("ExecForRowsAffected", suite.db, expectedSql, idFixture).
		Return(int64(-1), mockErr)

	// Exercise SUT
	err := suite.sut.DeleteByID(idFixture)

	// Verify results
	suite.EqualError(err, expectedErr)
}

func (suite *InventoryRepositoryTestSuite) TestDeleteByID_WhenNoRowsAffected_ShouldFail() {
	// Setup fixture
	idFixture := entity.ID(101)

	// Setup expectations
	expectedSql := `
	DELETE FROM inventory_item
	WHERE 
		id=$1;`
	expectedErr := "entity not found: type=[inventory item]"

	// Setup mocks
	suite.mockDbService.On("Get").Return(suite.db)
	suite.mockHelperService.On("ExecForRowsAffected", suite.db, expectedSql, idFixture).
		Return(int64(0), nil)

	// Exercise SUT
	err := suite.sut.DeleteByID(idFixture)

	// Verify results
	suite.EqualError(err, expectedErr)
}

func (suite *InventoryRepositoryTestSuite) TestDeleteByID_WhenMoreThanOneRowAffected_ShouldFail() {
	// Setup fixture
	idFixture := entity.ID(101)

	// Setup expectations
	expectedSql := `
	DELETE FROM inventory_item
	WHERE 
		id=$1;`
	expectedErr := "exec error: expected 1 entity to be affected, but was: 2"

	// Setup mocks
	suite.mockDbService.On("Get").Return(suite.db)
	suite.mockHelperService.On("ExecForRowsAffected", suite.db, expectedSql, idFixture).
		Return(int64(2), nil)

	// Exercise SUT
	err := suite.sut.DeleteByID(idFixture)

	// Verify results
	suite.EqualError(err, expectedErr)
}

func (suite *InventoryRepositoryTestSuite) TestDeleteByID_WhenOneRowAffected_ShouldPass() {
	// Setup fixture
	idFixture := entity.ID(101)

	// Setup expectations
	expectedSql := `
	DELETE FROM inventory_item
	WHERE 
		id=$1;`

	// Setup mocks
	suite.mockDbService.On("Get").Return(suite.db)
	suite.mockHelperService.On("ExecForRowsAffected", suite.db, expectedSql, idFixture).
		Return(int64(1), nil)

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
	expectedErr := "cannot execute exec - db exec error: mock.error"

	// Setup mocks
	mockEntity := &entityMocks.MockInventoryItem{}
	mockErr := fmt.Errorf("mock.error")
	suite.mockDbService.On("Get").Return(suite.db)
	mockEntity.On("ID").Return(entity.ID(101)).
		On("Name").Return("some.name").
		On("Location").Return("some.location").
		On("IsAvailable").Return(true)
	suite.mockHelperService.On("ExecForRowsAffected", suite.db, expectedSql,
		"some.name",
		"some.location",
		true,
		entity.ID(101),
	).Return(int64(-1), mockErr)

	// Exercise SUT
	err := suite.sut.Update(mockEntity)

	// Verify results
	suite.EqualError(err, expectedErr)
}

func (suite *InventoryRepositoryTestSuite) TestUpdate_WhenNoRowsAffected_ShouldFail() {
	// Setup expectations
	expectedSql := `
	UPDATE inventory_item
	SET
		name=$1, location=$2, available=$3
	WHERE 
		id=$4;`
	expectedErr := "entity not found: type=[inventory item]"

	// Setup mocks
	mockEntity := &entityMocks.MockInventoryItem{}
	suite.mockDbService.On("Get").Return(suite.db)
	mockEntity.On("ID").Return(entity.ID(101)).
		On("Name").Return("some.name").
		On("Location").Return("some.location").
		On("IsAvailable").Return(true)
	suite.mockHelperService.On("ExecForRowsAffected", suite.db, expectedSql,
		"some.name",
		"some.location",
		true,
		entity.ID(101),
	).Return(int64(0), nil)

	// Exercise SUT
	err := suite.sut.Update(mockEntity)

	// Verify results
	suite.EqualError(err, expectedErr)
}

func (suite *InventoryRepositoryTestSuite) TestUpdate_WhenMoreThanOneRowAffected_ShouldFail() {
	// Setup expectations
	expectedSql := `
	UPDATE inventory_item
	SET
		name=$1, location=$2, available=$3
	WHERE 
		id=$4;`
	expectedErr := "exec error: expected 1 entity to be affected, but was: 2"

	// Setup mocks
	mockEntity := &entityMocks.MockInventoryItem{}
	suite.mockDbService.On("Get").Return(suite.db)
	mockEntity.On("ID").Return(entity.ID(101)).
		On("Name").Return("some.name").
		On("Location").Return("some.location").
		On("IsAvailable").Return(true)
	suite.mockHelperService.On("ExecForRowsAffected", suite.db, expectedSql,
		"some.name",
		"some.location",
		true,
		entity.ID(101),
	).Return(int64(2), nil)

	// Exercise SUT
	err := suite.sut.Update(mockEntity)

	// Verify results
	suite.EqualError(err, expectedErr)
}

func (suite *InventoryRepositoryTestSuite) TestUpdate_WhenOneRowAffected_ShouldPass() {
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
	suite.mockHelperService.On("ExecForRowsAffected", suite.db, expectedSql,
		"some.name",
		"some.location",
		true,
		entity.ID(101),
	).Return(int64(1), nil)

	// Exercise SUT
	err := suite.sut.Update(mockEntity)

	// Verify results
	suite.NoError(err)
}
