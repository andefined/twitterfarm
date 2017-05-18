package commands

import (
	"fmt"

	"github.com/andefined/twitterfarm/projects"
	"github.com/andefined/twitterfarm/utils"
	"github.com/urfave/cli"
)

// Test : Test a specific project
func Test(c *cli.Context) error {
	if c.Args().Get(0) == "" {
		cli.ShowSubcommandHelp(c)
		return nil
	}

	path := utils.GetHomeDir() + "/" + c.Args().Get(0) + ".yml"
	// Create a temp project
	project := &projects.Project{}
	// Assign values from file
	project.Read(path)
	// Force create elasticsearch index
	createIndex := false
	if c.Bool("create-index") {
		createIndex = true
	}
	// Test if there is a living connection for elasticsearch && if index created
	isESHost, isESIndex := project.TestElasticsearch(createIndex)

	// Output
	fmt.Printf("%-12s | %-12v | %-12v | %v\n", "ID", "TWITTER API", "ELASTIC HOST", "ELASTIC INDEX")
	fmt.Printf("%14s %14s %14s\n", "-", "-", "-")
	fmt.Printf("%-12s | %-12v | %-12v | %v\n",
		project.ID,
		project.TestTwitter(),
		isESHost,
		isESIndex,
	)

	return nil
}
