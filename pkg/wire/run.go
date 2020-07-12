package wire

import "fmt"

// Run is the entrypoint of the application. This function
// is called by main().
func Run(args []string, env []string) error {

	fmt.Println("Hello world!")
	return nil
}
