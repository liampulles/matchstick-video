package inventory_test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	adapterMocks "github.com/liampulles/matchstick-video/test/mock/pkg/adapter"
	entityMocks "github.com/liampulles/matchstick-video/test/mock/pkg/domain/entity"

	"github.com/liampulles/matchstick-video/pkg/adapter/inventory"
	"github.com/liampulles/matchstick-video/pkg/domain/entity"
)

type SQLRepositoryTestSuite struct {
	suite.Suite
	db                *sql.DB
	mockDb            sqlmock.Sqlmock
	mockHelperService *adapterMocks.SQLDbHelperServiceMock
	mockConstructor   *entityMocks.InventoryItemConstructorMock
	sut               *inventory.SQLRepositoryImpl
}

func TestSQLRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(SQLRepositoryTestSuite))
}

func (suite *SQLRepositoryTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}
	suite.db = db
	suite.mockDb = mock
	suite.mockHelperService = &adapterMocks.SQLDbHelperServiceMock{}
	suite.mockConstructor = &entityMocks.InventoryItemConstructorMock{}
	suite.sut = inventory.NewSQLRepositoryImpl(
		db, suite.mockHelperService, suite.mockConstructor,
	)
}

func (suite *SQLRepositoryTestSuite) TestFindByID_WhenHelperServiceFails_ShouldFail() {
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
		On("SingleRowQuery", suite.db, expectedSql, sql.Named("id", idFixture)).
		Return(nil, mockErr)

	// Exercise SUT
	actual, err := suite.sut.FindByID(idFixture)

	// Verify results
	suite.Nil(actual)
	suite.EqualError(err, expectedErr)
}

func (suite *SQLRepositoryTestSuite) TestFindByID_WhenScanFails_ShouldFail() {
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
	mockRow := &adapterMocks.SQLRowMock{}
	mockErr := fmt.Errorf("mock.error")
	suite.mockHelperService.
		On("SingleRowQuery", suite.db, expectedSql, sql.Named("id", idFixture)).
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

func (suite *SQLRepositoryTestSuite) TestFindByID_WhenDBSucceeds_ShouldReturnEntity() {
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
	mockRow := &adapterMocks.SQLRowMock{}
	mockEntity := &entityMocks.InventoryItemMock{Data: "some.data"}
	suite.mockHelperService.
		On("SingleRowQuery", suite.db, expectedSql, sql.Named("id", idFixture)).
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

func (suite *SQLRepositoryTestSuite) TestCreate_WhenHelperServiceFails_ShouldFail() {
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
		sql.Named("name", "some.name"),
		sql.Named("location", "some.location"),
		sql.Named("available", true),
	).Return(entity.InvalidID, mockErr)

	// Exercise SUT
	actual, err := suite.sut.Create(mockEntity)

	// Verify results
	suite.Equal(actual, entity.InvalidID)
	suite.EqualError(err, expectedErr)
}

func (suite *SQLRepositoryTestSuite) TestCreate_WhenHelperServiceSucceeds_ShouldReturnID() {
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
		sql.Named("name", "some.name"),
		sql.Named("location", "some.location"),
		sql.Named("available", true),
	).Return(expectedID, nil)

	// Exercise SUT
	actual, err := suite.sut.Create(mockEntity)

	// Verify results
	suite.NoError(err)
	suite.Equal(expectedID, actual)
}

func (suite *SQLRepositoryTestSuite) TestDeleteByID_WhenHelperServiceFails_ShouldFail() {
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
		sql.Named("id", idFixture),
	).Return(mockErr)

	// Exercise SUT
	err := suite.sut.DeleteByID(idFixture)

	// Verify results
	suite.EqualError(err, expectedErr)
}

func (suite *SQLRepositoryTestSuite) TestDeleteByID_WhenHelperServiceSucceeds_ShouldPass() {
	// Setup fixture
	idFixture := entity.ID(101)

	// Setup expectations
	expectedSql := `
	DELETE FROM inventory_item
	WHERE 
		id=@id;`

	// Setup mocks
	suite.mockHelperService.On("ExecForError", suite.db, expectedSql,
		sql.Named("id", idFixture),
	).Return(nil)

	// Exercise SUT
	err := suite.sut.DeleteByID(idFixture)

	// Verify results
	suite.NoError(err)
}

func (suite *SQLRepositoryTestSuite) TestUpdate_WhenHelperServiceFails_ShouldFail() {
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
		sql.Named("name", "some.name"),
		sql.Named("location", "some.location"),
		sql.Named("available", true),
		sql.Named("id", entity.ID(101)),
	).Return(mockErr)

	// Exercise SUT
	err := suite.sut.Update(mockEntity)

	// Verify results
	suite.EqualError(err, expectedErr)
}

func (suite *SQLRepositoryTestSuite) TestUpdate_WhenHelperServiceSucceeds_ShouldPass() {
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
		sql.Named("name", "some.name"),
		sql.Named("location", "some.location"),
		sql.Named("available", true),
		sql.Named("id", entity.ID(101)),
	).Return(nil)

	// Exercise SUT
	err := suite.sut.Update(mockEntity)

	// Verify results
	suite.NoError(err)
}
