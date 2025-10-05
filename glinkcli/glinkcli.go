package glinkcli

import (
	"github.com/urfave/cli/v2"
	"github.com/waldirborbajr/glink/cmds"
)

// App returns the Glink CLI application
func App(version string) cli.App {
	return cli.App{
		Name:    "glink",
		Version: version,
		Usage:   "GOLang symlink manager ",
		Commands: []*cli.Command{
			cmds.List(),
			cmds.Link(),
			cmds.Remove(),
			// cmds.Ignore(),
			cmds.Version(),
		},
	}
}
