package inventory

import "github.com/liampulles/matchstick-video/pkg/domain/entity"

// EntityFactory defines methods for creating
// an entity.InventoryItem
type EntityFactory interface {
	CreateFromVO(*CreateItemVO) (*entity.InventoryItem, error)
}
