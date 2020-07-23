package dto

import "github.com/liampulles/matchstick-video/pkg/domain/entity"

// Factory constructs DTOs from common formats.
type Factory interface {
	CreateInventoryItemViewFromEntity(entity.InventoryItem) *InventoryItemView
}

// FactoryImpl implements Factory
type FactoryImpl struct{}

var _ Factory = &FactoryImpl{}

// NewFactoryImpl is a constructor
func NewFactoryImpl() *FactoryImpl {
	return &FactoryImpl{}
}

// CreateInventoryItemViewFromEntity implements Factory
func (f *FactoryImpl) CreateInventoryItemViewFromEntity(e entity.InventoryItem) *InventoryItemView {
	return &InventoryItemView{
		ID:        e.ID(),
		Name:      e.Name(),
		Location:  e.Location(),
		Available: e.IsAvailable(),
	}
}
