package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"syscall"

	"github.com/andefined/twitterfarm/utils"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli"
)

// List ...
func List(c *cli.Context) error {
	home := utils.GetHomeDir()
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
		project := utils.ReadProject(path)
		fmt.Printf("%s\n", project.ID)
	}
}

func renderTable(paths chan string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "PID", "NAME", "STATUS", "KEYWORDS"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.SetAutoWrapText(false)
	for path := range paths {
		project := utils.ReadProject(path)
		status := "stopped"
		if project.PID > 0 {
			status = "running"
		}
		proc, _ := os.FindProcess(project.PID)
		testproc := proc.Signal(syscall.Signal(0))
		table.Append([]string{project.ID, testproc.Error() + " (" + strconv.Itoa(project.PID) + ")", project.Name, status, project.Keywords})
	}

	table.Render()

}
