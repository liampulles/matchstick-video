package inventory

import "github.com/liampulles/matchstick-video/pkg/domain/entity"

// CreateItemVO defines data needed to create an inventory item.
type CreateItemVO struct {
	Name string
}

// UpdateItemVO defines data that may be used to update an inventory item.
type UpdateItemVO struct {
	ID   entity.ID
	Name string
}
