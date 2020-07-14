package fetch_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	goConfig "github.com/liampulles/go-config"

	"github.com/liampulles/matchstick-video/pkg/wire/config"
)

func TestFetch_GivenValidConfigSource_ShouldReturnExpectedConfig(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		fixture  map[string]string
		expected *config.Config
	}{
		{
			map[string]string{},
			&config.Config{
				LogLevel: "INFO",
				Port:     8080,
			},
		},
		{
			map[string]string{
				"LOGLEVEL": "ERROR",
				"PORT":     "9000",
			},
			&config.Config{
				LogLevel: "ERROR",
				Port:     9000,
			},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			// Exercise SUT
			actual, err := config.Fetch(goConfig.MapSource(test.fixture))

			// Verify result
			assert.Nil(t, err)
			assert.Equal(t, actual, test.expected)
		})
	}
}

func TestFetch_GivenInvalidConfigSource_ShouldReturnError(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		fixture         map[string]string
		expectedMessage string
	}{
		{
			map[string]string{
				"LOGLEVEL": "ERROR",
				"PORT":     "not.an.int",
			},
			"could not fetch config: value of PORT property can not be converted to int (is not.an.int)",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			// Exercise SUT
			actual, err := config.Fetch(goConfig.MapSource(test.fixture))

			// Verify result
			assert.Nil(t, actual)
			assert.EqualError(t, err, test.expectedMessage)
		})
	}
}
