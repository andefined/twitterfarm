package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/andefined/twitterfarm/utils"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli"
)

// List ...
func List(c *cli.Context) error {
	path, err := utils.GetHomeDir()
	if err != nil {
		return err
	}

	paths := make(chan string, 100)
	go (func() error {
		defer close(paths)
		return filepath.Walk(path, func(p string, f os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if f.IsDir() {
				return nil
			}
			select {
			case paths <- p:
			}
			return nil
		})
	})()

	if c.Bool("quiet") {
		renderIDs(paths)
	} else {
		renderTable(paths)
	}

	return nil
}

func renderIDs(paths chan string) {
	for path := range paths {
		project, err := utils.ReadFile(path)
		if err != nil {
			fmt.Print(err)
		} else {
			fmt.Printf("%s\n", project.ID)
		}
	}
}

func renderTable(paths chan string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "PID", "NAME", "STATUS", "KEYWORDS"})
	for path := range paths {
		project, err := utils.ReadFile(path)
		if err != nil {
			fmt.Print(err)
		} else {
			status := "stopped"
			if project.PID > 0 {
				status = "running"
			}
			table.Append([]string{project.ID, strconv.Itoa(project.PID), project.Name, status, project.Keywords})
		}
	}

	table.Render()

}
