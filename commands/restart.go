package commands

import (
	"github.com/urfave/cli"
)

// Restart ...
func Restart(c *cli.Context) error {
	if c.Args().Get(0) == "" {
		cli.ShowSubcommandHelp(c)
		return nil
	}

	c.App.Run([]string{c.App.Name, "stop", c.Args().Get(0)})
	c.App.Run([]string{c.App.Name, "start", c.Args().Get(0)})

	return nil
}
