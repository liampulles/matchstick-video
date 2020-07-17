package inventory

import (
	"database/sql"
	"fmt"

	"github.com/liampulles/matchstick-video/pkg/adapter"
	"github.com/liampulles/matchstick-video/pkg/domain/entity"
	usecaseInventory "github.com/liampulles/matchstick-video/pkg/usecase/inventory"
)

// SQLRepositoryImpl implements Repository to make use
// of SQL databases which have an associated driver.
type SQLRepositoryImpl struct {
	db            *sql.DB
	helperService adapter.SQLDbHelperService
	constructor   entity.InventoryItemConstructor
}

var _ usecaseInventory.Repository = &SQLRepositoryImpl{}

// NewSQLRepositoryImpl is a constructor
func NewSQLRepositoryImpl(
	db *sql.DB,
	helperService adapter.SQLDbHelperService,
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
	return s.singleEntityQuery(query, sql.Named("id", id))
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
		sql.Named("name", e.Name()),
		sql.Named("location", e.Location()),
		sql.Named("available", e.IsAvailable()),
	)
}

// DeleteByID implements the Repository interface
func (s *SQLRepositoryImpl) DeleteByID(id entity.ID) error {
	query := `
	DELETE FROM inventory_item
	WHERE 
		id=@id;`
	return s.helperService.ExecForError(s.db, query,
		sql.Named("id", id),
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
		sql.Named("name", e.Name()),
		sql.Named("location", e.Location()),
		sql.Named("available", e.IsAvailable()),
		sql.Named("id", e.ID()),
	)
}

// func (s *SQLRepositoryImpl) execForError(query string, args ...interface{}) error {
// 	_, err := s.exec(query, args...)
// 	if err != nil {
// 		return fmt.Errorf("cannot execute exec - db exec error: %w", err)
// 	}
// 	return nil
// }

// func (s *SQLRepositoryImpl) execForID(query string, args ...interface{}) (entity.ID, error) {
// 	res, err := s.exec(query, args...)
// 	if err != nil {
// 		return entity.InvalidID, fmt.Errorf("cannot execute exec - db exec error: %w", err)
// 	}

// 	id, err := res.LastInsertId()
// 	if err != nil {
// 		return entity.InvalidID, fmt.Errorf("cannot execute exec - result id error: %w", err)
// 	}
// 	return entity.ID(id), nil
// }

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

// func (s *SQLRepositoryImpl) exec(query string, args ...interface{}) (sql.Result, error) {
// 	ctx := context.TODO()
// 	stmt, err := s.db.PrepareContext(ctx, query)
// 	if err != nil {
// 		return nil, err
// 	}

// 	res, err := stmt.ExecContext(ctx, args...)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return res, nil
// }

// func (s *SQLRepositoryImpl) singleRowQuery(query string, args ...interface{}) (*sql.Row, error) {
// 	ctx := context.TODO()
// 	stmt, err := s.db.PrepareContext(ctx, query)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return stmt.QueryRowContext(ctx, args...), nil
// }

func (s *SQLRepositoryImpl) scanInventoryItem(row adapter.SQLRow) (entity.InventoryItem, error) {
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
