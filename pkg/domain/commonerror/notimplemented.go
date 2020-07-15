package commonerror

import "fmt"

// NotImplemented is returned when a method
// has no implementation
type NotImplemented struct {
	_package string
	_struct  string
	method   string
}

var _ error = &NotImplemented{}

// NewNotImplemented is a constructor
func NewNotImplemented(_package string, _struct string, method string) *NotImplemented {
	return &NotImplemented{
		_package: _package,
		_struct:  _struct,
		method:   method,
	}
}

func (n *NotImplemented) Error() string {
	return fmt.Sprintf(
		"method not implemented for package=[%s], struct=[%s], method=[%s]",
		n._package, n._struct, n.method)
}
