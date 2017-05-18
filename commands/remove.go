package commands

import (
	"os"

	"github.com/andefined/twitterfarm/projects"
	"github.com/andefined/twitterfarm/utils"
	"github.com/urfave/cli"
)

// Remove ...
func Remove(c *cli.Context) error {
	if c.Args().Get(0) == "" && !c.Bool("all") {
		cli.ShowSubcommandHelp(c)
		return nil
	}

	// Remove by Project ID
	if c.Args().Get(0) != "" {
		path := utils.GetHomeDir() + "/" + c.Args().Get(0) + ".yml"
		// Stop Project (if running)
		c.App.Run([]string{c.App.Name, "stop", c.Args().Get(0)})
		// Delete project configuration file
		err := os.Remove(path)
		utils.ExitOnError(err)
	}

	// Remove All
	if c.Bool("all") {
		paths := make(chan string, 5)
		utils.GetAllConfigs(paths)

		for path := range paths {
			// Create a temp project
			project := &projects.Project{}
			// Assign values from file
			project.Read(path)
			// Stop Project (if running)
			c.App.Run([]string{c.App.Name, "stop", project.ID})
			// Delete project configuration file
			err := os.Remove(path)
			utils.ExitOnError(err)
		}
	}

	return nil
}
