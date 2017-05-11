package commands

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/andefined/twitterfarm/utils"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all projects",
	Long:  ``,
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Print("\n")
	},
	Run: func(cmd *cobra.Command, args []string) {
		l := log.New(os.Stdout, "[TwitterFarm] ", log.Ldate|log.Ltime)

		home, err := homedir.Dir()
		if err != nil {
			l.Fatal(err)
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
				select {
				case paths <- p:
				}
				return nil
			})
		})()
		fmt.Print("ID         NAME \tKEYWORDS\n")
		for p := range paths {
			var c = utils.Project{}
			c.ReadFile(p)
			fmt.Printf("%s %s \t%s \n", c.ID, c.Name, c.Keyword)
			// l.Print(p)
		}

		fmt.Print("\n")
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
