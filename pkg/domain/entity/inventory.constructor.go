package entity

// InventoryItemConstructor constructs InventoryItems
type InventoryItemConstructor interface {
	Reincarnate(id ID, name string, location string, available bool) InventoryItem
	NewAvailable(name string, location string) (InventoryItem, error)
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
func (i *InventoryItemConstructorImpl) NewAvailable(name string, location string) (InventoryItem, error) {
	result, err := newBaseInventoryItem(name, location)
	if err != nil {
		return nil, err
	}
	result.available = true
	return result, nil
}

func newBaseInventoryItem(name string, location string) (*InventoryItemImpl, error) {
	result := &InventoryItemImpl{
		id:        InvalidID,
		available: true,
	}

	if err := result.ChangeName(name); err != nil {
		return nil, err
	}
	if err := result.ChangeLocation(location); err != nil {
		return nil, err
	}

	return result, nil
}
