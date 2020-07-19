package dto

import "github.com/liampulles/matchstick-video/pkg/domain/entity"

// Factory constructs DTOs from common formats.
type Factory interface {
	CreateInventoryItemViewFromEntity(entity.InventoryItem) *InventoryItemView
}
