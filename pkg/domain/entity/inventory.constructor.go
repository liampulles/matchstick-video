package entity

// InventoryItemConstructor constructs InventoryItems
type InventoryItemConstructor interface {
	NewAvailable(name string, location string) InventoryItem
}

// InventoryItemConstructorImpl implements InventoryItemConstructor
type InventoryItemConstructorImpl struct{}

var _ InventoryItemConstructor = &InventoryItemConstructorImpl{}

// NewInventoryItemConstructorImpl is a constructor
func NewInventoryItemConstructorImpl() *InventoryItemConstructorImpl {
	return &InventoryItemConstructorImpl{}
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
