package dto

import "github.com/liampulles/matchstick-video/pkg/domain/entity"

// InventoryItemView describes a JSON view of
// entity.InventoryItem
type InventoryItemView struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Location  string    `json:"location"`
	Available bool      `json:"available"`
}
