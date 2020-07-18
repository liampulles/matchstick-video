package commonerror

import "fmt"

// Validation is returned when a field
// do not meet certain requirements
type Validation struct {
	Field   string
	Problem string
}

// Check we implement the interface
var _ error = &Validation{}

// NewValidation is a constructor
func NewValidation(field string, problem string) *Validation {
	return &Validation{
		Field:   field,
		Problem: problem,
	}
}

func (v *Validation) Error() string {
	return fmt.Sprintf(
		"validation error: field=[%s], problem=[%s]",
		v.Field, v.Problem,
	)
}
