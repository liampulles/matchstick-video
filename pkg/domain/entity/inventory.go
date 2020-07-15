package entity

// InventoryItem defines a unique InventoryItem
type InventoryItem struct {
	ID        ID
	Name      string
	Location  string
	Available bool
}
