package commonerror_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/liampulles/matchstick-video/pkg/domain/commonerror"
)

func TestValidationError_Error_ShouldReturnCorrectMessage(t *testing.T) {
	// Setup fixture
	err := commonerror.NewValidation("some.field", "some.problem")

	// Setup expectations
	expected := "validation error: field=[some.field], problem=[some.problem]"

	// Verify results
	assert.EqualError(t, err, expected)
}
