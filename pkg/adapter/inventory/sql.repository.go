package inventory

import (
	goSql "database/sql"
	"fmt"

	"github.com/liampulles/matchstick-video/pkg/adapter/db/sql"
	"github.com/liampulles/matchstick-video/pkg/domain/entity"
	usecaseInventory "github.com/liampulles/matchstick-video/pkg/usecase/inventory"
)

// SQLRepositoryImpl implements Repository to make use
// of SQL databases which have an associated driver.
type SQLRepositoryImpl struct {
	db            *goSql.DB
	helperService sql.HelperService
	constructor   entity.InventoryItemConstructor
}

var _ usecaseInventory.Repository = &SQLRepositoryImpl{}

// NewSQLRepositoryImpl is a constructor
func NewSQLRepositoryImpl(
	db *goSql.DB,
	helperService sql.HelperService,
	constructor entity.InventoryItemConstructor,
) *SQLRepositoryImpl {
	return &SQLRepositoryImpl{
		db:            db,
		helperService: helperService,
		constructor:   constructor,
	}
}

// FindByID implements the Repository interface
func (s *SQLRepositoryImpl) FindByID(id entity.ID) (entity.InventoryItem, error) {
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
func (s *SQLRepositoryImpl) Create(e entity.InventoryItem) (entity.ID, error) {
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
	return s.helperService.ExecForID(s.db, query,
		goSql.Named("name", e.Name()),
		goSql.Named("location", e.Location()),
		goSql.Named("available", e.IsAvailable()),
	)
}

// DeleteByID implements the Repository interface
func (s *SQLRepositoryImpl) DeleteByID(id entity.ID) error {
	query := `
	DELETE FROM inventory_item
	WHERE 
		id=@id;`
	return s.helperService.ExecForError(s.db, query,
		goSql.Named("id", id),
	)
}

// Update implements the Repository interface
func (s *SQLRepositoryImpl) Update(e entity.InventoryItem) error {
	query := `
	UPDATE inventory_item
	SET
		name=@name, location=@location, available=@available
	WHERE 
		id=@id;`
	return s.helperService.ExecForError(s.db, query,
		goSql.Named("name", e.Name()),
		goSql.Named("location", e.Location()),
		goSql.Named("available", e.IsAvailable()),
		goSql.Named("id", e.ID()),
	)
}

func (s *SQLRepositoryImpl) singleEntityQuery(query string, args ...interface{}) (entity.InventoryItem, error) {
	row, err := s.helperService.SingleRowQuery(s.db, query, args...)
	if err != nil {
		return nil, fmt.Errorf("cannot execute query - db get row error: %w", err)
	}

	res, err := s.scanInventoryItem(row)
	if err != nil {
		return nil, fmt.Errorf("cannot execute query - db scan error: %w", err)
	}
	return res, nil
}

func (s *SQLRepositoryImpl) scanInventoryItem(row sql.Row) (entity.InventoryItem, error) {
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
