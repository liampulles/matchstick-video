//go:build integration
// +build integration

package integration_test

import (
	"testing"

	goConfig "github.com/liampulles/go-config"

	"github.com/stretchr/testify/suite"

	"github.com/liampulles/matchstick-video/pkg/adapter/config"
	"github.com/liampulles/matchstick-video/pkg/adapter/db/sql/inventory"
	"github.com/liampulles/matchstick-video/pkg/domain/entity"
)

type InventoryRepositoryTestSuite struct {
	suite.Suite
}

func TestInventoryRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(InventoryRepositoryTestSuite))
}

func (suite *InventoryRepositoryTestSuite) SetupTest() {
	config.Load(goConfig.MapSource{
		"PORT":             "9010",
		"MIGRATION_SOURCE": "file://../../migrations",
		"DB_USER":          "integration",
		"DB_PASSWORD":      "integration",
		"DB_NAME":          "integration",
		"DB_PORT":          "5050",
	})
}

func (suite *InventoryRepositoryTestSuite) TestFindByID_WhenDoesExist_ShouldPass() {
	// Setup fixture
	e := entity.TestInventoryItemImplConstructor(
		entity.InvalidID, "some.find.name", "some.find.location", true,
	)
	id, err := inventory.Create(e)
	suite.NoError(err)

	// Exercise SUT
	_, err = inventory.FindByID(id)

	// Verify results
	suite.NoError(err)
}

func (suite *InventoryRepositoryTestSuite) TestFindAll_ShouldPass() {
	// Exercise SUT
	_, err := inventory.FindAll()

	// Verify results
	suite.NoError(err)
}

func (suite *InventoryRepositoryTestSuite) TestCreate_ShouldPass() {
	// Setup fixture
	e := entity.TestInventoryItemImplConstructor(
		entity.InvalidID, "some.create.name", "some.create.location", true,
	)

	// Exercise SUT
	_, err := inventory.Create(e)

	// Verify results
	suite.NoError(err)
}

func (suite *InventoryRepositoryTestSuite) TestDeleteById_WhenDoesExist_ShouldPass() {
	// Setup fixture
	e := entity.TestInventoryItemImplConstructor(
		entity.InvalidID, "some.delete.name", "some.delete.location", true,
	)
	id, err := inventory.Create(e)
	suite.NoError(err)

	// Exercise SUT
	err = inventory.DeleteByID(id)

	// Verify results
	suite.NoError(err)
}
