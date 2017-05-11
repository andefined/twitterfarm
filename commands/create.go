package commands

import (
	"fmt"
	"log"
	"os"
	"time"

	yaml "gopkg.in/yaml.v2"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	utils "github.com/andefined/twitterfarm/utils"
)

var (
	// createCmd represents the create command
	createCmd = &cobra.Command{
		Use:   "create",
		Short: "Create a new project",
		Long:  ``,
		PreRun: func(cmd *cobra.Command, args []string) {
			fmt.Print("\n")
		},
		Run: func(cmd *cobra.Command, args []string) {
			l := log.New(os.Stdout, "[TwitterFarm] ", log.Ldate|log.Ltime)
			project := utils.Project{
				ID:                 utils.ID(5),
				Name:               viper.GetString("name"),
				ConsumerKey:        viper.GetString("consumer-key"),
				ConsumerSecret:     viper.GetString("consumer-secret"),
				AccessToken:        viper.GetString("access-token"),
				AccessTokenSecret:  viper.GetString("access-token-secret"),
				ElasticsearchHost:  viper.GetString("elasticsearch-host"),
				ElasticsearchIndex: viper.GetString("elasticsearch-index"),
				Keyword:            viper.GetString("keyword"),
				DateCreated:        time.Now(),
			}

			if project.Name == "" {
				project.Name = project.ID
			}

			if project.ConsumerKey == "" || project.ConsumerSecret == "" || project.AccessToken == "" || project.AccessTokenSecret == "" {
				l.Fatal("Consumer key/secret and Access token/secret required")
			}

			if project.ElasticsearchHost == "" {
				l.Fatal("Elasticsearch Host required")
			}

			if project.ElasticsearchIndex == "" {
				project.ElasticsearchIndex = "twitterfarm" + "-" + project.Name + "-" + project.ID
			}

			if project.Keyword == "" {
				l.Fatal("Keyword required")
			}

			y, err := yaml.Marshal(project)
			if err != nil {
				l.Fatal(err)
			}

			home, err := homedir.Dir()
			if err != nil {
				l.Fatal(err)
			}

			path := home + "/.twitterfarm"
			if _, err = os.Stat(path); os.IsNotExist(err) {
				os.Mkdir(path, os.ModePerm)
			}

			config := path + "/" + project.Name + ".yml"
			if _, err = os.Stat(config); err == nil {
				l.Fatalf("Project `%s` allready exists\n", project.Name)
			}

			err = utils.CreateFile(config, y)
			if err != nil {
				l.Fatalf("An error occured while saving `%s`\n", project.Name)
			}
			l.Printf("Project `%s` created`", project.ID)
			l.Printf("You can start by running `twitterfarm run %s`\n\n", project.Name)
		},
	}
)

func init() {
	RootCmd.AddCommand(createCmd)

	// Parse the flags
	createCmd.PersistentFlags().StringP("name", "n", "", "Project name")
	createCmd.PersistentFlags().StringP("consumer-key", "c", "", "Twitter Consumer Key")
	createCmd.PersistentFlags().StringP("consumer-secret", "s", "", "Twitter Consumer Secret")
	createCmd.PersistentFlags().StringP("access-token", "t", "", "Twitter Access Token")
	createCmd.PersistentFlags().StringP("access-token-secret", "a", "", "Twitter Access Secret")
	createCmd.PersistentFlags().StringP("elasticsearch-host", "e", "", "Comma Seperated Hosts (ex. `user:pass@es-1.clu:9200,user:pass@es-2.clu:9200`)")
	createCmd.PersistentFlags().StringP("elasticsearch-index", "i", "", "Elasticsearch Index (default Project Name)")
	createCmd.PersistentFlags().StringP("keyword", "k", "", "Keyword(s) to stream")

	viper.BindPFlag("name", createCmd.PersistentFlags().Lookup("name"))
	viper.BindPFlag("consumer-key", createCmd.PersistentFlags().Lookup("consumer-key"))
	viper.BindPFlag("consumer-secret", createCmd.PersistentFlags().Lookup("consumer-secret"))
	viper.BindPFlag("access-token", createCmd.PersistentFlags().Lookup("access-token"))
	viper.BindPFlag("access-token-secret", createCmd.PersistentFlags().Lookup("access-token-secret"))
	viper.BindPFlag("elasticsearch-host", createCmd.PersistentFlags().Lookup("elasticsearch-host"))
	viper.BindPFlag("elasticsearch-index", createCmd.PersistentFlags().Lookup("elasticsearch-index"))
	viper.BindPFlag("keyword", createCmd.PersistentFlags().Lookup("keyword"))
}
