package dto

// InventoryItemView describes a JSON view of
// entity.InventoryItem
type InventoryItemView struct {
	ID        string
	Name      string
	Location  string
	Available bool
}
