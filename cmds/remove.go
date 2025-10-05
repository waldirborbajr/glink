package cmds

import "github.com/urfave/cli/v2"

func Remove() *cli.Command {
	return &cli.Command{
		Name:    "remove",
		Aliases: []string{"r", "rm"},
		Usage:   "Remove Symlinks",
		Action: func(context *cli.Context) error {
			//TODO
			return nil
		},
	}
}
