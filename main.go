package main

import (
	"fmt"
	"os"

	"github.com/liampulles/matchstick-video/pkg/wire"
)

func main() {
	// Delegate logic to Run, since we can't easily
	// test this function.
	if err := wire.Run(os.Args, os.Environ()); err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}
