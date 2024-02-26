package cmds

import (
	"fmt"

	cli "github.com/urfave/cli/v2"
)

func Link() *cli.Command {
	return &cli.Command{
		Name:                   "link",
		Aliases:                []string{"l"},
		Usage:                  "Create symlinks",
		UseShortOptionHandling: true,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "force",
				Aliases: []string{"f"},
				Usage:   "force overwrite symlinkl",
			},
		},
		Action: func(cCtx *cli.Context) error {
			fmt.Println("symlinks created")
			return nil
		},
	}
}
