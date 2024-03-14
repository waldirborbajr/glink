package cmds

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// Version returns a cli.Command which displays the version of the glink app
func Version() *cli.Command {
	return &cli.Command{
		Name:                   "version",
		Aliases:                []string{"v"},
		Usage:                  "Display glink version",
		UseShortOptionHandling: true,
		Action: func(cCtx *cli.Context) error {
			fmt.Printf("%s v%s \n", cCtx.App.Name, cCtx.App.Version)
			return nil
		},
	}
}
