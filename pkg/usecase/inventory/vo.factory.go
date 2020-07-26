package inventory

import (
	"github.com/liampulles/matchstick-video/pkg/domain/entity"
)

// VOFactory is used to create inventory VOs
type VOFactory interface {
	CreateViewVOFromEntity(entity.InventoryItem) *ViewVO
}

// VOFactoryImpl implements VOFactory
type VOFactoryImpl struct{}

// Check we implement the interface
var _ VOFactory = &VOFactoryImpl{}

// NewVOFactoryImpl is a constructor
func NewVOFactoryImpl() *VOFactoryImpl {
	return &VOFactoryImpl{}
}

// CreateViewVOFromEntity maps an entity to a view vo
func (v *VOFactoryImpl) CreateViewVOFromEntity(e entity.InventoryItem) *ViewVO {
	return &ViewVO{
		ID:        e.ID(),
		Name:      e.Name(),
		Location:  e.Location(),
		Available: e.IsAvailable(),
	}
}
