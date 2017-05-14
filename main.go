package main

import (
	"os"

	"github.com/andefined/twitterfarm/commands"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "twitterfarm"
	app.Version = "0.0.1"
	app.Usage = "Collect data from Twitter"
	app.Commands = []cli.Command{
		{
			Name:   "create",
			Usage:  "Create a new project",
			Action: commands.Create,
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
					Usage: "Comma Separated Hosts (ex. `http://elastic:changeme@host-a:9200,http://elastic:changeme@host-b:9200`)",
				},
				cli.StringFlag{
					Name:  "elasticsearch-index, i",
					Usage: "Elasticsearch Index (Default Lowercase: twitterfarm_${project_nane}_${project_id})",
				},
				cli.StringFlag{
					Name:  "keywords, k",
					Usage: "Keyword to stream",
				},
			},
		},

		{
			Name:   "list",
			Usage:  "List all projects",
			Action: commands.List,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "quiet, q",
					Usage: "Print only ID",
				},
			},
		},

		{
			Name:   "test",
			Usage:  "Test project configuration, connections etc..",
			Action: commands.Test,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "fix, f",
					Usage: "Try to fix errors (ex. create the index)",
				},
			},
		},

		{
			Name:   "remove",
			Usage:  "Remove a project",
			Action: commands.Remove,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "all, a",
					Usage: "Remove all projects",
				},
			},
		},

		{
			Name:   "start",
			Usage:  "Start a project",
			Action: commands.Start,
		},

		{
			Name:   "exec",
			Usage:  "Execute a project",
			Action: commands.Exec,
		},
	}
	app.Run(os.Args)
}
