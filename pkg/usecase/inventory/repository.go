package inventory

import (
	"github.com/liampulles/matchstick-video/pkg/domain/entity"
)

// Repository handles persisting inventory entities
// and retrieving persisted entities
type Repository interface {
	Create(entity.InventoryItem) (entity.ID, error)
	FindByID(entity.ID) (entity.InventoryItem, error)
	Update(entity.InventoryItem) error
	DeleteByID(entity.ID) error
}
