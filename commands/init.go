package commands

import (
	"github.com/andefined/twitterfarm/utils"
	"github.com/urfave/cli"
)

// Init ...
func Init(c *cli.Context) error {
	utils.SetHomeDir()
	return nil
}
