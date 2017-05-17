package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"

	"github.com/andefined/twitterfarm/utils"
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
	fmt.Print("\n")
	for path := range paths {
		project := utils.ReadProject(path)
		proc, _ := os.FindProcess(project.PID)
		testproc := proc.Signal(syscall.Signal(0))
		fmt.Printf("ID: %s | PID: %s | Status: %s | Name: %.10s | Track: %.20s\n", project.ID, strconv.Itoa(project.PID), strings.Split(testproc.Error(), "os: process ")[1], project.Name, project.Track)
		fmt.Print("----------------\n")
	}
}
