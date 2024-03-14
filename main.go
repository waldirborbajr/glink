package main

import (
	"log"
	"os"

	"github.com/waldirborbajr/glink/glinkcli"
)

// BuildVersion is provided to be overridden at build time. Eg. go build -ldflags -X 'main.BuildVersion=...'
var BuildVersion = "(development build)"

func main() {
	app := glinkcli.App(BuildVersion)
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
