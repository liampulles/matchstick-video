package validation_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/liampulles/matchstick-video/pkg/domain/validation"
)

func TestIsBlank_WhenStringIsNotBlank_ShouldReturnFalse(t *testing.T) {
	// Setup fixture
	var fixtures = []string{
		// Blank
		"a",
		" a",
		"a ",
		" a ",
		"good boi",
	}

	for _, fixture := range fixtures {
		t.Run(fmt.Sprintf("\"%s\"", fixture), func(t *testing.T) {
			// Exercise SUT
			actual := validation.IsBlank(fixture)

			// Verify result
			assert.False(t, actual)
		})
	}
}

func TestIsBlank_WhenStringIsBlank_ShouldReturnTrue(t *testing.T) {
	// Setup fixture
	var fixtures = []string{
		"",
		" ",
		"  ",
		"\t",
		" \t",
		// -> We will assume the other whitespace
		// characters will work as well.
	}

	for _, fixture := range fixtures {
		t.Run(fmt.Sprintf("\"%s\"", fixture), func(t *testing.T) {
			// Exercise SUT
			actual := validation.IsBlank(fixture)

			// Verify result
			assert.True(t, actual)
		})
	}
}

func TestIsTrimmed_WhenStringIsNotTrimmed_ShouldReturnFalse(t *testing.T) {
	// Setup fixture
	var fixtures = []string{
		" ",
		"  ",
		"\t",
		" \t",
		" a",
		"a ",
		" a ",
	}

	for _, fixture := range fixtures {
		t.Run(fmt.Sprintf("\"%s\"", fixture), func(t *testing.T) {
			// Exercise SUT
			actual := validation.IsTrimmed(fixture)

			// Verify result
			assert.False(t, actual)
		})
	}
}

func TestIsTrimmed_WhenStringIsTrimmed_ShouldReturnTrue(t *testing.T) {
	// Setup fixture
	var fixtures = []string{
		"",
		"a",
		"good boi",
	}

	for _, fixture := range fixtures {
		t.Run(fmt.Sprintf("\"%s\"", fixture), func(t *testing.T) {
			// Exercise SUT
			actual := validation.IsTrimmed(fixture)

			// Verify result
			assert.True(t, actual)
		})
	}
}
