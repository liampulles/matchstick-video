package main

import (
	"fmt"
	"os"

	goConfig "github.com/liampulles/go-config"

	"github.com/liampulles/matchstick-video/pkg/wire"
)

func main() {
	// Delegate most logic elsewhere, since we can't
	// test this function.
	app := wire.CreateApp(goConfig.NewEnvSource())
	err := app()
	if err != nil {
		fmt.Printf("APP ERROR - PANICKING: %s\n", err.Error())
		panic(err)
	}

	os.Exit(0)
}
