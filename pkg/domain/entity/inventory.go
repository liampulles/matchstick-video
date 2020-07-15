package entity

import (
	"fmt"

	"github.com/liampulles/matchstick-video/pkg/domain/commonerror"
)

// InventoryItem defines a unique entity
type InventoryItem interface {
	ID() ID
	Name() string
	Location() string
	IsAvailable() bool
	Checkout() error
	CheckIn() error
	ChangeName(string) error
	ChangeLocation(string) error
}

// InventoryItemImpl implements InventoryItem
type InventoryItemImpl struct {
	id        ID
	name      string
	location  string
	available bool
}

// Check interface is implemented
var _ InventoryItem = &InventoryItemImpl{}

// NewAvailableInventoryItem is a constructor
func NewAvailableInventoryItem(id ID, name string, location string) *InventoryItemImpl {
	base := newBaseInventoryItem(id, name, location)
	base.available = true
	return base
}

// NewUnavailableInventoryItem is a constructor
func NewUnavailableInventoryItem(id ID, name string, location string) *InventoryItemImpl {
	base := newBaseInventoryItem(id, name, location)
	base.available = false
	return base
}

func newBaseInventoryItem(id ID, name string, location string) *InventoryItemImpl {
	return &InventoryItemImpl{
		id:       id,
		name:     name,
		location: location,
	}
}

// ID returns the id.
func (i *InventoryItemImpl) ID() ID {
	return i.id
}

// Name returns the name.
func (i *InventoryItemImpl) Name() string {
	return i.name
}

// Location returns the name.
func (i *InventoryItemImpl) Location() string {
	return i.location
}

// IsAvailable will return true if the inventory item may
// be checked out - false otherwise.
func (i *InventoryItemImpl) IsAvailable() bool {
	return i.available
}

// Checkout will mark the inventory item as unavilable.
// If the inventory item is not available,
// then an error is returned.
func (i *InventoryItemImpl) Checkout() error {
	if !i.available {
		return fmt.Errorf("cannot check out inventory item - it is unavailable")
	}
	i.available = false
	return nil
}

// CheckIn will mark the inventory item as available.
// If the inventory item is available, then an
// error is returned.
func (i *InventoryItemImpl) CheckIn() error {
	if i.available {
		return fmt.Errorf("cannot check in inventory item - it is already checked in")
	}
	i.available = true
	return nil
}

// ChangeName will change the name of the inventory item,
// if it is valid. If it is not valid, it will return
// an error
func (i *InventoryItemImpl) ChangeName(name string) error {
	return commonerror.NewNotImplemented("entity", "InventoryItemImpl", "ChangeName")
}

// ChangeLocation will change the location of the inventory item,
// if it is valid. If it is not valid, it will return
// an error
func (i *InventoryItemImpl) ChangeLocation(location string) error {
	return commonerror.NewNotImplemented("entity", "InventoryItemImpl", "ChangeLocation")
}
