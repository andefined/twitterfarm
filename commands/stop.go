package commands

import (
	"os"

	"github.com/andefined/twitterfarm/projects"
	"github.com/andefined/twitterfarm/utils"
	"github.com/urfave/cli"
)

// Stop ...
func Stop(c *cli.Context) error {
	if c.Args().Get(0) == "" {
		cli.ShowSubcommandHelp(c)
		return nil
	}

	path := utils.GetHomeDir() + "/" + c.Args().Get(0) + ".yml"
	// Create a temp project
	project := &projects.Project{}
	// Assign values from file
	project.Read(path)

	proc, _ := os.FindProcess(project.PID)
	proc.Kill()

	return nil
}
