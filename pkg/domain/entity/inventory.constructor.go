package entity

// InventoryItemConstructor constructs InventoryItems
type InventoryItemConstructor interface {
	Reincarnate(id ID, name string, location string, available bool) InventoryItem
	NewAvailable(name string, location string) InventoryItem
}

// InventoryItemConstructorImpl implements InventoryItemConstructor
type InventoryItemConstructorImpl struct{}

var _ InventoryItemConstructor = &InventoryItemConstructorImpl{}

// NewInventoryItemConstructorImpl is a constructor
func NewInventoryItemConstructorImpl() *InventoryItemConstructorImpl {
	return &InventoryItemConstructorImpl{}
}

// Reincarnate implements the InventoryItemConstructor interface
func (i *InventoryItemConstructorImpl) Reincarnate(id ID, name string, location string, available bool) InventoryItem {
	return &InventoryItemImpl{
		id:        id,
		name:      name,
		location:  location,
		available: available,
	}
}

// NewAvailable implements the InventoryItemConstructor interface
func (i *InventoryItemConstructorImpl) NewAvailable(name string, location string) InventoryItem {
	result := newBaseInventoryItem(name, location)
	result.available = true
	return result
}

func newBaseInventoryItem(name string, location string) *InventoryItemImpl {
	return &InventoryItemImpl{
		id:       InvalidID,
		name:     name,
		location: location,
	}
}
