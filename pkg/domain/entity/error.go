package entity

import "fmt"

// NotFoundError is returned when an entity
// cannot be found for a given id and type.
type NotFoundError struct {
	Type string
	ID   ID
}

var _ error = &NotFoundError{}

func (n *NotFoundError) Error() string {
	return fmt.Sprintf("entity not found for type=[%s] and id=[%d]", n.Type, n.ID)
}
