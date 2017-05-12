package commands

import (
	"fmt"

	"github.com/urfave/cli"
)

// Remove ...
func Remove(c *cli.Context) error {
	// c.App.Run([]string{"app", "list", "--quiet"})
	if c.String("name") != "" {
		fmt.Print(c.String("name"))
	}
	if c.String("id") != "" {
		fmt.Print(c.String("id"))
	}
	if c.Bool("all") {
		fmt.Printf("%#v\n", c.Bool("all"))
	}
	return nil
}

func removeByName(name string) error {
	return nil
}

func removeByID(id string) error {
	return nil
}

func removeAll() error {
	return nil
}
