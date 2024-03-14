package util

import (
	"fmt"
	"os"
)

func ExitWithError(message string, err error) {
	fmt.Fprintf(os.Stderr, "%s: %v\n", message, err)
	os.Exit(1)
}
