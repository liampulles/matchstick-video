package inventory

import (
	"github.com/liampulles/matchstick-video/pkg/domain/entity"
)

// VOFactory is used to create inventory VOs
type VOFactory interface {
	CreateViewVOFromEntity(entity.InventoryItem) *ViewVO
	CreateThinViewVOsFromEntities([]entity.InventoryItem) []ThinViewVO
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

// CreateThinViewVOsFromEntities maps an entity to a view vo
func (v *VOFactoryImpl) CreateThinViewVOsFromEntities(entities []entity.InventoryItem) []ThinViewVO {
	var results []ThinViewVO
	for _, e := range entities {
		view := v.createThinViewVOFromEntity(e)
		results = append(results, *view)
	}
	return results
}

func (v *VOFactoryImpl) createThinViewVOFromEntity(e entity.InventoryItem) *ThinViewVO {
	return &ThinViewVO{
		ID:   e.ID(),
		Name: e.Name(),
	}
}
