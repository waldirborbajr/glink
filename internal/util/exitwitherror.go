package util

import (
	"fmt"
	"os"
)

func ExitWithError(message string, err error) {
	if err == nil {
		fmt.Fprint(os.Stderr, message)
	} else {
		fmt.Fprintf(os.Stderr, "%s: %v\n", message, err)
	}
	os.Exit(1)
}
