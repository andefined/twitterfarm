package commands

import (
	"github.com/andefined/twitterfarm/utils"
	"github.com/urfave/cli"
)

// Init ...
func Init(c *cli.Context) {
	utils.SetHomeDir()
}
