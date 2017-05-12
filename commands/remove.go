package commands

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/andefined/twitterfarm/utils"
	"github.com/urfave/cli"
)

// Remove ...
func Remove(c *cli.Context) error {
	if c.Args().Get(0) == "" && !c.Bool("all") {
		cli.ShowSubcommandHelp(c)
		return nil
	}

	if c.Args().Get(0) != "" {
		_, err := removeByID(c.Args().Get(0))
		if err != nil {
			return err
		}
	}

	if c.Bool("all") {
		removeAll()
	}

	return nil
}

func removeByID(id string) (string, error) {
	home, err := utils.GetHomeDir()
	if err != nil {
		return "", err
	}

	config := home + "/" + id + ".yml"
	project := utils.ReadFile(config)

	kill(project.PID)
	delete(config)

	return id, nil
}

func removeAll() (string, error) {
	home, err := utils.GetHomeDir()
	if err != nil {
		return "", err
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

	for path := range paths {
		project := utils.ReadFile(path)
		removeByID(project.ID)
	}

	return "", nil
}

func delete(path string) error {
	err := os.Remove(path)
	if err != nil {
		fmt.Printf("Can't find project: %s exiting\n", path)
		return err
	}
	return nil
}

func kill(pid int) error {
	if pid <= 1 {
		fmt.Printf("Can't find process: %x exiting\n", pid)
		return nil
	}
	proc, err := os.FindProcess(pid)
	if err != nil {
		return err
	}
	proc.Kill()
	return nil
}
