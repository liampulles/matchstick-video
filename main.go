package main

import (
	"os"

	"github.com/liampulles/matchstick-video/pkg/wire"
)

func main() {
	// Delegate logic elsewhere, since we can't
	// test this function.
	os.Exit(wire.CreateMain()())
}
