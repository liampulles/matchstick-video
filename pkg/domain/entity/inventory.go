package entity

import "fmt"

// InventoryItem defines a unique InventoryItem
type InventoryItem struct {
	Name     string
	Location string

	id        ID
	available bool
}

// NewAvailableInventoryItem is a constructor
func NewAvailableInventoryItem(id ID, name string, location string) *InventoryItem {
	base := newBaseInventoryItem(id, name, location)
	base.available = true
	return base
}

// NewUnavailableInventoryItem is a constructor
func NewUnavailableInventoryItem(id ID, name string, location string) *InventoryItem {
	base := newBaseInventoryItem(id, name, location)
	base.available = false
	return base
}

func newBaseInventoryItem(id ID, name string, location string) *InventoryItem {
	return &InventoryItem{
		Name:     name,
		Location: location,

		id: id,
	}
}

// ID returns the id.
func (i *InventoryItem) ID() ID {
	return i.id
}

// IsAvailable will return true if the inventory item may
// be checked out - false otherwise.
func (i *InventoryItem) IsAvailable() bool {
	return i.available
}

// Checkout will mark the inventory item as unavilable.
// If the inventory item is not available,
// then an error is returned.
func (i *InventoryItem) Checkout() error {
	if !i.available {
		return fmt.Errorf("cannot check out inventory item - it is unavailable")
	}
	i.available = false
	return nil
}

// CheckIn will mark the inventory item as available.
// If the inventory item is available, then an
// error is returned.
func (i *InventoryItem) CheckIn() error {
	if i.available {
		return fmt.Errorf("cannot check in inventory item - it is already checked in")
	}
	i.available = true
	return nil
}
