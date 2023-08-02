package sql_test

import (
	goSql "database/sql"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"

	entityMocks "github.com/liampulles/matchstick-video/test/mock/pkg/domain/entity"

	"github.com/liampulles/matchstick-video/pkg/adapter/db/sql"
	"github.com/liampulles/matchstick-video/pkg/domain/entity"
)

type InventoryRepositoryTestSuite struct {
	suite.Suite
	db     *goSql.DB
	mockDb sqlmock.Sqlmock
	sut    *sql.InventoryRepositoryImpl
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
	suite.sut = sql.NewInventoryRepositoryImpl()
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
	sql.SingleRowQuery = func(query string, scanFunc sql.ScanFunc, _type string, args ...interface{}) error {
		suite.Equal(expectedSql, query)
		suite.Equal(_type, "inventory item")
		suite.Equal(args[0], idFixture)
		return mockErr
	}

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
	sql.ManyRowsQuery = func(query string, scanFunc sql.ScanFunc, _type string, args ...interface{}) error {
		suite.Equal(expectedSql, query)
		suite.Equal(_type, "inventory item")
		return mockErr
	}

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
	sql.ManyRowsQuery = func(query string, scanFunc sql.ScanFunc, _type string, args ...interface{}) error {
		suite.Equal(expectedSql, query)
		suite.Equal(_type, "inventory item")
		return nil
	}

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
	mockEntity.On("Name").Return("some.name").
		On("Location").Return("some.location").
		On("IsAvailable").Return(true)
	sql.SingleQueryForID = func(query, _type string, args ...interface{}) (entity.ID, error) {
		suite.Equal(expectedSql, query)
		suite.Equal(_type, "inventory item")
		suite.Equal("some.name", args[0])
		suite.Equal("some.location", args[1])
		suite.Equal(true, args[2])
		return entity.InvalidID, mockErr
	}

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
	mockEntity.On("Name").Return("some.name").
		On("Location").Return("some.location").
		On("IsAvailable").Return(true)
	sql.SingleQueryForID = func(query, _type string, args ...interface{}) (entity.ID, error) {
		suite.Equal(expectedSql, query)
		suite.Equal(_type, "inventory item")
		suite.Equal("some.name", args[0])
		suite.Equal("some.location", args[1])
		suite.Equal(true, args[2])
		return expectedID, nil
	}

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
	sql.ExecForSingleItem = func(query string, args ...interface{}) error {
		suite.Equal(expectedSql, query)
		suite.Equal(idFixture, args[0])
		return mockErr
	}

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
	sql.ExecForSingleItem = func(query string, args ...interface{}) error {
		suite.Equal(expectedSql, query)
		suite.Equal(idFixture, args[0])
		return nil
	}

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
	mockEntity.On("ID").Return(entity.ID(101)).
		On("Name").Return("some.name").
		On("Location").Return("some.location").
		On("IsAvailable").Return(true)
	sql.ExecForSingleItem = func(query string, args ...interface{}) error {
		suite.Equal(expectedSql, query)
		suite.Equal("some.name", args[0])
		suite.Equal("some.location", args[1])
		suite.Equal(true, args[2])
		suite.Equal(entity.ID(101), args[3])
		return mockErr
	}

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
	mockEntity.On("ID").Return(entity.ID(101)).
		On("Name").Return("some.name").
		On("Location").Return("some.location").
		On("IsAvailable").Return(true)
	sql.ExecForSingleItem = func(query string, args ...interface{}) error {
		suite.Equal(expectedSql, query)
		suite.Equal("some.name", args[0])
		suite.Equal("some.location", args[1])
		suite.Equal(true, args[2])
		suite.Equal(entity.ID(101), args[3])
		return nil
	}

	// Exercise SUT
	err := suite.sut.Update(mockEntity)

	// Verify results
	suite.NoError(err)
}
