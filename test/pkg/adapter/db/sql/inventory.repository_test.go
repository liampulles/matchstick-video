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
	mockHelperService *sqlMocks.HelperServiceMock
	mockConstructor   *entityMocks.InventoryItemConstructorMock
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
	suite.mockHelperService = &sqlMocks.HelperServiceMock{}
	suite.mockConstructor = &entityMocks.InventoryItemConstructorMock{}
	suite.sut = sql.NewInventoryRepositoryImpl(
		db, suite.mockHelperService, suite.mockConstructor,
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
		id=@id;`
	expectedErr := "cannot execute query - db get row error: mock.error"

	// Setup mocks
	mockErr := fmt.Errorf("mock.error")
	suite.mockHelperService.
		On("SingleRowQuery", suite.db, expectedSql, goSql.Named("id", idFixture)).
		Return(nil, mockErr)

	// Exercise SUT
	actual, err := suite.sut.FindByID(idFixture)

	// Verify results
	suite.Nil(actual)
	suite.EqualError(err, expectedErr)
}

func (suite *InventoryRepositoryTestSuite) TestFindByID_WhenScanFails_ShouldFail() {
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
		id=@id;`
	expectedErr := "cannot execute query - db scan error: mock.error"

	// Setup mocks
	// -> missing element
	mockRow := &sqlMocks.RowMock{}
	mockErr := fmt.Errorf("mock.error")
	suite.mockHelperService.
		On("SingleRowQuery", suite.db, expectedSql, goSql.Named("id", idFixture)).
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
		id=@id;`

	// Setup mocks
	// -> missing element
	mockRow := &sqlMocks.RowMock{}
	mockEntity := &entityMocks.InventoryItemMock{Data: "some.data"}
	suite.mockHelperService.
		On("SingleRowQuery", suite.db, expectedSql, goSql.Named("id", idFixture)).
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
	VALUES 
		(
			@name, 
			@location, 
			@available
		);`
	expectedErr := "mock.error"

	// Setup mocks
	mockEntity := &entityMocks.InventoryItemMock{}
	mockErr := fmt.Errorf(expectedErr)
	mockEntity.On("Name").Return("some.name").
		On("Location").Return("some.location").
		On("IsAvailable").Return(true)
	suite.mockHelperService.On("ExecForID", suite.db, expectedSql,
		goSql.Named("name", "some.name"),
		goSql.Named("location", "some.location"),
		goSql.Named("available", true),
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
	VALUES 
		(
			@name, 
			@location, 
			@available
		);`
	expectedID := entity.ID(101)

	// Setup mocks
	mockEntity := &entityMocks.InventoryItemMock{}
	mockEntity.On("Name").Return("some.name").
		On("Location").Return("some.location").
		On("IsAvailable").Return(true)
	suite.mockHelperService.On("ExecForID", suite.db, expectedSql,
		goSql.Named("name", "some.name"),
		goSql.Named("location", "some.location"),
		goSql.Named("available", true),
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
		id=@id;`
	expectedErr := "mock.error"

	// Setup mocks
	mockErr := fmt.Errorf(expectedErr)
	suite.mockHelperService.On("ExecForError", suite.db, expectedSql,
		goSql.Named("id", idFixture),
	).Return(mockErr)

	// Exercise SUT
	err := suite.sut.DeleteByID(idFixture)

	// Verify results
	suite.EqualError(err, expectedErr)
}

func (suite *InventoryRepositoryTestSuite) TestDeleteByID_WhenHelperServiceSucceeds_ShouldPass() {
	// Setup fixture
	idFixture := entity.ID(101)

	// Setup expectations
	expectedSql := `
	DELETE FROM inventory_item
	WHERE 
		id=@id;`

	// Setup mocks
	suite.mockHelperService.On("ExecForError", suite.db, expectedSql,
		goSql.Named("id", idFixture),
	).Return(nil)

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
		name=@name, location=@location, available=@available
	WHERE 
		id=@id;`
	expectedErr := "mock.error"

	// Setup mocks
	mockEntity := &entityMocks.InventoryItemMock{}
	mockErr := fmt.Errorf(expectedErr)
	mockEntity.On("ID").Return(entity.ID(101)).
		On("Name").Return("some.name").
		On("Location").Return("some.location").
		On("IsAvailable").Return(true)
	suite.mockHelperService.On("ExecForError", suite.db, expectedSql,
		goSql.Named("name", "some.name"),
		goSql.Named("location", "some.location"),
		goSql.Named("available", true),
		goSql.Named("id", entity.ID(101)),
	).Return(mockErr)

	// Exercise SUT
	err := suite.sut.Update(mockEntity)

	// Verify results
	suite.EqualError(err, expectedErr)
}

func (suite *InventoryRepositoryTestSuite) TestUpdate_WhenHelperServiceSucceeds_ShouldPass() {
	// Setup expectations
	expectedSql := `
	UPDATE inventory_item
	SET
		name=@name, location=@location, available=@available
	WHERE 
		id=@id;`

	// Setup mocks
	mockEntity := &entityMocks.InventoryItemMock{}
	mockEntity.On("ID").Return(entity.ID(101)).
		On("Name").Return("some.name").
		On("Location").Return("some.location").
		On("IsAvailable").Return(true)
	suite.mockHelperService.On("ExecForError", suite.db, expectedSql,
		goSql.Named("name", "some.name"),
		goSql.Named("location", "some.location"),
		goSql.Named("available", true),
		goSql.Named("id", entity.ID(101)),
	).Return(nil)

	// Exercise SUT
	err := suite.sut.Update(mockEntity)

	// Verify results
	suite.NoError(err)
}