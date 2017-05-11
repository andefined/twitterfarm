package commands

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var config string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "twitterfarm",
	Short: "Quickly collect data from Twitter Streaming API",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

func init() {
	cobra.OnInitialize(initConfig)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	RootCmd.PersistentFlags().StringVar(&config, "config", "", "config file (default is $HOME/.twitterfarm/config.yml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if config != "" {
		// Use config file from the flag.
		viper.SetConfigFile(config)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(home)
			os.Exit(1)
		}
		// Create conf folder if not exists
		/*path := home + "/.twitterfarm"
		if _, err := os.Stat(path); os.IsNotExist(err) {
			os.Mkdir(path, os.ModePerm)
		}

		config = path + "/config.yml"
		if _, err := os.Stat(path + "/config.yml"); os.IsNotExist(err) {
			_, err := os.Create(config)
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(0)
			}
		}*/
		// Search config in home directory with name ".twitterfarm" (without extension).
		// viper.AddConfigPath(path)
		viper.SetConfigType("yml")
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
