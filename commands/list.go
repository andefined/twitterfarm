package commands

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"syscall"

	"github.com/andefined/twitterfarm/projects"
	"github.com/andefined/twitterfarm/utils"
	"github.com/urfave/cli"
)

// List : List all projects
func List(c *cli.Context) error {
	// Read concurrent
	paths := make(chan string, 5)
	utils.GetAllConfigs(paths)

	if c.Bool("quiet") {
		// Output only IDs
		renderIDs(paths)
	} else {
		// Output Normal
		renderTable(paths)
	}

	return nil
}

func renderIDs(paths chan string) {
	for path := range paths {
		// Create a temp project
		project := &projects.Project{}
		// Assign values from file
		project.Read(path)
		// Output
		fmt.Printf("%s\n", project.ID)
	}
}

func renderTable(paths chan string) {
	fmt.Printf("%-10s | %-5s | %-16s | %-16s | %s \n", "ID", "PID", "STATUS", "NAME", "TRACK")
	fmt.Printf("%12s %7s %18s %18s\n", "-", "-", "-", "-")
	for path := range paths {
		// Create a temp project
		project := &projects.Project{}
		// Assign values from file
		project.Read(path)
		// Find project process
		proc, _ := os.FindProcess(project.PID)
		// Send Signal(0) to process, to test if running
		status := "running"
		err := proc.Signal(syscall.Signal(0))
		if err != nil {
			status = strings.Split(err.Error(), "os: process ")[1]
		}
		// Output
		fmt.Printf("%-10s | %-5s | %-16s | %-16s | %s\n",
			project.ID,
			strconv.Itoa(project.PID),
			status, // testproc.Error(), //strings.Split(testproc.Error(), "os: process ")[0],
			utils.TruncateString(16, project.Name),
			utils.TruncateString(24, project.Track),
		)
	}
}
