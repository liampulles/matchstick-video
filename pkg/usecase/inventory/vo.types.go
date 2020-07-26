package inventory

import "github.com/liampulles/matchstick-video/pkg/domain/entity"

// CreateItemVO defines data needed to create an inventory item.
type CreateItemVO struct {
	Name     string
	Location string
}

// UpdateItemVO defines data that may be used to update an inventory item.
type UpdateItemVO struct {
	Name     string
	Location string
}

// ViewVO describes an inventory item in full
// (or at least, to the greatest degree we want users
// to see them).
type ViewVO struct {
	ID        entity.ID
	Name      string
	Location  string
	Available bool
}
