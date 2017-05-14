package commands

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"

	yaml "gopkg.in/yaml.v2"

	"github.com/andefined/twitterfarm/utils"
	"github.com/urfave/cli"
)

// Start ...
func Start(c *cli.Context) error {
	if c.Args().Get(0) == "" {
		cli.ShowSubcommandHelp(c)
		return nil
	}

	config := utils.GetHomeDir() + "/" + c.Args().Get(0) + ".yml"
	project := utils.ReadProject(config)
	fmt.Printf("%s\n", project.ID)

	cmd := exec.Command("twitterfarm", "exec", config)
	cmd.Stdin = strings.NewReader("")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
		return err
	}
	project.PID = cmd.Process.Pid

	y, err := yaml.Marshal(project)
	if err != nil {
		log.Fatal(err)
		return err
	}

	utils.SaveFile(config, y)

	return nil
}
