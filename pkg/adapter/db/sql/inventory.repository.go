package sql

import (
	goSql "database/sql"
	"fmt"

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

// FindByID implements the Repository interface
func (s *InventoryRepositoryImpl) FindByID(id entity.ID) (entity.InventoryItem, error) {
	query := `
	SELECT 
		id, 
		name, 
		location, 
		available 
	FROM inventory_item
	WHERE 
		id=@id;`
	return s.singleEntityQuery(query, goSql.Named("id", id))
}

// Create implements the Repository interface
func (s *InventoryRepositoryImpl) Create(e entity.InventoryItem) (entity.ID, error) {
	query := `
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
	return s.helperService.ExecForID(s.dbService.Get(), query,
		goSql.Named("name", e.Name()),
		goSql.Named("location", e.Location()),
		goSql.Named("available", e.IsAvailable()),
	)
}

// DeleteByID implements the Repository interface
func (s *InventoryRepositoryImpl) DeleteByID(id entity.ID) error {
	query := `
	DELETE FROM inventory_item
	WHERE 
		id=@id;`
	return s.helperService.ExecForError(s.dbService.Get(), query,
		goSql.Named("id", id),
	)
}

// Update implements the Repository interface
func (s *InventoryRepositoryImpl) Update(e entity.InventoryItem) error {
	query := `
	UPDATE inventory_item
	SET
		name=@name, location=@location, available=@available
	WHERE 
		id=@id;`
	return s.helperService.ExecForError(s.dbService.Get(), query,
		goSql.Named("name", e.Name()),
		goSql.Named("location", e.Location()),
		goSql.Named("available", e.IsAvailable()),
		goSql.Named("id", e.ID()),
	)
}

func (s *InventoryRepositoryImpl) singleEntityQuery(query string, args ...interface{}) (entity.InventoryItem, error) {
	row, err := s.helperService.SingleRowQuery(s.dbService.Get(), query, args...)
	if err != nil {
		return nil, fmt.Errorf("cannot execute query - db get row error: %w", err)
	}

	res, err := s.scanInventoryItem(row)
	if err != nil {
		return nil, fmt.Errorf("cannot execute query - db scan error: %w", err)
	}
	return res, nil
}

func (s *InventoryRepositoryImpl) scanInventoryItem(row Row) (entity.InventoryItem, error) {
	var id entity.ID
	var name string
	var location string
	var available bool

	if err := row.Scan(&id, &name, &location, &available); err != nil {
		return nil, err
	}

	result := s.constructor.Reincarnate(id, name, location, available)
	return result, nil
}
