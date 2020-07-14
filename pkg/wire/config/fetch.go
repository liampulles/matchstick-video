package config

import (
	"fmt"

	goConfig "github.com/liampulles/go-config"
)

// Fetch resolves configuration properties from
// command line arguments and env variables into
// a struct.
func Fetch(source goConfig.Source) (*Config, error) {
	typedSource := goConfig.NewTypedSource(source)
	config := &Config{
		LogLevel: "INFO",
		Port:     8080,
	}

	if err := goConfig.LoadProperties(typedSource,
		goConfig.StrProp("LOGLEVEL", &config.LogLevel, false),
		goConfig.IntProp("PORT", &config.Port, false),
	); err != nil {
		return nil, fmt.Errorf("could not fetch config: %w", err)
	}

	return config, nil
}
