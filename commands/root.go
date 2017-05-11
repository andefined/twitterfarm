package commands

import (
	"github.com/spf13/cobra"
)

var config string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "twitterfarm",
	Short: "Quickly collect data from Twitter Streaming API",
	Long:  ``,
}

func init() {
	cobra.OnInitialize()
}
