package db

import (
	"fmt"
)

// NotFoundError is returned when an entity cannot be retrieved
// from persistence.
type NotFoundError struct {
	Type string
}

// Check we implement the interface
var _ error = &NotFoundError{}

// NewNotFoundError is a constructor
func NewNotFoundError(_type string) *NotFoundError {
	return &NotFoundError{
		Type: _type,
	}
}

func (n *NotFoundError) Error() string {
	return fmt.Sprintf("entity not found: type=[%s]", n.Type)
}
