package commands

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	yaml "gopkg.in/yaml.v2"

	"github.com/andefined/twitterfarm/utils"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli"
)

// List ...
func List(c *cli.Context) error {
	home, err := utils.GetHomeDir()
	if err != nil {
		return err
	}

	paths := make(chan string, 100)
	go (func() error {
		defer close(paths)
		return filepath.Walk(home, func(p string, f os.FileInfo, err error) error {
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
		project := utils.ReadFile(path)
		fmt.Printf("%s\n", project.ID)
	}
}

func renderTable(paths chan string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "PID", "NAME", "STATUS", "KEYWORDS"})
	for path := range paths {
		project := utils.ReadFile(path)
		status := "stopped"
		if project.PID > 0 {
			status = "running"
		}
		proc, err := os.FindProcess(project.PID)
		if err != nil {
			project.PID = 0
			status = "stopped"
			y, err := yaml.Marshal(project)
			if err != nil {
				log.Fatal(err)
			}

			utils.SaveFile(path, y)
		}
		fmt.Print(proc.Pid)
		table.Append([]string{project.ID, strconv.Itoa(project.PID), project.Name, status, project.Keywords})
	}

	table.Render()

}
