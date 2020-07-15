package entity

import "fmt"

// NotFoundError is returned when an entity
// cannot be found for a given id and type.
type NotFoundError struct {
	_type string
	id    ID
}

var _ error = &NotFoundError{}

// NewNotFoundError is a constructor
func NewNotFoundError(_type string, id ID) *NotFoundError {
	return &NotFoundError{
		_type: _type,
		id:    id,
	}
}

func (n *NotFoundError) Error() string {
	return fmt.Sprintf("entity not found for type=[%s] and id=[%d]", n._type, n.id)
}
