package domain

import "fmt"

// NotImplementedError is returned when a method
// has no implementation
type NotImplementedError struct {
	Package string
	Struct  string
	Method  string
}

var _ error = &NotImplementedError{}

func (n *NotImplementedError) Error() string {
	return fmt.Sprintf(
		"method not implemented for package=[%s], struct=[%s], method=[%s]",
		n.Package, n.Struct, n.Method)
}
