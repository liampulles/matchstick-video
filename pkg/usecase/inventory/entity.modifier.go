package inventory

import (
	"fmt"

	"github.com/liampulles/matchstick-video/pkg/domain/entity"
)

// EntityModifier encapsulates methods which make mass
// updates to an entity
type EntityModifier interface {
	ModifyWithUpdateItemVO(entity.InventoryItem, *UpdateItemVO) error
}

// EntityModifierImpl implements EntityModifier
type EntityModifierImpl struct{}

var _ EntityModifier = &EntityModifierImpl{}

// NewEntityModifierImpl is a constructor
func NewEntityModifierImpl() *EntityModifierImpl {
	return &EntityModifierImpl{}
}

// ModifyWithUpdateItemVO modidies an existing entity as directed by an update vo
func (e *EntityModifierImpl) ModifyWithUpdateItemVO(ent entity.InventoryItem, vo *UpdateItemVO) error {
	err := ent.ChangeName(vo.Name)
	if err != nil {
		return fmt.Errorf("could not modify entity with update vo - entity name change error: %w", err)
	}

	err = ent.ChangeLocation(vo.Location)
	if err != nil {
		return fmt.Errorf("could not modify entity with update vo - entity location change error: %w", err)
	}

	return nil
}
