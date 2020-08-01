package sql

import (
	"fmt"

	"github.com/liampulles/matchstick-video/pkg/adapter/db"
	"github.com/liampulles/matchstick-video/pkg/domain/entity"
	usecaseInventory "github.com/liampulles/matchstick-video/pkg/usecase/inventory"
)

// InventoryRepositoryImpl implements Repository to make use
// of SQL databases which have an associated driver.
type InventoryRepositoryImpl struct {
	dbService     DatabaseService
	helperService HelperService
	constructor   entity.InventoryItemConstructor
}

// Check we implement the interface
var _ usecaseInventory.Repository = &InventoryRepositoryImpl{}

// NewInventoryRepositoryImpl is a constructor
func NewInventoryRepositoryImpl(
	dbService DatabaseService,
	helperService HelperService,
	constructor entity.InventoryItemConstructor,
) *InventoryRepositoryImpl {
	return &InventoryRepositoryImpl{
		dbService:     dbService,
		helperService: helperService,
		constructor:   constructor,
	}
}

// FindByID finds an inventory item matching the given id
func (s *InventoryRepositoryImpl) FindByID(id entity.ID) (entity.InventoryItem, error) {
	query := `
	SELECT 
		id, 
		name, 
		location, 
		available 
	FROM inventory_item
	WHERE 
		id=$1;`
	return s.singleEntityQuery(query, id)
}

// Create persists a new entity. The ID is ignored in the input entity, and the
// generated id is then returned.
func (s *InventoryRepositoryImpl) Create(e entity.InventoryItem) (entity.ID, error) {
	// TODO: Return specific errors for unique constraint violations
	// and parse those as 400's. Do for update as well.
	query := `
	INSERT INTO inventory_item
		(
			name, 
			location, 
			available
		)
	VALUES ($1, $2, $3)
	RETURNING id;`
	return s.helperService.SingleQueryForID(s.dbService.Get(), query,
		e.Name(),
		e.Location(),
		e.IsAvailable(),
	)
}

// DeleteByID deletes the inventory id matching the id. If there
// isn't an entry corresponding to the id - an error is returned.
func (s *InventoryRepositoryImpl) DeleteByID(id entity.ID) error {
	query := `
	DELETE FROM inventory_item
	WHERE 
		id=$1;`
	return s.execExpectingSingleRowAffected(query, id)
}

// Update persists new data for all fields in the given inventory item,
// excluding the id.
func (s *InventoryRepositoryImpl) Update(e entity.InventoryItem) error {
	query := `
	UPDATE inventory_item
	SET
		name=$1, location=$2, available=$3
	WHERE 
		id=$4;`
	return s.execExpectingSingleRowAffected(query,
		e.Name(),
		e.Location(),
		e.IsAvailable(),
		e.ID(),
	)
}

func (s *InventoryRepositoryImpl) execExpectingSingleRowAffected(query string, args ...interface{}) error {
	// Run exec to get rows affected
	rows, err := s.helperService.ExecForRowsAffected(s.dbService.Get(), query, args...)
	if err != nil {
		return fmt.Errorf("cannot execute exec - db exec error: %w", err)
	}

	// Verify rows affected is 1
	if rows == 0 {
		return db.NewNotFoundError("inventory item")
	}
	if rows != 1 {
		return fmt.Errorf("exec error: expected 1 entity to be affected, but was: %d", rows)
	}
	return nil
}

func (s *InventoryRepositoryImpl) singleEntityQuery(query string, args ...interface{}) (entity.InventoryItem, error) {
	// Run the query to get a row
	row, err := s.helperService.SingleRowQuery(s.dbService.Get(), query, args...)
	if err != nil {
		return nil, fmt.Errorf("cannot execute query - db get row error: %w", err)
	}

	// Extract data from the row
	res, err := s.scanInventoryItem(row)
	if err != nil {
		return nil, db.NewNotFoundError("inventory item")
	}
	return res, nil
}

func (s *InventoryRepositoryImpl) scanInventoryItem(row Row) (entity.InventoryItem, error) {
	var id entity.ID
	var name string
	var location string
	var available bool

	// Extract data from the row
	if err := row.Scan(&id, &name, &location, &available); err != nil {
		return nil, err
	}

	// Restore the entity from the extracted data (bypassing validations).
	result := s.constructor.Reincarnate(id, name, location, available)
	return result, nil
}
