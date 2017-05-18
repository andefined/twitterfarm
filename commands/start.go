package commands

import (
	"os/exec"

	"github.com/andefined/twitterfarm/projects"
	"github.com/andefined/twitterfarm/utils"
	"github.com/urfave/cli"
)

// Start ...
func Start(c *cli.Context) error {
	if c.Args().Get(0) == "" {
		cli.ShowSubcommandHelp(c)
		return nil
	}

	path := utils.GetHomeDir() + "/" + c.Args().Get(0) + ".yml"
	// Create a temp project
	project := &projects.Project{}
	// Assign values from file
	project.Read(path)

	/**
	  Unfortunatelly urfave/cli doesn't support "directly" background processes,
	  so `start` command only executes the `exec` command in the background via os/exec.
	  The idea is to keep a history of the `pid` of the process
	  in order to restart, stop, and start without overlapping processes.
	  The problem is that i am not sure if this workaround going to work in NON-Unix systems.
	*/

	// cmd := exec.Command("go run /home/andefined/go/src/github.com/andefined/twitterfarm/main.go", "exec", path)
	cmd := exec.Command(c.App.Name, "exec", path)

	// var out bytes.Buffer
	// cmd.Stdout = os.Stdout

	err := cmd.Start()
	if err != nil {
		utils.ExitOnError(err)
	}

	project.PID = cmd.Process.Pid
	project.Save(path)

	return nil
}
