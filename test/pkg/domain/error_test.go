package domain_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/liampulles/matchstick-video/pkg/domain"
)

func TestNotImplementedError_Error_ShouldReturnCorrectMessage(t *testing.T) {
	// Setup fixture
	err := &domain.NotImplementedError{
		Package: "some.package",
		Struct:  "some.struct",
		Method:  "some.method",
	}

	// Setup expectations
	expected := "method not implemented for package=[some.package], struct=[some.struct], method=[some.method]"

	// Verify results
	assert.EqualError(t, err, expected)
}
