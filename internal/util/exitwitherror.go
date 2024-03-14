package util

import (
	"fmt"
	"os"
)

// ExitWithError dsplay error message and exit with error code 1
func ExitWithError(message string, err error) {
	fmt.Fprintf(os.Stderr, "%s: %v\n", message, err)
	os.Exit(1)
}
