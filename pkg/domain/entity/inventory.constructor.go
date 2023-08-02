package entity

// Reincarnate creates an entity which was already tested and accepted - but
// just needs to be restored. Thus, this method bypasses validation. It should
// be used from system-sources, e.g. a database, and not user sources, e.g.
// a request.
var Reincarnate = func(id ID, name string, location string, available bool) InventoryItem {
	return &InventoryItemImpl{
		id:        id,
		name:      name,
		location:  location,
		available: available,
	}
}

// NewAvailable creates a brand new entity from the given parameters. The input
// is validated and will fail if appropriate. The resulting entity will not have
// a valid id (you will probably want to persist it to get one).
var NewAvailable = func(name string, location string) (InventoryItem, error) {
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
