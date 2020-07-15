package commonerror_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/liampulles/matchstick-video/pkg/domain/commonerror"
)

func TestNotImplementedError_Error_ShouldReturnCorrectMessage(t *testing.T) {
	// Setup fixture
	err := commonerror.NewNotImplemented("some.package", "some.struct", "some.method")

	// Setup expectations
	expected := "method not implemented for package=[some.package], struct=[some.struct], method=[some.method]"

	// Verify results
	assert.EqualError(t, err, expected)
}
