package inventory

import "github.com/liampulles/matchstick-video/pkg/domain/entity"

// EntityFactory defines methods for creating
// an entity.InventoryItem from VOs
type EntityFactory interface {
	CreateFromVO(*CreateItemVO) (entity.InventoryItem, error)
}

// EntityFactoryImpl implements EntityFactory
type EntityFactoryImpl struct {
	constructor entity.InventoryItemConstructor
}

// Check we implement the interface
var _ EntityFactory = &EntityFactoryImpl{}

// NewEntityFactoryImpl is a constructor
func NewEntityFactoryImpl(constructor entity.InventoryItemConstructor) *EntityFactoryImpl {
	return &EntityFactoryImpl{
		constructor: constructor,
	}
}

// CreateFromVO creates a new entity from a vo
func (e *EntityFactoryImpl) CreateFromVO(vo *CreateItemVO) (entity.InventoryItem, error) {
	return e.constructor.NewAvailable(vo.Name, vo.Location)
}
