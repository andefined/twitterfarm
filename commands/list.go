package commands

import (
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/andefined/twitterfarm/utils"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all projects",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		home, err := homedir.Dir()
		if err != nil {
			log.Fatal(err)
		}

		path := home + "/.twitterfarm"
		if _, err = os.Stat(path); os.IsNotExist(err) {
			os.Mkdir(path, os.ModePerm)
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

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "PID", "NAME", "STATUS", "KEYWORDS"})

		for p := range paths {
			var c = utils.Project{}
			c.ReadFile(p)
			status := "stopped"
			if c.PID > 0 {
				status = "running"
			}
			table.Append([]string{c.ID, strconv.Itoa(c.PID), c.Name, status, c.Keyword})
		}

		table.Render()
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
