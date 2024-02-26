package main

import (
	"log"
	"os"

	"github.com/waldirborbajr/glink/glinkcli"
)

var (
	version  = "0.1.0"
	revision = "dev"
)

func main() {
	app := glinkcli.App(version, revision)
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
