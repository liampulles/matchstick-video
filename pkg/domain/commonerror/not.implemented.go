package commonerror

import "fmt"

// NotImplemented is returned when a method
// has no implementation
type NotImplemented struct {
	Package string
	Struct  string
	Method  string
}

var _ error = &NotImplemented{}

// NewNotImplemented is a constructor
func NewNotImplemented(_package string, _struct string, method string) *NotImplemented {
	return &NotImplemented{
		Package: _package,
		Struct:  _struct,
		Method:  method,
	}
}

func (n *NotImplemented) Error() string {
	return fmt.Sprintf(
		"method not implemented for package=[%s], struct=[%s], method=[%s]",
		n.Package, n.Struct, n.Method)
}
