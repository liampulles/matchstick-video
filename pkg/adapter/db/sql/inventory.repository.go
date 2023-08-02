package sql

import (
	"github.com/liampulles/matchstick-video/pkg/domain/entity"
)

// FindByID finds an inventory item matching the given id
var FindByID = func(id entity.ID) (entity.InventoryItem, error) {
	query := `
	SELECT 
		id, 
		name, 
		location, 
		available 
	FROM inventory_item
	WHERE 
		id=$1;`
	return singleEntityQuery(query, id)
}

// FindAll retrieves all the inventory items in the database
var FindAll = func() ([]entity.InventoryItem, error) {
	query := `
	SELECT 
		id, 
		name, 
		location, 
		available 
	FROM inventory_item;`
	return manyEntityQuery(query)
}

// Create persists a new entity. The ID is ignored in the input entity, and the
// generated id is then returned.
var Create = func(e entity.InventoryItem) (entity.ID, error) {
	query := `
	INSERT INTO inventory_item
		(
			name, 
			location, 
			available
		)
	VALUES ($1, $2, $3)
	RETURNING id;`
	return SingleQueryForID(query, "inventory item",
		e.Name(),
		e.Location(),
		e.IsAvailable(),
	)
}

// DeleteByID deletes the inventory id matching the id. If there
// isn't an entry corresponding to the id - an error is returned.
var DeleteByID = func(id entity.ID) error {
	query := `
	DELETE FROM inventory_item
	WHERE 
		id=$1;`
	return ExecForSingleItem(query, id)
}

// Update persists new data for all fields in the given inventory item,
// excluding the id.
var Update = func(e entity.InventoryItem) error {
	query := `
	UPDATE inventory_item
	SET
		name=$1, location=$2, available=$3
	WHERE 
		id=$4;`
	return ExecForSingleItem(query,
		e.Name(),
		e.Location(),
		e.IsAvailable(),
		e.ID(),
	)
}

func singleEntityQuery(query string, args ...interface{}) (entity.InventoryItem, error) {
	var result entity.InventoryItem

	// Run the query to get a row
	err := SingleRowQuery(query, func(row Row) error {
		res, err := scanInventoryItem(row)
		result = res
		return err
	}, "inventory item", args...)

	return result, err
}

func manyEntityQuery(query string, args ...interface{}) ([]entity.InventoryItem, error) {
	var results []entity.InventoryItem

	// Run the query to get a row
	err := ManyRowsQuery(query, func(row Row) error {
		res, err := scanInventoryItem(row)
		if res != nil {
			results = append(results, res)
		}
		return err
	}, "inventory item", args...)

	if err != nil {
		return nil, err
	}
	return results, nil
}

func scanInventoryItem(row Row) (entity.InventoryItem, error) {
	var id entity.ID
	var name string
	var location string
	var available bool

	// Extract data from the row
	if err := row.Scan(&id, &name, &location, &available); err != nil {
		return nil, err
	}

	// Restore the entity from the extracted data (bypassing validations).
	result := entity.ReincarnateInventory(id, name, location, available)
	return result, nil
}
