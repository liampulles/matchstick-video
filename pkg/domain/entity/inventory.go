package entity

import "fmt"

// InventoryItem defines a unique InventoryItem
type InventoryItem struct {
	ID       ID
	Name     string
	Location string

	available bool
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
