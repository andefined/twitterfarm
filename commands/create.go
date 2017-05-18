package commands

import (
	"fmt"

	"github.com/andefined/twitterfarm/projects"
	"github.com/urfave/cli"
)

// Create : Create a new project
func Create(c *cli.Context) error {
	// New project
	project := &projects.Project{}

	if c.String("config") != "" {
		// Init project values from --config file
		project.Init(c.String("config"))
	} else {
		// Init project values from flags
		project.Name = c.String("name")
		project.Track = c.String("track")
		project.FilterLevel = c.String("filter-level")
		project.Language = c.String("language")
		project.Location = c.String("location")
		project.ConsumerKey = c.String("consumer-key")
		project.ConsumerSecret = c.String("consumer-secret")
		project.AccessToken = c.String("access-token")
		project.AccessTokenSecret = c.String("access-token-secret")
		project.ElasticsearchHost = c.String("elasticsearch-host")
		project.ElasticsearchIndex = c.String("elasticsearch-index")
		project.Init("")
	}

	// Test required values and fallback to help
	if project.ConsumerKey == "" ||
		project.ConsumerSecret == "" ||
		project.AccessToken == "" ||
		project.AccessTokenSecret == "" ||
		project.ElasticsearchHost == "" ||
		project.Track == "" {
		cli.ShowSubcommandHelp(c)
		return nil
	}

	// Output project id if created
	fmt.Printf("%s\n", project.ID)

	return nil
}
