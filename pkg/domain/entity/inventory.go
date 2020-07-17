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
	InitID(ID) error
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

// TestInventoryItemImplConstructor allows you to
// create an InventoryItemImpl, directly - bypassing
// the constructor service. It should ONLY be used
// in tests.
func TestInventoryItemImplConstructor(
	id ID,
	name string,
	location string,
	available bool) *InventoryItemImpl {

	return &InventoryItemImpl{
		id:        id,
		name:      name,
		location:  location,
		available: available,
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

// InitID will allow the caller to set the ID - but only if
// it has not already been set. Otherwise, an error is
// returned.
func (i *InventoryItemImpl) InitID(id ID) error {
	if i.id != InvalidID {
		return fmt.Errorf("could not init id - id has already been set [%d]", i.id)
	}
	i.id = id
	return nil
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
