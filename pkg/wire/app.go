package wire

import (
	"fmt"

	goConfig "github.com/liampulles/go-config"

	"github.com/liampulles/matchstick-video/pkg/domain"
	"github.com/liampulles/matchstick-video/pkg/wire/config"
)

// Run is the entrypoint for the application
func Run() int {
	// Delegate nearly all this logic,
	// since we can't easily test os functions
	// and the ongoing server.
	app := CreateApp(goConfig.NewEnvSource())
	err := app()
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return 1
	}
	return 0
}

// CreateApp reads configuration, wires the application, and
// returns a runnable which will start the application.
func CreateApp(source goConfig.Source) domain.Runnable {
	cfg, err := config.Fetch(source)
	if err != nil {
		return func() error {
			return fmt.Errorf("could not construct config: %w", err)
		}
	}

	return wireApp(cfg)
}

func wireApp(cfg *config.Config) domain.Runnable {
	// TODO: Select drivers based on config (strategy pattern?)
	// TODO:

	return func() error {
		fmt.Println("Hello world!")
		return nil
	}
}
