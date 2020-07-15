package inventory

import "github.com/liampulles/matchstick-video/pkg/domain/entity"

// EntityModifier encapsulates methods which make mass
// updates to an entity
type EntityModifier interface {
	ModifyWithUpdateItemVO(entity.InventoryItem, *UpdateItemVO) error
}
