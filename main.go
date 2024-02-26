package main

import (
	"fmt"
	"log"
	"os"

	"github.com/waldirborbajr/glink/glinkcli"
)

var version = "dev"

func main() {
	app := glinkcli.App(version)
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
	fmt.Println("glink v 0.1.0")
}
