package main

import (
	"os"

	"github.com/andefined/twitterfarm/commands"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "twitterfarm"
	app.Version = "0.0.2"
	app.Usage = "Collect data from Twitter"
	app.Commands = []cli.Command{
		{
			Name:      "init",
			Usage:     "Initialize twitterfarm. Will create a folder under $HOME/.twitterfarm",
			Action:    commands.Init,
			ArgsUsage: " ",
		},
		{
			Name:      "create",
			Usage:     "Create a new project",
			Action:    commands.Create,
			ArgsUsage: " ",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "config",
					Usage: "Path to your .yml configuration file",
				},
				cli.StringFlag{
					Name:  "name",
					Usage: "Project Name",
				},

				cli.StringFlag{
					Name:  "track",
					Usage: "Tracking Keywords",
				},
				cli.StringFlag{
					Name:  "filter-level",
					Usage: "Filter Level",
					Value: "none",
				},
				cli.StringFlag{
					Name:  "language",
					Usage: "Language",
					Value: "en",
				},
				cli.StringFlag{
					Name:  "location",
					Usage: "Location",
				},

				cli.StringFlag{
					Name:  "consumer-key",
					Usage: "Twitter Consumer Key",
				},
				cli.StringFlag{
					Name:  "consumer-secret",
					Usage: "Twitter Consumer Secret",
				},
				cli.StringFlag{
					Name:  "access-token",
					Usage: "Twitter Access Token",
				},
				cli.StringFlag{
					Name:  "access-token-secret",
					Usage: "Twitter Access Secret",
				},

				cli.StringFlag{
					Name:  "elasticsearch-host",
					Usage: "Comma Separated Elasticsearch Hosts",
				},
				cli.StringFlag{
					Name:  "elasticsearch-index",
					Usage: "Elasticsearch Index",
				},
			},
		},
		{
			Name:      "list",
			Usage:     "List all projects",
			Action:    commands.List,
			ArgsUsage: " ",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "quiet, q",
					Usage: "Print only ID",
				},
			},
		},
		{
			Name:      "test",
			Usage:     "Test project configuration",
			Action:    commands.Test,
			ArgsUsage: `[project id]`,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "create-index, c",
					Usage: "Create the Elasticsearch Index",
				},
			},
		},

		/*{
			Name:   "rm",
			Usage:  "Remove a project",
			Action: commands.Remove,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "all, a",
					Usage: "Remove all projects",
				},
			},
		},*/

		{
			Name:      "start",
			Usage:     "Start a project",
			Action:    commands.Start,
			ArgsUsage: `[project id]`,
		},

		{
			Name:      "stop",
			Usage:     "Stop a project",
			Action:    commands.Stop,
			ArgsUsage: `[project id]`,
		},

		{
			Name:      "restart",
			Usage:     "Restart a project",
			Action:    commands.Restart,
			ArgsUsage: `[project id]`,
		},

		{
			Name:      "exec",
			Usage:     "Execute a project",
			Action:    commands.Exec,
			ArgsUsage: `[project configuration file]`,
		},
	}

	app.Run(os.Args)
}
