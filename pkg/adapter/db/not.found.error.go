package db

import (
	"fmt"

	"github.com/liampulles/matchstick-video/pkg/domain/entity"
)

// NotFoundError is returned when an entity cannot be retrieved
// from persistence.
type NotFoundError struct {
	Type string
	ID   entity.ID
}

// Check we implement the interface
var _ error = &NotFoundError{}

// NewNotFoundError is a constructor
func NewNotFoundError(_type string, id entity.ID) *NotFoundError {
	return &NotFoundError{
		Type: _type,
		ID:   id,
	}
}

func (n *NotFoundError) Error() string {
	return fmt.Sprintf("entity not found: type=[%s], id=[%d]", n.Type, n.ID)
}
