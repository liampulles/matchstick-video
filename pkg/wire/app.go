package wire

import (
	"fmt"

	goConfig "github.com/liampulles/go-config"
)

// MainFunc returns an int, non-zero
// if there was an error.
type MainFunc func() int

// CreateMain creates the entrypoint for the application
func CreateMain() MainFunc {
	// Delegate nearly all this logic,
	// since we can't easily test os functions
	// and the ongoing server.
	server, err := CreateServer(goConfig.NewEnvSource())
	if err != nil {
		return failedMain(err)
	}
	// TODO: should maybe return runnable directly instead.
	app := server.Create()
	return func() int {
		// This is effectively the main function.
		fmt.Println("Now serving...")
		err = app()
		if err != nil {
			return failedMain(err)()
		}
		return 0
	}
}

func failedMain(err error) MainFunc {
	return func() int {
		fmt.Printf("ERROR: %v\n", err)
		return 1
	}
}
