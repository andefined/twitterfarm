package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the list command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Twitterfarm",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Twitterfarm v0.1")
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
