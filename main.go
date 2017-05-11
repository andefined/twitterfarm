package main

import (
	"fmt"
	"os"
	"time"

	yaml "gopkg.in/yaml.v2"

	"github.com/andefined/twitterfarm/utils"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "twitterfarm"
	app.Version = "0.0.1"
	app.EnableBashCompletion = true
	app.Usage = "Quickly collect data from Twitter Streaming API"
	app.Commands = []cli.Command{
		{
			Name:  "create",
			Usage: "Create a new project",
			Action: func(c *cli.Context) error {
				project := utils.Project{
					ID:                 utils.ID(5),
					Name:               c.String("name"),
					ConsumerKey:        c.String("consumer-key"),
					ConsumerSecret:     c.String("consumer-secret"),
					AccessToken:        c.String("access-token"),
					AccessTokenSecret:  c.String("access-token-secret"),
					ElasticsearchHost:  c.String("elasticsearch-host"),
					ElasticsearchIndex: c.String("elasticsearch-index"),
					Keywords:           c.String("keywords"),
					DateCreated:        time.Now(),
					PID:                0,
				}
				if project.Name == "" {
					project.Name = project.ID
				}

				if project.ConsumerKey == "" || project.ConsumerSecret == "" || project.AccessToken == "" || project.AccessTokenSecret == "" {
					cli.ShowSubcommandHelp(c)
					os.Exit(1)
				}

				if project.ElasticsearchHost == "" {
					cli.ShowSubcommandHelp(c)
					os.Exit(1)
				}

				if project.ElasticsearchIndex == "" {
					project.ElasticsearchIndex = "twitterfarm" + "-" + project.Name + "-" + project.ID
				}

				if project.Keywords == "" {
					cli.ShowSubcommandHelp(c)
					os.Exit(1)
				}

				y, err := yaml.Marshal(project)
				if err != nil {
					return err
				}

				home, err := homedir.Dir()
				if err != nil {
					return err
				}

				path := home + "/.twitterfarm"
				if _, err = os.Stat(path); os.IsNotExist(err) {
					os.Mkdir(path, os.ModePerm)
				}

				config := path + "/" + project.ID + ".yml"
				if _, err = os.Stat(config); err == nil {
					return err
				}

				err = utils.CreateFile(config, y)
				if err != nil {
					return err
				}

				fmt.Printf("Project created: %s\n", project.ID)

				return nil
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name, n",
					Usage: "Project name",
				},
				cli.StringFlag{
					Name:  "consumer-key, c",
					Usage: "Twitter Consumer Key",
				},
				cli.StringFlag{
					Name:  "consumer-secret, s",
					Usage: "Twitter Consumer Secret",
				},
				cli.StringFlag{
					Name:  "access-token, t",
					Usage: "Twitter Access Token",
				},
				cli.StringFlag{
					Name:  "access-token-secret, a",
					Usage: "Twitter Access Secret",
				},
				cli.StringFlag{
					Name:  "elasticsearch-host, e",
					Usage: "Comma Seperated Hosts (ex. `user:pass@es-1.clu:9200,user:pass@es-2.clu:9200`)",
				},
				cli.StringFlag{
					Name:  "elasticsearch-index, i",
					Usage: "Elasticsearch Index",
				},
				cli.StringFlag{
					Name:  "keywords, k",
					Usage: "Keyword to stream",
				},
			},
		},

		{
			Name:  "list",
			Usage: "List all projects",
			Action: func(c *cli.Context) error {
				return nil
			},
		},

		{
			Name:  "run",
			Usage: "Run a project",
			Action: func(c *cli.Context) error {
				return nil
			},
		},

		{
			Name:  "remove",
			Usage: "Remove a project",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
	}
	app.Run(os.Args)
}
