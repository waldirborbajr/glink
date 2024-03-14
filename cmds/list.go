package cmds

import (
	cli "github.com/urfave/cli/v2"
)

// List returns a cli.Command for the list subcommand
func List() *cli.Command {
	return &cli.Command{
		Name:                   "list",
		Aliases:                []string{"s"},
		Usage:                  "List symlinks created",
		UseShortOptionHandling: true,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "json",
				Aliases: []string{"j"},
				Usage:   "output as json",
			},
			&cli.BoolFlag{
				Name:    "tmux",
				Aliases: []string{"t"},
				Usage:   "show tmux sessions",
			},
			&cli.BoolFlag{
				Name:    "zoxide",
				Aliases: []string{"z"},
				Usage:   "show zoxide results",
			},
			&cli.BoolFlag{
				Name:    "hide-attached",
				Aliases: []string{"H"},
				Usage:   "don't show currently attached sessions",
			},
			&cli.BoolFlag{
				Name:    "icons",
				Aliases: []string{"i"},
				Usage:   "show Nerd Font icons",
			},
		},
		Action: func(_ *cli.Context) error {
			return nil
		},
	}
}
