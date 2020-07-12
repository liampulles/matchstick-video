package main

import (
	"fmt"
	"os"

	"github.com/liampulles/matchstick-video/pkg/run"
)

func main() {
	if err := run.Run(os.Args); err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}
