package inventory

import (
	"github.com/liampulles/matchstick-video/pkg/domain/entity"
)

// Repository handles persisting inventory entities
// and retrieving persisted entities
type Repository interface {
	FindByID(entity.ID) (*entity.InventoryItem, error)
	Save(*entity.InventoryItem) (*entity.InventoryItem, error)
}
