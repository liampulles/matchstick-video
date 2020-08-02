package db

import "fmt"

// UniqueConstraintError is returned when a transaction fails a uniqueness
// constraint set in the database.
type UniqueConstraintError struct {
	Cause error
}

// Check we implement the interface
var _ error = &UniqueConstraintError{}

// NewUniqueConstraintError is a constructor
func NewUniqueConstraintError(cause error) *UniqueConstraintError {
	return &UniqueConstraintError{
		Cause: cause,
	}
}

func (u *UniqueConstraintError) Error() string {
	return fmt.Sprintf("uniqueness constraint error: %s", u.Cause.Error())
}
